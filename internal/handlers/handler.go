package handlers

import "github.com/g0r0d3tsky/DSSDutyBot/internal/usecase"

type Handler struct {
	services *service.ServiceUsecase
}

func NewHandler(services *service.ServiceUsecase) *Handler {
	return &Handler{services: services}
}
