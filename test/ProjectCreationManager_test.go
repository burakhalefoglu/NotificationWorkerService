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

func Test_ProjectCreationSendMessageToClient_SuccessIsTrue(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var testLog = new (mocklog.MockLogger)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket
	IoC.Logger = testLog

	var projectCreation = concrete.ProjectCreationManagerConstructor()

	m := models.ProjectCreationResultModel{
		CustomerId:                "TestCustomerId",
		ProjectId:               "TestProjectId",
		Token: "FakeToken",
	}
	projectCreationResultDto := models.ProjectCreationResultDto{
		Token: "FakeToken",
	}
	responseModel, _ := (*projectCreation.JsonParser).EncodeJson(&projectCreationResultDto)
	testWebsocket.On("SendMessageToCustomer",
		responseModel,
		"TestCustomerId",
		"TestProjectId",
		"ProjectCreationResultChannel").Return(nil)

	rawModel, _ :=(*projectCreation.JsonParser).EncodeJson(&m)

	//Act
	success, err:= projectCreation.SendMessageToCustomer(rawModel)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_ProjectCreationSendMessageToClient_SuccessIsFalse(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var testLog = new (mocklog.MockLogger)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket
	IoC.Logger = testLog

	var projectCreation = concrete.ProjectCreationManagerConstructor()

	m := models.ProjectCreationResultModel{
		CustomerId:                "TestCustomerId",
		ProjectId:               "TestProjectId",
		Token: "FakeToken",
	}
	projectCreationResultDto := models.ProjectCreationResultDto{
		Token: "FakeToken",
	}
	responseModel, _ := (*projectCreation.JsonParser).EncodeJson(&projectCreationResultDto)
	testWebsocket.On("SendMessageToCustomer",
		responseModel,
		"TestCustomerId",
		"TestProjectId",
		"ProjectCreationResultChannel").Return(errors.New("fakeError"))

	rawModel, _ := (*projectCreation.JsonParser).EncodeJson(&m)

	//Act
	success, err:= projectCreation.SendMessageToCustomer(rawModel)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "fakeError", err)

}
