package pkg

import (
	"net/rpc"
)

type Client struct {
	c *rpc.Client
}

func DialServer() (*Client, error) {
	if c, err := rpc.Dial("tcp", Addr); err != nil {
		return nil, err
	} else {
		return &Client{c: c}, nil
	}
}

func (c *Client) Request(rsp *String) error {
	if err := c.c.Call(ServerName+".Server", String{Value: "hello"}, rsp); err != nil {
		return err
	}

	return nil
}
