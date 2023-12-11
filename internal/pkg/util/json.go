package util

import (
	"encoding/json"
	"fmt"
)

func ToJSON(structure any) ([]byte, error) {
	json, err := json.Marshal(structure)

	if err != nil {
		fmt.Println("Error when marshalling JSON: ", err)
		return nil, err
	}

	fmt.Printf("\nJSON: %s\n", string(json))

	return json, nil
}
