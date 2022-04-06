package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	jsonParser "NotificationWorkerService/pkg/jsonParser"

	"github.com/appneuroncompany/light-logger/clogger"
)

type churnBlockerManager struct {
	WebSocket  *IWebSocket.IWebsocket
	JsonParser *jsonParser.IJsonParser
}

func ChurnBlockerManagerConstructor() *churnBlockerManager {
	return &churnBlockerManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser}
}

func (c *churnBlockerManager) SendMessageToClient(data *[]byte) (success bool, message string) {
	m := models.ChurnBlockerResultModel{}
	if err := (*c.JsonParser).DecodeJson(data, &m); err != nil {
		clogger.Error(&map[string]interface{}{
			"byte array to ChurnBlockerResultModel, Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&map[string]interface{}{
		"ChurnBlockerManager": m.ClientId + m.ProjectId,
	})

	difficultyServerResultResponseModel := models.ChurnBlockerResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	v, err := (*c.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
	if err != nil {
		clogger.Error(&map[string]interface{}{
			"difficultyServerResultResponseModel to byte array Json Parser Encode Err: ": err,
		})
		return false, err.Error()
	}

	WebSocketErr := (*c.WebSocket).SendMessageToClient(v,
		m.ClientId,
		m.ProjectId,
		"ChurnBlockerResultChannel")
	if WebSocketErr != nil {
		clogger.Error(&map[string]interface{}{
			"WebSocket error: ": err,
		})
		return false, WebSocketErr.Error()
	}

	return true, ""
}
