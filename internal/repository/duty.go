package repository

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"time"
)

func (rw rw) GetDutyByID(ctx context.Context, dutyID uuid.UUID) (*domain.Duty, error) {
	duty := &domain.Duty{}

	if err := rw.store.QueryRow(
		ctx,
		`SELECT id, date, user_id1, user_id2, amount FROM "DUTY" u WHERE u.id = $1`, dutyID,
	).Scan(&duty.Id, &duty.Date, &duty.UserId.First, &duty.UserId.Second, &duty.Amount); err != nil {
		return nil, err
	}

	return duty, nil
}

func (rw rw) CreateDuty(ctx context.Context, duty *domain.Duty) error {
	_, err := rw.store.Exec(ctx,
		`INSERT INTO "DUTY" (id, date, user_id1, user_id2) VALUES($1, $2, $3, $4)`,
		duty.Id, duty.Date, duty.UserId.First, duty.UserId.Second,
	)
	if err != nil {
		return err
	}
	return nil
}
func (rw rw) GetDutyByPeriod(ctx context.Context, userID uuid.UUID,
	start time.Time, end time.Time) ([]*domain.Duty, error) {
	var duties []*domain.Duty

	rows, err := rw.store.Query(
		ctx,
		`SELECT  id, date, user_id1, user_id2, amount FROM "DUTY" WHERE (user_id1=$1 OR user_id2=$2)
				AND date>=$3
				AND date<=$4`,
		userID, userID, start, end)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		duty := &domain.Duty{}

		if err := rows.Scan(&duty.Id, &duty.Date, &duty.UserId.First, &duty.UserId.Second, &duty.Amount); err != nil {
			return nil, err
		}

		duties = append(duties, duty)
	}

	return duties, nil
}

func (rw rw) GetDutyByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Duty, error) {

	var duties []*domain.Duty

	rows, err := rw.store.Query(
		ctx,
		`SELECT id, date, user_id1, user_id2 FROM "DUTY" WHERE (user_id1=$1 OR user_id2=$2)`,
		userID, userID)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		duty := &domain.Duty{}

		if err := rows.Scan(&duty.Id, &duty.UserId.First, &duty.UserId.Second, &duty.Date); err != nil {
			return nil, err
		}

		duties = append(duties, duty)
	}

	return duties, nil
}
func (rw rw) UpdateDuty(ctx context.Context, duty *domain.Duty) error {
	if _, err := rw.store.Exec(
		ctx,
		`UPDATE "DUTY" SET user_id1=$2, user_id2=$3, date=$4 WHERE id=$1`,
		duty.Id, duty.UserId.First, duty.UserId.Second, duty.Date,
	); err != nil {
		return err
	}
	return nil
}
func (rw rw) DeleteDuty(ctx context.Context, dutyID uuid.UUID) error {
	if _, err := rw.store.Exec(ctx,
		`DELETE FROM "DUTY" WHERE id=$1`,
		dutyID,
	); err != nil {
		return err
	}
	return nil
}
