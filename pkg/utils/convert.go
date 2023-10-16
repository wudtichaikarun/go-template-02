package utils

import (
	"encoding/json"
	"fmt"
)

func ConvertInterfaceToStruct(data interface{}, obj interface{}) {
	// Convert map to json string
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Print(err)
	}

	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, obj); err != nil {
		fmt.Print(err)
	}
}
