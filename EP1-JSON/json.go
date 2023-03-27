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
	encodeJson() // converting a data structure to JSON
	decodeJson() // converting JSON to a data structure
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

	fmt.Println("Encode JSON:")
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

	fmt.Println("\nDecode JSON:")
	fmt.Printf("Name: %s\nAge: %d\n", person.Name, person.Age)
}

func decodeIntoMap() {
	// Create a JSON string
	jsonData := []byte(`{"name":"John Doe","age":32}`)

	// Decode the JSON string to a map using the json.Unmarshal function
	// The interface{} type is used to represent values of any type. This is useful when you don't know the type of the JSON data
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data) // & is used to pass a pointer to the data variable to change directly
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nDecode JSON into Map:")
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

	fmt.Println("\nEncoding slice of structs:")
	fmt.Println(string(jsonData))
}

func encodingJsonWithIndent() {
	// Create a new instance of the struct and populate it
	person := Person{
		Name: "John Doe",
		Age:  32,
	}

	// Encode the struct to JSON using the json.MarshalIndent function
	// The first argument is the value to encode
	// The second argument is a prefix to prepend to each line of output
	// The third argument is the indentation string to use
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nEncoding JSON with indent:")
	fmt.Println(string(jsonData))
}

func decodeIntoSlice() {
	// Create a JSON string
	jsonData := []byte(`[{"name":"John Doe","age":32},{"name":"Jane Doe","age":28}]`)

	// Decode the JSON string to a slice using the json.Unmarshal function
	var people []Person                      // The slice is initialized to an empty slice of Person structs
	err := json.Unmarshal(jsonData, &people) // & is used to pass a pointer to the people variable to change directly
	if err != nil {
		fmt.Println(err)
		return
	}

	// Loop through the slice and print the values
	fmt.Println("\nDecoding JSON into slice:")
	for _, person := range people {
		fmt.Printf("Name: %s\nAge: %d\n", person.Name, person.Age)
	}
}

func encodeStructWithOmitEmpty() {
	// Create a new instance of the struct and populate it with only 1 field
	person := Person{
		Name: "John Doe",
	}

	// Encode the struct to JSON using the json.Marshal function
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	// The Age field is not included in the JSON output because it is empty
	fmt.Println("\nEncode struct with OmitEmpty:")
	fmt.Println(string(jsonData))
}
