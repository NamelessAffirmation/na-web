package cmd

import (
	"github.com/morpheuszero/go-heimdall"
	"github.com/snowlynxsoftware/na-api/config"
)

type MigrationCommand struct {
	AppConfig *config.AppConfig
}

func NewMigrationCommand() *MigrationCommand {
	return &MigrationCommand{
		AppConfig: config.NewAppConfig(),
	}
}

func (c *MigrationCommand) RunCommand() error {
	h := heimdall.NewHeimdall(c.AppConfig.DBConnectionString, "migration_history", "./migrations", true)
	return h.RunMigrations()
}
