package service

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"time"
)

func (uc uc) CreateUser(ctx context.Context, user *domain.User) error {
	return uc.serviceRepo.CreateUser(ctx, user)
}
func (uc uc) CreateStimulation(ctx context.Context, userID uuid.UUID, stimul *domain.Stimulation) error {
	return uc.serviceRepo.CreateStimulation(ctx, userID, stimul)
}

func (uc uc) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return uc.serviceRepo.GetUserByID(ctx, userID)
}
func (uc uc) GetStimulationForOneMonth(ctx context.Context, userID uuid.UUID) ([]*domain.Stimulation, error) {
	currentTime := time.Now()
	firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0,
		currentTime.Location())
	nextMonth := currentTime.AddDate(0, 1, 0)
	firstDayOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	lastDayOfMonth := firstDayOfNextMonth.Add(-time.Hour * 24)

	st, err := uc.serviceRepo.GetStimulationByPeriod(ctx, firstDayOfMonth, lastDayOfMonth, userID)
	if err != nil {
		return nil, err
	}
	return st, nil

}
func (uc uc) UpdateUser(ctx context.Context, user *domain.User) error {
	return uc.serviceRepo.UpdateUser(ctx, user)
}
func (uc uc) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return uc.serviceRepo.DeleteUser(ctx, userID)
}
func (uc uc) DeleteStimulation(ctx context.Context, rewSanId uuid.UUID) error {
	return uc.DeleteStimulation(ctx, rewSanId)
}
