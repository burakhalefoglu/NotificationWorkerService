package concrete

import (
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)


type InterstitialManager struct {
	WebSocket  IWebSocket.IWebsocket
	JsonParser IJsonParser.IJsonParser
}


func (i *InterstitialManager) SendMessageToClient(data *[]byte)(success bool, message string) {

	m := models.InterstitialAdModel{}
	err := i.JsonParser.DecodeJson(data, &m)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	for _, clientId := range m.ClientIdList{

		interstitialAdResponseModel := models.InterstitialAdDto{
			IsAdvSettingsActive:    m.IsAdvSettingsActive,
			AdvFrequencyStrategies: m.AdvFrequencyStrategies,
		}
		v, err := i.JsonParser.EncodeJson(&interstitialAdResponseModel)
		if err != nil{
			return false, err.Error()
		}
		websocketErr := i.WebSocket.SendMessageToClient(v,
			clientId,
			m.ProjectId,
			"InterstitialAdChannel")
		if websocketErr != nil {
			return false, websocketErr.Error()
		}
	}
	return true, ""
}
