package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

// HmacSha1 加密
func HmacSha1(data string, secret string) string {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}
