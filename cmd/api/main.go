package main

import (
	"fmt"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/config"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	c, err := config.Read()
	fmt.Printf("tg:  %v    _+_+_+_+_+__++", c.TelegramToken)

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
	//router := mux.NewRouter()
	//srv := &http.Server{
	//	Addr:         c.ServerAddress(),
	//	Handler:      router,
	//	ReadTimeout:  4,
	//	WriteTimeout: 4,
	//	IdleTimeout:  60,
	//}
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil {
	//		log.Fatal("failed to start server")
	//	}
	//}()
	//
	//log.Println("server started")
	//<-ctx.Done()
	//fmt.Println(c.TelegramToken)
	bot, err := tgbotapi.NewBotAPI(c.TelegramToken)
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)

	bot.Debug = true

}
