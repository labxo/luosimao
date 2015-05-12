package luosimao

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	Error int
	Msg   string
}

func (err Error) Check() error {
	if err.Error == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("%d-%s", err.Error, err.Msg))
	}
}

type BasicAuthTransport struct {
	Username string
	Password string
}

func (bat BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(bat.Username, bat.Password)
	return http.DefaultTransport.RoundTrip(req)
}

type Client struct {
	*http.Client
}

func NewClient(key string) *Client {
	client := &Client{}
	client.Client = &http.Client{Transport: &BasicAuthTransport{
		Username: "api",
		Password: "key-" + key,
	}}
	return client
}
