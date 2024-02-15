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

func (uc *DutyUseCase) GetAvailableDuty(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error) {
	duties, err := uc.GetDutiesByMonth(ctx, userID)
	var filteredDuties []*domain.Duty
	if err != nil {
		return nil, err
	}
	for _, v := range duties {
		switch v.Amount {
		case 2:
			if v.UserId.Second == uuid.Nil && v.UserId.First != userID ||
				v.UserId.First == uuid.Nil && v.UserId.Second != userID {
				filteredDuties = append(filteredDuties, v)
			}
		case 1:
			if v.UserId.Second == uuid.Nil && v.UserId.First == uuid.Nil {
				filteredDuties = append(filteredDuties, v)
			}
		default:
			return nil, err
		}
	}
	return filteredDuties, nil
}

func (uc *DutyUseCase) GetDutyByID(ctx context.Context, dutyID uuid.UUID) (*domain.Duty, error) {
	return uc.Repo.GetDutyByID(ctx, dutyID)
}
func (uc *DutyUseCase) CreateDuty(ctx context.Context, duty *domain.Duty) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	duty.Id = id
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
