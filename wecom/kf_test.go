package wecom

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestKfModule(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/cgi-bin/gettoken", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "access_token": "T", "expires_in": 7200})
	})
	mux.HandleFunc("/cgi-bin/kf/account/list", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "account_list": []map[string]any{{"open_kfid": "kf_123", "name": "客服1"}}})
	})
	mux.HandleFunc("/cgi-bin/kf/send_msg", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "msgid": "m-kf"})
	})
	mux.HandleFunc("/cgi-bin/kf/sync_msg", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "msg_list": []map[string]any{{"msgid": "m1"}}, "has_more": false})
	})
	mux.HandleFunc("/cgi-bin/kf/service_state/get", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "service_state": 1, "servicer_userid": "u1"})
	})
	mux.HandleFunc("/cgi-bin/kf/service_state/trans", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
	})
	mux.HandleFunc("/cgi-bin/kf/add_contact_way", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "url": "https://kf", "qr_code": "q"})
	})
	mux.HandleFunc("/cgi-bin/kf/get_contact_way", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "url": "https://kf", "qr_code": "q"})
	})
	mux.HandleFunc("/cgi-bin/kf/del_contact_way", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	c, err := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	var lst kfAccountListResp
	if err := c.GetJSON(ctx, "/cgi-bin/kf/account/list", nil, &lst); err != nil {
		t.Fatal(err)
	}
	if len(lst.AccountList) == 0 {
		t.Fatal("no kf account")
	}
	var sendResp kfSendMsgResp
	reqSend := kfSendMsgReq{OpenKfID: "kf_123", ToUser: "ext_user", MsgType: "text"}
	reqSend.Text.Content = "hello"
	if err := c.PostJSON(ctx, "/cgi-bin/kf/send_msg", reqSend, &sendResp); err != nil {
		t.Fatal(err)
	}
	var syncResp kfSyncMsgResp
	if err := c.PostJSON(ctx, "/cgi-bin/kf/sync_msg", kfSyncMsgReq{Limit: 1}, &syncResp); err != nil {
		t.Fatal(err)
	}
	var sGet kfServiceStateGetResp
	if err := c.PostJSON(ctx, "/cgi-bin/kf/service_state/get", kfServiceStateGetReq{OpenKfID: "kf_123", ExternalUserID: "ext"}, &sGet); err != nil {
		t.Fatal(err)
	}
	var sTrans kfServiceStateTransResp
	if err := c.PostJSON(ctx, "/cgi-bin/kf/service_state/trans", kfServiceStateTransReq{OpenKfID: "kf_123", ExternalUserID: "ext", ServiceState: 2, ServicerUserID: "u2"}, &sTrans); err != nil {
		t.Fatal(err)
	}
	var add kfAddContactWayResp
	if err := c.PostJSON(ctx, "/cgi-bin/kf/add_contact_way", kfAddContactWayReq{OpenKfID: "kf_123"}, &add); err != nil {
		t.Fatal(err)
	}
	var get kfGetContactWayResp
	if err := c.PostJSON(ctx, "/cgi-bin/kf/get_contact_way", kfGetContactWayReq{OpenKfID: "kf_123"}, &get); err != nil {
		t.Fatal(err)
	}
	var del kfDelContactWayResp
	if err := c.PostJSON(ctx, "/cgi-bin/kf/del_contact_way", kfDelContactWayReq{OpenKfID: "kf_123"}, &del); err != nil {
		t.Fatal(err)
	}
}
