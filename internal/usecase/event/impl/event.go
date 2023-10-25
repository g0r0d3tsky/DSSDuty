package impl

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/google/uuid"
	"time"
)

type EventUseCase struct {
	repo repository.ServiceRepository
}

func (uc *EventUseCase) CreateEvent(ctx context.Context, event *domain.Event) error {
	return uc.repo.CreateEvent(ctx, event)
}
func (uc *EventUseCase) GetEventsByType(ctx context.Context, eType string) ([]*domain.Event, error) {
	return uc.repo.GetEventsByType(ctx, eType)
}
func (uc *EventUseCase) GetEventsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Event, error) {
	return uc.repo.GetEventsByUserID(ctx, userID)
}
func (uc *EventUseCase) GetEventsByOneMonth(ctx context.Context) ([]*domain.Event, error) {
	currentTime := time.Now()
	firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0,
		currentTime.Location())
	nextMonth := currentTime.AddDate(0, 1, 0)
	firstDayOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	lastDayOfMonth := firstDayOfNextMonth.Add(-time.Hour * 24)

	es, err := uc.repo.GetEventsByPeriod(ctx, firstDayOfMonth, lastDayOfMonth)
	if err != nil {
		return nil, err
	}
	return es, nil
}
func (uc *EventUseCase) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	return uc.repo.DeleteEvent(ctx, eventID)
}
