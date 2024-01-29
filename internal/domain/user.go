package domain

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	ADMIN  = "ADMIN"
	WORKER = "WORKER"
)

type Stimulation struct {
	Id          uuid.UUID `json:"id"`
	Stimulation bool      `json:"stimulation"`
	Info        string    `json:"info"`
	DateTime    time.Time `json:"date_time"`
}

type User struct {
	Id          uuid.UUID     `json:"id"`
	Role        string        `json:"role"`
	DutyAmount  int           `json:"duty_amount"`
	FullName    string        `json:"full_name"`
	Course      int           `json:"course"`
	Stimulation []Stimulation `json:"stimulation"`
	CreatedAt   time.Time     `json:"created_at"`
	Email       string        `json:"email"`
	Password    password      `json:"-"`
	Activated   bool          `json:"activated"`
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
