package KafkaListener

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/manager/abstract"
	"NotificationWorkerService/pkg/helper"
	IKafka "NotificationWorkerService/pkg/kafka"
	"sync"
)

type kafkaListener struct {
	Kafka                  *IKafka.IKafka
	RemoteOfferM           *abstract.IRemoteOfferService
	InterstitialAdM        *abstract.IInterstitialAdService
	ChurnBlockerResultM    *abstract.IChurnBlockerService
	ChurnPredictionResultM *abstract.IChurnPredictionService
	ProjectCreationResultM *abstract.IProjectCreationService
}

func KafkaListenerConstructor() *kafkaListener {
	return &kafkaListener{Kafka: &IoC.Kafka,
		RemoteOfferM:           &IoC.RemoteOfferM,
		InterstitialAdM:        &IoC.InterstitialAdM,
		ChurnBlockerResultM:    &IoC.ChurnBlockerResultM,
		ChurnPredictionResultM: &IoC.ChurnPredictionResultM,
		ProjectCreationResultM: &IoC.ProjectCreationResultM}
}

func (k *kafkaListener) StartKafkaListening(wg *sync.WaitGroup) {

	wg.Add(5)
	helper.CreateHealthFile()
	go (*k.Kafka).Consume("RemoteOfferEventModel",
		"RemoteOfferModel_ConsumerGroup",
		wg,
		(*k.RemoteOfferM).SendMessageToClient)

	go (*k.Kafka).Consume("InterstitialAdEventModel",
		"InterstitialAdModel_ConsumerGroup",
		wg,
		(*k.InterstitialAdM).SendMessageToClient)

	go (*k.Kafka).Consume("ChurnBlockerResultModel",
		"ChurnBlockerResult_ConsumerGroup",
		wg,
		(*k.ChurnBlockerResultM).SendMessageToClient)

	go (*k.Kafka).Consume("ChurnPredictionResultModel",
		"ChurnPredictionResult_ConsumerGroup",
		wg,
		(*k.ChurnPredictionResultM).SendMessageToClient)

	go (*k.Kafka).Consume("ProjectCreationResult",
		"ProjectCreationResultModel_ConsumerGroup",
		wg,
		(*k.ProjectCreationResultM).SendMessageToCustomer)
}
