package read

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type declaration.
type Flock struct {
	Flock []Duck `json:"flock, omitempty"`
}

// type declaration
type Duck struct {
	Id        int    `json:"id, omitempty"`
	Firstname string `json:"first_name, omitempty"`
	Lastname  string `json:"last_name, omitempty"`
}

// interface declaration.
type FlockInterfce interface {
	Load(jsonFileName string)
	Save(jsonFileName string)
}

// interface declaration.
type DuckInterface interface {
	GetId() int
	SetId(id int)
	GetFirstname() string
	SetFirstname(firstname string)
	GetLastname() string
	SetLastname(lastname string)
}

// getters and setters.
func (duck Duck) GetId() int {
	return duck.Id
}

func (duck Duck) SetId(id int) {
	duck.Id = id
}

func (duck Duck) GetFirstname() string {
	return duck.Firstname
}

func (duck Duck) SetFirstname(firstname string) {
	duck.Firstname = firstname
}

func (duck Duck) GetLastname() string {
	return duck.Lastname
}

func (duck Duck) SetLastname(lastname string) {
	duck.Lastname = lastname
}

// Load JSON file to Flock.
func (flock *Flock) Load(jsonFileName string) {
	// read json file.
	jsonFile, readError := ioutil.ReadFile(jsonFileName)
	// file reading error.
	if readError != nil {
		// print error message.
		fmt.Printf("Reading JSON file '%s' failed. %s.\n", jsonFileName, readError.Error())
		// exit with error.
		os.Exit(1)
	}
	// unmarhsall json file content ([]byte) to struct content.
	unmarshalError := json.Unmarshal(jsonFile, &flock)
	// unmarshalling error.
	if unmarshalError != nil {
		// print error message.
		fmt.Printf("Unmarshaling JSON file '%s' failed. %s.\n", jsonFileName, unmarshalError.Error())
		// exit with error.
		os.Exit(1)
	}
}

// Save Flock to JSON file.
func (flock Flock) Save(jsonFileName string) {
	// variables.
	var jsonFile []byte
	// marshal struct content to json file conent ([]byte).
	jsonFile, marshalError := json.Marshal(flock)
	// marshaling error.
	if marshalError != nil {
		// print error message.
		fmt.Printf("Marshaling struct failed. %s.\n", marshalError.Error())
		// exit with error.
		os.Exit(1)
	}
	// write json file.
	writeError := ioutil.WriteFile(jsonFileName, jsonFile, 0644)
	// file writing error.
	if writeError != nil {
		// print error message.
		fmt.Printf("Writing JSON file '%s' failed. %s.\n", jsonFileName, writeError.Error())
		// exit with error.
		os.Exit(1)
	}
}
