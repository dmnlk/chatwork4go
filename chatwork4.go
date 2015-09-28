package chatwork4go

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


func (client *Client) PostToRoom(roomId string, postMessage string) {

}
