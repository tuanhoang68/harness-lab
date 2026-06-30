package shop

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
	if err := db.AutoMigrate(&Shop{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return NewRepository(db), db
}

// F1 Scenario 1 — cài mới: lưu được shop + token, đọc lại đúng.
func TestUpsert_CreatesNewShop(t *testing.T) {
	repo, _ := newTestRepo(t)

	s, err := repo.Upsert("test-shop.myshopify.com", "token-1")
	if err != nil {
		t.Fatalf("upsert: %v", err)
	}
	if s.ID == 0 {
		t.Fatal("expected persisted shop to have an ID")
	}

	got, err := repo.GetByDomain("test-shop.myshopify.com")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.AccessToken != "token-1" {
		t.Fatalf("access token = %q, want token-1", got.AccessToken)
	}
}

// F1 Scenario 3 — cài lại shop đã có: cập nhật token, KHÔNG tạo bản ghi trùng.
func TestUpsert_ExistingUpdatesTokenNoDuplicate(t *testing.T) {
	repo, db := newTestRepo(t)

	if _, err := repo.Upsert("test-shop.myshopify.com", "old-token"); err != nil {
		t.Fatalf("first upsert: %v", err)
	}
	s, err := repo.Upsert("test-shop.myshopify.com", "new-token")
	if err != nil {
		t.Fatalf("second upsert: %v", err)
	}
	if s.AccessToken != "new-token" {
		t.Fatalf("access token = %q, want new-token", s.AccessToken)
	}

	var n int64
	db.Model(&Shop{}).Count(&n)
	if n != 1 {
		t.Fatalf("expected exactly 1 shop row, got %d", n)
	}
}
