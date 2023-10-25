package duty

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"time"
)

type DutyUseCase interface {
	CreateDuty(ctx context.Context, duty *domain.Duty) error
	GetDutyByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error)
	GetDutyByPeriod(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]*domain.Duty, error)
	UpdateDuty(ctx context.Context, duty *domain.Duty) error
	DeleteDuty(ctx context.Context, dutyID uuid.UUID) error
}
