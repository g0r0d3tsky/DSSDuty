package repository

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
)

func (rw rw) CreateUser(ctx context.Context, u *domain.User) error {
	if u == nil {
		return nil
	}
	if _, err := rw.store.Exec(
		ctx,
		`INSERT INTO "USER" (id, username, duty_amount, role) VALUES ($1, $2, $3, $4)`,
		u.Id, u.Username, u.DutyAmount, u.Role,
	); err != nil {
		return err
	}
	return nil
}
func (rw rw) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user := &domain.User{}

	if err := rw.store.QueryRow(
		ctx,
		`SELECT * FROM "USER" u WHERE u.id = $1`, userID,
	).Scan(&user.Id, &user.Username, &user.DutyAmount, &user.Role); err != nil {
		return nil, err
	}

	return user, nil
}

func (rw rw) UpdateUser(ctx context.Context, userID uuid.UUID, role string, username string, dutyAmount int) error {
	if _, err := rw.store.Exec(
		ctx,
		`UPDATE "USER" SET role=$2, username=$3, duty_amount=$4 WHERE id=$1`,
		userID, role, username, dutyAmount,
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
