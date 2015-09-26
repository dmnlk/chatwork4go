package chatwork4go

type Client struct {
	apikey APIKEY
}

type APIKEY string

func NewClient() *Client {
	client := new (Client);
	return client;
}
