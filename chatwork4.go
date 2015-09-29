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
	fmt.Print("api call start")
	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	req, err := http.NewRequest("GET", END_POINT_URL+"/my/status ", buf)
	if err != nil {
		fmt.Errorf("occur error")
	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))
	hclient := &http.Client{Timeout: time.Duration(10 * time.Second)}
	hclient.Do(req)
	fmt.Print("api call end")
}
