package wecom

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha1"
    "encoding/base64"
    "encoding/binary"
    "encoding/hex"
    "errors"
    "io"
    "sort"
    "strings"
)

type CallbackCrypto struct {
    token     string
    aesKey    []byte
    receiveID string
}

func NewCallbackCrypto(token, encodingAESKey, receiveID string) (*CallbackCrypto, error) {
    k, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
    if err != nil { return nil, err }
    if len(k) != 32 { return nil, errors.New("invalid aes key") }
    return &CallbackCrypto{token: token, aesKey: k, receiveID: receiveID}, nil
}

func (c *CallbackCrypto) Signature(timestamp, nonce, encrypted string) string {
    arr := []string{c.token, timestamp, nonce, encrypted}
    sort.Strings(arr)
    h := sha1.Sum([]byte(strings.Join(arr, "")))
    return hex.EncodeToString(h[:])
}

func (c *CallbackCrypto) Verify(sig, timestamp, nonce, encrypted string) bool {
    return strings.EqualFold(sig, c.Signature(timestamp, nonce, encrypted))
}

func pkcs7Unpad(b []byte) ([]byte, error) {
    if len(b) == 0 { return nil, errors.New("empty") }
    p := int(b[len(b)-1])
    if p <= 0 || p > len(b) { return nil, errors.New("bad pad") }
    for i := len(b)-p; i < len(b); i++ { if b[i] != byte(p) { return nil, errors.New("bad pad") } }
    return b[:len(b)-p], nil
}

func pkcs7Pad(b []byte, block int) []byte {
    p := block - (len(b) % block)
    if p == 0 { p = block }
    pad := make([]byte, p)
    for i := range pad { pad[i] = byte(p) }
    return append(b, pad...)
}

func (c *CallbackCrypto) Decrypt(encrypted string) ([]byte, error) {
    data, err := base64.StdEncoding.DecodeString(encrypted)
    if err != nil { return nil, err }
    block, err := aes.NewCipher(c.aesKey)
    if err != nil { return nil, err }
    iv := c.aesKey[:16]
    mode := cipher.NewCBCDecrypter(block, iv)
    out := make([]byte, len(data))
    mode.CryptBlocks(out, data)
    out, err = pkcs7Unpad(out)
    if err != nil { return nil, err }
    if len(out) < 20 { return nil, errors.New("short") }
    msgLen := binary.BigEndian.Uint32(out[16:20])
    if int(20+msgLen) > len(out) { return nil, errors.New("bad len") }
    msg := out[20:20+msgLen]
    rid := out[20+msgLen:]
    if string(rid) != c.receiveID { return nil, errors.New("mismatch receiveid") }
    return msg, nil
}

func (c *CallbackCrypto) Encrypt(plain []byte) (string, error) {
    rand16 := make([]byte, 16)
    if _, err := io.ReadFull(rand.Reader, rand16); err != nil { return "", err }
    buf := make([]byte, 0, 16+4+len(plain)+len(c.receiveID))
    buf = append(buf, rand16...)
    l := make([]byte, 4)
    binary.BigEndian.PutUint32(l, uint32(len(plain)))
    buf = append(buf, l...)
    buf = append(buf, plain...)
    buf = append(buf, []byte(c.receiveID)...)
    buf = pkcs7Pad(buf, aes.BlockSize)
    block, err := aes.NewCipher(c.aesKey)
    if err != nil { return "", err }
    iv := c.aesKey[:16]
    mode := cipher.NewCBCEncrypter(block, iv)
    out := make([]byte, len(buf))
    mode.CryptBlocks(out, buf)
    return base64.StdEncoding.EncodeToString(out), nil
}

