package impl

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/google/uuid"
	"time"
)

type UserUseCase struct {
	repo repository.ServiceRepository
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	return uc.repo.CreateUser(ctx, user)
}
func (uc *UserUseCase) CreateStimulation(ctx context.Context, userID uuid.UUID, stimul *domain.Stimulation) error {
	return uc.repo.CreateStimulation(ctx, userID, stimul)
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return uc.repo.GetUserByID(ctx, userID)
}
func (uc *UserUseCase) GetStimulationForOneMonth(ctx context.Context, userID uuid.UUID) ([]*domain.Stimulation, error) {
	currentTime := time.Now()
	firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0,
		currentTime.Location())
	nextMonth := currentTime.AddDate(0, 1, 0)
	firstDayOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	lastDayOfMonth := firstDayOfNextMonth.Add(-time.Hour * 24)

	st, err := uc.repo.GetStimulationByPeriod(ctx, firstDayOfMonth, lastDayOfMonth, userID)
	if err != nil {
		return nil, err
	}
	return st, nil

}
func (uc *UserUseCase) UpdateUser(ctx context.Context, user *domain.User) error {
	return uc.repo.UpdateUser(ctx, user)
}
func (uc *UserUseCase) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return uc.repo.DeleteUser(ctx, userID)
}
func (uc *UserUseCase) DeleteStimulation(ctx context.Context, rewSanId uuid.UUID) error {
	return uc.DeleteStimulation(ctx, rewSanId)
}
