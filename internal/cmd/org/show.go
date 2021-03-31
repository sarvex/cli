package org

import (
	"errors"
	"fmt"
	"os"

	"github.com/planetscale/cli/internal/cmdutil"
	"github.com/planetscale/cli/internal/config"

	"github.com/spf13/cobra"
)

func ShowCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display the currently active organization",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile, err := config.ProjectConfigPath()
			if err != nil {
				return err
			}

			cfg, err := config.NewFileConfig(configFile)
			if os.IsNotExist(err) {
				configFile, err = config.DefaultConfigPath()
				if err != nil {
					return err
				}

				cfg, err = config.NewFileConfig(configFile)
				if os.IsNotExist(err) {
					return errors.New("not authenticated, please authenticate with: \"pscale auth login\"")
				}

				if err != nil {
					return err
				}
			}

			if err != nil {
				return err
			}

			if cfg.Organization == "" {
				return errors.New("config file exists, but organization is not set")
			}

			fmt.Printf("%s (from file: %s)", cmdutil.Bold(cfg.Organization), configFile)

			return nil
		},
	}

	return cmd
}
