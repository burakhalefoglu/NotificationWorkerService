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

func Test_InterstitialSendMessageToClient_SuccessIsTrue(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var testLog = new (mocklog.MockLogger)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket
	IoC.Logger = testLog

	var interstitialM =concrete.InterstitialManagerConstructor()

	m := models.InterstitialAdModel{
		ClientIdList:           []string{
			"TestId1",
			"TestId2",
		},
		ProjectId:              "TestProjectId",
		IsAdvSettingsActive:    false,
		AdvFrequencyStrategies: map[string]int{
			"strategy1":1,
			"strategy2":2,
		},
	}
	interstitialDto := models.InterstitialAdDto{
		IsAdvSettingsActive:    false,
		AdvFrequencyStrategies: map[string]int{
			"strategy1":1,
			"strategy2":2,
		},
	}
	responseModel, _ := (*interstitialM.JsonParser).EncodeJson(&interstitialDto)

	for _, clientId := range m.ClientIdList {
		testWebsocket.On("SendMessageToClient",
			responseModel,
			clientId,
			"TestProjectId",
			"InterstitialAdChannel").Return(nil)

	}
	rawModel, _ := (*interstitialM.JsonParser).EncodeJson(&m)

	//Act
	success, err:= interstitialM.SendMessageToClient(rawModel)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_InterstitialSendMessageToClient_SuccessIsFalse(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var testLog = new (mocklog.MockLogger)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket
	IoC.Logger = testLog

	var interstitialM =concrete.InterstitialManagerConstructor()

	m := models.InterstitialAdModel{
		ClientIdList:           []string{
			"TestId1",
			"TestId2",
		},
		ProjectId:              "TestProjectId",
		IsAdvSettingsActive:    false,
		AdvFrequencyStrategies: map[string]int{
			"strategy1":1,
			"strategy2":2,
		},
	}
	interstitialDto := models.InterstitialAdDto{
		IsAdvSettingsActive:    false,
		AdvFrequencyStrategies: map[string]int{
			"strategy1":1,
			"strategy2":2,
		},
	}
	responseModel, _ := (*interstitialM.JsonParser).EncodeJson(&interstitialDto)

	for _, clientId := range m.ClientIdList {
		testWebsocket.On("SendMessageToClient",
			responseModel,
			clientId,
			"TestProjectId",
			"InterstitialAdChannel").Return(errors.New("fakeError"))

	}
	rawModel, _ := (*interstitialM.JsonParser).EncodeJson(&m)

	//Act
	success, err:= interstitialM.SendMessageToClient(rawModel)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "fakeError", err)
}


