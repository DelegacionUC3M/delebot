package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	delegates "github.com/DelegacionUC3M/delebot/delegates"
)

const (
	tomlFile = "./config.toml"
)

func main() {

	privateConfig, err := ParseConfig(tomlFile)
	if err != nil {
		panic(err)
	}

	dbConfig := CreateDbInfo(privateConfig.BD)
	DB, err := sql.Open("postgres", dbConfig)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	Bot, err := tgbotapi.NewBotAPI(privateConfig.Bot.Token)
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := Bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for Update := range updates {

		if Update.Message == nil {
			continue
		}

		switch Update.Message.Command() {
		case "curso":
			course := delegates.ParseCourse(Update.Message.CommandArguments())
			delegates.QueryDelegatesFromCourse(course)
		} // End of switch
	}
}
