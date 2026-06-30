package web

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Collect nhận event do Web Pixel (storefront, sandbox strict) fetch về.
// Sandbox không load được fbq client-side (R1) → web pixel chỉ fetch event server-side.
// Bậc 2 (🅐): chứng minh pixel fire thật bằng cách backend nhận được event này.
func (h *Handler) Collect(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*") // cho fetch cross-origin từ storefront
	body, _ := io.ReadAll(c.Request.Body)
	log.Printf("PixelEvent received from storefront: %s", string(body))
	c.String(http.StatusOK, "ok")
}
