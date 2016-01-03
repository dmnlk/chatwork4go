package main

import (
	"os"

	"github.com/dmnlk/chatwork4go"
	"github.com/k0kubun/pp"
)

func main() {
	var apikey = os.Args[1]
	_, err := chatwork4go.NewClient(apikey)
	if err != nil {
		pp.Println(err)
	}
	// post sample
//	err = client.PostMesseages(41380405, "aaa")
//	if err != nil {
//		pp.Println(err)
//	}
}
