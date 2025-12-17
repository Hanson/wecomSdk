package wecom

import (
    "bytes"
    "context"
    "encoding/json"
    "mime/multipart"
    "net/http"
    "net/url"
)

type MediaUploadResponse struct { APIError; Type string `json:"type"`; MediaID string `json:"media_id"`; CreatedAt int64 `json:"created_at,omitempty"` }

func (c *Client) UploadMedia(ctx context.Context, typ, filename string, data []byte) (*MediaUploadResponse, error) {
    t, err := c.token.Get(ctx)
    if err != nil { return nil, err }
    u, _ := url.Parse(c.cfg.BaseURL)
    u.Path = "/cgi-bin/media/upload"
    q := u.Query(); q.Set("access_token", t); q.Set("type", typ); u.RawQuery = q.Encode()
    var buf bytes.Buffer
    w := multipart.NewWriter(&buf)
    fw, err := w.CreateFormFile("media", filename); if err != nil { return nil, err }
    if _, err := fw.Write(data); err != nil { return nil, err }
    _ = w.Close()
    req, _ := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), &buf)
    req.Header.Set("Content-Type", w.FormDataContentType())
    resp, err := c.cfg.HTTPClient.Do(req); if err != nil { return nil, err }
    defer resp.Body.Close()
    var r MediaUploadResponse
    dec := json.NewDecoder(resp.Body)
    if err := dec.Decode(&r); err != nil { return nil, err }
    if r.ErrCode != 0 { b, _ := json.Marshal(r); return nil, &Error{Code: r.ErrCode, Message: r.ErrMsg, Raw: b} }
    return &r, nil
}
