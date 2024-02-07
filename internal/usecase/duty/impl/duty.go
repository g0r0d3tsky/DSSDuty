package impl

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/google/uuid"
	"time"
)

type DutyUseCase struct {
	Repo repository.ServiceRepository
}

func (uc *DutyUseCase) GetDutyByPeriod(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]*domain.Duty, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *DutyUseCase) GetDutyByID(ctx context.Context, dutyID uuid.UUID) (*domain.Duty, error) {
	return uc.Repo.GetDutyByID(ctx, dutyID)
}
func (uc *DutyUseCase) CreateDuty(ctx context.Context, duty *domain.Duty) error {

	return uc.Repo.CreateDuty(ctx, duty)
}
func (uc *DutyUseCase) GetDutyByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error) {
	return uc.Repo.GetDutyByUserID(ctx, userID)
}
func (uc *DutyUseCase) GetDutiesByMonth(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error) {
	currentTime := time.Now()
	firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0,
		currentTime.Location())
	nextMonth := currentTime.AddDate(0, 1, 0)
	firstDayOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	lastDayOfMonth := firstDayOfNextMonth.Add(-time.Hour * 24)

	es, err := uc.Repo.GetDutyByPeriod(ctx, userID, firstDayOfMonth, lastDayOfMonth)
	if err != nil {
		return nil, err
	}
	return es, nil
}

func (uc *DutyUseCase) UpdateDuty(ctx context.Context, duty *domain.Duty) error {
	return uc.Repo.UpdateDuty(ctx, duty)
}
func (uc *DutyUseCase) DeleteDuty(ctx context.Context, dutyID uuid.UUID) error {
	return uc.Repo.DeleteDuty(ctx, dutyID)
}
