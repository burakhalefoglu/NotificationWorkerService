package InterstielAdManager

import (
	"NotificationWorkerService/internal/helper/jsonParser"
	"NotificationWorkerService/internal/models"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type InterstielAdKafkaModel struct {
	ClientIdList           []string
	ProjectId              string
	IsAdvSettingsActive    bool
	AdvFrequencyStrategies map[string]int
}

type InterstielAdResponseModel struct {
	IsAdvSettingsActive    bool
	AdvFrequencyStrategies map[string]int
}

func SendMessageToClient(reader *kafka.Reader, m kafka.Message) {

	interstielAdKafkaModel := InterstielAdKafkaModel{}
	err := jsonParser.DecodeJson(m.Value, &interstielAdKafkaModel)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, clientId := range interstielAdKafkaModel.ClientIdList {

		for _, client := range models.InterstielAdChannelModel.Clients {

			if clientId == client.Id && client.ProjectId == interstielAdKafkaModel.ProjectId {

				interstielAdResponseModel := InterstielAdResponseModel{
					IsAdvSettingsActive:    interstielAdKafkaModel.IsAdvSettingsActive,
					AdvFrequencyStrategies: interstielAdKafkaModel.AdvFrequencyStrategies,
				}
				v, err := jsonParser.EncodeJson((interstielAdResponseModel))
				if !err {
					client.Connection.WriteMessage(1, v)
				}
			}
		}

	}

	if err := reader.CommitMessages(context.Background(), m); err != nil {
		log.Fatal("failed to commit messages:", err)
	}
}
