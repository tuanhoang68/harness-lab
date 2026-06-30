// Package pixel lưu cấu hình Facebook Pixel ID của từng shop.
package pixel

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Config là cấu hình pixel của một shop (1—1 với Shop theo domain).
type Config struct {
	ID         uint   `gorm:"primaryKey"`
	ShopDomain string `gorm:"uniqueIndex;not null"`
	PixelID    string `gorm:"not null"`
	WebPixelID string // điền sau khi đăng ký Web Pixel (F3)
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Repository thao tác Config trên *gorm.DB (MySQL prod, SQLite khi test).
type Repository struct {
	db *gorm.DB
}

// NewRepository tạo repository từ một kết nối GORM.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Upsert lưu pixel cho shop theo domain: tạo mới hoặc cập nhật PixelID (không trùng).
func (r *Repository) Upsert(shopDomain, pixelID string) (*Config, error) {
	c := &Config{ShopDomain: shopDomain, PixelID: pixelID}
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "shop_domain"}},
		DoUpdates: clause.AssignmentColumns([]string{"pixel_id", "updated_at"}),
	}).Create(c).Error
	if err != nil {
		return nil, err
	}
	return r.GetByShop(shopDomain)
}

// SetWebPixelID lưu id của Web Pixel đã đăng ký (Shopify) cho shop.
func (r *Repository) SetWebPixelID(shopDomain, webPixelID string) error {
	return r.db.Model(&Config{}).
		Where("shop_domain = ?", shopDomain).
		Update("web_pixel_id", webPixelID).Error
}

// GetByShop đọc cấu hình pixel theo shop domain.
func (r *Repository) GetByShop(shopDomain string) (*Config, error) {
	var c Config
	if err := r.db.Where("shop_domain = ?", shopDomain).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// ValidatePixelID kiểm tra Pixel ID đúng định dạng MVP: chuỗi 10–20 chữ số.
// (Chỉ validate định dạng — KHÔNG verify pixel có thật với Facebook; xem App Spec R2.)
func ValidatePixelID(id string) bool {
	if len(id) < 10 || len(id) > 20 {
		return false
	}
	for _, r := range id {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
