package repository

import (
	"context"
	"errors"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

func (rw rw) CreateUser(ctx context.Context, u *domain.User) error {
	if _, err := rw.store.Exec(
		ctx,
		`INSERT INTO "USER" (id, email, password_hash, activated, full_name, duty_amount, role, course) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		&u.Id, &u.Email, &u.Password.Hash, &u.Activated, &u.FullName, &u.DutyAmount, &u.Role, &u.Course,
	); err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "USER_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}
func (rw rw) CreateStimulation(ctx context.Context, userId uuid.UUID, stimul *domain.Stimulation) error {
	if _, err := rw.store.Exec(ctx,
		`INSERT INTO STIMULATION (id, user_id, stimulation, info, timestamp)
 				VALUES ($1, $2, $3, $4, $5)`,
		stimul.Id, userId, stimul.Stimulation, stimul.Info, stimul.DateTime); err != nil {
		return err
	}
	return nil
}
func (rw rw) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user := &domain.User{}

	if err := rw.store.QueryRow(
		ctx,
		`SELECT id, email, duty_amount, role, full_name, course FROM "USER" u WHERE u.id = $1`, userID,
	).Scan(&user.Id, &user.Email, &user.DutyAmount, &user.Role, &user.FullName, &user.Course); err != nil {
		return nil, err
	}

	return user, nil
}
func (rw rw) GetUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User

	rows, err := rw.store.Query(
		ctx,
		`SELECT id, email, duty_amount, role, full_name, course FROM "USER"`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := &domain.User{}

		if err := rows.Scan(&user.Id, &user.Email, &user.DutyAmount, &user.Role, &user.FullName, &user.Course); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (rw rw) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	if err := rw.store.QueryRow(
		ctx,
		`SELECT id, email, duty_amount, role, full_name, course FROM "USER" u WHERE u.email = $1`, email,
	).Scan(&user.Id, &user.Email, &user.DutyAmount, &user.Role, &user.FullName, &user.Course); err != nil {
		return nil, err
	}

	return user, nil
}

//	func (rw rw) GetStimulationByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Stimulation, error) {
//		var stimuls []*domain.Stimulation
//
//		rows, err := rw.store.Query(
//			ctx,
//			`SELECT stimulation.id, stimulation.info, stimulation.timestamp, stimulation.stimulation
//					FROM "stimulation" WHERE user_id=$1`,
//			userID)
//
//		if err != nil {
//			return nil, err
//		}
//		for rows.Next() {
//			stimul := &domain.Stimulation{}
//
//			if err := rows.Scan(&stimul.Id, &stimul.Info, &stimul.DateTime, &stimul.Stimulation); err != nil {
//				return nil, err
//			}
//
//			stimuls = append(stimuls, stimul)
//		}
//
//		return stimuls, nil
//	}
func (rw rw) GetStimulationByPeriod(ctx context.Context, start time.Time, end time.Time, userId uuid.UUID) ([]*domain.Stimulation, error) {
	var stimuls []*domain.Stimulation

	rows, err := rw.store.Query(
		ctx,
		`SELECT stimulation.id, stimulation.info, stimulation.timestamp, 
       stimulation.stimulation FROM "STIMULATION" WHERE (timestamp>=$1 AND timestamp<=$2 AND user_id=&3)`,
		start, end, userId)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		stimul := &domain.Stimulation{}

		if err := rows.Scan(&stimul.Id, &stimul.Info, stimul.DateTime, &stimul.Stimulation); err != nil {
			return nil, err
		}

		stimuls = append(stimuls, stimul)
	}

	return stimuls, nil
}

func (rw rw) UpdateUser(ctx context.Context, user *domain.User) error {
	if _, err := rw.store.Exec(
		ctx,
		`UPDATE "USER" SET email=$2, duty_amount=$3, role=$4, full_name=$5, course=$6  WHERE id=$1`,
		user.Id, user.Email, user.DutyAmount, user.Role, user.FullName, user.Course,
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
