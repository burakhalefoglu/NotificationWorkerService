package IoC

import (
	"NotificationWorkerService/internal/controller"
	"NotificationWorkerService/internal/manager/abstract"
	"NotificationWorkerService/internal/websocket"
	cache "NotificationWorkerService/pkg/Cache"
	jsonParser "NotificationWorkerService/pkg/jsonParser"
	"NotificationWorkerService/pkg/kafka"
	"NotificationWorkerService/pkg/logger"
)

type IContainer interface {
	Inject()
}

func InjectContainers(container IContainer){
	container.Inject()
}

var RedisCache cache.ICache
var Logger logger.ILog
var Kafka kafka.IKafka
var JsonParser jsonParser.IJsonParser

var Controller controller.IController
var WebSocket websocket.IWebsocket
var RemoteOfferM    abstract.IRemoteOfferService
var InterstitialAdM abstract.IInterstitialAdService
var ChurnBlockerResultM abstract.IChurnBlockerService
var ChurnPredictionResultM abstract.IChurnPredictionService
var ProjectCreationResultM abstract.IProjectCreationService