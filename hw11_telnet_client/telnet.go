package main

import (
	"io"
	"net"
	"time"
)

// TelnetClient сигнатура заменена (io.Close) по причине линтера.
type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type Client struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *Client) Connect() (err error) {
	c.conn, err = net.DialTimeout("tcp", c.address, c.timeout)
	return err
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Send() error {
	_, err := io.Copy(c.conn, c.in)
	return err
}

func (c *Client) Receive() error {
	_, err := io.Copy(c.out, c.conn)
	return err
}
