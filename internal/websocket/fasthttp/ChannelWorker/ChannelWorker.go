package ChannelWorker

import (
	"log"

	hub "NotificationWorkerService/pkg/FastHttp"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
} 

func Work(ctx *fasthttp.RequestCtx,
	 id string,
	 projectKey string,
	 ch *hub.Channel) {

	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		defer ws.Close()
		
		client := hub.Client{
			Id: id,
			ProjectId: projectKey,
			Connection: ws,
		}
		ch.AddClient(client)
		log.Println("New Client is connected, total: ", len(ch.Clients))
		
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				ch.RemoveClient(client)
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



