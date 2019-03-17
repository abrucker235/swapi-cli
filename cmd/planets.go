package cmd

import (
	"bytes"
	context "context"
	"encoding/json"
	"log"
	"star-wars-api/model"
	pb "star-wars-api/swapi"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	planetsGRPC.Flags().Int32VarP(&id, "id", "i", 0, "Identifier to be used.")
	planetsGRPC.Flags().StringVarP(&name, "name", "n", "", "Name to be used to retrieve planet.")
	grpcCMD.AddCommand(planetsGRPC)
}

var planetsGRPC = &cobra.Command{
	Use: "planets",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		planets, err := Client.Planets(ctx, &pb.PlanetsRequest{})
		if err != nil {
			log.Fatalf("Error calling planets: %v", err)
		}
		defer cancel()

		people, err := Client.People(ctx, &pb.PeopleRequest{})

		prettyPlanets := butify(planets)
		prettyPeople := butify(people)

		log.Println(string(prettyPlanets.Bytes()))
		log.Println(string(prettyPeople.Bytes()))

		planetsMap := make(map[int32]model.Planet)
		for _, planet := range planets.Planets {
			if id > 0 && id == planet.Id {
				planetsMap[planet.Id] = mapPlanet(planet)
				break
			}

			if name != "" && name == planet.Name {
				planetsMap[planet.Id] = mapPlanet(planet)
				break
			}

			planetsMap[planet.Id] = mapPlanet(planet)
		}

		for _, person := range people.People {
			if planet, ok := planetsMap[person.Homeworld]; ok {
				planet.Characters = append(planet.Characters, person.Name)
				planetsMap[planet.ID] = planet
			}
		}

		v := make([]model.Planet, 0, len(planetsMap))
		for _, value := range planetsMap {
			v = append(v, value)
		}

		prettyCombined := butify(v)
		log.Println(string(prettyCombined.Bytes()))

		return nil
	},
}

func butify(i interface{}) bytes.Buffer {
	var pretty bytes.Buffer
	bytes, err := json.Marshal(&i)
	if err != nil {
		log.Fatal("Error marshalling JSON.")
	}
	err = json.Indent(&pretty, bytes, "", "  ")
	if err != nil {
		log.Fatal("Error butifying JSON.")
	}
	return pretty
}

func mapPlanet(planet *pb.Planet) model.Planet {
	return model.Planet{
		ID:             planet.Id,
		Name:           planet.Name,
		RotationPeriod: planet.RotationPeriod,
		OrbitalPeriod:  planet.OrbitalPeriod,
		Diameter:       planet.Diameter,
		Climate:        planet.Climate,
		Terrain:        planet.Terrain,
		SurfaceWater:   planet.SurfaceWater,
		Population:     planet.Population,
		URL:            planet.Url,
		Films:          planet.Films,
		Created:        planet.Created,
		Edited:         planet.Edited,
	}
}
