package cmd

import (
	"errors"
	"fmt"
	"os"
)

const SERVER_CMD = "server"
const MIGRATION_CMD = "migration"

type CommandService struct {
}

func NewCommandService() *CommandService {
	return &CommandService{}
}

func (s *CommandService) ParseCommand() error {

	activeCommand := ""

	// If no command is supplied, default to server.
	if len(os.Args) < 2 {
		activeCommand = SERVER_CMD
	} else {
		activeCommand = os.Args[1]
	}

	switch activeCommand {
	case SERVER_CMD:
		fmt.Println("-- RUNNING SERVER COMMAND --")
		NewServerCommand().RunCommand()
		return nil
	case MIGRATION_CMD:
		fmt.Println("-- RUNNING MIGRATION COMMAND --")
		return NewMigrationCommand().RunCommand()
	default:
		return errors.New("expected 'server' or 'migration' subcommands")
	}

	return nil
}
