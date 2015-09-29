package main
import (
	"github.com/dmnlk/chatwork4go"
)

func main() {
	client := chatwork4go.NewClient("a");
	client.GetMyStatus()
}
