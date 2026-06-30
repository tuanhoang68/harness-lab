// Package shop lưu trữ thông tin shop đã cài app (domain + OAuth access token).
package shop

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Shop là một shop Shopify đã cài app.
type Shop struct {
	ID          uint   `gorm:"primaryKey"`
	Domain      string `gorm:"uniqueIndex;not null"`
	AccessToken string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Repository thao tác Shop trên một *gorm.DB (MySQL ở prod, SQLite khi test).
type Repository struct {
	db *gorm.DB
}

// NewRepository tạo repository từ một kết nối GORM.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Upsert lưu shop theo domain: tạo mới nếu chưa có, cập nhật access token nếu đã có
// (không tạo bản ghi trùng — dựa trên uniqueIndex của Domain).
func (r *Repository) Upsert(domain, accessToken string) (*Shop, error) {
	s := &Shop{Domain: domain, AccessToken: accessToken}
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "domain"}},
		DoUpdates: clause.AssignmentColumns([]string{"access_token", "updated_at"}),
	}).Create(s).Error
	if err != nil {
		return nil, err
	}
	return r.GetByDomain(domain)
}

// GetByDomain đọc shop theo domain.
func (r *Repository) GetByDomain(domain string) (*Shop, error) {
	var s Shop
	if err := r.db.Where("domain = ?", domain).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
