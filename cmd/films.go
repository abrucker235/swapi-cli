package cmd

import (
	context "context"
	"log"
	pb "star-wars-api/swapi"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	grpcCMD.AddCommand(filmsGRPC)
}

var filmsGRPC = &cobra.Command{
	Use:   "films",
	Short: "Retrieve Films using GRPC",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			return err
		}

		client := pb.NewSwapiClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if resp, err := client.Films(ctx, &pb.FilmsRequest{Id: 6}); err != nil {
			log.Fatalf("could not greet: %v", err)
		} else {
			log.Printf("%v", resp)
		}

		return nil
	},
}
