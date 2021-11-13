package fiber

import (
	"NotificationWorkerService/internal/websocket/fiber/ChannelWorker/ClientWorker"
	"NotificationWorkerService/internal/websocket/fiber/ChannelWorker/CustomerWorker"
	hub "NotificationWorkerService/internal/websocket/fiber/hub"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
	"os"
	"sync"
)

type fiberWebsocket struct {
	Channel map[string]*hub.Channel
}

var RemoteOfferChannelModel = hub.Channel{
	Name: "RemoteOfferEventModel",
}

var InterstitialAdChannelModel = hub.Channel{
	Name: "InterstitialAdEventModel",
}

var ChurnPredictionResultChannel = hub.Channel{
	Name: "ChurnPredictionResultModel",
}

var ChurnBlockerResultChannel = hub.Channel{
	Name: "ChurnBlockerResultModel",
}

var ProjectCreationResultChannel = hub.Channel{
	Name: "ProjectCreationResult",
}


var FiberServer = &fiberWebsocket{
	Channel: map[string]*hub.Channel{
		"RemoteOfferChannel": &RemoteOfferChannelModel,
		"InterstitialAdChannel": &InterstitialAdChannelModel,
		"ChurnPredictionResultChannel": &ChurnPredictionResultChannel,
		"ChurnBlockerResultChannel": &ChurnBlockerResultChannel,
		"ProjectCreationResultChannel": &ProjectCreationResultChannel,
	},
}

func (f *fiberWebsocket) ListenServer(wgGroup *sync.WaitGroup) {

	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	wgGroup.Add(5)
	go ClientWorker.StartClientListener(wgGroup,
		app,
		FiberServer.Channel["RemoteOfferChannel"])
	go ClientWorker.StartClientListener(wgGroup,
		app,
		FiberServer.Channel["InterstitialAdChannel"])
	go ClientWorker.StartClientListener(wgGroup,
		app,
		FiberServer.Channel["ChurnPredictionResultChannel"])
	go ClientWorker.StartClientListener(wgGroup,
		app,
		FiberServer.Channel["ChurnBlockerResultChannel"])
	go CustomerWorker.StartCustomerListener(wgGroup,
		app,
		FiberServer.Channel["ProjectCreationResultChannel"])

	log.Fatal(app.Listen(os.Getenv("WEBSOCKET_CONN")))
}


func (f *fiberWebsocket) SendMessageToClient(message *[]byte, clientId string,
	projectId string,
	channelName string) error{

	for _, v := range f.Channel[channelName].Clients {
		if v.Id == clientId && v.ProjectId == projectId {
			log.Println(*message)
			err := v.Connection.WriteMessage(1, *message)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}

func (f *fiberWebsocket) SendMessageToAllClient(message *[]byte,
	projectId string,
	channelName string) error{

	for _, v := range f.Channel[channelName].Clients {
		if v.ProjectId == projectId {
			log.Println(*message)
			err := v.Connection.WriteMessage(1, *message)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}

func (f *fiberWebsocket) SendMessageToCustomer(message *[]byte, customerId string,
	projectId string,
	channelName string) error{

	for _, v := range f.Channel[channelName].Customers {
		if v.Id == customerId && v.ProjectId == projectId {
			err := v.Connection.WriteMessage(2, *message)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}

func (f *fiberWebsocket) SendMessageToAllCustomer(message *[]byte,
	projectId string,
	channelName string) error{

	for _, v := range f.Channel[channelName].Customers {
		if v.ProjectId == projectId {
			err := v.Connection.WriteMessage(2, *message)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}