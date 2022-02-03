package golobby

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/controller"
	"NotificationWorkerService/internal/controller/KafkaListener"
	"NotificationWorkerService/internal/manager/abstract"
	"NotificationWorkerService/internal/manager/concrete"
	"NotificationWorkerService/internal/websocket"
	"NotificationWorkerService/internal/websocket/fiber"
	cache "NotificationWorkerService/pkg/Cache"
	RedisCacheV8 "NotificationWorkerService/pkg/Cache/Redis/RedisV8"
	jsonParser "NotificationWorkerService/pkg/jsonParser"
	"NotificationWorkerService/pkg/jsonParser/gojson"
	"NotificationWorkerService/pkg/kafka"
	"NotificationWorkerService/pkg/kafka/kafkago"
	"github.com/golobby/container/v3"
)

type golobbyInjection struct {}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func (i *golobbyInjection) Inject(){
	injectKafka()
	injectJsonParser()
	injectCache()

	injectWebSocket()
	injectController()

	injectChurnBlockerManager()
	injectChurnPredictionManager()
	injectInterstitialAdManager()
	injectProjectCreationManager()
	injectRemoteOfferManager()

}

func injectJsonParser() {
	if err := container.Singleton(func() jsonParser.IJsonParser {
		return gojson.GoJsonConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.JsonParser); err != nil{
		panic(err)
	}
}

func injectKafka() {
	if err :=container.Transient(func() kafka.IKafka {
		return kafkago.KafkaGoConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.Kafka); err != nil{
		panic(err)
	}
}

func injectCache() {
	if err := container.Transient(func() cache.ICache {
		return RedisCacheV8.RedisCacheConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.RedisCache); err != nil{
		panic(err)
	}
}

func injectWebSocket() {
	if err := container.Transient(func() websocket.IWebsocket {
		return fiber.FiberWebsocketConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.WebSocket); err != nil{
		panic(err)
	}
}

func injectController() {
	if err := container.Transient(func() controller.IController {
		return KafkaListener.KafkaListenerConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.Controller); err != nil{
		panic(err)
	}
}

func injectChurnBlockerManager() {
	if err := container.Transient(func() abstract.IChurnBlockerService {
		return concrete.ChurnBlockerManagerConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.ChurnBlockerResultM); err != nil{
		panic(err)
	}
}

func injectChurnPredictionManager() {
	if err := container.Transient(func() abstract.IChurnPredictionService {
		return concrete.ChurnPredictionManagerConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.ChurnPredictionResultM); err != nil{
		panic(err)
	}
}

func injectInterstitialAdManager() {
	if err := container.Transient(func() abstract.IInterstitialAdService {
		return concrete.InterstitialManagerConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.InterstitialAdM); err != nil{
		panic(err)
	}
}

func injectProjectCreationManager() {
	if err := container.Transient(func() abstract.IProjectCreationService {
		return concrete.ProjectCreationManagerConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.ProjectCreationResultM); err != nil{
		panic(err)
	}
}

func injectRemoteOfferManager() {
	if err := container.Transient(func() abstract.IRemoteOfferService {
		return concrete.RemoteOfferManagerConstructor()
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.RemoteOfferM); err != nil{
		panic(err)
	}
}

