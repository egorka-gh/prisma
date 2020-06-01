package prisma

import (
	"fmt"
	"net"
	"time"
)

//Client udp client
type Client struct {
	pr      *Prisma
	host    string
	port    string
	udpaddr *net.UDPAddr
}

//NewClient create default server
func NewClient(host string) *Client {
	return &Client{
		pr:   DefaultPrisma(),
		host: host,
		port: "21845",
	}
}

//Send message to udp server
func (c *Client) Send(m *Message) error {
	data, err := c.pr.Ecode(m)
	if err != nil {
		return err
	}
	err = c.resolveUDP()
	if err != nil {
		return err
	}
	cnn, err := net.DialUDP("udp", nil, c.udpaddr)
	if err != nil {
		c.udpaddr = nil
		return err
	}
	defer cnn.Close()
	_, err = cnn.Write(data)
	fmt.Println(string(data))
	return err
}

//SendBatch messageы to udp server
func (c *Client) SendBatch(messages []*Message) error {
	err := c.resolveUDP()
	if err != nil {
		return err
	}
	cnn, err := net.DialUDP("udp", nil, c.udpaddr)
	if err != nil {
		c.udpaddr = nil
		return err
	}
	defer cnn.Close()
	var lastErr error
	for _, m := range messages {
		data, err := c.pr.Ecode(m)
		if err != nil {
			lastErr = err
			fmt.Println(err)
			continue
		}
		_, err = cnn.Write(data)
		if err != nil {
			lastErr = err
			fmt.Println(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	return lastErr
}

func (c *Client) resolveUDP() error {
	if c.udpaddr != nil {
		return nil
	}
	s, err := net.ResolveUDPAddr("udp", c.host+":"+c.port)
	if err != nil {
		return err
	}
	c.udpaddr = s
	return nil
}
