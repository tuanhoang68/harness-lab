package webpixel

import (
	"errors"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"fb-pixel-mvp/internal/pixel"
	"fb-pixel-mvp/internal/shop"
)

func newRepos(t *testing.T) (*shop.Repository, *pixel.Repository) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&shop.Shop{}, &pixel.Config{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return shop.NewRepository(db), pixel.NewRepository(db)
}

// F3 Scenario 1 (server-side) — có Pixel ID → đăng ký Web Pixel với đúng pixel id, lưu web_pixel_id.
func TestActivate_RegistersWhenPixelConfigured(t *testing.T) {
	shops, pixels := newRepos(t)
	shops.Upsert("test-shop.myshopify.com", "access-token")
	pixels.Upsert("test-shop.myshopify.com", "1234567890123456")

	var gotShop, gotToken, gotPixel string
	reg := func(shopDomain, token, pixelID string) (string, error) {
		gotShop, gotToken, gotPixel = shopDomain, token, pixelID
		return "gid://shopify/WebPixel/999", nil
	}
	svc := NewService(shops, pixels, reg)

	if err := svc.Activate("test-shop.myshopify.com"); err != nil {
		t.Fatalf("activate: %v", err)
	}
	if gotShop != "test-shop.myshopify.com" || gotToken != "access-token" || gotPixel != "1234567890123456" {
		t.Fatalf("registrar got shop=%q token=%q pixel=%q", gotShop, gotToken, gotPixel)
	}
	cfg, _ := pixels.GetByShop("test-shop.myshopify.com")
	if cfg.WebPixelID != "gid://shopify/WebPixel/999" {
		t.Fatalf("web_pixel_id = %q, want gid://shopify/WebPixel/999", cfg.WebPixelID)
	}
}

// F3 Scenario 3 — chưa cấu hình Pixel ID → KHÔNG đăng ký (không nhúng pixel rỗng).
func TestActivate_NoPixelConfigured_DoesNotRegister(t *testing.T) {
	shops, pixels := newRepos(t)
	shops.Upsert("test-shop.myshopify.com", "access-token")
	// KHÔNG cấu hình pixel.

	called := false
	reg := func(string, string, string) (string, error) { called = true; return "", nil }
	svc := NewService(shops, pixels, reg)

	err := svc.Activate("test-shop.myshopify.com")
	if !errors.Is(err, ErrNoPixelConfigured) {
		t.Fatalf("err = %v, want ErrNoPixelConfigured", err)
	}
	if called {
		t.Fatal("registrar must NOT be called when no pixel configured")
	}
}
