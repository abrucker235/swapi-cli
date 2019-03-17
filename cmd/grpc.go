package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var id int32

var name string

func init() {
	root.AddCommand(grpcCMD)
}

var grpcCMD = &cobra.Command{
	Use:   "grpc",
	Short: "Retrieve Star Wars Data using GRPC API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("please provide sub-command and options")
	},
}
