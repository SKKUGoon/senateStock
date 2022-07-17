package main

import (
	"fmt"
	"reflect"
	"senateStock/dataObject"
)

type Test struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Class       string   `json:"class"`
	Type        []string `json:"type"`
	Address     struct {
		City   string `json:"city"`
		Street string `json:"street"`
	} `json:"address"`
	Students  int      `json:"students"`
	Location  string   `json:"location"`
	StatusLog []string `json:"status_log"`
}

func main() {
	var tgt []Test
	conn, _ := dataObject.Conn("./token.json")
	result, _ := dataObject.JSONGet[[]Test](conn, "school_json:4", "$", &tgt)
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result[0].Students))
}
