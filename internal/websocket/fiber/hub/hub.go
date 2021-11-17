package hub

import (
	"github.com/gofiber/websocket/v2"
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
}

 
func (ps *Channel) RemoveClient(client Client) {

	for index, c := range ps.Clients {

		if c.Id == client.Id {
			ps.Clients = append(ps.Clients[:index], ps.Clients[index+1:]...)
		}

	}
}

func (ps *Channel) AddCustomer(customer Customer) {

	ps.Customers = append(ps.Customers, customer)
}


func (ps *Channel) RemoveCustomer(customer Customer) {

	for index, c := range ps.Customers {

		if c.Id == customer.Id {
			ps.Customers = append(ps.Customers[:index], ps.Customers[index+1:]...)
		}

	}
}