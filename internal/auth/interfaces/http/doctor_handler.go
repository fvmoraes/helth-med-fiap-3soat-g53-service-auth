package http

import (
	"helthmed-auth/internal/auth/domain"
	"helthmed-auth/internal/auth/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	usecase usecase.DoctorUseCase
}

func NewDoctorHandler(u usecase.DoctorUseCase) *DoctorHandler {
	return DoctorHandler{usecase: u}
}

func (h DoctorHandler) Register(c *gin.Context) {
	var doctor domain.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Register(doctor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, doctor)
}
