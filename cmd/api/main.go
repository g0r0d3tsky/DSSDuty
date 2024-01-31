package main

import (
	"fmt"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/config"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/mailer"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
)

const version = "1.0.0"

type app struct {
	UC     *usecase.Service
	config *config.Config
	logger *slog.Logger
	mailer mailer.Mailer
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	c, err := config.Read()

	if err != nil {
		log.Println("failed to read config:", err.Error())
		return
	}
	dbPool, err := repository.Connect(c)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		if dbPool != nil {
			dbPool.Close()
		}
	}()
	repo := repository.New(dbPool)

	service := usecase.New(repo)
	app := &app{
		UC:     service,
		config: c,
		logger: logger,
		mailer: mailer.New(c.SMTP.Host, c.SMTP.Port, c.SMTP.Username, c.SMTP.Password, c.SMTP.Sender),
	}

	err = app.serve()
	if err != nil {
		logger.Error("running server", err)
	}

}
