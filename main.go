package main

import (
	"context"
	"fmt"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/config"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	c, err := config.Read()
	fmt.Println("___________________________________________________")
	fmt.Println(c.TelegramToken)
	fmt.Println("___________________________________________________")

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
	rw := repository.New(dbPool)
	ctx := context.Background()
	var users []*domain.User
	users, _ = rw.GetUsers(ctx)
	for _, user := range users {
		fmt.Println(user)
	}
}
