// Package web chứa HTTP handler (Gin) cho luồng OAuth của app.
package web

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"fb-pixel-mvp/internal/pixel"
	"fb-pixel-mvp/internal/shop"
	"fb-pixel-mvp/internal/shopifyauth"
)

// Config gom cấu hình app cần cho luồng OAuth.
type Config struct {
	APIKey      string
	Secret      string
	Scopes      string
	RedirectURI string
}

// pixelActivator kích hoạt Web Pixel cho một shop (webpixel.Service hiện thực).
type pixelActivator interface {
	Activate(shopDomain string) error
}

// Handler giữ cấu hình + repo để phục vụ các route OAuth & settings.
type Handler struct {
	cfg       Config
	shops     *shop.Repository
	pixels    *pixel.Repository
	activator pixelActivator
	// exchange đổi authorization code lấy access token (inject để test offline).
	exchange tokenExchangeFunc
}

// NewHandler tạo handler. shops/pixels gán sau qua field nếu route cần (vd /install không cần).
func NewHandler(cfg Config, shops *shop.Repository) *Handler {
	h := &Handler{cfg: cfg, shops: shops}
	h.exchange = h.exchangeViaShopify
	return h
}

// WithPixels gán pixel repository (cho route /settings/pixel).
func (h *Handler) WithPixels(pixels *pixel.Repository) *Handler {
	h.pixels = pixels
	return h
}

// WithWebPixel gán bộ kích hoạt Web Pixel (gọi sau khi lưu Pixel ID).
func (h *Handler) WithWebPixel(a pixelActivator) *Handler {
	h.activator = a
	return h
}

// Install xử lý GET /install?shop=... → chuyển hướng tới màn OAuth grant của Shopify.
func (h *Handler) Install(c *gin.Context) {
	shopDomain := c.Query("shop")
	if !validShopDomain(shopDomain) {
		c.String(http.StatusBadRequest, "missing or invalid shop")
		return
	}

	state := shopifyauth.GenerateState(shopDomain, h.cfg.Secret)

	u := url.URL{Scheme: "https", Host: shopDomain, Path: "/admin/oauth/authorize"}
	q := u.Query()
	q.Set("client_id", h.cfg.APIKey)
	q.Set("scope", h.cfg.Scopes)
	q.Set("redirect_uri", h.cfg.RedirectURI)
	q.Set("state", state)
	u.RawQuery = q.Encode()

	c.Redirect(http.StatusFound, u.String())
}

// validShopDomain chặn open-redirect: chỉ chấp nhận "<tên>.myshopify.com".
func validShopDomain(shopDomain string) bool {
	if shopDomain == "" || strings.ContainsAny(shopDomain, "/:?") {
		return false
	}
	host := strings.TrimSuffix(shopDomain, ".myshopify.com")
	return host != shopDomain && host != "" && !strings.Contains(host, ".")
}
