package adapter

import "fmt"

type (
	Client struct {
	}
	computer interface {
		insertIntoLightningPort()
	}
)

func (c *Client) InsertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into bridge.")
	com.insertIntoLightningPort()
}
