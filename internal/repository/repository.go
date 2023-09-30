package repository

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type rw struct {
	store *pgxpool.Pool
}

type User interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, role string, username string, dutyAmount int) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

//type Duty interface {
//	CreateDuty(ctx context.Context, duty *domain.Duty) error
//	GetDutyByUserID(ctx context.Context, userID uuid.UUID) (*domain.Duty, error)
//	GetDutyByPeriod(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]*domain.Duty, error)
//	UpdateDuty(ctx context.Context, userID1 uuid.UUID, userID2 uuid.UUID, time time.Time) error
//	DeleteDuty(ctx context.Context, dutyID uuid.UUID) error
//}
//
//type Event interface {
//	CreateEvent(ctx context.Context, event *domain.Event) error
//	GetEventByType(ctx context.Context, eType string) (*domain.Event, error)
//	GetEventByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Event, error)
//	GetEventByPeriod(ctx context.Context, start time.Time, end time.Time) ([]*domain.Event, error)
//	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
//}
//type RewardsSanctions interface {
//	CreateRewardsSanctions(ctx context.Context, rewSan *domain.RewardsSanctions) error
//	GetRewardSanctionsByUserId(ctx context.Context, userId uuid.UUID) (*domain.RewardsSanctions, error)
//	UpdateRewardsSanctionsByUserId(ctx context.Context, userId uuid.UUID, rewards int, sanctions int, info string) error
//}

// go:generate mockery --name ServiceRepository
type ServiceRepository interface {
	User
	//Duty
	//Event
	//RewardsSanctions
}

func New(dbPool *pgxpool.Pool) ServiceRepository {
	return rw{
		store: dbPool,
	}
}
