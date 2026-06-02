package audit

import (
	"github.com/anandudevops/aegis/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(log models.AuditLog) error {
	return r.db.Create(&log).Error
}

func (r *Repository) FindAll(page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64
	offset := (page - 1) * limit

	if err := r.db.Model(&models.AuditLog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&logs).Error; err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}

func (r *Repository) FindByToken(token string) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	if err := r.db.Where("token = ?", token).Order("created_at DESC").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
