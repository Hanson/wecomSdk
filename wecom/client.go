package wecom

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	CorpID     string
	CorpSecret string
	BaseURL    string
	HTTPClient *http.Client
}

type Client struct {
	cfg   Config
	token TokenProvider
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://qyapi.weixin.qq.com"
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}
	tp := NewDefaultTokenProvider(cfg.CorpID, cfg.CorpSecret, cfg.BaseURL, cfg.HTTPClient)
	return &Client{cfg: cfg, token: tp}, nil
}

func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
	return c.token.Get(ctx)
}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any) ([]byte, error) {
	t, err := c.token.Get(ctx)
	if err != nil {
		return nil, err
	}
	u, _ := url.Parse(c.cfg.BaseURL)
	u.Path = path
	q := u.Query()
	if query != nil {
		for k := range query {
			for _, v := range query[k] {
				q.Add(k, v)
			}
		}
	}
	q.Set("access_token", t)
	u.RawQuery = q.Encode()
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rdr = bytes.NewReader(b)
	}
	req, _ := http.NewRequestWithContext(ctx, method, u.String(), rdr)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ae APIError
	_ = json.Unmarshal(b, &ae)
	if ae.ErrCode != 0 {
		return nil, &Error{Code: ae.ErrCode, Message: ae.ErrMsg, Raw: b}
	}
	return b, nil
}

func (c *Client) request(ctx context.Context, method string, path string, query url.Values, body any) ([]byte, APIError, error) {
	t, err := c.token.Get(ctx)
	if err != nil {
		return nil, APIError{}, err
	}
	u, _ := url.Parse(c.cfg.BaseURL)
	u.Path = path
	q := u.Query()
	if query != nil {
		for k := range query {
			for _, v := range query[k] {
				q.Add(k, v)
			}
		}
	}
	q.Set("access_token", t)
	u.RawQuery = q.Encode()
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, APIError{}, err
		}
		rdr = bytes.NewReader(b)
	}
	req, _ := http.NewRequestWithContext(ctx, method, u.String(), rdr)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, APIError{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, APIError{}, err
	}
	var ae APIError
	_ = json.Unmarshal(b, &ae)
	return b, ae, nil
}

func (c *Client) CallJSON(ctx context.Context, method string, path string, query url.Values, body any, out any) error {
	b, ae, err := c.request(ctx, method, path, query, body)
	if err != nil {
		return err
	}
	if out != nil {
		_ = json.Unmarshal(b, out)
	}
	if ae.ErrCode != 0 {
		return &Error{Code: ae.ErrCode, Message: ae.ErrMsg, Raw: b}
	}
	return nil
}

func (c *Client) GetJSON(ctx context.Context, path string, query url.Values, out any) error {
	return c.CallJSON(ctx, http.MethodGet, path, query, nil, out)
}

func (c *Client) PostJSON(ctx context.Context, path string, body any, out any) error {
	return c.CallJSON(ctx, http.MethodPost, path, nil, body, out)
}

func (c *Client) GetJSONWithReq(ctx context.Context, path string, req any, out any) error {
	q := toQueryValues(req)
	return c.GetJSON(ctx, path, q, out)
}

func (c *Client) GetRaw(ctx context.Context, path string, query url.Values) ([]byte, *Error, error) {
	b, ae, err := c.request(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return nil, nil, err
	}
	if ae.ErrCode != 0 {
		return b, &Error{Code: ae.ErrCode, Message: ae.ErrMsg, Raw: b}, nil
	}
	return b, nil, nil
}
