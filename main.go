package main

import (
	"bytes"
	"database/sql"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

const (
	basicError = "Ha habido un problema al ejecutar el comando."
	tomlFile   = "./config.toml"
)

func main() {

	privateConfig, err := ParseConfig(tomlFile)
	if err != nil {
		panic(err)
	}

	dbConfig := CreateDbInfo(privateConfig.BD)
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(privateConfig.Bot.Token)
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {

		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "curso":
			course := ParseCourse(update.Message.CommandArguments())

			// The course provided is not valid
			if course == -1 {
				BotReturnError(
					bot,
					update,
					"Comando incorrecto.\nPrueba con /course [1,2,3,4]")

			} else {

				query := CourseQuery(course)
				rows, err := db.Query(query)

				if err != nil {
					BotReturnError(
						bot,
						update,
						basicError)

				} else {
					// Append all delegates from the course to the list
					var delegatesList []Delegate

					for rows.Next() {
						delegate := Delegate{}
						err = rows.Scan(
							&delegate.Name,
							&delegate.Surname,
							&delegate.NIA)
						if err != nil {
							panic(err)
						}

						delegatesList = append(delegatesList, delegate)
					}

					// Format all delegates in a message
					var result bytes.Buffer
					for _, person := range delegatesList {
						result.WriteString(strings.Join(
							[]string{person.Name, person.Surname, person.NIA},
							"-",
						))
						result.WriteString("\n")
					}
					msg := tgbotapi.NewMessage(update.Message.Chat.ID,
						result.String())
					bot.Send(msg)
				}

			} // End of switch
		}
	}
}
