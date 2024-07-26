package http

import (
	"helthmed-auth/internal/auth/domain"
	"helthmed-auth/internal/auth/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	usecase usecase.PatientUseCase
}

func NewPatientHandler(u usecase.PatientUseCase) PatientHandler {
	return PatientHandler{usecase: u}
}

func (h PatientHandler) Register(c *gin.Context) {
	var patient domain.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Register(patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}
