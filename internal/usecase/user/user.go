package user

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *domain.User) error
	CreateStimulation(ctx context.Context, userID uuid.UUID, stimul *domain.Stimulation) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetStimulationByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Stimulation, error)
	GetStimulationForOneMonth(ctx context.Context, userID uuid.UUID) ([]*domain.Stimulation, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	DeleteStimulation(ctx context.Context, rewSanId uuid.UUID) error
}
