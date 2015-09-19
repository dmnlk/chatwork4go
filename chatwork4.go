package chatwork4go

import (
	"net/http"
	"bytes"
	"time"
	"github.com/codegangsta/cli"
)

type Client struct {
	apikey APIKEY
}

type APIKEY string

const (
	END_POINT_URL string = "https://api.chatwork.com/v1"
)

func NewClient(key string) *Client {
	client := new (Client);
	client.apikey = key;
	return client;
}

func (client *Client) getMyStatus() {
	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	req, err := http.NewRequest("GET", "", buf)
	if err != nil {

	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))
	hclient := &http.Client{Timeout: time.Duration(10 * time.Second)}
	hclient.Do(req)
}
