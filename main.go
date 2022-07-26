package main

import (
	"context"
	"fmt"
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

func reJson() {
	// TODO: Create it as a test
	var tgt []Test
	conn, _ := dataObject.Conn("./token.json")

	var s = dataObject.RedisConn[[]Test]{
		context.Background(),
		conn,
		&tgt,
	}
	_ = s.JsonGet("school_json:3", "$")
	fmt.Println(s.Container)
}

func main() {
	reJson()
}
