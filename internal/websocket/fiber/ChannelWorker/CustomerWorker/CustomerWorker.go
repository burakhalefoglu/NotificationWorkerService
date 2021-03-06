package CustomerWorker

import (
	"NotificationWorkerService/internal/websocket/fiber/hub"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
	"sync"
)



func StartCustomerListener(wgGroup *sync.WaitGroup,
	app *fiber.App,
ch *hub.Channel) {

	app.Get("/"+ch.Name, websocket.New(func(c *websocket.Conn) {
		var clientId = c.Query("clientId")
		var projectId = c.Query("projectId")

		customer := hub.Customer{
			Id:         clientId,
			ProjectId:  projectId,
			Connection: c,
		}
		ch.AddCustomer(customer)
		log.Print("CustomerWorker", "StartCustomerListener",
			"New Customer is connected, total: ",customer.Id ,len(ch.Clients))
		var (
			_   int //message type
			_ []byte // message
			err error
		)
		for {
			if _, _, err = c.ReadMessage(); err != nil {
				ch.RemoveCustomer(customer)
				log.Print("CustomerWorker", "StartCustomerListener",
					"Customer is disconnected: " ,err)
				break
			}
		}
		wgGroup.Done()
	}))

}