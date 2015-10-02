package main

import (
	"os"

	"github.com/dmnlk/chatwork4go"
)

func main() {
	var apikey = os.Args[1]

	client := chatwork4go.NewClient(apikey)
	client.GetMyStatus()
}
