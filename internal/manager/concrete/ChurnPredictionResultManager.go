package concrete

import (
	"NotificationWorkerService/internal/models"
	IwebSocket "NotificationWorkerService/internal/websocket"
	jsonparser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type ChurnPredictionManager struct {
	WebSocket IwebSocket.IWebsocket
	JsonParser jsonparser.IJsonParser
}

func (c *ChurnPredictionManager) SendMessageToClient(data *[]byte)(success bool, message string)  {
	m := models.ChurnPredictionResultModel{}
	err := c.JsonParser.DecodeJson(data, &m)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	difficultyServerResultResponseModel := models.ChurnPredictionResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	log.Println(difficultyServerResultResponseModel)
	v, err := c.JsonParser.EncodeJson(&difficultyServerResultResponseModel)
		if err != nil{
			return false, err.Error()
		}
	websocketErr := c.WebSocket.SendMessageToClient(v,
		m.ClientId,
		m.ProjectId,
		"ChurnPredictionResultChannel")
	if websocketErr != nil {
		return false, websocketErr.Error()
	}
	return true, ""
}



