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
		log.Println("New Client is connected, total: ", len(ch.Clients))

		var (
			_   int
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				ch.RemoveCustomer(customer)
				break
			}
			log.Printf("recv: %s", msg)
		}
		wgGroup.Done()
	}))

}