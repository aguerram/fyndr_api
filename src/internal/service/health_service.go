package service

import (
	"fyndr.com/api/src/internal/api/response/health"
	"gorm.io/gorm"
)

type HealthService interface {
	GetStatus() health.StatusResponse
}

type HealthServiceImpl struct {
	db *gorm.DB
}

func (h HealthServiceImpl) GetStatus() health.StatusResponse {
	db, err := h.db.DB()
	downResponse := health.StatusResponse{Status: "DOWN"}
	if err != nil {
		return downResponse
	}
	err = db.Ping()
	if err != nil {
		return downResponse
	}
	return health.StatusResponse{Status: "UP"}
}

func NewHealthService(db *gorm.DB) HealthService {
	return &HealthServiceImpl{db: db}
}
