package cmdutil

import (
	"errors"
	"fmt"
	"strings"

	"github.com/planetscale/cli/internal/config"
	"github.com/planetscale/cli/internal/printer"
	ps "github.com/planetscale/planetscale-go/planetscale"
	"github.com/spf13/cobra"
)

// Helper is passed to every single command and is used by individual
// subcommands.
type Helper struct {
	// Config contains globally sourced configuration
	Config *config.Config

	// Client returns the PlanetScale API client
	Client func() (*ps.Client, error)

	// Printer is used to print output of a command to stdout.
	Printer *printer.Printer
}

// RequiredArgs returns a short and actionable error message if the given
// required arguments are not available.
func RequiredArgs(reqArgs ...string) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		n := len(reqArgs)
		if len(args) >= n {
			return nil
		}

		missing := reqArgs[len(args):]

		a := fmt.Sprintf("arguments <%s>", strings.Join(missing, ", "))
		if len(missing) == 1 {
			a = fmt.Sprintf("argument <%s>", missing[0])
		}

		return fmt.Errorf("missing %s \n\n%s", a, cmd.UsageString())
	}
}

// CheckAuthentication checks whether the user is authenticated and returns a
// actionable error message.
func CheckAuthentication(cfg *config.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cfg.IsAuthenticated() {
			return nil
		}

		return errors.New("not authenticated yet. Please run 'pscale auth login'")
	}
}