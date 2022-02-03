package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type projectCreationManager struct {
	WebSocket *IWebSocket.IWebsocket
	JsonParser *IJsonParser.IJsonParser
}

func ProjectCreationManagerConstructor() *projectCreationManager {
	return &projectCreationManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser}
}

func (r *projectCreationManager) SendMessageToCustomer(data *[]byte)(success bool, message string) {

	m := models.ProjectCreationResultModel{}
	err := (*r.JsonParser).DecodeJson(data, &m)
	if err != nil {
		log.Fatal("projectCreationManager", "SendMessageToCustomer",
			"byte array to ProjectCreationResultModel", "Json Parser Decode Err: ", err)
		return false, err.Error()
	}

	defer log.Print("projectCreationManager", "SendMessageToCustomer",
		m.CustomerId, m.ProjectId)

		responseModel := models.ProjectCreationResultDto{
			Token: m.Token,
		}
		v, err := (*r.JsonParser).EncodeJson(&responseModel)
		if err != nil{
			log.Fatal("projectCreationManager", "SendMessageToCustomer",
				"ProjectCreationResultDto to byte array", "Json Parser Encode Err: ", err)
			return false, err.Error()
		}
	websocketErr := (*r.WebSocket).SendMessageToCustomer(v,
		m.CustomerId,
		m.ProjectId,
		"ProjectCreationResultChannel")
	if websocketErr != nil {
		log.Fatal("projectCreationManager", "SendMessageToCustomer",
			"WebSocket error: ", err)
		return false, websocketErr.Error()
	}
	return true, ""
}
