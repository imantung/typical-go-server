package typimain

import (
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typitask"
	"gopkg.in/urfave/cli.v1"
)

// TypicalTaskTool represent typical task tool application
type TypicalTaskTool struct {
	typitask.TypicalTask
}

// NewTypicalTaskTool return new instance of TypicalCli
func NewTypicalTaskTool(context typictx.Context) *TypicalTaskTool {
	return &TypicalTaskTool{typitask.TypicalTask{context}}
}

// Cli return the command line interface
func (t *TypicalTaskTool) Cli() *cli.App {
	app := cli.NewApp()
	app.Name = t.Name + " (TYPICAL)"
	app.Usage = ""
	app.Description = t.Description
	app.Version = t.Version
	app.Commands = t.StandardCommands()
	for key := range t.Modules {
		module := t.Modules[key]
		if module.Command != nil {
			app.Commands = append(app.Commands, *module.Command)
		}

	}

	for key := range t.Commands {
		command := t.Commands[key]
		app.Commands = append(app.Commands, command)
	}

	// NOTE: export the enviroment before run
	// exportEnviroment()

	return app
}