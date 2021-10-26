package IoC

import (
	"NotificationWorkerService/internal/controller/KafkaListener"
	"NotificationWorkerService/internal/manager/concrete"
	fastHttpServer "NotificationWorkerService/internal/websocket/fasthttp"
	"NotificationWorkerService/pkg/jsonParser/gojson"
	"NotificationWorkerService/pkg/kafka/confluent"
)

var KafkaController = &KafkaListener.KafkaListener{
	Kafka:                  &confluent.Kafka{},

	RemoteOfferM:           &concrete.RemoteOfferManager{
		WebSocket:  fastHttpServer.FastHttpServer,
		JsonParser: &gojson.GoJson{},
	},

	InterstitialAdM:        &concrete.InterstitialManager{
		WebSocket:  fastHttpServer.FastHttpServer,
		JsonParser: &gojson.GoJson{},
	},

	ChurnBlockerResultM:    &concrete.ChurnBlockerManager{
		WebSocket:  fastHttpServer.FastHttpServer,
		JsonParser: &gojson.GoJson{},
	},

	ChurnPredictionResultM: &concrete.ChurnPredictionManager{
		WebSocket:  fastHttpServer.FastHttpServer,
		JsonParser: &gojson.GoJson{},
	},

	ProjectCreationResultM: &concrete.ProjectCreationManager{
		WebSocket:  fastHttpServer.FastHttpServer,
		JsonParser: &gojson.GoJson{},
	},
}
