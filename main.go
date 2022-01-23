package main

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/IoC/golobby"
	"NotificationWorkerService/internal/controller"
	"NotificationWorkerService/internal/websocket"
	"NotificationWorkerService/pkg/helper"
	"github.com/joho/godotenv"
	"log"
	"runtime"
	"sync"
)

func main() {
	defer helper.DeleteHealthFile()
	runtime.MemProfileRate = 0
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	IoC.InjectContainers(golobby.InjectionConstructor())

	var wg sync.WaitGroup
	controller.StartListening(&IoC.Controller, &wg)
	websocket.ListenServer(&IoC.WebSocket, &wg)
	wg.Wait()
}
