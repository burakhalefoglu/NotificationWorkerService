package RemoteOfferManager

import (
	"NotificationWorkerService/internal/helper/jsonParser"
	"NotificationWorkerService/internal/models"
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)
type ProductModel struct{

	Name string
	Image []byte
	Count float32

}

type RemoteOfferKafkaModel struct{
	ProductModel []*ProductModel
	ClientIdList []string
	ProjectId string
	FirstPrice float32
	LastPrice float32
	OfferId int
	IsGift bool
	GiftTexture []byte
	StartTime time.Time
	FinishTime time.Time
}

type RemoteOfferResponseModel struct{
	ProductModel []*ProductModel
	FirstPrice float32
	LastPrice float32
	OfferId int
	IsGift bool
	GiftTexture []byte
	StartTime time.Time
	FinishTime time.Time
}
 
func SendMessageToClient(reader *kafka.Reader, m kafka.Message) {

	remoteOfferKafkaModel := RemoteOfferKafkaModel{}
	err := jsonParser.DecodeJson(m.Value, &remoteOfferKafkaModel)
	if(err != nil){
		log.Fatal(err)
		return
	}

	for _, clientId := range remoteOfferKafkaModel.ClientIdList{

		for _, client := range models.RemoteOfferChannelModel.Clients{

			if(clientId == client.Id && client.ProjectId == remoteOfferKafkaModel.ProjectId) {

				remoteOfferResponseModel := RemoteOfferResponseModel{
					ProductModel: remoteOfferKafkaModel.ProductModel,
					FirstPrice: remoteOfferKafkaModel.FirstPrice,
					LastPrice: remoteOfferKafkaModel.LastPrice,
					OfferId: remoteOfferKafkaModel.OfferId,
					IsGift: remoteOfferKafkaModel.IsGift,
					GiftTexture: remoteOfferKafkaModel.GiftTexture,
					StartTime: remoteOfferKafkaModel.StartTime,
					FinishTime: remoteOfferKafkaModel.FinishTime,

				}
				v, err := jsonParser.EncodeJson((remoteOfferResponseModel))
				if(!err){
					client.Connection.WriteMessage(1, v)
				}
			}
		}


	}

	if err := reader.CommitMessages(context.Background(), m); err != nil {
	log.Fatal("failed to commit messages:", err)
	}
}