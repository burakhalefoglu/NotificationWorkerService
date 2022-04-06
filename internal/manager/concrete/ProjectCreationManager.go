package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"

	"github.com/appneuroncompany/light-logger/clogger"
)

type projectCreationManager struct {
	WebSocket  *IWebSocket.IWebsocket
	JsonParser *IJsonParser.IJsonParser
}

func ProjectCreationManagerConstructor() *projectCreationManager {
	return &projectCreationManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser}
}

func (r *projectCreationManager) SendMessageToCustomer(data *[]byte) (success bool, message string) {

	m := models.ProjectCreationResultModel{}
	err := (*r.JsonParser).DecodeJson(data, &m)
	if err != nil {
		clogger.Error(&map[string]interface{}{
			"byte array to ProjectCreationResultModel, Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}
	defer clogger.Info(&map[string]interface{}{
		"ChurnPredictionManager": m.CustomerId + m.ProjectId,
	})

	responseModel := models.ProjectCreationResultDto{
		Token: m.Token,
	}
	v, err := (*r.JsonParser).EncodeJson(&responseModel)
	if err != nil {
		clogger.Error(&map[string]interface{}{
			"ProjectCreationResultDto to byte array Json Parser Encode Err: ": err,
		})
		return false, err.Error()
	}
	websocketErr := (*r.WebSocket).SendMessageToCustomer(v,
		m.CustomerId,
		m.ProjectId,
		"ProjectCreationResultChannel")
	if websocketErr != nil {
		clogger.Error(&map[string]interface{}{
			"WebSocket error: ": err,
		})
		return false, websocketErr.Error()
	}
	return true, ""
}
