package repository

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
)

func (rw rw) CreateUser(ctx context.Context, u *domain.User) error {
	if _, err := rw.store.Exec(
		ctx,
		`INSERT INTO "USER" (id, username, duty_amount, role) VALUES ($1, $2, $3, $4)`,
		u.Id, u.Username, u.DutyAmount, u.Role,
	); err != nil {
		return err
	}
	return nil
}
func (rw rw) CreateStimulation(ctx context.Context, user *domain.User) error {
	if _, err := rw.store.Exec(ctx,
		`INSERT INTO STIMULATION (id, user_id, rewards, sanctions, info, timestamp)
 				VALUES ($1, $2, $3, $4, $5, $6)`,
		user.Stimulation[len(user.Stimulation)-1].Id, user.Id, user.Stimulation[len(user.Stimulation)-1].Rewards, user.Stimulation[len(user.Stimulation)-1].Sanctions, user.Stimulation[len(user.Stimulation)-1].Info, user.Stimulation[len(user.Stimulation)-1].DateTime); err != nil {
		return err
	}
	return nil
}
func (rw rw) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user := &domain.User{}

	if err := rw.store.QueryRow(
		ctx,
		`SELECT id, username, duty_amount, role, full_name, course FROM "USER" u WHERE u.id = $1`, userID,
	).Scan(&user.Id, &user.Username, &user.DutyAmount, &user.Role, &user.FullName, &user.Course); err != nil {
		return nil, err
	}

	return user, nil
}

func (rw rw) UpdateUser(ctx context.Context, user *domain.User) error {
	if _, err := rw.store.Exec(
		ctx,
		`UPDATE "USER" SET username=$2, duty_amount=$3, role=$4, full_name=$5, course=$6  WHERE id=$1`,
		user.Id, user.Username, user.DutyAmount, user.Role, user.FullName, user.Course,
	); err != nil {
		return err
	}
	return nil
}

func (rw rw) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if _, err := rw.store.Exec(ctx,
		`DELETE FROM "USER" WHERE id=$1`,
		userID,
	); err != nil {
		return err
	}
	return nil
}
func (rw rw) DeleteStimulation(ctx context.Context, rewSanId uuid.UUID) error {
	if _, err := rw.store.Exec(ctx,
		`DELETE FROM STIMULATION WHERE id=$1`,
		rewSanId); err != nil {
		return err
	}
	return nil
}
