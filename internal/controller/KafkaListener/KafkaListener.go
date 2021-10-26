package KafkaListener

import (
	"NotificationWorkerService/internal/manager/abstract"
	IKafka "NotificationWorkerService/pkg/kafka"
	"sync"
)

type KafkaListener struct {
	Kafka IKafka.IKafka
	RemoteOfferM    abstract.IRemoteOfferService
	InterstitialAdM abstract.IInterstitialAdService
	ChurnBlockerResultM abstract.IChurnBlockerService
	ChurnPredictionResultM abstract.IChurnPredictionService
	ProjectCreationResultM abstract.IProjectCreationService
}

func (k *KafkaListener)StartKafkaListening(wg *sync.WaitGroup) {

	wg.Add(5)

	go k.Kafka.Consume("RemoteOfferModel",
		"RemoteOfferModel_ConsumerGroup",
		wg,
		k.RemoteOfferM.SendMessageToClient)

	go k.Kafka.Consume("InterstitialAdModel",
		"InterstitialAdModel_ConsumerGroup",
		wg,
		k.InterstitialAdM.SendMessageToClient)

	go k.Kafka.Consume("ChurnBlockerResultModel",
		"ChurnBlockerResult_ConsumerGroup",
		wg,
		k.ChurnBlockerResultM.SendMessageToClient)

	go k.Kafka.Consume("ChurnPredictionResultModel",
		"ChurnPredictionResult_ConsumerGroup",
		wg,
		k.ChurnPredictionResultM.SendMessageToClient)

	go k.Kafka.Consume("ProjectCreationResultModel",
		"ProjectCreationResultModel_ConsumerGroup",
		wg,
		k.ProjectCreationResultM.SendMessageToCustomer)
}