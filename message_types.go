package wecom

import (
    "context"
    "encoding/json"
    "net/http"
)

type MarkdownMessage struct { ToUser string `json:"touser,omitempty"`; ToParty string `json:"toparty,omitempty"`; ToTag string `json:"totag,omitempty"`; MsgType string `json:"msgtype"`; AgentID int `json:"agentid"`; Markdown struct { Content string `json:"content"` } `json:"markdown"` }
type ImageMessage struct { ToUser string `json:"touser,omitempty"`; ToParty string `json:"toparty,omitempty"`; ToTag string `json:"totag,omitempty"`; MsgType string `json:"msgtype"`; AgentID int `json:"agentid"`; Image struct { MediaID string `json:"media_id"` } `json:"image"` }
type FileMessage struct { ToUser string `json:"touser,omitempty"`; ToParty string `json:"toparty,omitempty"`; ToTag string `json:"totag,omitempty"`; MsgType string `json:"msgtype"`; AgentID int `json:"agentid"`; File struct { MediaID string `json:"media_id"` } `json:"file"` }
type NewsMessage struct { ToUser string `json:"touser,omitempty"`; ToParty string `json:"toparty,omitempty"`; ToTag string `json:"totag,omitempty"`; MsgType string `json:"msgtype"`; AgentID int `json:"agentid"`; News struct { Articles []struct { Title string `json:"title"`; Description string `json:"description,omitempty"`; URL string `json:"url"`; PicURL string `json:"picurl,omitempty"` } `json:"articles"` } `json:"news"` }

func (c *Client) SendMarkdownMessage(ctx context.Context, msg MarkdownMessage) (*MessageSendResponse, error) {
    if msg.MsgType == "" { msg.MsgType = "markdown" }
    b, err := c.do(ctx, http.MethodPost, "/cgi-bin/message/send", nil, msg)
    if err != nil { return nil, err }
    var r MessageSendResponse
    if err := json.Unmarshal(b, &r); err != nil { return nil, err }
    return &r, nil
}

func (c *Client) SendImageMessage(ctx context.Context, msg ImageMessage) (*MessageSendResponse, error) {
    if msg.MsgType == "" { msg.MsgType = "image" }
    b, err := c.do(ctx, http.MethodPost, "/cgi-bin/message/send", nil, msg)
    if err != nil { return nil, err }
    var r MessageSendResponse
    if err := json.Unmarshal(b, &r); err != nil { return nil, err }
    return &r, nil
}

func (c *Client) SendFileMessage(ctx context.Context, msg FileMessage) (*MessageSendResponse, error) {
    if msg.MsgType == "" { msg.MsgType = "file" }
    b, err := c.do(ctx, http.MethodPost, "/cgi-bin/message/send", nil, msg)
    if err != nil { return nil, err }
    var r MessageSendResponse
    if err := json.Unmarshal(b, &r); err != nil { return nil, err }
    return &r, nil
}

func (c *Client) SendNewsMessage(ctx context.Context, msg NewsMessage) (*MessageSendResponse, error) {
    if msg.MsgType == "" { msg.MsgType = "news" }
    b, err := c.do(ctx, http.MethodPost, "/cgi-bin/message/send", nil, msg)
    if err != nil { return nil, err }
    var r MessageSendResponse
    if err := json.Unmarshal(b, &r); err != nil { return nil, err }
    return &r, nil
}
