package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"fb-pixel-mvp/internal/pixel"
)

func newPixelRepo(t *testing.T) *pixel.Repository {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&pixel.Config{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return pixel.NewRepository(db)
}

func servePostPixel(h *Handler, form url.Values) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/settings/pixel", h.SavePixel)
	req := httptest.NewRequest(http.MethodPost, "/settings/pixel", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// F2 Scenario 1 — Pixel ID hợp lệ → lưu được, đọc lại đúng.
func TestSavePixel_ValidSaves(t *testing.T) {
	h := testHandler()
	h.pixels = newPixelRepo(t)

	form := url.Values{}
	form.Set("shop", "test-shop.myshopify.com")
	form.Set("pixel_id", "1234567890123456")

	if w := servePostPixel(h, form); w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	got, err := h.pixels.GetByShop("test-shop.myshopify.com")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.PixelID != "1234567890123456" {
		t.Fatalf("pixel = %q, want 1234567890123456", got.PixelID)
	}
}

func serveGetPixel(h *Handler, shop string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/settings/pixel", h.GetPixel)
	target := "/settings/pixel?shop=" + url.QueryEscape(shop)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// F4 Scenario 1 — shop đã cấu hình → 200 + Pixel ID (plain text).
func TestGetPixel_ConfiguredReturnsID(t *testing.T) {
	h := testHandler()
	h.pixels = newPixelRepo(t)
	if _, err := h.pixels.Upsert("test-shop.myshopify.com", "1234567890123456"); err != nil {
		t.Fatalf("seed: %v", err)
	}

	w := serveGetPixel(h, "test-shop.myshopify.com")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	if got := w.Body.String(); got != "1234567890123456" {
		t.Fatalf("body = %q, want 1234567890123456", got)
	}
}

// F4 Scenario 2 — shop chưa cấu hình → 404 + "chưa cấu hình".
func TestGetPixel_NotConfigured404(t *testing.T) {
	h := testHandler()
	h.pixels = newPixelRepo(t)

	w := serveGetPixel(h, "fresh-shop.myshopify.com")
	if w.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", w.Code)
	}
	if !strings.Contains(w.Body.String(), "chưa cấu hình") {
		t.Fatalf("body = %q, want contains \"chưa cấu hình\"", w.Body.String())
	}
}

// F4 Scenario 3 — shop domain sai → 400 "invalid shop".
func TestGetPixel_InvalidShop400(t *testing.T) {
	h := testHandler()
	h.pixels = newPixelRepo(t)

	w := serveGetPixel(h, "evil.com")
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", w.Code)
	}
}

// F2 Scenario 2 — định dạng sai → 400, KHÔNG lưu.
func TestSavePixel_InvalidFormat400(t *testing.T) {
	h := testHandler()
	h.pixels = newPixelRepo(t)

	form := url.Values{}
	form.Set("shop", "test-shop.myshopify.com")
	form.Set("pixel_id", "abc")

	if w := servePostPixel(h, form); w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", w.Code)
	}
	if _, err := h.pixels.GetByShop("test-shop.myshopify.com"); err == nil {
		t.Fatal("expected nothing saved on invalid format")
	}
}
