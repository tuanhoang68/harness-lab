package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fb-pixel-mvp/internal/pixel"
)

// SavePixel xử lý POST /settings/pixel: validate định dạng Pixel ID rồi lưu cho shop.
//
// Lưu ý bảo mật (ngoài MVP): shop hiện lấy từ form; app thật phải xác thực shop qua
// Shopify session token, không tin tham số client.
func (h *Handler) SavePixel(c *gin.Context) {
	shopDomain := c.PostForm("shop")
	pixelID := c.PostForm("pixel_id")

	if !validShopDomain(shopDomain) {
		c.String(http.StatusBadRequest, "invalid shop")
		return
	}
	if !pixel.ValidatePixelID(pixelID) {
		c.String(http.StatusBadRequest, "invalid pixel id format")
		return
	}
	if _, err := h.pixels.Upsert(shopDomain, pixelID); err != nil {
		c.String(http.StatusInternalServerError, "save failed")
		return
	}

	// Pixel vừa lưu → kích hoạt Web Pixel (best-effort; lỗi đăng ký không chặn lưu).
	if h.activator != nil {
		_ = h.activator.Activate(shopDomain)
	}

	c.String(http.StatusOK, "saved")
}

// GetPixel xử lý GET /settings/pixel?shop=<domain>: trả Pixel ID đang lưu của shop
// (đối xứng đọc với SavePixel). Chưa cấu hình → 404 "chưa cấu hình".
//
// Lưu ý bảo mật (ngoài MVP): như SavePixel, shop lấy từ query; app thật phải xác thực
// shop qua Shopify session token, không tin tham số client.
func (h *Handler) GetPixel(c *gin.Context) {
	shopDomain := c.Query("shop")
	if !validShopDomain(shopDomain) {
		c.String(http.StatusBadRequest, "invalid shop")
		return
	}

	cfg, err := h.pixels.GetByShop(shopDomain)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusNotFound, "chưa cấu hình")
		return
	}
	if err != nil {
		c.String(http.StatusInternalServerError, "read failed")
		return
	}

	c.String(http.StatusOK, cfg.PixelID)
}
