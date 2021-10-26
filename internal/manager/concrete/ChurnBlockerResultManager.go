package concrete

import (
	"NotificationWorkerService/internal/Models"
	IwebSocket "NotificationWorkerService/internal/websocket"
	jsonparser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type ChurnBlockerManager struct {
	WebSocket IwebSocket.IWebsocket
	JsonParser jsonparser.IJsonParser
}

func (c *ChurnBlockerManager) SendMessageToClient(data *[]byte)(success bool, message string)  {
	m := Models.ChurnBlockerResultModel{}
	err := c.JsonParser.DecodeJson(data, &m)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	difficultyServerResultResponseModel := Models.ChurnBlockerResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	v, err := c.JsonParser.EncodeJson(&difficultyServerResultResponseModel)
	if err != nil{
		return false, err.Error()
	}
	c.WebSocket.SendMessageToClient(v,
		m.ClientId,
		m.ProjectId,
		"ChurnBlockerResultChannel")
	return true, ""
}
