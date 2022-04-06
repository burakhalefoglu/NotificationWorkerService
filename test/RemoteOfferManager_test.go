package test

import (
	"NotificationWorkerService/internal/IoC"
	"NotificationWorkerService/internal/manager/concrete"
	"NotificationWorkerService/internal/models"
	"NotificationWorkerService/pkg/jsonParser/gojson"
	"NotificationWorkerService/test/Mock/mockwebsocket"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_RemoteOfferSendMessageToClient_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket

	var remoteOffer = concrete.RemoteOfferManagerConstructor()

	m := models.RemoteOfferModel{
		ClientIdList: []string{
			"TestId1",
			"TestId2",
		},
		ProjectId:    "TestProjectId",
		ProductModel: nil,
		FirstPrice:   12,
		LastPrice:    8,
		OfferId:      1,
		IsGift:       false,
		GiftTexture:  nil,
		StartTime:    time.Time{},
		FinishTime:   time.Time{},
	}

	remoteOfferDto := models.RemoteOfferDto{
		ProductModel: nil,
		FirstPrice:   12,
		LastPrice:    8,
		OfferId:      1,
		IsGift:       false,
		GiftTexture:  nil,
		StartTime:    time.Time{},
		FinishTime:   time.Time{},
	}
	responseModel, _ := (*remoteOffer.JsonParser).EncodeJson(&remoteOfferDto)

	for _, clientId := range m.ClientIdList {
		testWebsocket.On("SendMessageToClient",
			responseModel,
			clientId,
			"TestProjectId",
			"RemoteOfferChannel").Return(nil)

	}
	rawModel, _ := (*remoteOffer.JsonParser).EncodeJson(&m)

	//Act
	success, err := remoteOffer.SendMessageToClient(rawModel)

	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)

}

func Test_RemoteOfferSendMessageToClient_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testWebsocket = new(mockwebsocket.MockWebSocket)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.WebSocket = testWebsocket

	var remoteOffer = concrete.RemoteOfferManagerConstructor()

	m := models.RemoteOfferModel{
		ClientIdList: []string{
			"TestId1",
			"TestId2",
		},
		ProjectId:    "TestProjectId",
		ProductModel: nil,
		FirstPrice:   12,
		LastPrice:    8,
		OfferId:      1,
		IsGift:       false,
		GiftTexture:  nil,
		StartTime:    time.Time{},
		FinishTime:   time.Time{},
	}

	remoteOfferDto := models.RemoteOfferDto{
		ProductModel: nil,
		FirstPrice:   12,
		LastPrice:    8,
		OfferId:      1,
		IsGift:       false,
		GiftTexture:  nil,
		StartTime:    time.Time{},
		FinishTime:   time.Time{},
	}
	responseModel, _ := (*remoteOffer.JsonParser).EncodeJson(&remoteOfferDto)

	for _, clientId := range m.ClientIdList {
		testWebsocket.On("SendMessageToClient",
			responseModel,
			clientId,
			"TestProjectId",
			"RemoteOfferChannel").Return(errors.New("fakeError"))

	}
	rawModel, _ := (*remoteOffer.JsonParser).EncodeJson(&m)

	//Act
	success, err := remoteOffer.SendMessageToClient(rawModel)

	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "fakeError", err)
}
