package main

import (
	"fmt"
	"github.com/g0r0d3tsky/DSSDutyBot/internal"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/config"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type app struct {
	UC     *internal.Service
	config *config.Config
	logger *log.Logger
}

func main() {
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
	logger := log.New(os.Stdout, "", log.Ldate|log.LUTC)
	app := &app{
		config: c,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
