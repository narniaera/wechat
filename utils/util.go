package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func WxSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

// HmacSha256ToHex 将加密后的二进制转16进制字符串
func HmacSha256ToHex(key string, data string) string {
	return hex.EncodeToString(HmacSha256(key, data))
}

// HmacSha256ToHex 将加密后的二进制转Base64字符串
func HmacSha256ToBase64(key string, data string) string {
	return base64.URLEncoding.EncodeToString(HmacSha256(key, data))
}
