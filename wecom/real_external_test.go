package wecom

import (
	"context"
	"encoding/json"
	"net/url"
	"os"
	"testing"
)

func TestRealExternalContact(t *testing.T) {
	corpID := os.Getenv("WECOM_CORP_ID")
	secret := os.Getenv("WECOM_APP_SECRET")
	userID := os.Getenv("WECOM_USER_ID")
	if userID == "" {
		userID = os.Getenv("WECOM_TO_USER")
	}
	if corpID == "" || secret == "" || userID == "" {
		t.Skip("missing env: WECOM_CORP_ID/WECOM_APP_SECRET/WECOM_USER_ID")
	}
	c, err := NewClient(Config{CorpID: corpID, CorpSecret: secret})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	ctx := context.Background()
	var listOut externalContactListResp
	err = c.GetJSON(ctx, "/cgi-bin/externalcontact/list", url.Values{"userid": []string{userID}}, &listOut)
	if err != nil {
		if e, ok := err.(*Error); ok && (e.Code == 60020 || e.Code == 48002) {
			t.Skip("ip or permission not ready")
		}
		t.Fatalf("list external contacts: %v", err)
	}
	if len(listOut.ExternalUserID) == 0 {
		t.Skip("no external contacts bound to user: " + userID)
	}
	b1, _ := json.Marshal(listOut.ExternalUserID)
	t.Logf("external_userid: %s", string(b1))
	// get the first external contact
	var getOut externalContactGetResp
	err = c.GetJSON(ctx, "/cgi-bin/externalcontact/get", url.Values{"external_userid": []string{listOut.ExternalUserID[0]}}, &getOut)
	if err != nil {
		if e, ok := err.(*Error); ok && (e.Code == 60020 || e.Code == 48002) {
			t.Skip("ip or permission not ready")
		}
		t.Fatalf("get external contact: %v", err)
	}
	if getOut.ExternalContact.ExternalUserID == "" {
		t.Fatalf("empty external user id")
	}
	b2, _ := json.Marshal(getOut.ExternalContact)
	t.Logf("external_contact: %s", string(b2))
}
