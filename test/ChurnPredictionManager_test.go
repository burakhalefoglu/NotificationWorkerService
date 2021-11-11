package test

import (
	"NotificationWorkerService/internal/manager/concrete"
	"NotificationWorkerService/internal/models"
	"NotificationWorkerService/pkg/jsonParser/gojson"
	"NotificationWorkerService/test/Mock/mockwebsocket"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ChurnPredictionSendMessageToClient_SuccessIsTrue(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var churnPredcition = concrete.ChurnPredictionManager{
		JsonParser:      &gojson.GoJson{},
		WebSocket: testWebsocket,
	}
	m := models.ChurnPredictionResultModel{
		ClientId:                "TestClientId",
		ProjectId:               "TestProjectId",
		CenterOfDifficultyLevel: 8,
		RangeCount:              2,
	}
	difficultyServerResultResponseModel := models.ChurnPredictionResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	responseModel, _ := churnPredcition.JsonParser.EncodeJson(&difficultyServerResultResponseModel)
	testWebsocket.On("SendMessageToClient",
		responseModel,
		"TestClientId",
		"TestProjectId",
		"ChurnPredictionResultChannel").Return(nil)

	rawModel, _ := churnPredcition.JsonParser.EncodeJson(&m)

	//Act
	success, err:= churnPredcition.SendMessageToClient(rawModel)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_ChurnPredictionSendMessageToClient_SuccessIsFalse(t *testing.T){
	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var churnPrediction = concrete.ChurnPredictionManager{
		JsonParser:      &gojson.GoJson{},
		WebSocket: testWebsocket,
	}
	m := models.ChurnPredictionResultModel{
		ClientId:                "TestClientId",
		ProjectId:               "TestProjectId",
		CenterOfDifficultyLevel: 8,
		RangeCount:              2,
	}
	difficultyServerResultResponseModel := models.ChurnPredictionResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	responseModel, _ := churnPrediction.JsonParser.EncodeJson(&difficultyServerResultResponseModel)
	testWebsocket.On("SendMessageToClient",
		responseModel,
		"TestClientId",
		"TestProjectId",
		"ChurnPredictionResultChannel").Return(errors.New("FakeError"))

	rawModel, _ := churnPrediction.JsonParser.EncodeJson(&m)

	//Act
	success, err:= churnPrediction.SendMessageToClient(rawModel)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)

}

