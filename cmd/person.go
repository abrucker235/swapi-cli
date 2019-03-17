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
	http.AddCommand(personHTTP)
	grpcCMD.AddCommand(personGRPC)
}

var personHTTP = &cobra.Command{
	Use: "person",
}

var personGRPC = &cobra.Command{
	Use: "person",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			return err
		}

		client := pb.NewSwapiClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if resp, err := client.People(ctx, &pb.PeopleRequest{Id: 1}); err != nil {
			log.Fatalf("Error calling people: %v", err)
		} else {
			log.Printf("%v", resp)
		}

		return nil
	},
}
