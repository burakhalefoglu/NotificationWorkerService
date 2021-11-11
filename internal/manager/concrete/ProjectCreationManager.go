package concrete

import (
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type ProjectCreationManager struct {
	WebSocket IWebSocket.IWebsocket
	JsonParser IJsonParser.IJsonParser
}

func (r *ProjectCreationManager) SendMessageToCustomer(data *[]byte)(success bool, message string) {

	m := models.ProjectCreationResultModel{}
	err := r.JsonParser.DecodeJson(data, &m)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
		responseModel := models.ProjectCreationResultDto{
			Token: m.Token,
		}
		v, err := r.JsonParser.EncodeJson(&responseModel)
		if err != nil{
			return false, err.Error()
		}
	websocketErr := r.WebSocket.SendMessageToCustomer(v,
		m.CustomerId,
		m.ProjectId,
		"ProjectCreationResultChannel")
	if websocketErr != nil {
		return false, websocketErr.Error()
	}
	return true, ""
}
