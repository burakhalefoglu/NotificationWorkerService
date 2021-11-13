package main

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/controller"
	websocket "NotificationWorkerService/internal/websocket"
	fiber "NotificationWorkerService/internal/websocket/fiber"
	"runtime"
	"sync"
)

func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0

	var wg sync.WaitGroup

	controller.StartListening(IoC.KafkaController, &wg)
	websocket.ListenServer(fiber.FiberServer, &wg)

	wg.Wait()
}