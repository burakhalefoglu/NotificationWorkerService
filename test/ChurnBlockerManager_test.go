package test

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/manager/concrete"
	"NotificationWorkerService/internal/models"
	"NotificationWorkerService/pkg/jsonParser/gojson"
	"NotificationWorkerService/test/Mock/mockwebsocket"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendMessageToClient_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket

	var churnBlockerManager = concrete.ChurnBlockerManagerConstructor()

	m := models.ChurnBlockerResultModel{
		ClientId:                "TestClientId",
		ProjectId:               "TestProjectId",
		CenterOfDifficultyLevel: 8,
		RangeCount:              2,
	}
	difficultyServerResultResponseModel := models.ChurnBlockerResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	responseModel, _ := (*churnBlockerManager.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
	testWebsocket.On("SendMessageToClient",
		responseModel,
		"TestClientId",
		"TestProjectId",
		"ChurnBlockerResultChannel").Return(nil)

	rawModel, _ := (*churnBlockerManager.JsonParser).EncodeJson(&m)

	//Act
	success, err := churnBlockerManager.SendMessageToClient(rawModel)

	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_SendMessageToClient_SuccessIsFalse(t *testing.T) {
	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket

	var churnBlockerManager = concrete.ChurnBlockerManagerConstructor()

	m := models.ChurnBlockerResultModel{
		ClientId:                "TestClientId",
		ProjectId:               "TestProjectId",
		CenterOfDifficultyLevel: 8,
		RangeCount:              2,
	}
	difficultyServerResultResponseModel := models.ChurnBlockerResultDto{
		CenterOfDifficultyLevel: m.CenterOfDifficultyLevel,
		RangeCount:              m.RangeCount,
	}
	responseModel, _ := (*churnBlockerManager.JsonParser).EncodeJson(&difficultyServerResultResponseModel)
	testWebsocket.On("SendMessageToClient",
		responseModel,
		"TestClientId",
		"TestProjectId",
		"ChurnBlockerResultChannel").Return(errors.New("FakeError"))

	rawModel, _ := (*churnBlockerManager.JsonParser).EncodeJson(&m)

	//Act
	success, err := churnBlockerManager.SendMessageToClient(rawModel)

	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)

}
