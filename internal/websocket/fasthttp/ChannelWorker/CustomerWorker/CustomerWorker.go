package CustomerWorker

import (
	hub "NotificationWorkerService/internal/websocket/fasthttp"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"log"
)

var clientUpgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StartCustomerListener(ctx *fasthttp.RequestCtx,
	id string,
	projectKey string,
	ch *hub.Channel) {

	err := clientUpgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		defer ws.Close()

		customer := hub.Customer{
			Id: id,
			ProjectId: projectKey,
			Connection: ws,
		}
		ch.AddCustomer(customer)
		log.Println("New Client is connected, total: ", len(ch.Clients))

		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				ch.RemoveCustomer(customer)
				break
			}
			log.Println(message)
		}
	})

	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}



