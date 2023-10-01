package repository

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"time"
)

func (rw rw) CreateDuty(ctx context.Context, duty *domain.Duty) error {
	if duty == nil {
		return nil
	}
	_, err := rw.store.Exec(ctx,
		`INSERT INTO "DUTY" (id, date, user_id1, user_id2) VALUES($1, $2, $3, $4)`,
		duty.Id, duty.Date, duty.UserId1, duty.UserId2,
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
		`SELECT * FROM "DUTY" WHERE (user_id1=$1 OR user_id2=$2)
				AND date>=$3
				AND date<=$4`,
		userID, userID, start, end)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		duty := &domain.Duty{}

		if err := rows.Scan(&duty.Id, &duty.UserId1, &duty.UserId2, &duty.Date); err != nil {
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
		`SELECT * FROM "DUTY" WHERE (user_id1=$1 OR user_id2=$2)`,
		userID, userID)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		duty := &domain.Duty{}

		if err := rows.Scan(&duty.Id, &duty.UserId1, &duty.UserId2, &duty.Date); err != nil {
			return nil, err
		}

		duties = append(duties, duty)
	}

	return duties, nil
}
func (rw rw) UpdateDuty(ctx context.Context, userID1 uuid.UUID, userID2 uuid.UUID, time time.Time, dutyId uuid.UUID) error {
	if _, err := rw.store.Exec(
		ctx,
		`UPDATE "DUTY" SET user_id1=$2, user_id2=$3, date=$4 WHERE id=$1`,
		dutyId, userID1, userID2, time,
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
