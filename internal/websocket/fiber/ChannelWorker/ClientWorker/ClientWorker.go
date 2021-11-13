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
		log.Println("New Client is connected, total: ", len(ch.Clients))


		var (
			_   int
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				ch.RemoveClient(client)
				break
			}
			log.Printf("recv: %s", msg)
		}
		wgGroup.Done()
	}))
}



