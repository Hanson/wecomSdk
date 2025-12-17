package wecom

import (
    "context"
    "os"
    "strconv"
    "testing"
)

func TestRealSendTextMessage(t *testing.T) {
    corpID := os.Getenv("WECOM_CORP_ID")
    secret := os.Getenv("WECOM_APP_SECRET")
    agentStr := os.Getenv("WECOM_AGENT_ID")
    toUser := os.Getenv("WECOM_TO_USER")
    if toUser == "" { toUser = "demo_user" }
    if corpID == "" || secret == "" || agentStr == "" || toUser == "" { t.Skip("missing env: WECOM_CORP_ID/WECOM_APP_SECRET/WECOM_AGENT_ID/WECOM_TO_USER") }
    agentID, err := strconv.Atoi(agentStr); if err != nil { t.Fatalf("invalid WECOM_AGENT_ID: %v", err) }
    c, err := NewClient(Config{CorpID: corpID, CorpSecret: secret}); if err != nil { t.Fatalf("new client: %v", err) }
    ctx := context.Background()
    msg := TextMessage{ToUser: toUser, AgentID: agentID}
    msg.Text.Content = "hello from real test"
    r, err := c.SendTextMessage(ctx, msg)
    if err != nil {
        if e, ok := err.(*Error); ok && (e.Code == 60020 || e.Code == 48002) { t.Skip("ip or permission not ready") }
        t.Fatalf("send failed: %v", err)
    }
    if r.InvalidUser != "" { t.Skip("invaliduser: " + r.InvalidUser) }
    if r.MsgID == "" { t.Fatalf("empty msgid") }
}
