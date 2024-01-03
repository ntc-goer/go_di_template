package cmd

import (
	"github.com/spf13/cobra"
	"go_di_template/config"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Run RootCmd ",
		Long:  `Using Cobra Cmd to run many child cmd depend on your command. Refer at https://github.com/spf13/cobra/blob/main/site/content/user_guide.md`,
	}
	cfg *config.Config
)

func Execute() error {
	cfg = config.Load()
	rootCmd.AddCommand(serverCmd())
	rootCmd.AddCommand(workerCmd())
	return rootCmd.Execute()
}
