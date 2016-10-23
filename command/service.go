package command

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/mitchellh/cli"
)

// ServiceCommand is a Command implementation that checks local or distant service avaibility
type ServiceCommand struct {
	Ui cli.Ui
}

// Run is the function mapped to the ServiceCommand implementation.
// This is invoked upon calling `consul-check service ...`
func (m *ServiceCommand) Run(args []string) int {
	var protocolService, nameService, hostService string
	var timeoutService, exitCode int

	cmdFlags := flag.NewFlagSet("service", flag.ContinueOnError)
	cmdFlags.Usage = func() { m.Ui.Output(m.Help()) }
	cmdFlags.StringVar(&protocolService, "proto", "tcp", "protocol [udp]|tcp")
	cmdFlags.StringVar(&nameService, "port", "22", "service ip or name")
	cmdFlags.StringVar(&hostService, "host", "127.0.0.1", "hostname or ip")
	cmdFlags.IntVar(&timeoutService, "timeout", 3, "timeout in second")

	if err := cmdFlags.Parse(args); err != nil {
		return 0
	}
	service := hostService + ":" + nameService
	conn, err := net.DialTimeout(protocolService, service, (time.Duration(timeoutService) * time.Second))

	if err != nil {
		fmt.Println("Connection error:", err)
		exitCode = 2
	} else {
		fmt.Println(service, "up")
		exitCode = 0
		defer conn.Close()
	}

	return exitCode
}

// Help returns a string that is the usage for the ServiceCommand
func (m *ServiceCommand) Help() string {
	helpText := `
Usage: consul-check service udp|tcp port <options>

	Check a service of the local or distant system

Options:
  -proto=tcp or udp, default to tcp
  -port= port number or service name
  -host=127.0.0.1   Host ip or name, default localhost
  -timeout=3   Timeout before fail
`

	return strings.TrimSpace(helpText)
}

// Synopsis of the ServiceCommand implementation.
func (m *ServiceCommand) Synopsis() string {
	return "Checks the local service"
}
