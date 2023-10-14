package repository

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"time"
)

func (rw rw) CreateEvent(ctx context.Context, event *domain.Event) error {
	if event == nil {
		return nil
	}
	if _, err := rw.store.Exec(ctx,
		`INSERT INTO "EVENT" (id, user_id, type, timestamp) VALUES($1, $2, $3, $4)`,
		event.Id, event.UserId, event.Type, event.TimeStamp,
	); err != nil {
		return err
	}
	return nil
}
func (rw rw) GetEventsByType(ctx context.Context, eType string) ([]*domain.Event, error) {
	var events []*domain.Event

	rows, err := rw.store.Query(
		ctx,
		`SELECT * FROM "EVENT" WHERE type=$1`,
		eType)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		event := &domain.Event{}

		if err := rows.Scan(&event.Id, &event.UserId, &event.Type, &event.TimeStamp); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
func (rw rw) GetEventsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Event, error) {
	var events []*domain.Event

	rows, err := rw.store.Query(
		ctx,
		`SELECT * FROM "EVENT" WHERE user_id=$1`,
		userID)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		event := &domain.Event{}

		if err := rows.Scan(&event.Id, &event.UserId, &event.Type, &event.TimeStamp); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
func (rw rw) GetEventsByPeriod(ctx context.Context, start time.Time, end time.Time) ([]*domain.Event, error) {
	var events []*domain.Event

	rows, err := rw.store.Query(
		ctx,
		`SELECT * FROM "EVENT" WHERE (timestamp>=$1 AND timestamp<=$2)`,
		start, end)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		event := &domain.Event{}

		if err := rows.Scan(&event.Id, &event.UserId, &event.Type, &event.TimeStamp); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
func (rw rw) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	if _, err := rw.store.Exec(ctx,
		`DELETE FROM "EVENT" WHERE id=$1`,
		eventID); err != nil {
		return err
	}
	return nil
}
