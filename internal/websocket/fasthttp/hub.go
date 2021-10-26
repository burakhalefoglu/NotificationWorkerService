package fastHttpServer

import (
	"fmt"
	"log"

	"github.com/fasthttp/websocket"
)

type Channel struct {
	Clients []Client
	Customers []Customer
	Name string
}

type Client struct {
	Id         string
	ProjectId string
	Connection *websocket.Conn

}

type Customer struct {
	Id         string
	ProjectId string
	Connection *websocket.Conn
}


func (ps *Channel) AddClient(client Client) {

	ps.Clients = append(ps.Clients, client)
	fmt.Println("adding new client to the list", client.Id, len(ps.Clients))
	payload := []byte("Hello Client ID:" +
		client.Id)

	client.Connection.WriteMessage(1, payload)
}

 
func (ps *Channel) RemoveClient(client Client) {

	for index, c := range ps.Clients {

		if c.Id == client.Id {
			ps.Clients = append(ps.Clients[:index], ps.Clients[index+1:]...)
			log.Println(c.Id ,"is removed")
		}

	}
}

func (ps *Channel) AddCustomer(customer Customer) {

	ps.Customers = append(ps.Customers, customer)
	payload := []byte("Hello Client ID:" +
		customer.Id)
	customer.Connection.WriteMessage(1, payload)
}


func (ps *Channel) RemoveCustomer(customer Customer) {

	for index, c := range ps.Customers {

		if c.Id == customer.Id {
			ps.Customers = append(ps.Customers[:index], ps.Customers[index+1:]...)
			log.Println(c.Id ,"is removed")
		}

	}
}