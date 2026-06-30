// Package webpixel đăng ký Shopify Web Pixel để nhúng Facebook Pixel vào storefront.
package webpixel

import (
	"errors"

	"fb-pixel-mvp/internal/pixel"
	"fb-pixel-mvp/internal/shop"
)

// ErrNoPixelConfigured: shop chưa cấu hình Pixel ID → không đăng ký (không nhúng pixel rỗng).
var ErrNoPixelConfigured = errors.New("no pixel configured for shop")

// Registrar gọi Shopify tạo Web Pixel với pixelID, trả về web pixel id.
// Inject để test offline; bản thật gọi Shopify Admin API (GraphQL webPixelCreate).
type Registrar func(shopDomain, accessToken, pixelID string) (webPixelID string, err error)

// Service ráp shop + pixel repo + registrar để kích hoạt Web Pixel cho một shop.
type Service struct {
	shops    *shop.Repository
	pixels   *pixel.Repository
	register Registrar
}

// NewService tạo service.
func NewService(shops *shop.Repository, pixels *pixel.Repository, register Registrar) *Service {
	return &Service{shops: shops, pixels: pixels, register: register}
}

// Activate đăng ký Web Pixel cho shop NẾU đã cấu hình Pixel ID; ngược lại trả
// ErrNoPixelConfigured. Thành công thì lưu web_pixel_id.
func (s *Service) Activate(shopDomain string) error {
	cfg, err := s.pixels.GetByShop(shopDomain)
	if err != nil || cfg.PixelID == "" {
		return ErrNoPixelConfigured
	}
	sh, err := s.shops.GetByDomain(shopDomain)
	if err != nil {
		return err
	}
	webPixelID, err := s.register(shopDomain, sh.AccessToken, cfg.PixelID)
	if err != nil {
		return err
	}
	return s.pixels.SetWebPixelID(shopDomain, webPixelID)
}
