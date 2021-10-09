package DifficultyServerResultManager

import (
	"NotificationWorkerService/internal/helper/jsonParser"
	"NotificationWorkerService/internal/models"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type DifficultyServerResultKafkaModel struct {
	ClientId                string
	ProjectId               string
	CenterOfDifficultyLevel int
	RangeCount              int
}

type DifficultyServerResultResponseModel struct {
	CenterOfDifficultyLevel int
	RangeCount              int
}

func SendMessageToClient(reader *kafka.Reader, m kafka.Message) {

	difficultyServerResultKafkaModel := DifficultyServerResultKafkaModel{}
	err := jsonParser.DecodeJson(m.Value, &difficultyServerResultKafkaModel)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, client := range models.DifficultyServerResultChannelModel.Clients {

		if difficultyServerResultKafkaModel.ClientId == client.Id &&
			client.ProjectId == difficultyServerResultKafkaModel.ProjectId {

			difficultyServerResultResponseModel := DifficultyServerResultResponseModel{
				CenterOfDifficultyLevel: difficultyServerResultKafkaModel.CenterOfDifficultyLevel,
				RangeCount:              difficultyServerResultKafkaModel.RangeCount,
			}
			v, err := jsonParser.EncodeJson((difficultyServerResultResponseModel))
			if !err {
				client.Connection.WriteMessage(1, v)
			}
			break
		}
	}

	if err := reader.CommitMessages(context.Background(), m); err != nil {
		log.Fatal("failed to commit messages:", err)
	}
}
