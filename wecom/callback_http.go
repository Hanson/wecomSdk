package wecom

import (
    "encoding/xml"
    "io"
    "net/http"
)

type xmlEncrypt struct {
    XMLName xml.Name `xml:"xml"`
    ToUserName string `xml:"ToUserName"`
    Encrypt string `xml:"Encrypt"`
}

func ValidateURL(w http.ResponseWriter, r *http.Request, cc *CallbackCrypto) {
    qs := r.URL.Query()
    sig := qs.Get("msg_signature")
    ts := qs.Get("timestamp")
    nonce := qs.Get("nonce")
    echostr := qs.Get("echostr")
    if !cc.Verify(sig, ts, nonce, echostr) { w.WriteHeader(http.StatusForbidden); return }
    msg, err := cc.Decrypt(echostr)
    if err != nil { w.WriteHeader(http.StatusBadRequest); return }
    _, _ = w.Write(msg)
}

func ReceiveMessage(w http.ResponseWriter, r *http.Request, cc *CallbackCrypto, handle func([]byte) []byte) {
    qs := r.URL.Query()
    sig := qs.Get("msg_signature")
    ts := qs.Get("timestamp")
    nonce := qs.Get("nonce")
    b, _ := io.ReadAll(r.Body)
    _ = r.Body.Close()
    var x xmlEncrypt
    _ = xml.Unmarshal(b, &x)
    if x.Encrypt == "" { w.WriteHeader(http.StatusBadRequest); return }
    if !cc.Verify(sig, ts, nonce, x.Encrypt) { w.WriteHeader(http.StatusForbidden); return }
    msg, err := cc.Decrypt(x.Encrypt)
    if err != nil { w.WriteHeader(http.StatusBadRequest); return }
    if handle != nil {
        resp := handle(msg)
        if len(resp) > 0 { _, _ = w.Write(resp) }
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

