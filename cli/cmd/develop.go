package cmd

import (
	"bytes"
	"context"
	"os"
	"os/exec"

	"github.com/omni-network/omni/lib/errors"

	"github.com/spf13/cobra"
)

func newDeveloperCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "developer",
		Short: "XApp development commands",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		newForgeProjectCmd(),
	)

	return cmd
}

func newForgeProjectCmd() *cobra.Command {
	var cfg developerForgeProjectConfig

	cmd := &cobra.Command{
		Use:   "new",
		Short: "Scaffold a Forge project",
		Long:  `Scaffold a new Forge project accompanied by simple mocked testing and a multi-chain deployment script.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return newForgeProjectTemplate(cmd.Context(), cfg)
		},
	}

	bindDeveloperForgeProjectConfig(cmd, &cfg)

	return cmd
}

type developerForgeProjectConfig struct {
	templateName string
}

// newForgeProjectTemplate creates a new project using the forge template.
func newForgeProjectTemplate(_ context.Context, cfg developerForgeProjectConfig) error {
	// Check if forge is installed
	if !checkForgeInstalled() {
		// Forge is not installed, return an error with a suggestion.
		return &cliError{
			Msg:     "forge is not installed.",
			Suggest: "You can install foundry by visiting https://github.com/foundry-rs/foundry",
		}
	}

	// Execute the `forge init` command.
	// #nosec G204
	cmd := exec.Command("forge", "init", "--template", "https://github.com/omni-network/"+cfg.templateName+".git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to run forge init")
	}

	return nil
}

// checkForgeInstalled checks if forge is installed by attempting to run 'forge --version'.
func checkForgeInstalled() bool {
	cmd := exec.Command("forge", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return err == nil // If there is no error, forge is installed.
}
