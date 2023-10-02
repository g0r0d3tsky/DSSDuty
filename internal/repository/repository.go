package repository

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type rw struct {
	store *pgxpool.Pool
}

// TODO: mock
type User interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
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
	GetEventsByPeriod(ctx context.Context, start time.Time, end time.Time) ([]*domain.Event, error)
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}

type RewardsSanctions interface {
	CreateRewardsSanctions(ctx context.Context, rewSan *domain.RewardsSanctions) error
	GetRewardSanctionsByUserId(ctx context.Context, userID uuid.UUID) ([]*domain.RewardsSanctions, error)
	UpdateRewardsSanctionsByUserId(ctx context.Context, rewSan *domain.RewardsSanctions) error
	DeleteRewardsSanctions(ctx context.Context, rewSanId uuid.UUID) error
}

// go:generate mockery --name ServiceRepository
type ServiceRepository interface {
	User
	Duty
	Event
	RewardsSanctions
}

func New(dbPool *pgxpool.Pool) ServiceRepository {
	return rw{
		store: dbPool,
	}
}
