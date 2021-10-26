package concrete

import (
	"NotificationWorkerService/internal/Models"
	IWebSocket "NotificationWorkerService/internal/websocket"
	IJsonParser "NotificationWorkerService/pkg/jsonParser"
	"log"
)

type RemoteOfferManager struct {
	WebSocket IWebSocket.IWebsocket
	JsonParser IJsonParser.IJsonParser
}
 
func (r *RemoteOfferManager) SendMessageToClient(data *[]byte)(success bool, message string) {

	remoteOfferKafkaModel := Models.RemoteOfferModel{}
	err := r.JsonParser.DecodeJson(data, &remoteOfferKafkaModel)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	for _, clientId := range remoteOfferKafkaModel.ClientIdList{

		remoteOfferResponseModel := Models.RemoteOfferDto{
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
		r.WebSocket.SendMessageToClient(v,
			clientId,
			remoteOfferKafkaModel.ProjectId,
			"RemoteOfferChannel")
	}
	return true, ""
}