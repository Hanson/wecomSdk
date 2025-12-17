package wecom

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"
)

func newTestServer(token string, expires int, requireToken bool, msgResp any, tokenCounter *int32) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/cgi-bin/gettoken", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(tokenCounter, 1)
		u, _ := url.Parse(r.URL.String())
		q := u.Query()
		if q.Get("corpid") == "" || q.Get("corpsecret") == "" {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 400, "errmsg": "missing params"})
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "access_token": token, "expires_in": expires})
	})
	mux.HandleFunc("/cgi-bin/message/send", func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.URL.String())
		q := u.Query()
		if requireToken && q.Get("access_token") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]any{"errcode": 401, "errmsg": "no token"})
			return
		}
		_ = json.NewEncoder(w).Encode(msgResp)
	})
	return httptest.NewServer(mux)
}

func TestGetAccessTokenCaching(t *testing.T) {
	var cnt int32
	srv := newTestServer("TOKEN1", 7200, false, map[string]any{"errcode": 0, "errmsg": "ok"}, &cnt)
	defer srv.Close()
	c, err := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	ctx := context.Background()
	tok1, err := c.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("get token: %v", err)
	}
	tok2, err := c.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("get token 2: %v", err)
	}
	if tok1 != "TOKEN1" || tok2 != "TOKEN1" {
		t.Fatalf("unexpected tokens: %s %s", tok1, tok2)
	}
	if atomic.LoadInt32(&cnt) != 1 {
		t.Fatalf("gettoken called %d times", cnt)
	}
}

func TestSendTextMessage(t *testing.T) {
	var cnt int32
	srv := newTestServer("TOKEN2", 7200, true, map[string]any{"errcode": 0, "errmsg": "ok", "msgid": "123"}, &cnt)
	defer srv.Close()
	c, err := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	msg := TextMessage{ToUser: "zhangsan", AgentID: 1000001}
	msg.Text.Content = "hello"
	ctx := context.Background()
	r, err := c.SendTextMessage(ctx, msg)
	if err != nil {
		t.Fatalf("send: %v", err)
	}
	if r.MsgID != "123" {
		t.Fatalf("unexpected msgid: %s", r.MsgID)
	}
}

func TestTokenAutoRefresh(t *testing.T) {
	var cnt int32
	srv := newTestServer("T1", 1, false, map[string]any{"errcode": 0, "errmsg": "ok"}, &cnt)
	defer srv.Close()
	c, err := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	ctx := context.Background()
	tok1, err := c.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("get token: %v", err)
	}
	time.Sleep(1100 * time.Millisecond)
	tok2, err := c.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("get token 2: %v", err)
	}
	if tok1 != "T1" || tok2 != "T1" {
		t.Fatalf("unexpected tokens: %s %s", tok1, tok2)
	}
	if atomic.LoadInt32(&cnt) < 2 {
		t.Fatalf("expected refresh, calls: %d", cnt)
	}
}
