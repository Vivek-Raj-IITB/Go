package main

import (
	"encoding/json"
	"fmt"
)

type ErrorResponse struct {
}
type course struct {
	Name  string        `json:"coursename"`
	Error ErrorResponse `json:"error"`
}

func main() {
	fmt.Println("Welcome to json")
	// EncodeJson()
	DecodeJson()

}

// func EncodeJson() {
// 	lcoCourse := []course{
// 		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"},
// 			ErrorResponse{"code": 404,
// 				"message": "Course not found"}}}

// 	//package this data as JSON data

// 	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", finalJson)
// }

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"],
		"error": {
			"code": 404,
			"message": "Course not found"
		}
	}
	`)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	fmt.Println("-----------------------------------------")
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
	fmt.Println("-----------------------------------------")
	fmt.Printf("%v", myOnlineData["error"].(map[string]interface{})["message"])
}
