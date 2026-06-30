package web

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sort"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"fb-pixel-mvp/internal/shop"
	"fb-pixel-mvp/internal/shopifyauth"
)

func newShopRepo(t *testing.T) *shop.Repository {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&shop.Shop{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return shop.NewRepository(db)
}

// signQuery đóng vai Shopify: ký HMAC các param (bỏ hmac), giống canonicalization của server.
func signQuery(q url.Values, secret string) string {
	keys := make([]string, 0, len(q))
	for k := range q {
		if k == "hmac" || k == "signature" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, k+"="+q.Get(k))
	}
	m := hmac.New(sha256.New, []byte(secret))
	m.Write([]byte(strings.Join(parts, "&")))
	return hex.EncodeToString(m.Sum(nil))
}

func serveCallback(h *Handler, target string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/auth/callback", h.Callback)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, target, nil))
	return w
}

// F1 Scenario 1 — callback hợp lệ: đổi code lấy token, lưu shop, redirect.
func TestCallback_ExchangesTokenAndSavesShop(t *testing.T) {
	h := testHandler()
	h.shops = newShopRepo(t)
	called := false
	h.exchange = func(shopDomain, code string) (string, error) {
		called = true
		if shopDomain != "test-shop.myshopify.com" || code != "auth-code-123" {
			t.Errorf("exchange got shop=%q code=%q", shopDomain, code)
		}
		return "access-token-xyz", nil
	}

	q := url.Values{}
	q.Set("shop", "test-shop.myshopify.com")
	q.Set("code", "auth-code-123")
	q.Set("state", shopifyauth.GenerateState("test-shop.myshopify.com", secret))
	q.Set("timestamp", "1700000000")
	q.Set("hmac", signQuery(q, secret))

	w := serveCallback(h, "/auth/callback?"+q.Encode())

	if w.Code != http.StatusFound {
		t.Fatalf("status = %d, want 302", w.Code)
	}
	if !called {
		t.Fatal("expected token exchange to be called")
	}
	saved, err := h.shops.GetByDomain("test-shop.myshopify.com")
	if err != nil {
		t.Fatalf("get saved shop: %v", err)
	}
	if saved.AccessToken != "access-token-xyz" {
		t.Fatalf("saved token = %q, want access-token-xyz", saved.AccessToken)
	}
}

// F1 Scenario 4 — hmac sai → 401, không đổi token, không lưu.
func TestCallback_InvalidHMAC_401(t *testing.T) {
	h := testHandler()
	h.shops = newShopRepo(t)
	h.exchange = func(string, string) (string, error) {
		t.Fatal("token exchange must NOT run on invalid hmac")
		return "", nil
	}

	q := url.Values{}
	q.Set("shop", "test-shop.myshopify.com")
	q.Set("code", "c")
	q.Set("state", shopifyauth.GenerateState("test-shop.myshopify.com", secret))
	q.Set("timestamp", "1700000000")
	q.Set("hmac", "deadbeef") // sai

	if w := serveCallback(h, "/auth/callback?"+q.Encode()); w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", w.Code)
	}
}

// Bảo mật — state ký cho shop KHÁC (không khớp param shop) → 401.
func TestCallback_StateShopMismatch_401(t *testing.T) {
	h := testHandler()
	h.shops = newShopRepo(t)
	h.exchange = func(string, string) (string, error) { return "x", nil }

	q := url.Values{}
	q.Set("shop", "test-shop.myshopify.com")
	q.Set("code", "c")
	q.Set("state", shopifyauth.GenerateState("other-shop.myshopify.com", secret)) // shop khác
	q.Set("timestamp", "1700000000")
	q.Set("hmac", signQuery(q, secret)) // hmac hợp lệ

	if w := serveCallback(h, "/auth/callback?"+q.Encode()); w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", w.Code)
	}
}
