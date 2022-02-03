package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type remoteOfferManager struct {
	WebSocket *IWebSocket.IWebsocket
	JsonParser *IJsonParser.IJsonParser
}

func RemoteOfferManagerConstructor() *remoteOfferManager {
	return &remoteOfferManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser}
}
 
func (r *remoteOfferManager) SendMessageToClient(data *[]byte)(success bool, message string) {

	m := models.RemoteOfferModel{}
	err := (*r.JsonParser).DecodeJson(data, &m)
	if err != nil {
		log.Fatal("RemoteOfferManager", "SendMessageToClient",
			"byte array to RemoteOfferModel", "Json Parser Decode Err: ", err)
		return false, err.Error()
	}

	defer log.Print("RemoteOfferManager", "SendMessageToClient",
		m.ClientIdList, m.ProjectId)


	for _, clientId := range m.ClientIdList{

		remoteOfferResponseModel := models.RemoteOfferDto{
			ProductModel: m.ProductModel,
			FirstPrice:   m.FirstPrice,
			LastPrice:    m.LastPrice,
			OfferId:      m.OfferId,
			IsGift:       m.IsGift,
			GiftTexture:  m.GiftTexture,
			StartTime:    m.StartTime,
			FinishTime:   m.FinishTime,

		}
		v, err := (*r.JsonParser).EncodeJson(&remoteOfferResponseModel)
		if err != nil{
			log.Fatal("RemoteOfferManager", "SendMessageToClient",
				"RemoteOfferResponseModel to byte array", "Json Parser Encode Err: ", err)
			return false, err.Error()
		}
		websocketErr := (*r.WebSocket).SendMessageToClient(v,
			clientId,
			m.ProjectId,
			"RemoteOfferChannel")
		if websocketErr != nil {
			log.Fatal("RemoteOfferManager", "SendMessageToClient",
				"WebSocket error: ", err)
			return false, websocketErr.Error()
		}
	}
	return true, ""
}