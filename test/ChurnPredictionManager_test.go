package test

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/manager/concrete"
	"NotificationWorkerService/internal/models"
	"NotificationWorkerService/pkg/jsonParser/gojson"
	mocklog "NotificationWorkerService/test/Mock/Log"
	"NotificationWorkerService/test/Mock/mockwebsocket"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ChurnPredictionSendMessageToClient_SuccessIsTrue(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var testLog = new (mocklog.MockLogger)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket
	IoC.Logger = testLog

	var churnPrediction = concrete.ChurnPredictionManagerConstructor()

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
	responseModel, _ := (*churnPrediction.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
	testWebsocket.On("SendMessageToClient",
		responseModel,
		"TestClientId",
		"TestProjectId",
		"ChurnPredictionResultChannel").Return(nil)

	rawModel, _ := (*churnPrediction.JsonParser).EncodeJson(&m)

	//Act
	success, err:= churnPrediction.SendMessageToClient(rawModel)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_ChurnPredictionSendMessageToClient_SuccessIsFalse(t *testing.T){
	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var testLog = new (mocklog.MockLogger)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket
	IoC.Logger = testLog

	var churnPrediction = concrete.ChurnPredictionManagerConstructor()

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
	responseModel, _ := (*churnPrediction.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
	testWebsocket.On("SendMessageToClient",
		responseModel,
		"TestClientId",
		"TestProjectId",
		"ChurnPredictionResultChannel").Return(errors.New("FakeError"))

	rawModel, _ :=(*churnPrediction.JsonParser).EncodeJson(&m)

	//Act
	success, err:= churnPrediction.SendMessageToClient(rawModel)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)

}

