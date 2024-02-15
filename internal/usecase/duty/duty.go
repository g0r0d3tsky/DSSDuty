package duty

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
)

type DutyUseCase interface {
	CreateDuty(ctx context.Context, duty *domain.Duty) error
	GetDutyByID(ctx context.Context, dutyID uuid.UUID) (*domain.Duty, error)
	GetAvailableDuty(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error)
	GetDutyByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error)
	GetDutiesByMonth(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error)
	UpdateDuty(ctx context.Context, duty *domain.Duty) error
	DeleteDuty(ctx context.Context, dutyID uuid.UUID) error
}
