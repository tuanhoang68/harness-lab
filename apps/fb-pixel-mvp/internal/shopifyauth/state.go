package shopifyauth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// GenerateState tạo một OAuth state token KHÔNG cần lưu server (stateless):
// payload = "<shop>:<nonce ngẫu nhiên>", ký HMAC-SHA256 bằng secret.
// Token = base64url(payload) + "." + hex(sig). Mỗi lần gọi ra một token khác nhau.
func GenerateState(shop, secret string) string {
	nonce := make([]byte, 16)
	_, _ = rand.Read(nonce)
	payload := shop + ":" + hex.EncodeToString(nonce)
	enc := base64.RawURLEncoding.EncodeToString([]byte(payload))
	return enc + "." + sign(enc, secret)
}

// VerifyState kiểm tra chữ ký của state token và trả lại shop nếu hợp lệ.
func VerifyState(token, secret string) (shop string, ok bool) {
	enc, sig, found := strings.Cut(token, ".")
	if !found {
		return "", false
	}
	if !hmac.Equal([]byte(sign(enc, secret)), []byte(sig)) {
		return "", false
	}
	payload, err := base64.RawURLEncoding.DecodeString(enc)
	if err != nil {
		return "", false
	}
	shop, _, found = strings.Cut(string(payload), ":")
	if !found {
		return "", false
	}
	return shop, true
}

func sign(msg, secret string) string {
	m := hmac.New(sha256.New, []byte(secret))
	m.Write([]byte(msg))
	return hex.EncodeToString(m.Sum(nil))
}
