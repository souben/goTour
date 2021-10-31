package unmarshaling

import (
	"encoding/json"
	"fmt"
)

/* Unmarshal raw json data
Unmashal func: parse raw JSON data in the form of []byte variables
*/

// Structured Data (Decoding JSON Into Structs)

type Bird struct {
	Species     string `json:"birdType"` // explicitly tell our code which JSON property to map to which attribute
	Description string `json:"what it does"`
	Dimensions  Dimensions
}

type Dimensions struct {
	Height int
	Width  int
}

func UnmarshalWithStructuredData() {

	// single item
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)

	//array of items
	birdArrayJson := `[{"species": "pigeon","description": "likes to perch on rocks"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdArrayJson), &birds)
	fmt.Printf("Birds: %+v\n", birds)

	//primitive types
	numberJson := "3"
	floatJson := "3.1412"
	stringJson := `"bird"`

	var n int
	var pi float64
	var str string

	json.Unmarshal([]byte(numberJson), &n)
	fmt.Println(n)
	// 3

	json.Unmarshal([]byte(floatJson), &pi)
	fmt.Println(pi)
	// 3.1412

	json.Unmarshal([]byte(stringJson), &str)
	fmt.Println(str)
	// bird
}

// Decoding JSON to Maps - Unstructured Data
// we can use maps whr the json body have no structure

func UnmarshalWithUnStructuredData() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	// otherwise it going to be treated as interface{}
	birds := result["birds"].(map[string]interface{})

	for key, value := range birds {
		fmt.Println(key, value.(string))
	}
}
