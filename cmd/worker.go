package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func workerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "Run Worker Service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print("Run Worker Service")
			return nil
		},
	}
}
