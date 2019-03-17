package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(http)
}

var http = &cobra.Command{
	Use:   "http",
	Short: "Retrieve Star Wars Data using HTTP REST API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("please provide sub-command and options")
	},
}
