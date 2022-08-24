// Package unmarshal demonstrates unmarshalling json in go
//
// Json is mapped to a struct as detailed in encoding/json:
//	type people []struct {
// 		First 	string `json:"First"`
// 		Last  	string `json:"Last"`
// 		Age   	int    `json:"Age,omitempty"`
// 		Option  string `json:"-"`
// 	}
// Optional fields are denoted by:
//	`json:"Age,omitempty"
// Ignored fields with:
//	`json:"-"`
// Keys are handled using:
//	`json:"-,"`
// JSON field FirstName could be renamed to First for application using:
// 	First 	string `json:"FirstName"`
package unmarshal

import (
	"encoding/json"
	"fmt"
)

type people []struct {
	First  string `json:"First"`
	Last   string `json:"Last"`
	Age    int    `json:"Age,omitempty"`
	Option string `json:"-"`
}

// Unmarshal takes the slice of byte data and unmarshals to the people []struct
//
// Errors print to console and are returned with people []struct
func Unmarshal(data []byte, ppl people) (people, error) {
	err := json.Unmarshal(data, &ppl)
	if err != nil {
		fmt.Printf("The was an error: %v", err)
	}
	return ppl, err
}
