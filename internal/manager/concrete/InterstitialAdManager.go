package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)


type interstitialManager struct {
	WebSocket  *IWebSocket.IWebsocket
	JsonParser *IJsonParser.IJsonParser
}

func InterstitialManagerConstructor() *interstitialManager {
	return &interstitialManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser}
}


func (i *interstitialManager) SendMessageToClient(data *[]byte)(success bool, message string) {

	m := models.InterstitialAdModel{}
	err := (*i.JsonParser).DecodeJson(data, &m)
	if err != nil {
		log.Fatal("InterstitialManager", "SendMessageToClient",
			"byte array to InterstitialAdModel", "Json Parser Decode Err: ", err)
		return false, err.Error()
	}

	defer log.Print("InterstitialManager", "SendMessageToClient",
		m.ClientIdList, m.ProjectId)

	for _, clientId := range m.ClientIdList{

		interstitialAdResponseModel := models.InterstitialAdDto{
			IsAdvSettingsActive:    m.IsAdvSettingsActive,
			AdvFrequencyStrategies: m.AdvFrequencyStrategies,
		}
		v, err := (*i.JsonParser).EncodeJson(&interstitialAdResponseModel)
		if err != nil{
			log.Fatal("InterstitialManager", "SendMessageToClient",
				"interstitialAdResponseModel to byte array", "Json Parser Encode Err: ", err)
			return false, err.Error()
		}
		websocketErr := (*i.WebSocket).SendMessageToClient(v,
			clientId,
			m.ProjectId,
			"InterstitialAdChannel")
		if websocketErr != nil {
			log.Fatal("InterstitialManager", "SendMessageToClient",
				"WebSocket error: ", err)
			return false, websocketErr.Error()
		}
	}
	return true, ""
}
