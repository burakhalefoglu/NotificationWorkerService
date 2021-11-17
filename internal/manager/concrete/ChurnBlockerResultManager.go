package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	jsonParser "NotificationWorkerService/pkg/jsonParser"
	"NotificationWorkerService/pkg/logger"
)

type churnBlockerManager struct {
	WebSocket  *IWebSocket.IWebsocket
	JsonParser *jsonParser.IJsonParser
	Logg *logger.ILog
}

func ChurnBlockerManagerConstructor() *churnBlockerManager {
	return &churnBlockerManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser,
		Logg: &IoC.Logger}
}

func (c *churnBlockerManager) SendMessageToClient(data *[]byte)(success bool, message string)  {
	m := models.ChurnBlockerResultModel{}
	if err := (*c.JsonParser).DecodeJson(data, &m); err != nil {
		(*c.Logg).SendErrorLog("ChurnBlockerManager", "SendMessageToClient",
			"byte array to ChurnBlockerResultModel", "Json Parser Decode Err: ", err)

		return false, err.Error()
	}

	defer (*c.Logg).SendInfoLog("ChurnBlockerManager", "SendMessageToClient",
		m.ClientId, m.ProjectId)

	difficultyServerResultResponseModel := models.ChurnBlockerResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	v, err := (*c.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
	if err != nil{
		(*c.Logg).SendErrorLog("ChurnBlockerManager", "SendMessageToClient",
			"difficultyServerResultResponseModel to byte array", "Json Parser Encode Err: ", err)
		return false, err.Error()
	}

	WebSocketErr := (*c.WebSocket).SendMessageToClient(v,
		m.ClientId,
		m.ProjectId,
		"ChurnBlockerResultChannel")
	if WebSocketErr != nil {
		(*c.Logg).SendErrorLog("ChurnBlockerManager", "SendMessageToClient",
			"WebSocket error: ", err)
		return false, WebSocketErr.Error()
	}

	return true, ""
}

