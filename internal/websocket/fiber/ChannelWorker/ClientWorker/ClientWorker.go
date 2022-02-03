package ClientWorker

import (
	"NotificationWorkerService/internal/websocket/fiber/hub"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
	"sync"
)

func StartClientListener(wgGroup *sync.WaitGroup,
	 app *fiber.App,
	 ch *hub.Channel) {

	app.Get("/"+ch.Name, websocket.New(func(c *websocket.Conn) {
		var clientId = c.Query("clientId")
		var projectId = c.Query("projectId")

		client := hub.Client{
			Id: clientId,
			ProjectId: projectId,
			Connection: c,
		}
		ch.AddClient(client)
		log.Print("ClientWorker", "StartClientListener",
			"New Client is connected, total: ",client.Id ,len(ch.Clients))


		var (
			_   int //message type
			_ []byte // message
			err error
		)
		for {
			if _, _, err = c.ReadMessage(); err != nil {
				ch.RemoveClient(client)
				log.Print("CustomerWorker", "StartCustomerListener",
					"Customer is disconnected: " ,err)
				break
			}
		}
		wgGroup.Done()
	}))
}



