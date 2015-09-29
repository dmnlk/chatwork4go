package chatwork4go

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	END_POINT_URL string = "https://api.chatwork.com/v1"
)

type Client struct {
	apikey APIKEY
}

type APIKEY string

func NewClient(key string) *Client {
	client := new(Client)
	client.apikey = APIKEY(key)
	return client
}

func (client *Client) GetMyStatus() {
	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	fmt.Print("a")
	req, err := http.NewRequest("GET", "", buf)
	if err != nil {

	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))
	hclient := &http.Client{Timeout: time.Duration(10 * time.Second)}
	hclient.Do(req)
}
