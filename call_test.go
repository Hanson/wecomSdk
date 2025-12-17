package wecom

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGenericPostJSON(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/cgi-bin/gettoken", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "access_token": "T", "expires_in": 7200})
    })
    mux.HandleFunc("/cgi-bin/user/update", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
    })
    srv := httptest.NewServer(mux)
    defer srv.Close()
    c, err := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
    if err != nil { t.Fatal(err) }
    ctx := context.Background()
    var out UserUpdateResp
    err = c.PostJSON(ctx, "/cgi-bin/user/update", UserUpdateReq{UserID: "u1", Name: "n1"}, &out)
    if err != nil { t.Fatal(err) }
}

func TestGenericPostJSONErrorOutFilled(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/cgi-bin/gettoken", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "access_token": "T", "expires_in": 7200})
    })
    mux.HandleFunc("/cgi-bin/user/update", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 400, "errmsg": "bad request", "extra": "x"})
    })
    srv := httptest.NewServer(mux)
    defer srv.Close()
    c, err := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
    if err != nil { t.Fatal(err) }
    ctx := context.Background()
    var out struct { APIError; Extra string `json:"extra"` }
    err = c.PostJSON(ctx, "/cgi-bin/user/update", UserUpdateReq{UserID: "u1"}, &out)
    if err == nil { t.Fatal("expected error") }
    if out.Extra != "x" { t.Fatalf("out not filled: %+v", out) }
}

func TestGenericGetJSONWithReq(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/cgi-bin/gettoken", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "access_token": "T", "expires_in": 7200})
    })
    mux.HandleFunc("/cgi-bin/user/simplelist", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "userlist": []map[string]any{{"userid": "u1", "name": "n1"}}})
    })
    srv := httptest.NewServer(mux)
    defer srv.Close()
    c, err := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
    if err != nil { t.Fatal(err) }
    ctx := context.Background()
    var out UserSimpleListResp
    err = c.GetJSONWithReq(ctx, "/cgi-bin/user/simplelist", UserSimpleListReq{DepartmentID: 1, FetchChild: 1}, &out)
    if err != nil { t.Fatal(err) }
    if len(out.UserList) != 1 || out.UserList[0].UserID != "u1" { t.Fatalf("unexpected: %+v", out) }
}
