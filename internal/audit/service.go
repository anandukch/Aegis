package audit

import (
	"log"

	"github.com/anandudevops/aegis/internal/models"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Log(actorID, action, token, fieldType, accessLevel, ip string, success bool, reason string) {
	entry := models.AuditLog{
		Action:        action,
		Token:         token,
		FieldType:     fieldType,
		AccessLevel:   accessLevel,
		IPAddress:     ip,
		Success:       success,
		FailureReason: reason,
	}

	if actorID != "" {
		if id, err := uuid.Parse(actorID); err == nil {
			entry.ActorID = &id
		}
	}

	if err := s.repo.Create(entry); err != nil {
		log.Printf("audit log failed: %v", err)
	}
}

func (s *Service) GetAll(page, limit int) ([]models.AuditLog, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *Service) GetByToken(token string) ([]models.AuditLog, error) {
	return s.repo.FindByToken(token)
}
