package chatwork4go

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
		return nil, err
	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()

	val, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res *Status
	err = json.Unmarshal(val, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (client *Client) GetMyTasks() (*Task, error) {
	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	req, err := http.NewRequest("GET", END_POINT_URL+"/my/tasks", buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-ChatWorkToken", string(client.apikey))

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()

	val, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res *Task
	err = json.Unmarshal(val, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *Client) PostMesseages(roomId int, message string) error {
	data := url.Values{"body": {message}}
	req, err := http.NewRequest("POST", END_POINT_URL+"/rooms/"+strconv.Itoa(roomId)+"/messages", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("X-ChatWorkToken", string(client.apikey))

	resp, err := client.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}
	defer resp.Body.Close()

	return nil
}
