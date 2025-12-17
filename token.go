package wecom

import (
    "context"
    "encoding/json"
    "net/http"
    "net/url"
    "sync"
    "time"
)

type TokenProvider interface { Get(ctx context.Context) (string, error) }

type DefaultTokenProvider struct {
    corpID     string
    corpSecret string
    baseURL    string
    httpClient *http.Client
    mu         sync.RWMutex
    token      string
    expiresAt  time.Time
}

type tokenResp struct { APIError; AccessToken string `json:"access_token"`; ExpiresIn int `json:"expires_in"` }

func NewDefaultTokenProvider(corpID, corpSecret, baseURL string, httpClient *http.Client) *DefaultTokenProvider {
    return &DefaultTokenProvider{corpID: corpID, corpSecret: corpSecret, baseURL: baseURL, httpClient: httpClient}
}

func (p *DefaultTokenProvider) Get(ctx context.Context) (string, error) {
    p.mu.RLock()
    if p.token != "" && time.Now().Before(p.expiresAt) { t := p.token; p.mu.RUnlock(); return t, nil }
    p.mu.RUnlock()
    p.mu.Lock(); defer p.mu.Unlock()
    if p.token != "" && time.Now().Before(p.expiresAt) { return p.token, nil }
    u, _ := url.Parse(p.baseURL)
    u.Path = "/cgi-bin/gettoken"
    q := u.Query(); q.Set("corpid", p.corpID); q.Set("corpsecret", p.corpSecret); u.RawQuery = q.Encode()
    req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
    resp, err := p.httpClient.Do(req); if err != nil { return "", err }
    defer resp.Body.Close()
    var tr tokenResp
    dec := json.NewDecoder(resp.Body)
    if err := dec.Decode(&tr); err != nil { return "", err }
    if tr.ErrCode != 0 { b, _ := json.Marshal(tr); return "", &Error{Code: tr.ErrCode, Message: tr.ErrMsg, Raw: b} }
    if tr.ExpiresIn <= 0 { tr.ExpiresIn = 7200 }
    buf := 60; if tr.ExpiresIn <= buf+1 { buf = 0 }
    p.token = tr.AccessToken
    p.expiresAt = time.Now().Add(time.Duration(tr.ExpiresIn-buf) * time.Second)
    return p.token, nil
}
