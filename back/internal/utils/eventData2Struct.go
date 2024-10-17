package utils

import (
	"encoding/json"
	"fmt"
)

func EventDataToStruct[T any](data interface{}) *T {
	var result T
	mapData, ok := data.(map[string]interface{})
	if ok {
		jsonData, err := json.Marshal(mapData)
		if err != nil {
			fmt.Println("error marshalling map to json:", err)
			return nil
		}

		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			fmt.Println("error unmarshalling json to struct:", err)
			return nil
		}

		return &result
	} else {
		return data.(*T)
	}
}
