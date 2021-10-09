package ClientListenerWorker

import (
	kafkaAdapter "NotificationWorkerService/internal/helper/kafka"
	"NotificationWorkerService/internal/manager/DifficultyServerResultManager"
	"NotificationWorkerService/internal/manager/InterstielAdManager"
	"NotificationWorkerService/internal/manager/RemoteOfferManager"
)

func StartKafkaListening() {
	kafkaAdapter.Consume("RemoteOfferModel", "RemoteOfferModel_ConsumerGroup", RemoteOfferManager.SendMessageToClient)
	kafkaAdapter.Consume("InterstielAdModel", "InterstielAdModel_ConsumerGroup", InterstielAdManager.SendMessageToClient)
	kafkaAdapter.Consume("DifficultyServerResultModel", "DifficultyServerResultModel_ConsumerGroup", DifficultyServerResultManager.SendMessageToClient)

}
