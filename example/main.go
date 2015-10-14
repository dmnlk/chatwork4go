package main

import (
	"os"

	"github.com/dmnlk/chatwork4go"
	"github.com/k0kubun/pp"
)

func main() {
	var apikey = os.Args[1]
	client := chatwork4go.NewClient(apikey)
	pp.Println(client.GetMyStatus())
	client.GetMyTasks()
}
