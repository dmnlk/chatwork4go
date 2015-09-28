package chatwork4go

type Client struct {
	apikey APIKEY
}

type APIKEY string

func NewClient(key string) *Client {
	client := new (Client);
	client.apikey = key;
	return client;
}


func (client *Client) PostToRoom(roomId string, postMessage string) {

}
