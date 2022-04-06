package main

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/IoC/golobby"
	"NotificationWorkerService/internal/controller"
	"NotificationWorkerService/internal/websocket"
	"NotificationWorkerService/pkg/helper"
	"log"
	"runtime"
	"sync"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/joho/godotenv"
)

func main() {
	defer helper.DeleteHealthFile()
	logger.Log.App = "NotificationWorkerService"
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
