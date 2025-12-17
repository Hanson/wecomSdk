package wecom

import (
    "context"
    "encoding/json"
    "net/http"
)

type TextMessage struct {
    ToUser  string `json:"touser,omitempty"`
    ToParty string `json:"toparty,omitempty"`
    ToTag   string `json:"totag,omitempty"`
    MsgType string `json:"msgtype"`
    AgentID int    `json:"agentid"`
    Text    struct { Content string `json:"content"` } `json:"text"`
    Safe    int `json:"safe,omitempty"`
}

func (c *Client) SendTextMessage(ctx context.Context, msg TextMessage) (*MessageSendResponse, error) {
    if msg.MsgType == "" { msg.MsgType = "text" }
    b, err := c.do(ctx, http.MethodPost, "/cgi-bin/message/send", nil, msg)
    if err != nil { return nil, err }
    var r MessageSendResponse
    if err := json.Unmarshal(b, &r); err != nil { return nil, err }
    return &r, nil
}
