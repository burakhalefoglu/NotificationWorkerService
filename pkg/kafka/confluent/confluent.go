package confluent

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
	"sync"
)

type Kafka struct {
}

func (k *Kafka) Produce(key *[]byte, value *[]byte, topic string) (err error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.1.1:9092"})
	if err != nil {
		panic(err)
	}
		p.Produce(&kafka.Message{
			Key: *key,
			TopicPartition: kafka.TopicPartition{Topic: &topic},
			Value:          *value,
		}, nil)
	p.Flush(15 * 1000)
	return nil
}

func (k *Kafka) Consume(topic string, groupId string, waitGroup *sync.WaitGroup, callback func( data *[]byte) (success bool, message string)) {

	defer waitGroup.Done()
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    "192.168.1.1:9092",
		"group.id":             groupId,
		"auto.offset.reset":    "smallest",
		topic: 					topic,
	})
	if err != nil{
		panic(err)
	}
	consumer.SubscribeTopics([]string{topic}, nil)

	var run = true
	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			success, _ := callback(&e.Value)
			if success == true{
				go func() {
					offsets, err := consumer.Commit()
					if err != nil{
						log.Printf("%% Reached %v\n", err)
					}
					log.Printf("%% Reached %v\n", offsets)
				}()
			}

		case kafka.PartitionEOF:
			log.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			log.Printf("Ignored %v\n", e)
		}
	}
}

