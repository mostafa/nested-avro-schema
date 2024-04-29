package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/linkedin/goavro/v2"
)

func main() {
	// Check if the schema file is provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <schema-file> <data-file>")
		os.Exit(1)
	}

	// Load the schema from a file
	var (
		schema []byte
		data   []byte
		err    error
	)
	if schema, err = os.ReadFile(os.Args[1]); err != nil {
		panic(err)
	}

	// Create a new codec for the Avro schema
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		panic(err)
	}

	// Load the data from a file
	if data, err = os.ReadFile(os.Args[2]); err != nil {
		panic(err)
	}

	// Unmarshal the JSON payload
	var native map[string]interface{}
	err = json.Unmarshal(data, &native)
	if err != nil {
		panic(err)
	}

	// Encode the data using the Avro codec and schema
	binary, err := codec.BinaryFromNative(nil, native)
	if err != nil {
		panic(err)
	}

	// Decode the data back to JSON (map[string]interfac{})
	decoded, _, err := codec.NativeFromBinary(binary)
	if err != nil {
		panic(err)
	}

	// Print out the original and decoded data for comparison
	fmt.Println("Original:\n", native)
	fmt.Println("Decoded:\n", decoded)

	// Check if the original and decoded data are equal
	fmt.Println("Original and Decoded are equal: ", fmt.Sprint(native) == fmt.Sprint(decoded))
}
