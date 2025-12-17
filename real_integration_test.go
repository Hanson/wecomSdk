package wecom

import (
    "context"
    "os"
    "strconv"
    "testing"
)

func TestRealTokenAndAgent(t *testing.T) {
    corpID := os.Getenv("WECOM_CORP_ID")
    secret := os.Getenv("WECOM_APP_SECRET")
    agentStr := os.Getenv("WECOM_AGENT_ID")
    if corpID == "" || secret == "" || agentStr == "" { t.Skip("missing WECOM_CORP_ID / WECOM_APP_SECRET / WECOM_AGENT_ID") }
    agentID, err := strconv.Atoi(agentStr); if err != nil { t.Fatalf("invalid WECOM_AGENT_ID: %v", err) }
    c, err := NewClient(Config{CorpID: corpID, CorpSecret: secret}); if err != nil { t.Fatalf("new client: %v", err) }
    ctx := context.Background()
    tok, err := c.GetAccessToken(ctx)
    if err != nil || tok == "" {
        if e, ok := err.(*Error); ok && e.Code == 60020 { t.Skip("可信IP未配置，跳过真实接口测试") }
        t.Fatalf("get token failed: %v", err)
    }
    a, err := c.GetAgent(ctx, agentID)
    if err != nil {
        if e, ok := err.(*Error); ok && e.Code == 60020 { t.Skip("可信IP未配置，跳过真实接口测试") }
        t.Fatalf("get agent failed: %v", err)
    }
    if a.AgentID != agentID { t.Fatalf("agentid mismatch: %d", a.AgentID) }
}
