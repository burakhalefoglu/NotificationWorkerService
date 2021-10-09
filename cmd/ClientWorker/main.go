package main

import (
	websocketAdapter "NotificationWorkerService/internal/websocket"
	fastHttpServer "NotificationWorkerService/internal/websocket/fasthttp"
	"runtime"
)

func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0 
	
	// ClientListenerWorker.KafkaStartListening()

	conn := fastHttpServer.FasthttpConn{
		ConnString : "localhost:8081",
	}
	websocketAdapter.ListenServer(conn)
}

