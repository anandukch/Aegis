package vault

import (
	"net/http"

	"github.com/anandudevops/aegis/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Tokenize(c *gin.Context) {
	var req struct {
		FieldType string `json:"field_type" binding:"required"`
		Value     string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	record, err := h.svc.Tokenize(req.FieldType, req.Value, userID.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "tokenization failed")
		return
	}

	response.Success(c, http.StatusCreated, gin.H{
		"token":      record.Token,
		"field_type": record.FieldType,
		"created_at": record.CreatedAt,
	})
}

func (h *Handler) Detokenize(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	role, _ := c.Get("role")
	value, accessLevel, err := h.svc.Detokenize(req.Token, role.(string))
	if err != nil {
		if accessLevel == "DENIED" {
			response.Error(c, http.StatusForbidden, "access denied for this field type")
		} else {
			response.Error(c, http.StatusNotFound, err.Error())
		}
		return
	}

	record, _ := h.svc.GetMetadata(req.Token)
	response.Success(c, http.StatusOK, gin.H{
		"token":        req.Token,
		"field_type":   record.FieldType,
		"value":        value,
		"access_level": accessLevel,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	token := c.Param("token")
	if err := h.svc.Delete(token); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"message": "record deleted"})
}

func (h *Handler) GetMetadata(c *gin.Context) {
	token := c.Param("token")
	record, err := h.svc.GetMetadata(token)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{
		"token":      record.Token,
		"field_type": record.FieldType,
		"created_at": record.CreatedAt,
		"created_by": record.CreatedBy,
	})
}
