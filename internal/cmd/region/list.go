package region

import (
	"fmt"

	"github.com/planetscale/cli/internal/cmdutil"
	"github.com/planetscale/cli/internal/printer"
	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/spf13/cobra"
)

// ListCmd is the command for listing all regions for an organization.
func ListCmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List regions",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client, err := ch.Client()
			if err != nil {
				return err
			}

			org := ch.Config.Organization // --org flag
			if org == "" {
				cfg, err := ch.ConfigFS.DefaultConfig()
				if err != nil {
					return nil, cobra.ShellCompDirectiveNoFileComp
				}

				org = cfg.Organization
			}

			end := ch.Printer.PrintProgress("Fetching regions...")
			defer end()
			regions, err := client.Organizations.ListRegions(ctx, &planetscale.ListOrganizationRegionsRequest{
				Organization: org,
			})
			if err != nil {
				switch cmdutil.ErrCode(err) {
				case planetscale.ErrNotFound:
					return fmt.Errorf("organization %s does not exist", printer.BoldBlue(org))
				default:
					return cmdutil.HandleError(err)
				}
			}

			end()

			if len(regions) == 0 && ch.Printer.Format() == printer.Human {
				ch.Printer.Println("No regions have been created yet.")
				return nil
			}

			return ch.Printer.PrintResource(toRegions(regions))
		},
		TraverseChildren: true,
	}

	return cmd
}
