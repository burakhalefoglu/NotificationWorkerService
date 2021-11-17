package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"NotificationWorkerService/pkg/logger"
)


type interstitialManager struct {
	WebSocket  *IWebSocket.IWebsocket
	JsonParser *IJsonParser.IJsonParser
	Logg *logger.ILog
}

func InterstitialManagerConstructor() *interstitialManager {
	return &interstitialManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser,
		Logg: &IoC.Logger}
}


func (i *interstitialManager) SendMessageToClient(data *[]byte)(success bool, message string) {

	m := models.InterstitialAdModel{}
	err := (*i.JsonParser).DecodeJson(data, &m)
	if err != nil {
		(*i.Logg).SendErrorLog("InterstitialManager", "SendMessageToClient",
			"byte array to InterstitialAdModel", "Json Parser Decode Err: ", err)
		return false, err.Error()
	}

	defer (*i.Logg).SendInfoLog("InterstitialManager", "SendMessageToClient",
		m.ClientIdList, m.ProjectId)

	for _, clientId := range m.ClientIdList{

		interstitialAdResponseModel := models.InterstitialAdDto{
			IsAdvSettingsActive:    m.IsAdvSettingsActive,
			AdvFrequencyStrategies: m.AdvFrequencyStrategies,
		}
		v, err := (*i.JsonParser).EncodeJson(&interstitialAdResponseModel)
		if err != nil{
			(*i.Logg).SendErrorLog("InterstitialManager", "SendMessageToClient",
				"interstitialAdResponseModel to byte array", "Json Parser Encode Err: ", err)
			return false, err.Error()
		}
		websocketErr := (*i.WebSocket).SendMessageToClient(v,
			clientId,
			m.ProjectId,
			"InterstitialAdChannel")
		if websocketErr != nil {
			(*i.Logg).SendErrorLog("InterstitialManager", "SendMessageToClient",
				"WebSocket error: ", err)
			return false, websocketErr.Error()
		}
	}
	return true, ""
}
