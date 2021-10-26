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

	m := Models.InterstitialAdModel{}
	err := i.JsonParser.DecodeJson(data, &m)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	for _, clientId := range m.ClientIdList{

		interstitialAdResponseModel := Models.InterstitialAdDto{
			IsAdvSettingsActive:    m.IsAdvSettingsActive,
			AdvFrequencyStrategies: m.AdvFrequencyStrategies,
		}
		v, err := i.JsonParser.EncodeJson(&interstitialAdResponseModel)
		if err != nil{
			return false, err.Error()
		}
		i.WebSocket.SendMessageToClient(v,
			clientId,
			m.ProjectId,
			"InterstitialAdChannel")
	}
	return true, ""
}
