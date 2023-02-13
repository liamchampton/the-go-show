package main

import (
	"encoding/json"
	"fmt"
)

// Define the struct type to be encoded
type Person struct {
	Name string `json:"name"` // The json tag is used to specify the name of the field in the JSON output but is not required
	Age  int    `json:"age"`
}

func main() {
	encodeJson()
	decodeJson()
	decodeIntoMap()
	encodingSliceOfStructs()
	encodingJsonWithIndent()
	decodeIntoSlice()
	encodeStructWithOmitEmpty()
}

func encodeJson() {
	// Create a new instance of the struct and populate it
	person := Person{
		Name: "John Doe",
		Age:  32,
	}

	// Encode the struct to JSON using the json.Marshal function
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}

func decodeJson() {
	// Create a JSON string
	jsonData := []byte(`{"name":"John Doe","age":32}`)

	// Decode the JSON string to a struct using the json.Unmarshal function
	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Name: %s\nAge: %d\n", person.Name, person.Age)
}

func decodeIntoMap() {
	// Create a JSON string
	jsonData := []byte(`{"name":"John Doe","age":32}`)

	// Decode the JSON string to a map using the json.Unmarshal function
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Name:", data["name"]) // The map keys are strings and the values are of type interface{}
	fmt.Println("Age:", data["age"])
}

func encodingSliceOfStructs() {
	// Create a slice of Person structs
	people := []Person{
		{Name: "John Doe", Age: 32},
		{Name: "Jane Doe", Age: 28},
	}

	// Encode the slice to JSON using the json.Marshal function
	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}

func encodingJsonWithIndent() {
	// Create a new instance of the struct and populate it
	person := Person{
		Name: "John Doe",
		Age:  32,
	}

	// Encode the struct to JSON using the json.MarshalIndent function
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}

func decodeIntoSlice() {
	// Create a JSON string
	jsonData := []byte(`[{"name":"John Doe","age":32},{"name":"Jane Doe","age":28}]`)

	// Decode the JSON string to a slice using the json.Unmarshal function
	var people []Person
	err := json.Unmarshal(jsonData, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, person := range people {
		fmt.Printf("Name: %s\nAge: %d\n", person.Name, person.Age)
	}
}

func encodeStructWithOmitEmpty() {
	// Create a new instance of the struct and populate it
	person := Person{
		Name: "John Doe",
	}

	// Encode the struct to JSON using the json.Marshal function
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
