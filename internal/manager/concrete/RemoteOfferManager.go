package concrete

import (
	"NotificationWorkerService/internal/models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type RemoteOfferManager struct {
	WebSocket IWebSocket.IWebsocket
	JsonParser IJsonParser.IJsonParser
}
 
func (r *RemoteOfferManager) SendMessageToClient(data *[]byte)(success bool, message string) {

	remoteOfferKafkaModel := models.RemoteOfferModel{}
	err := r.JsonParser.DecodeJson(data, &remoteOfferKafkaModel)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	for _, clientId := range remoteOfferKafkaModel.ClientIdList{

		remoteOfferResponseModel := models.RemoteOfferDto{
			ProductModel: remoteOfferKafkaModel.ProductModel,
			FirstPrice: remoteOfferKafkaModel.FirstPrice,
			LastPrice: remoteOfferKafkaModel.LastPrice,
			OfferId: remoteOfferKafkaModel.OfferId,
			IsGift: remoteOfferKafkaModel.IsGift,
			GiftTexture: remoteOfferKafkaModel.GiftTexture,
			StartTime: remoteOfferKafkaModel.StartTime,
			FinishTime: remoteOfferKafkaModel.FinishTime,

		}
		v, err := r.JsonParser.EncodeJson(&remoteOfferResponseModel)
		if err != nil{
			return false, err.Error()
		}
		websocketErr := r.WebSocket.SendMessageToClient(v,
			clientId,
			remoteOfferKafkaModel.ProjectId,
			"RemoteOfferChannel")
		if websocketErr != nil {
			return false, websocketErr.Error()
		}
	}
	return true, ""
}