package main

import (
	"encoding/json"
	"fmt"
)

var content = `
{
	"species": "ハト",
	"description": "鳩は、鳩目、ハト科に属する鳥類の総称である。",
	"dimensions": {
		"height": 30,
		"width": 40
	}
}
`

type Dimensions struct {
	Width int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species string `json:"species"`
	Description string `json:"description"`
	Dimensions Dimensions `json:"dimensions"`
}

func main() {
	var data Data

	err := json.Unmarshal(([]byte(content)), &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Species: %s\n", data.Species)
	fmt.Printf("Description: %s\n", data.Description)
	fmt.Printf("Dimensions:\n")
	fmt.Printf("  Height: %d\n", data.Dimensions.Height)
}