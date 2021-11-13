package main

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/controller"
	websocket "NotificationWorkerService/internal/websocket"
	"NotificationWorkerService/internal/websocket/fiber"
	"github.com/joho/godotenv"
	"log"
	"runtime"
	"sync"
)

func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	var wg sync.WaitGroup

	controller.StartListening(IoC.KafkaController, &wg)
	websocket.ListenServer(fiber.FiberServer, &wg)

	wg.Wait()
}