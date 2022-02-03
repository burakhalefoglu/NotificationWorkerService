package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	jsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type churnBlockerManager struct {
	WebSocket  *IWebSocket.IWebsocket
	JsonParser *jsonParser.IJsonParser
}

func ChurnBlockerManagerConstructor() *churnBlockerManager {
	return &churnBlockerManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser}
}

func (c *churnBlockerManager) SendMessageToClient(data *[]byte)(success bool, message string)  {
	m := models.ChurnBlockerResultModel{}
	if err := (*c.JsonParser).DecodeJson(data, &m); err != nil {
		log.Fatal("ChurnBlockerManager", "SendMessageToClient",
			"byte array to ChurnBlockerResultModel", "Json Parser Decode Err: ", err)

		return false, err.Error()
	}

	defer log.Print("ChurnBlockerManager", "SendMessageToClient",
		m.ClientId, m.ProjectId)

	difficultyServerResultResponseModel := models.ChurnBlockerResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	v, err := (*c.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
	if err != nil{
		log.Fatal("ChurnBlockerManager", "SendMessageToClient",
			"difficultyServerResultResponseModel to byte array", "Json Parser Encode Err: ", err)
		return false, err.Error()
	}

	WebSocketErr := (*c.WebSocket).SendMessageToClient(v,
		m.ClientId,
		m.ProjectId,
		"ChurnBlockerResultChannel")
	if WebSocketErr != nil {
		log.Fatal("ChurnBlockerManager", "SendMessageToClient",
			"WebSocket error: ", err)
		return false, WebSocketErr.Error()
	}

	return true, ""
}

