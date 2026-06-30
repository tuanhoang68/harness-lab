package shopifyauth

import "testing"

const testShop = "test-shop.myshopify.com"

// State hợp lệ ký bằng đúng secret thì verify được và lấy lại đúng shop.
func TestState_RoundTrip(t *testing.T) {
	token := GenerateState(testShop, testSecret)
	shop, ok := VerifyState(token, testSecret)
	if !ok {
		t.Fatal("expected valid state to verify")
	}
	if shop != testShop {
		t.Fatalf("expected shop %q, got %q", testShop, shop)
	}
}

// Sửa token (giả mạo) thì trượt.
func TestState_TamperedFails(t *testing.T) {
	token := GenerateState(testShop, testSecret)
	b := []byte(token)
	b[len(b)-1] ^= 0x01 // lật 1 bit ở chữ ký
	if _, ok := VerifyState(string(b), testSecret); ok {
		t.Fatal("expected tampered state to fail")
	}
}

// Sai secret thì trượt (token ký bằng secret khác).
func TestState_WrongSecretFails(t *testing.T) {
	token := GenerateState(testShop, testSecret)
	if _, ok := VerifyState(token, "other_secret"); ok {
		t.Fatal("expected wrong secret to fail")
	}
}

// Mỗi lần sinh ra một token khác nhau (có nonce ngẫu nhiên) → chống replay.
func TestState_Unique(t *testing.T) {
	a := GenerateState(testShop, testSecret)
	b := GenerateState(testShop, testSecret)
	if a == b {
		t.Fatal("expected unique state tokens per call")
	}
}
