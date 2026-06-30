package pixel

import "testing"

// F3 — sau khi đăng ký Web Pixel, lưu lại web_pixel_id cho shop.
func TestSetWebPixelID(t *testing.T) {
	repo, _ := newTestRepo(t)
	if _, err := repo.Upsert("test-shop.myshopify.com", "1234567890123456"); err != nil {
		t.Fatalf("upsert: %v", err)
	}

	if err := repo.SetWebPixelID("test-shop.myshopify.com", "gid://shopify/WebPixel/123"); err != nil {
		t.Fatalf("set web pixel id: %v", err)
	}

	got, err := repo.GetByShop("test-shop.myshopify.com")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.WebPixelID != "gid://shopify/WebPixel/123" {
		t.Fatalf("web_pixel_id = %q, want gid://shopify/WebPixel/123", got.WebPixelID)
	}
}
