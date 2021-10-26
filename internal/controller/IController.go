package controller

import "sync"

type IController interface {
	StartKafkaListening(wg *sync.WaitGroup)
}


func StartListening(controller IController, wg *sync.WaitGroup){
	controller.StartKafkaListening(wg)
}