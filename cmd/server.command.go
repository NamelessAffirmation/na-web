package cmd

import (
	"github.com/snowlynxsoftware/na-api/server"
)

type ServerCommand struct {
}

func NewServerCommand() *ServerCommand {
	return &ServerCommand{}
}

func (c *ServerCommand) RunCommand() {
	server := server.NewAppServer()
	server.Start()
}
