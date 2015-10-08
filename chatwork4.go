package chatwork4go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/k0kubun/pp"
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

func (client *Client) GetMyStatus() (*Status, error) {
	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	req, err := http.NewRequest("GET", END_POINT_URL+"/my/status ", buf)
	if err != nil {
		fmt.Errorf("occur error")
		return nil, err
	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))
	hclient := &http.Client{Timeout: time.Duration(10 * time.Second)}
	resp, err := hclient.Do(req)
	if err != nil {
		fmt.Errorf("err")
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf(resp.Status)
		return nil, err
	}
	defer resp.Body.Close()

	val, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("err")
		return nil, err
	}

	var res Status
	err = json.Unmarshal(val, &res)
	if err != nil {
		fmt.Errorf("json err")
		return nil, err
	}

	pp.Println(res)
	return *res, nil
}
