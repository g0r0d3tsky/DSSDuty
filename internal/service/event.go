package service

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"time"
)

func (uc uc) CreateEvent(ctx context.Context, event *domain.Event) error {
	return uc.serviceRepo.CreateEvent(ctx, event)
}
func (uc uc) GetEventsByType(ctx context.Context, eType string) ([]*domain.Event, error) {
	return uc.serviceRepo.GetEventsByType(ctx, eType)
}
func (uc uc) GetEventsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Event, error) {
	return uc.serviceRepo.GetEventsByUserID(ctx, userID)
}
func (uc uc) GetEventsByOneMonth(ctx context.Context) ([]*domain.Event, error) {
	currentTime := time.Now()
	firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0,
		currentTime.Location())
	nextMonth := currentTime.AddDate(0, 1, 0)
	firstDayOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	lastDayOfMonth := firstDayOfNextMonth.Add(-time.Hour * 24)

	es, err := uc.serviceRepo.GetEventsByPeriod(ctx, firstDayOfMonth, lastDayOfMonth)
	if err != nil {
		return nil, err
	}
	return es, nil
}
func (uc uc) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	return uc.serviceRepo.DeleteEvent(ctx, eventID)
}
