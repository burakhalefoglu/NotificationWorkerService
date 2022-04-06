package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
)

type interstitialManager struct {
	WebSocket  *IWebSocket.IWebsocket
	JsonParser *IJsonParser.IJsonParser
}

func InterstitialManagerConstructor() *interstitialManager {
	return &interstitialManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser}
}

func (i *interstitialManager) SendMessageToClient(data *[]byte) (success bool, message string) {

	m := models.InterstitialAdModel{}
	err := (*i.JsonParser).DecodeJson(data, &m)
	if err != nil {
		clogger.Error(&logger.Messages{
			"byte array to InterstitialAdModel, Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	defer clogger.Info(&logger.Messages{
		"ChurnPredictionManager": m.ClientIdList,
		"projectId":              m.ProjectId,
	})

	for _, clientId := range m.ClientIdList {

		interstitialAdResponseModel := models.InterstitialAdDto{
			IsAdvSettingsActive:    m.IsAdvSettingsActive,
			AdvFrequencyStrategies: m.AdvFrequencyStrategies,
		}
		v, err := (*i.JsonParser).EncodeJson(&interstitialAdResponseModel)
		if err != nil {
			clogger.Error(&logger.Messages{
				"interstitialAdResponseModel to byte array Json Parser Encode Err: ": err,
			})
			return false, err.Error()
		}
		websocketErr := (*i.WebSocket).SendMessageToClient(v,
			clientId,
			m.ProjectId,
			"InterstitialAdChannel")
		if websocketErr != nil {
			clogger.Error(&logger.Messages{
				"WebSocket error: ": err,
			})
			return false, websocketErr.Error()
		}
	}
	return true, ""
}
