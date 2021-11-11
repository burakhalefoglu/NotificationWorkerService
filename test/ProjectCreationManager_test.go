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

func Test_ProjectCreationSendMessageToClient_SuccessIsTrue(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var churnBlockerManager = concrete.ProjectCreationManager{
		JsonParser:      &gojson.GoJson{},
		WebSocket: testWebsocket,
	}
	m := models.ProjectCreationResultModel{
		CustomerId:                "TestCustomerId",
		ProjectId:               "TestProjectId",
		Token: "FakeToken",
	}
	projectCreationResultDto := models.ProjectCreationResultDto{
		Token: "FakeToken",
	}
	responseModel, _ := churnBlockerManager.JsonParser.EncodeJson(&projectCreationResultDto)
	testWebsocket.On("SendMessageToCustomer",
		responseModel,
		"TestCustomerId",
		"TestProjectId",
		"ProjectCreationResultChannel").Return(nil)

	rawModel, _ := churnBlockerManager.JsonParser.EncodeJson(&m)

	//Act
	success, err:= churnBlockerManager.SendMessageToCustomer(rawModel)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_ProjectCreationSendMessageToClient_SuccessIsFalse(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var churnBlockerManager = concrete.ProjectCreationManager{
		JsonParser:      &gojson.GoJson{},
		WebSocket: testWebsocket,
	}
	m := models.ProjectCreationResultModel{
		CustomerId:                "TestCustomerId",
		ProjectId:               "TestProjectId",
		Token: "FakeToken",
	}
	projectCreationResultDto := models.ProjectCreationResultDto{
		Token: "FakeToken",
	}
	responseModel, _ := churnBlockerManager.JsonParser.EncodeJson(&projectCreationResultDto)
	testWebsocket.On("SendMessageToCustomer",
		responseModel,
		"TestCustomerId",
		"TestProjectId",
		"ProjectCreationResultChannel").Return(errors.New("fakeError"))

	rawModel, _ := churnBlockerManager.JsonParser.EncodeJson(&m)

	//Act
	success, err:= churnBlockerManager.SendMessageToCustomer(rawModel)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "fakeError", err)

}
