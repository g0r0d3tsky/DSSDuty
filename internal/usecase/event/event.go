package event

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, event *domain.Event) error
	GetEventsByType(ctx context.Context, eType string) ([]*domain.Event, error)
	GetEventsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Event, error)
	GetEventsByOneMonth(ctx context.Context) ([]*domain.Event, error)
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}
