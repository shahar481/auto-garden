package cli

import (
	"auto-garden/db"
	"fmt"
	"github.com/c-bata/go-prompt"
	"strings"
)

type commandType struct {
	handler func(args []string) error
	argCount int
}


var (
	commands = map[string]commandType {
		"add-user": {
		addUserCommand,
		4,
		},
		"delete-user": {
			removeUserCommand,
			1,
		},
	}
	suggestionsInstance *[]prompt.Suggest
)


func addUserCommand(args []string) error {
	user := args[0]
	password := args[1]
	writePermission := args[2] == "true"
	readPermission := args[3] == "true"
	return db.AddUser(user, password, writePermission, readPermission)
}

func removeUserCommand(args []string) error {
	user := args[0]
	return db.DeleteUser(user)
}


func commandExecutor(command string) {
	splitCommand := strings.Split(strings.TrimSpace(command), " ")
	if command, ok := commands[splitCommand[0]]; ok {
		if len(splitCommand[1:]) == command.argCount {
			err := command.handler(splitCommand[1:])
			if err != nil {
				fmt.Printf("%+v\n", err)
			}
		} else {
			println("Not enough args!")
		}

	}
}

func setSuggestions() {
	suggestions := make([]prompt.Suggest, 0)
	for command, _ := range commands {
		suggestions = append(suggestions, prompt.Suggest{
			Text:command,
		})
	}
	suggestionsInstance = &suggestions
}


func commandSuggester(t prompt.Document) []prompt.Suggest {
	if suggestionsInstance == nil {
		setSuggestions()
	}
	return *suggestionsInstance
}


func StartCli() {
	p := prompt.New(
		commandExecutor,
		commandSuggester)
	p.Run()
}