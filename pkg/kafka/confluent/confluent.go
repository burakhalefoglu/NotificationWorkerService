package confluent

import (
	"NotificationWorkerService/pkg/helper"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"sync"
)

type confluentKafka struct {
}

func ConfluentKafkaConstructor() *confluentKafka {
	return &confluentKafka{}
}

func (k *confluentKafka) Produce(key *[]byte, value *[]byte, topic string) (err error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": helper.ResolvePath("KAFKA_HOST", "KAFKA_PORT")})
	if err != nil {
		log.Fatal("ConfluentKafka", "Produce Connection Failed: ", err.Error())
	}
	pErr := p.Produce(&kafka.Message{
		Key:            *key,
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          *value,
	}, nil)
	if pErr != nil {
		return pErr
	}
	log.Print("ConfluentKafka", "Producer", topic, key)
	p.Flush(15 * 1000)
	return nil
}

func (k *confluentKafka) Consume(topic string, groupId string, waitGroup *sync.WaitGroup, callback func(data *[]byte) (bool, string)) {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": helper.ResolvePath("KAFKA_HOST", "KAFKA_PORT"),
		"group.id":          groupId,
		"auto.offset.reset": "smallest"})
	if err != nil {
		//k.Log.SendPanicLog("ConfluentKafka", "Consumer Connection Failed: ", err.Error())
		panic(err)
	}

	var run = true
	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			isSuccess, _ := callback(&e.Value)
			log.Print("ConfluentKafka", "Consumer", topic, groupId)
			if isSuccess {
				go func() {
					offsets, err := consumer.Commit()
					if err != nil {
						log.Fatal("ConfluentKafka",
							"Consumer", "%% Commit failed %v\n", offsets, err.Error())
					}
				}()
			}

		case kafka.PartitionEOF:
			log.Fatal("ConfluentKafka",
				"Consumer", "%% PartitionEOF %v\n", e, err.Error())
		case kafka.Error:
			log.Fatal("ConfluentKafka",
				"Consumer", "%% Kafka Error: %v\n", e, err.Error())
			run = false
		default:
		}
	}
	waitGroup.Done()
}
