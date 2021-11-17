package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	JsonParser "NotificationWorkerService/pkg/jsonParser"
	"NotificationWorkerService/pkg/logger"
)

type churnPredictionManager struct {
	WebSocket *IWebSocket.IWebsocket
	JsonParser *JsonParser.IJsonParser
	Logg *logger.ILog
}

func ChurnPredictionManagerConstructor() *churnPredictionManager {
	return &churnPredictionManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser,
		Logg: &IoC.Logger}
}

func (c *churnPredictionManager) SendMessageToClient(data *[]byte)(success bool, message string)  {
	m := models.ChurnPredictionResultModel{}
	err := (*c.JsonParser).DecodeJson(data, &m)
	if err != nil {
		(*c.Logg).SendErrorLog("ChurnPredictionManager", "SendMessageToClient",
			"byte array to ChurnPredictionResultModel", "Json Parser Decode Err: ", err)
		return false, err.Error()
	}

	defer (*c.Logg).SendInfoLog("ChurnPredictionManager", "SendMessageToClient",
		m.ClientId, m.ProjectId)

	difficultyServerResultResponseModel := models.ChurnPredictionResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}

	v, err :=(*c.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
		if err != nil{
			(*c.Logg).SendErrorLog("ChurnPredictionManager", "SendMessageToClient",
				"difficultyServerResultResponseModel to byte array", "Json Parser Encode Err: ", err)
			return false, err.Error()
		}
	websocketErr := (*c.WebSocket).SendMessageToClient(v,
		m.ClientId,
		m.ProjectId,
		"ChurnPredictionResultChannel")
	if websocketErr != nil {
		(*c.Logg).SendErrorLog("ChurnPredictionManager", "SendMessageToClient",
			"WebSocket error: ", err)
		return false, websocketErr.Error()
	}
	return true, ""
}



