package websocket

import "sync"

type IWebsocket interface {
	ListenServer(group *sync.WaitGroup)

	SendMessageToClient(message *[]byte, clientId string,
		projectId string,
		channelName string) error

	SendMessageToAllClient(message *[]byte,
		projectId string,
		channelName string) error

	SendMessageToCustomer(message *[]byte, customerId string,
		projectId string,
		channelName string) error

	SendMessageToAllCustomer(message *[]byte,
		projectId string,
		channelName string) error
}

func ListenServer(d *IWebsocket, group *sync.WaitGroup) {
	(*d).ListenServer(group)
}