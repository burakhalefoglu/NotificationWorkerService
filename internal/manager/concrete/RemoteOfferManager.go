package concrete

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"NotificationWorkerService/pkg/logger"
)

type remoteOfferManager struct {
	WebSocket *IWebSocket.IWebsocket
	JsonParser *IJsonParser.IJsonParser
	Logg *logger.ILog
}

func RemoteOfferManagerConstructor() *remoteOfferManager {
	return &remoteOfferManager{WebSocket: &IoC.WebSocket,
		JsonParser: &IoC.JsonParser,
		Logg: &IoC.Logger}
}
 
func (r *remoteOfferManager) SendMessageToClient(data *[]byte)(success bool, message string) {

	m := models.RemoteOfferModel{}
	err := (*r.JsonParser).DecodeJson(data, &m)
	if err != nil {
		(*r.Logg).SendErrorLog("RemoteOfferManager", "SendMessageToClient",
			"byte array to RemoteOfferModel", "Json Parser Decode Err: ", err)
		return false, err.Error()
	}

	defer (*r.Logg).SendInfoLog("RemoteOfferManager", "SendMessageToClient",
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
			(*r.Logg).SendErrorLog("RemoteOfferManager", "SendMessageToClient",
				"RemoteOfferResponseModel to byte array", "Json Parser Encode Err: ", err)
			return false, err.Error()
		}
		websocketErr := (*r.WebSocket).SendMessageToClient(v,
			clientId,
			m.ProjectId,
			"RemoteOfferChannel")
		if websocketErr != nil {
			(*r.Logg).SendErrorLog("RemoteOfferManager", "SendMessageToClient",
				"WebSocket error: ", err)
			return false, websocketErr.Error()
		}
	}
	return true, ""
}