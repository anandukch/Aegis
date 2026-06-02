package rbac

import (
	"net/http"

	"github.com/anandudevops/aegis/internal/auth"
	"github.com/anandudevops/aegis/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	authSvc *auth.Service
}

func NewHandler(authSvc *auth.Service) *Handler {
	return &Handler{authSvc: authSvc}
}

func (h *Handler) GetRoles(c *gin.Context) {
	roles := map[string]map[string]string{
		"ADMIN":   {"all fields": "FULL"},
		"ANALYST": {"email": "MASKED", "name": "MASKED", "card_number": "DENIED", "default": "MASKED"},
		"SERVICE": {"card_number": "FULL", "default": "MASKED"},
		"VIEWER":  {"all fields": "MASKED"},
	}
	response.Success(c, http.StatusOK, roles)
}

func (h *Handler) AssignRole(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	var req struct {
		Role string `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authSvc.AssignRole(userID, req.Role); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"message": "role updated"})
}
