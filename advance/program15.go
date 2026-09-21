package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StringList []string

// UnmarshalJSON handles both a single string or an array of strings
func (sl *StringList) UnmarshalJSON(data []byte) error {
	// Check if it's a single string
	if len(data) > 0 && data[0] == '"' {
		var single string
		if err := json.Unmarshal(data, &single); err != nil {
			return err
		}
		*sl = StringList{single}
		return nil
	}
	
	// Otherwise, process it as a slice of strings
	var multiple []string
	if err := json.Unmarshal(data, &multiple); err != nil {
		return err
	}
	*sl = multiple
	return nil
}

func main() {
	json1 := `{"items": "apple"}`
	json2 := `{"items": ["apple", "banana"]}`

	var result1, result2 struct {
		Items StringList `json:"items"`
	}

	json.Unmarshal([]byte(json1), &result1)
	json.Unmarshal([]byte(json2), &result2)

	fmt.Printf("Single: %s\n", strings.Join(result1.Items, ", "))
	fmt.Printf("Multiple: %s\n", strings.Join(result2.Items, ", "))
}