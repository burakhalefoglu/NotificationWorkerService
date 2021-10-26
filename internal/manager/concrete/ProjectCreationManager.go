package concrete

import (
	"NotificationWorkerService/internal/Models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type ProjectCreationManager struct {
	WebSocket IWebSocket.IWebsocket
	JsonParser IJsonParser.IJsonParser
}

func (r *ProjectCreationManager) SendMessageToCustomer(data *[]byte)(success bool, message string) {

	m := Models.ProjectCreationResultModel{}
	err := r.JsonParser.DecodeJson(data, &m)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
		responseModel := Models.ProjectCreationResultDto{
			Token: m.Token,
		}
		v, err := r.JsonParser.EncodeJson(&responseModel)
		if err != nil{
			return false, err.Error()
		}
		r.WebSocket.SendMessageToCustomer(v,
			m.CustomerId,
			m.ProjectId,
			"ProjectCreationResultChannel")
	return true, ""
}
