package main

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/controller"
	IWebSocket "NotificationWorkerService/internal/websocket"
	fastHttpServer "NotificationWorkerService/internal/websocket/fasthttp"
	"runtime"
	"sync"
)

func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0

	var wg sync.WaitGroup

	IWebSocket.ListenServer(fastHttpServer.FastHttpServer, &wg)
	controller.StartListening(IoC.KafkaController, &wg)

	wg.Wait()
}