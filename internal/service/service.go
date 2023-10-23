package service

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/google/uuid"
	"time"
)

type Auth interface {
	AuthUser(ctx context.Context, username string) (string, error)
}
type User interface {
	CreateUser(ctx context.Context, user *domain.User) error
	CreateStimulation(ctx context.Context, userID uuid.UUID) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetStimulationByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Stimulation, error)
	GetStimulationForOneMonth(ctx context.Context, userID uuid.UUID) ([]*domain.Stimulation, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	DeleteStimulation(ctx context.Context, rewSanId uuid.UUID) error
}

type Duty interface {
	CreateDuty(ctx context.Context, duty *domain.Duty) error
	GetDutyByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error)
	GetDutyByPeriod(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]*domain.Duty, error)
	UpdateDuty(ctx context.Context, duty *domain.Duty) error
	DeleteDuty(ctx context.Context, dutyID uuid.UUID) error
}

type Event interface {
	CreateEvent(ctx context.Context, event *domain.Event) error
	GetEventsByType(ctx context.Context, eType string) ([]*domain.Event, error)
	GetEventsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Event, error)
	GetEventsByOneMonth(ctx context.Context, start time.Time, end time.Time) ([]*domain.Event, error)
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}
type ServiceUsecase interface {
	User
	Duty
	Event
}
type uc struct {
	serviceRepo repository.ServiceRepository
}

func New(mr repository.ServiceRepository) *uc {
	return &uc{
		serviceRepo: mr,
	}
}
