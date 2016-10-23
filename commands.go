package main

import (
	"os"

	"github.com/F0rth/consul-check/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Consul check commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"disk": func() (cli.Command, error) {
			return &command.DiskCommand{
				Ui: ui,
			}, nil
		},
		"memory": func() (cli.Command, error) {
			return &command.MemoryCommand{
				Ui: ui,
			}, nil
		},
		"cpuload": func() (cli.Command, error) {
			return &command.CpuloadCommand{
				Ui: ui,
			}, nil
		},
		"service": func() (cli.Command, error) {
			return &command.ServiceCommand{
				Ui: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Version: Version,
				Ui:      ui,
			}, nil
		},
	}
}
