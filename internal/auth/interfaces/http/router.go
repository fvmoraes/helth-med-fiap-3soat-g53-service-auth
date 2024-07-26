package http

import (
	"helthmed-auth/internal/auth/infrastructure/repository"
	"helthmed-auth/internal/auth/usecase"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	doctorRepo := repository.NewDoctorRepository()
	doctorUsecase := usecase.NewDoctorUseCase(doctorRepo)
	doctorHandler := NewDoctorHandler(doctorUsecase)

	patientRepo := repository.NewPatientRepository()
	patientUsecase := usecase.NewPatientUseCase(patientRepo)
	patientHandler := NewPatientHandler(patientUsecase)

	api := router.Group("/api")
	{
		api.POST("/doctors", doctorHandler.Register)
		api.POST("/patients", patientHandler.Register)
	}
}
