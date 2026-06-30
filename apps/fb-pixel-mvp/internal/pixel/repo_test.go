package pixel

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func newTestRepo(t *testing.T) (*Repository, *gorm.DB) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&Config{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return NewRepository(db), db
}

// F2 Scenario 1 — lưu pixel cho shop, đọc lại đúng.
func TestUpsert_SavesPixel(t *testing.T) {
	repo, _ := newTestRepo(t)
	if _, err := repo.Upsert("test-shop.myshopify.com", "1234567890123456"); err != nil {
		t.Fatalf("upsert: %v", err)
	}
	got, err := repo.GetByShop("test-shop.myshopify.com")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.PixelID != "1234567890123456" {
		t.Fatalf("pixel = %q, want 1234567890123456", got.PixelID)
	}
}

// F2 Scenario 3 — cập nhật pixel, KHÔNG tạo bản ghi trùng.
func TestUpsert_UpdatesNoDuplicate(t *testing.T) {
	repo, db := newTestRepo(t)
	if _, err := repo.Upsert("test-shop.myshopify.com", "1111111111111111"); err != nil {
		t.Fatalf("first upsert: %v", err)
	}
	c, err := repo.Upsert("test-shop.myshopify.com", "2222222222222222")
	if err != nil {
		t.Fatalf("second upsert: %v", err)
	}
	if c.PixelID != "2222222222222222" {
		t.Fatalf("pixel = %q, want 2222222222222222", c.PixelID)
	}
	var n int64
	db.Model(&Config{}).Count(&n)
	if n != 1 {
		t.Fatalf("rows = %d, want 1", n)
	}
}
