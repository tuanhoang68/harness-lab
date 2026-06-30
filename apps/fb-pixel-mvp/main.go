// Command fb-pixel-mvp khởi động HTTP server cho app Facebook Pixel (MVP).
//
// Composition root: đọc config từ env, mở DB (SQLite để chạy thử zero-setup; prod đổi
// sang GORM MySQL driver theo Design Doc), ráp handler, đăng ký route.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"fb-pixel-mvp/internal/pixel"
	"fb-pixel-mvp/internal/shop"
	"fb-pixel-mvp/internal/web"
	"fb-pixel-mvp/internal/webpixel"
)

func main() {
	// Nạp .env nếu có (không bắt buộc) — để `go run .` đọc config từ file thay vì truyền param.
	if err := godotenv.Load(); err == nil {
		log.Println("loaded config from .env")
	}

	appURL := env("APP_URL", "http://localhost:8080")
	cfg := web.Config{
		APIKey:      env("SHOPIFY_API_KEY", "dev-api-key"),
		Secret:      env("SHOPIFY_API_SECRET", "dev-secret"),
		Scopes:      env("SHOPIFY_SCOPES", "write_pixels,read_customer_events"), // Bậc 2: web pixel
		RedirectURI: appURL + "/auth/callback",
	}

	db, err := gorm.Open(sqlite.Open(env("DB_PATH", "fb-pixel.db")), &gorm.Config{})
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&shop.Shop{}, &pixel.Config{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	shopRepo := shop.NewRepository(db)
	pixelRepo := pixel.NewRepository(db)
	webPixel := webpixel.NewService(shopRepo, pixelRepo, webpixel.NewShopifyRegistrar(appURL+"/collect"))
	h := web.NewHandler(cfg, shopRepo).WithPixels(pixelRepo).WithWebPixel(webPixel)

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	r.GET("/install", h.Install)
	r.GET("/auth/callback", h.Callback)
	r.GET("/settings/pixel", h.GetPixel) // đọc Pixel ID đang lưu (đối xứng POST)
	r.POST("/settings/pixel", h.SavePixel)
	r.POST("/collect", h.Collect) // web pixel storefront → backend nhận event
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "fb-pixel-mvp running (shop=%s)", c.Query("shop"))
	})

	addr := ":" + env("PORT", "8080")
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
