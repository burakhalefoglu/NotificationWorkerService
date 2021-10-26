package IwebSocket

import "sync"

type IWebsocket interface {
	ListenServer(group *sync.WaitGroup)

	SendMessageToClient(message *[]byte, clientId string,
		projectId string,
		channelName string)

	SendMessageToAllClient(message *[]byte,
		projectId string,
		channelName string)

	SendMessageToCustomer(message *[]byte, customerId string,
		projectId string,
		channelName string)

	SendMessageToAllCustomer(message *[]byte,
		projectId string,
		channelName string)
}

func ListenServer(d IWebsocket, group *sync.WaitGroup) {
	d.ListenServer(group)
}