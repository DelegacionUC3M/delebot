package delegates

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/DelegacionUC3M/delebot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	DB     *sql.DB          // To access the delegates' database
	Bot    *tgbotapi.BotAPI // To send messages
	Update *tgbotapi.Update // To access the content of the message
)

// Delegate has all the info required to contact a delegate
type Delegate struct {
	Name    string
	Surname string
	NIA     string
}

// ParseCourse returns the course queried as an int
func ParseCourse(course string) int {
	switch course {
	case "primero":
	case "1":
		return 1
	case "segundo":
	case "2":
		return 2
	case "tercero":
	case "3":
		return 3
	case "cuarto":
	case "4":
		return 4
	}

	// None of the courses match
	return -1
}

// CourseQuery returns the query depending on the course provided
func CourseQuery(course int) string {
	query := fmt.Sprintf("SELECT name, surname, nia FROM delegates WHERE course=%d",
		course)

	return query
}

// QueryDelegatesFromCourse obtains all of the delegates from the course provided
func QueryDelegatesFromCourse(course int) {
	// The course provided is not valid
	if course == -1 {
		tools.BotReturnError("Comando incorrecto.\nPrueba con /course [1,2,3,4]")
	}

	rows, err := DB.Query(CourseQuery(course))
	if err != nil {
		tools.BotReturnError(tools.BasicError)
	}

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
		msgFormatted := fmt.Sprintf("%s %s - %s", person.Name, person.Surname, person.NIA)
		result.WriteString(msgFormatted)
		result.WriteString("\n")
	}

	msg := tgbotapi.NewMessage(Update.Message.Chat.ID,
		result.String())
	Bot.Send(msg)
}
