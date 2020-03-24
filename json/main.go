package json

import (
	"fmt"

	"workshop/json/read"
)

// main.
func main() {
	// variables.
	var ducks read.Flock
	var inputFileName = "./input.json"
	var outputFileName = "./output.json"
	// load ducks from input xml file.
	ducks.Load(inputFileName)
	// print loaded ducks.
	for _, duck := range ducks.Flock {
		fmt.Printf("%d - %s - %s\n", duck.GetId(), duck.GetFirstname(), duck.GetLastname())
	}
	// save ducks to output xml file.
	ducks.Save(outputFileName)
	// reset ducks.
	ducks.Flock = nil
	// load ducks from output xml file.
	ducks.Load(outputFileName)
	// print loaded ducks.
	for _, duck := range ducks.Flock {
		fmt.Printf("%d - %s - %s\n", duck.GetId(), duck.GetFirstname(), duck.GetLastname())
	}
}
