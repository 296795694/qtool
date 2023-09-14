package encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// HmacSha256 加密
func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
