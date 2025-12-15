package main

import (
	"encoding/json"
	"fmt"
)

func fixedType() {
	jsonString := `{"name": "John Doe", "age": 30, "city": "New York", "isGeek": true}`

	var person Person

	if err := json.Unmarshal([]byte(jsonString), &person); err != nil {
		fmt.Println("Error parsing JSON into struct:", err)
		return
	}

	fmt.Printf("%+v\n\n", person)
}
func main() {

	fixedType()

	dynamicType()

	serialize()
}

func serialize() {
	person := Person{
		Name:   "Victor Corleone",
		Age:    30,
		City:   "New York",
		IsGeek: false,
	}

	str, _ := json.Marshal(person)
	fmt.Println(string(str))

}

func dynamicType() {

	dynamicJsonString := `{"username": "jane.doe", "metrics": {"logins": 105, "last_login_at": "2025-12-14T15:04:05Z"}, "active": true}`

	var data map[string]interface{}

	if err := json.Unmarshal([]byte(dynamicJsonString), &data); err != nil {
		fmt.Println("Error parsing JSON into map:", err)
		return
	}
	fmt.Println(data["username"].(string))
	fmt.Println(data["active"].(bool))
	fmt.Println(data["metrics"].(map[string]interface{})["logins"].(float64))
}
