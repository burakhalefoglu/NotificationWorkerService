package test

import (
	"NotificationWorkerService/internal/manager/concrete"
	"NotificationWorkerService/internal/models"
	"NotificationWorkerService/pkg/jsonParser/gojson"
	"NotificationWorkerService/test/Mock/mockwebsocket"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_RemoteOfferSendMessageToClient_SuccessIsTrue(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var interstitial = concrete.RemoteOfferManager{
		JsonParser:      &gojson.GoJson{},
		WebSocket: testWebsocket,
	}
	m := models.RemoteOfferModel{
		ClientIdList: []string{
			"TestId1",
			"TestId2",
		},
		ProjectId:   "TestProjectId",
		ProductModel: nil,
		FirstPrice:  12,
		LastPrice:   8,
		OfferId:     1,
		IsGift:      false,
		GiftTexture: nil,
		StartTime:   time.Time{},
		FinishTime:  time.Time{},
	}

	remoteOfferDto := models.RemoteOfferDto{
		ProductModel: nil,
		FirstPrice:  12,
		LastPrice:   8,
		OfferId:     1,
		IsGift:      false,
		GiftTexture: nil,
		StartTime:   time.Time{},
		FinishTime:  time.Time{},
	}
	responseModel, _ := interstitial.JsonParser.EncodeJson(&remoteOfferDto)

	for _, clientId := range m.ClientIdList {
		testWebsocket.On("SendMessageToClient",
			responseModel,
			clientId,
			"TestProjectId",
			"RemoteOfferChannel").Return(nil)

	}
	rawModel, _ := interstitial.JsonParser.EncodeJson(&m)

	//Act
	success, err:= interstitial.SendMessageToClient(rawModel)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_RemoteOfferSendMessageToClient_SuccessIsFalse(t *testing.T){

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var interstitial = concrete.RemoteOfferManager{
		JsonParser:      &gojson.GoJson{},
		WebSocket: testWebsocket,
	}
	m := models.RemoteOfferModel{
		ClientIdList: []string{
			"TestId1",
			"TestId2",
		},
		ProjectId:   "TestProjectId",
		ProductModel: nil,
		FirstPrice:  12,
		LastPrice:   8,
		OfferId:     1,
		IsGift:      false,
		GiftTexture: nil,
		StartTime:   time.Time{},
		FinishTime:  time.Time{},
	}

	remoteOfferDto := models.RemoteOfferDto{
		ProductModel: nil,
		FirstPrice:  12,
		LastPrice:   8,
		OfferId:     1,
		IsGift:      false,
		GiftTexture: nil,
		StartTime:   time.Time{},
		FinishTime:  time.Time{},
	}
	responseModel, _ := interstitial.JsonParser.EncodeJson(&remoteOfferDto)

	for _, clientId := range m.ClientIdList {
		testWebsocket.On("SendMessageToClient",
			responseModel,
			clientId,
			"TestProjectId",
			"RemoteOfferChannel").Return(errors.New("fakeError"))

	}
	rawModel, _ := interstitial.JsonParser.EncodeJson(&m)

	//Act
	success, err:= interstitial.SendMessageToClient(rawModel)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "fakeError", err)
}


