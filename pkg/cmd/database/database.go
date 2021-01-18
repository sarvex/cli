package database

import (
	"github.com/planetscale/cli/config"
	"github.com/spf13/cobra"
)

// DatabaseCmd encapsulates the commands for creating a database
func DatabaseCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "database <command>",
		Short:   "Create, read, destroy, and update databases",
		Aliases: []string{"db"},
		Long:    "TODO",
	}

	cmd.PersistentFlags().Bool("json", false, "Show output as JSON")

	cmd.AddCommand(CreateCmd(cfg))
	cmd.AddCommand(ListCmd(cfg))
	cmd.AddCommand(DeleteCmd(cfg))
	cmd.AddCommand(GetCmd(cfg))

	return cmd
}
