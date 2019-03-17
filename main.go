package main

import (
	"log"
	"star-wars-api/cmd"
	"star-wars-api/model"

	"star-wars-api/swapi"

	"google.golang.org/grpc"
)

var Characters map[string]model.Character

var Planets map[string]model.Planet

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to establish connection with server.")
	}
	defer conn.Close()

	client := swapi.NewSwapiClient(conn)

	cmd.Client = client

	cmd.Execute()
}

// func retrieveCharacters(client *http.Client) map[string]model.Character {
// 	charactersMap := make(map[string]model.Character)

// 	path := charactersURL
// 	for path != "" {
// 		var characterResponse *model.CharactersResponse = retrieveCharacter(client, path)
// 		for _, character := range characterResponse.Characters {
// 			charactersMap[character.Name] = character
// 		}
// 		path = characterResponse.Next
// 	}

// 	return charactersMap
// }

// func retrieveCharacter(client *http.Client, path string) *model.CharactersResponse {
// 	if resp, err := client.Get(path); err != nil {
// 		panic(err)
// 	} else {
// 		body, _ := ioutil.ReadAll(resp.Body)
// 		var characterResponse model.CharactersResponse
// 		if err := json.Unmarshal(body, &characterResponse); err != nil {
// 			panic(err)
// 		}
// 		return &characterResponse
// 	}
// }

// func retrievePlanets(client *http.Client) map[string]model.Planet {
// 	planetMap := make(map[string]model.Planet)

// 	path := planetsURL
// 	for path != "" {
// 		var planetsResponse *model.PlanetsResponse = retrievePlanet(client, path)
// 		for _, planet := range planetsResponse.Planets {
// 			planetMap[planet.Name] = planet
// 		}
// 		path = planetsResponse.Next
// 	}

// 	return planetMap
// }

// func retrievePlanet(client *http.Client, path string) *model.PlanetsResponse {
// 	if resp, err := client.Get(path); err != nil {
// 		panic(err)
// 	} else {
// 		body, _ := ioutil.ReadAll(resp.Body)
// 		var planetResponse model.PlanetsResponse
// 		if err := json.Unmarshal(body, &planetResponse); err != nil {
// 			panic(err)
// 		}
// 		return &planetResponse
// 	}
// }
