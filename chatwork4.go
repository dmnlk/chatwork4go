package chatwork4go

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/k0kubun/pp"
	"io/ioutil"
	"encoding/json"
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
	req, err := http.NewRequest("GET", END_POINT_URL+"/my/status ", buf)
	if err != nil {
		fmt.Errorf("occur error")
	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))
	hclient := &http.Client{Timeout: time.Duration(10 * time.Second)}
	resp, err :=hclient.Do(req)
	if err != nil {
		fmt.Errorf("err")
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()



	val, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("err")
	}

	var res Status
	err = json.Unmarshal(val, &res)
	if err != nil {
		fmt.Errorf("json err")
	}

	pp.Println(res)
}
