package wecom

import (
    "encoding/base64"
    "net/http"
    "net/http/httptest"
    "net/url"
    "strings"
    "testing"
)

func TestCallbackCryptoRoundtrip(t *testing.T) {
    raw := make([]byte, 32)
    for i := range raw { raw[i] = byte(i) }
    encKey := strings.TrimRight(base64.StdEncoding.EncodeToString(raw), "=")
    cc, err := NewCallbackCrypto("TOKEN", encKey, "rid")
    if err != nil { t.Fatal(err) }
    enc, err := cc.Encrypt([]byte("hello"))
    if err != nil { t.Fatal(err) }
    sig := cc.Signature("1","2", enc)
    if !cc.Verify(sig, "1","2", enc) { t.Fatal("verify failed") }
    msg, err := cc.Decrypt(enc)
    if err != nil || string(msg) != "hello" { t.Fatalf("decrypt: %v %s", err, string(msg)) }
}

func TestValidateURLAndReceiveMessage(t *testing.T) {
    raw := make([]byte, 32)
    for i := range raw { raw[i] = byte(32+i) }
    encKey := strings.TrimRight(base64.StdEncoding.EncodeToString(raw), "=")
    cc, err := NewCallbackCrypto("T", encKey, "rid")
    if err != nil { t.Fatal(err) }
    enc, _ := cc.Encrypt([]byte("world"))
    sig := cc.Signature("1","2", enc)
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            ValidateURL(w, r, cc)
            return
        }
        ReceiveMessage(w, r, cc, func(b []byte) []byte { return b })
    }))
    defer srv.Close()
    u, _ := url.Parse(srv.URL)
    q := u.Query()
    q.Set("msg_signature", sig)
    q.Set("timestamp", "1")
    q.Set("nonce", "2")
    q.Set("echostr", enc)
    u.RawQuery = q.Encode()
    resp, err := http.Get(u.String())
    if err != nil { t.Fatal(err) }
    b := make([]byte, 5)
    _, _ = resp.Body.Read(b)
    _ = resp.Body.Close()
    if string(b) != "world" { t.Fatalf("get: %s", string(b)) }
    body := "<xml><ToUserName><![CDATA[rid]]></ToUserName><Encrypt><![CDATA[" + enc + "]]></Encrypt></xml>"
    req, _ := http.NewRequest(http.MethodPost, srv.URL+"?msg_signature="+sig+"&timestamp=1&nonce=2", strings.NewReader(body))
    resp2, err := http.DefaultClient.Do(req)
    if err != nil { t.Fatal(err) }
    b2 := make([]byte, 5)
    _, _ = resp2.Body.Read(b2)
    _ = resp2.Body.Close()
    if string(b2) != "world" { t.Fatalf("post: %s", string(b2)) }
}
