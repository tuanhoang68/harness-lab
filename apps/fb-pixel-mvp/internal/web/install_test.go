package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"

	"fb-pixel-mvp/internal/shopifyauth"
)

const secret = "shpss_testsecret"

func testHandler() *Handler {
	return NewHandler(Config{
		APIKey:      "test-api-key",
		Secret:      secret,
		Scopes:      "read_products,write_pixels",
		RedirectURI: "https://app.example.com/auth/callback",
	}, nil)
}

func doGET(h *Handler, target string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/install", h.Install)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, target, nil))
	return w
}

// F1 — /install chuyển hướng tới màn OAuth grant của Shopify, kèm state ký được.
func TestInstall_RedirectsToShopifyOAuth(t *testing.T) {
	w := doGET(testHandler(), "/install?shop=test-shop.myshopify.com")

	if w.Code != http.StatusFound {
		t.Fatalf("status = %d, want 302", w.Code)
	}
	u, err := url.Parse(w.Header().Get("Location"))
	if err != nil {
		t.Fatalf("parse location: %v", err)
	}
	if u.Host != "test-shop.myshopify.com" || u.Path != "/admin/oauth/authorize" {
		t.Fatalf("redirect base = %s", u)
	}
	q := u.Query()
	if q.Get("client_id") != "test-api-key" {
		t.Errorf("client_id = %q", q.Get("client_id"))
	}
	if q.Get("scope") != "read_products,write_pixels" {
		t.Errorf("scope = %q", q.Get("scope"))
	}
	if q.Get("redirect_uri") != "https://app.example.com/auth/callback" {
		t.Errorf("redirect_uri = %q", q.Get("redirect_uri"))
	}
	gotShop, ok := shopifyauth.VerifyState(q.Get("state"), secret)
	if !ok || gotShop != "test-shop.myshopify.com" {
		t.Errorf("state invalid: ok=%v shop=%q", ok, gotShop)
	}
}

// Thiếu shop → 400.
func TestInstall_MissingShopReturns400(t *testing.T) {
	if w := doGET(testHandler(), "/install"); w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", w.Code)
	}
}

// Domain không phải *.myshopify.com → 400 (chống open redirect).
func TestInstall_InvalidShopReturns400(t *testing.T) {
	if w := doGET(testHandler(), "/install?shop=evil.com"); w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", w.Code)
	}
}
