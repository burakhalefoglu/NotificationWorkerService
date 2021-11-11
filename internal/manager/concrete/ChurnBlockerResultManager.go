package concrete

import (
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	jsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type ChurnBlockerManager struct {
	WebSocket  IWebSocket.IWebsocket
	JsonParser jsonParser.IJsonParser
}

func (c *ChurnBlockerManager) SendMessageToClient(data *[]byte)(success bool, message string)  {
	m := models.ChurnBlockerResultModel{}
	err := c.JsonParser.DecodeJson(data, &m)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	difficultyServerResultResponseModel := models.ChurnBlockerResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	v, err := c.JsonParser.EncodeJson(&difficultyServerResultResponseModel)
	if err != nil{
		return false, err.Error()
	}
	WebSocketErr := c.WebSocket.SendMessageToClient(v,
		m.ClientId,
		m.ProjectId,
		"ChurnBlockerResultChannel")
	if WebSocketErr != nil {
		return false, WebSocketErr.Error()
	}

	return true, ""
}

