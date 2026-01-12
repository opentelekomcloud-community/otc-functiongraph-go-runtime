package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	instance_id := "your_instance_id"
	var jsonStr map[string]interface{}

	jsonStr = map[string]interface{}{
		"os-start": map[string]interface{}{
			"servers": []map[string]string{
				{"id": instance_id},
			},
		},
	}

	jsonData, err := json.Marshal(jsonStr)
	if err != nil {
		fmt.Errorf("JSON marshal failed: %v", err)
		return
	}

	fmt.Printf("Request JSON: %s\n", jsonStr)
	fmt.Printf("Request JSON: %s\n", string(jsonData))

}
