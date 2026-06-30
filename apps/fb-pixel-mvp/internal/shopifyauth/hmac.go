// Package shopifyauth xử lý xác thực OAuth callback của Shopify.
package shopifyauth

import (
	"crypto/hmac"
	"net/url"
	"sort"
	"strings"
)

// VerifyHMAC kiểm tra HMAC-SHA256 của một Shopify OAuth callback.
//
// Quy tắc: bỏ param "hmac" (và "signature"), sắp xếp các param còn lại theo key,
// ghép thành "key=value" nối bằng "&", ký bằng app secret, rồi so sánh
// hằng-thời-gian với hmac nhận được. Trả false nếu thiếu hmac hoặc không khớp.
func VerifyHMAC(query url.Values, secret string) bool {
	provided := query.Get("hmac")
	if provided == "" {
		return false
	}

	keys := make([]string, 0, len(query))
	for k := range query {
		if k == "hmac" || k == "signature" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, k+"="+query.Get(k))
	}
	msg := strings.Join(parts, "&")

	return hmac.Equal([]byte(sign(msg, secret)), []byte(provided))
}
