package marshaling

import (
	"encoding/json"
	"fmt"
)

// Same rules can be applied to Marshaling data

// Structured data

type Bird struct {
	Species     string `json:"birdType"`
	Descreption string `json:"waht is does,omitempty"`
	Dimensions  string `json:"-"` // always ignoer it
}

func MarshalingWithStructuredData() {
	fmt.Println("Marshaling ------------------")
	pigeon := &Bird{
		Species:     "Pigeon",
		Descreption: "likes to eat seed",
	}

	canary := &Bird{
		Species:    "Canary",
		Dimensions: "test if dimesion will be ignored",
	}

	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	canaryData, _ := json.Marshal(canary)
	//to print data , we can typecast it t s a string
	fmt.Println(string(data))

	// OMIT & IGNORE : we use omitempty to ignore emoty fields
	fmt.Println(string(canaryData))

	// SLICES : let's try to create a json array y using slices
	arrayData, _ := json.Marshal([]*Bird{pigeon, canary})
	fmt.Println(string(arrayData))
}

func MarshalingWithUnStructuredData() {
	// The keys need to be strings, the values can be
	// any serializable value

	birdData := map[string]interface{}{
		"birdsSounds": map[string]interface{}{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	//JSON encoding is done the same as before
	data, _ := json.Marshal(birdData)
	fmt.Println(string(data))
}
