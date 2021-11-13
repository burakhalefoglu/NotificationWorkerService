package IoC

import (
	"NotificationWorkerService/internal/controller/KafkaListener"
	"NotificationWorkerService/internal/manager/concrete"
	fastHttpServer "NotificationWorkerService/internal/websocket/fiber"
	"NotificationWorkerService/pkg/jsonParser/gojson"
	"NotificationWorkerService/pkg/kafka/kafkago"
)

var KafkaController = &KafkaListener.KafkaListener{
	Kafka:                  &kafkago.KafkaGo{},

	RemoteOfferM:           &concrete.RemoteOfferManager{
		WebSocket:  fastHttpServer.FiberServer,
		JsonParser: &gojson.GoJson{},
	},

	InterstitialAdM:        &concrete.InterstitialManager{
		WebSocket:  fastHttpServer.FiberServer,
		JsonParser: &gojson.GoJson{},
	},

	ChurnBlockerResultM:    &concrete.ChurnBlockerManager{
		WebSocket:  fastHttpServer.FiberServer,
		JsonParser: &gojson.GoJson{},
	},

	ChurnPredictionResultM: &concrete.ChurnPredictionManager{
		WebSocket:  fastHttpServer.FiberServer,
		JsonParser: &gojson.GoJson{},
	},

	ProjectCreationResultM: &concrete.ProjectCreationManager{
		WebSocket:  fastHttpServer.FiberServer,
		JsonParser: &gojson.GoJson{},
	},
}
