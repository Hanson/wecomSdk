package wecom

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type AgentInfo struct {
	APIError
	AgentID       int    `json:"agentid"`
	Name          string `json:"name"`
	SquareLogoURL string `json:"square_logo_url,omitempty"`
	Description   string `json:"description,omitempty"`
}

func (c *Client) GetAgent(ctx context.Context, agentid int) (*AgentInfo, error) {
	q := url.Values{"agentid": []string{intToString(agentid)}}
	b, err := c.do(ctx, http.MethodGet, "/cgi-bin/agent/get", q, nil)
	if err != nil {
		return nil, err
	}
	var r AgentInfo
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
