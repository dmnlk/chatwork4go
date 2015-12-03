package chatwork4go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	END_POINT_URL string = "https://api.chatwork.com/v1"
)

type Client struct {
	apikey APIKEY
	client *http.Client
}

type APIKEY string

func NewClient(key string) (*Client, error) {
	client := new(Client)
	client.apikey = APIKEY(key)
	//client.client = http.DefaultClient
	client.client = &http.Client{Timeout: time.Duration(10 * time.Second)}
	// check apikey
	_, err := client.GetMyStatus()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *Client) GetMyStatus() (*Status, error) {
	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	req, err := http.NewRequest("GET", END_POINT_URL+"/my/status", buf)
	if err != nil {
		fmt.Errorf("occur error")
		return nil, err
	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))
	//hclient := &http.Client{Timeout: time.Duration(10 * time.Second)}
	resp, err := client.client.Do(req)
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

	var res *Status
	err = json.Unmarshal(val, &res)
	if err != nil {
		fmt.Errorf("json err")
		return nil, err
	}
	return res, nil
}

func (client *Client) GetMyTasks() (*Task, error) {
	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	req, err := http.NewRequest("GET", END_POINT_URL+"/my/tasks", buf)
	if err != nil {
		fmt.Errorf("occur error")
		return nil, err
	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))
	//hclient := &http.Client{Timeout: time.Duration(10 * time.Second)}
	resp, err := client.client.Do(req)
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

	var res *Task
	err = json.Unmarshal(val, &res)
	if err != nil {
		fmt.Errorf("error")
		return nil, err
	}

	return res, nil
}

func (client *Client) PostMesseage(roomId int, message string) {

}
