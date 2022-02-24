package utils

import (
	"encoding/json"
	"fmt"
)

func ResolveJson(jsonString string) error {
	jsonData := []byte(jsonString)
	var value interface{}
	err := json.Unmarshal(jsonData, &value)
	if err != nil {
		return err
	}
	data := value.(map[string]interface{})
	for k, v := range data {
		switch v := v.(type) {
			case string:
				fmt.Println(k, v, "(string)")
			case float64:
				//fmt.Println(k, v, "(float64)")
			case []interface{}:
				fmt.Println(k, "(array):")
				for i, u := range v {
					fmt.Println("    ", i, u)
				}
			default:
				fmt.Println(k, "(unknown)")
			}
	}
	return nil
}
