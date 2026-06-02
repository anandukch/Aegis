package audit

import (
	"net/http"
	"strconv"

	"github.com/anandudevops/aegis/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	logs, total, err := h.svc.GetAll(page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch audit logs")
		return
	}

	response.Success(c, http.StatusOK, gin.H{
		"logs":  logs,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) GetLogsByToken(c *gin.Context) {
	token := c.Param("token")
	logs, err := h.svc.GetByToken(token)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch audit logs")
		return
	}
	response.Success(c, http.StatusOK, gin.H{"logs": logs})
}
