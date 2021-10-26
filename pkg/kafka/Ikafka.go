package Ikafka

import (
	"sync"
)

type IKafka interface {
	Produce(key *[]byte, value *[]byte, topic string) (err error)
	Consume(topic string, groupId string, waitGroup *sync.WaitGroup, callback func(data *[]byte)(success bool, message string))
}
