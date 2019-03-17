package cmd

import (
	"fmt"
	"os"
	pb "star-wars-api/swapi"

	"github.com/spf13/cobra"
)

var Client pb.SwapiClient

var root = &cobra.Command{Use: "starwars", Short: "CLI for Star Wars API"}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
