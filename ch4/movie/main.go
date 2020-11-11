package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Color:  false,
			Actors: []string{"Humprey Bogart", "Ingrid Bergman"},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Color:  true,
			Actors: []string{"Paul Newman"},
		},
		{
			Title:  "Bullitt",
			Year:   1969,
			Color:  true,
			Actors: []string{"Steve McQueen", "Jacquelline Bisset"},
		},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("Data: %s\n", data)
	fmt.Println("")
	data, err = json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("JSON marshalingIndent failed: %s", err)
	}
	fmt.Printf("Data: %s\n", data)
}
