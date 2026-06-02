package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ActorID       *uuid.UUID `gorm:"type:uuid" json:"actor_id"`
	Action        string     `gorm:"not null" json:"action"`
	Token         string     `json:"token"`
	FieldType     string     `json:"field_type"`
	AccessLevel   string     `json:"access_level"`
	IPAddress     string     `json:"ip_address"`
	Success       bool       `gorm:"not null" json:"success"`
	FailureReason string     `json:"failure_reason,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
}
