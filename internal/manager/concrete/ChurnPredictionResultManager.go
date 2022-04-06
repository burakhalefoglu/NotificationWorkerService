package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	JsonParser "NotificationWorkerService/pkg/jsonParser"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
)

type churnPredictionManager struct {
	WebSocket  *IWebSocket.IWebsocket
	JsonParser *JsonParser.IJsonParser
}

func ChurnPredictionManagerConstructor() *churnPredictionManager {
	return &churnPredictionManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser}
}

func (c *churnPredictionManager) SendMessageToClient(data *[]byte) (success bool, message string) {
	m := models.ChurnPredictionResultModel{}
	err := (*c.JsonParser).DecodeJson(data, &m)
	if err != nil {
		clogger.Error(&logger.Messages{
			"byte array to ChurnPredictionResultModel, Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&logger.Messages{
		"ChurnPredictionManager": m.ClientId + m.ProjectId,
	})

	difficultyServerResultResponseModel := models.ChurnPredictionResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}

	v, err := (*c.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
	if err != nil {
		clogger.Error(&logger.Messages{
			"difficultyServerResultResponseModel to byte array Json Parser Encode Err: ": err,
		})
		return false, err.Error()
	}
	websocketErr := (*c.WebSocket).SendMessageToClient(v,
		m.ClientId,
		m.ProjectId,
		"ChurnPredictionResultChannel")
	if websocketErr != nil {
		clogger.Error(&logger.Messages{
			"WebSocket error: ": err,
		})
		return false, websocketErr.Error()
	}
	return true, ""
}
