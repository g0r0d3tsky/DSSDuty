package service

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"time"
)

func (uc uc) CreateDuty(ctx context.Context, duty *domain.Duty) error {
	return uc.serviceRepo.CreateDuty(ctx, duty)
}
func (uc uc) GetDutyByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error) {
	return uc.serviceRepo.GetDutyByUserID(ctx, userID)
}
func (uc uc) GetDutiesByMonth(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error) {
	currentTime := time.Now()
	firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0,
		currentTime.Location())
	nextMonth := currentTime.AddDate(0, 1, 0)
	firstDayOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	lastDayOfMonth := firstDayOfNextMonth.Add(-time.Hour * 24)

	es, err := uc.serviceRepo.GetDutyByPeriod(ctx, userID, firstDayOfMonth, lastDayOfMonth)
	if err != nil {
		return nil, err
	}
	return es, nil
}

func (uc uc) UpdateDuty(ctx context.Context, duty *domain.Duty) error {
	return uc.serviceRepo.UpdateDuty(ctx, duty)
}
func (uc uc) DeleteDuty(ctx context.Context, dutyID uuid.UUID) error {
	return uc.serviceRepo.DeleteDuty(ctx, dutyID)
}
