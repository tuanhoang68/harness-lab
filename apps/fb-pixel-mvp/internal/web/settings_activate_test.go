package web

import (
	"net/http"
	"net/url"
	"testing"
)

type fakeActivator struct {
	calledShop string
	err        error
}

func (f *fakeActivator) Activate(shopDomain string) error {
	f.calledShop = shopDomain
	return f.err
}

// F2→F3 — lưu Pixel ID thành công thì kích hoạt Web Pixel cho đúng shop.
func TestSavePixel_TriggersWebPixelActivation(t *testing.T) {
	h := testHandler()
	h.pixels = newPixelRepo(t)
	fa := &fakeActivator{}
	h.WithWebPixel(fa)

	form := url.Values{}
	form.Set("shop", "test-shop.myshopify.com")
	form.Set("pixel_id", "1234567890123456")

	if w := servePostPixel(h, form); w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	if fa.calledShop != "test-shop.myshopify.com" {
		t.Fatalf("activator called with %q, want test-shop.myshopify.com", fa.calledShop)
	}
}
