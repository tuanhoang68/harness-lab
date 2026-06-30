package shopifyauth

import (
	"net/url"
	"testing"
)

const (
	testSecret = "shpss_testsecret"
	// Vector đóng băng: HMAC-SHA256 của
	// "code=abc123&shop=test-shop.myshopify.com&state=nonce123&timestamp=1700000000"
	// ký bằng testSecret (sinh 1 lần bằng script throwaway).
	validHMAC = "caa5ebb6ac5a4f2310bb33ad789c80562141212d86e01cc32b2140295579b36e"
)

func baseQuery() url.Values {
	q := url.Values{}
	q.Set("code", "abc123")
	q.Set("shop", "test-shop.myshopify.com")
	q.Set("state", "nonce123")
	q.Set("timestamp", "1700000000")
	return q
}

// F1 Scenario 4 — callback hợp lệ thì qua.
func TestVerifyHMAC_ValidPasses(t *testing.T) {
	q := baseQuery()
	q.Set("hmac", validHMAC)
	if !VerifyHMAC(q, testSecret) {
		t.Fatal("expected valid HMAC to pass")
	}
}

// F1 Scenario 4 — đổi param sau khi ký (giả mạo) thì trượt.
func TestVerifyHMAC_TamperedFails(t *testing.T) {
	q := baseQuery()
	q.Set("hmac", validHMAC)
	q.Set("shop", "evil-shop.myshopify.com")
	if VerifyHMAC(q, testSecret) {
		t.Fatal("expected tampered request to fail")
	}
}

// F1 Scenario 4 — thiếu hmac thì trượt.
func TestVerifyHMAC_MissingHMACFails(t *testing.T) {
	q := baseQuery()
	if VerifyHMAC(q, testSecret) {
		t.Fatal("expected missing HMAC to fail")
	}
}
