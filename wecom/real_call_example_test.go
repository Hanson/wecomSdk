package wecom

import (
	"context"
	"encoding/json"
	"os"
	"testing"
)

// 演示显式路径+参数对象的通用调用方式：更新成员
func TestRealUpdateUserCallJSON(t *testing.T) {
	corpID := os.Getenv("WECOM_CORP_ID")
	secret := os.Getenv("WECOM_APP_SECRET")
	uid := os.Getenv("WECOM_UPDATE_USER_ID")
	if uid == "" {
		uid = os.Getenv("WECOM_TO_USER")
	}
	if uid == "" {
		uid = "demo_user"
	}
	if corpID == "" || secret == "" || uid == "" {
		t.Skip("missing env: WECOM_CORP_ID/WECOM_APP_SECRET/WECOM_UPDATE_USER_ID")
	}
	c, err := NewClient(Config{CorpID: corpID, CorpSecret: secret})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	ctx := context.Background()
	path := "/cgi-bin/user/update"
	payload := User{UserID: uid}
	bPayload, _ := json.Marshal(payload)
	t.Logf("POST %s body=%s", path, string(bPayload))
	var out APIError
	err = c.PostJSON(ctx, path, payload, &out)
	if err != nil {
		if e, ok := err.(*Error); ok && (e.Code == 60020 || e.Code == 48002) {
			t.Skip("ip or permission not ready")
		}
		t.Fatalf("call failed: %v", err)
	}
	bOut, _ := json.Marshal(out)
	t.Logf("response=%s", string(bOut))
}
