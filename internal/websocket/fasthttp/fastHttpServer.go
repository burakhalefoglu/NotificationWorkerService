package fastHttpServer

import (
	"NotificationWorkerService/internal/websocket/fasthttp/ChannelWorker/ClientWorker"
	"NotificationWorkerService/internal/websocket/fasthttp/ChannelWorker/CustomerWorker"
	"log"
	"sync"

	"github.com/valyala/fasthttp"
)

type FastHttp struct {
	ConnString string
	Channel map[string]*Channel
}

var FastHttpServer = &FastHttp{
ConnString : "localhost:8081",
}

var RemoteOfferChannelModel = Channel{
	Name: "RemoteOfferModel",
}

var InterstitialAdChannelModel = Channel{
	Name: "InterstitialAdModel",
}

var ChurnPredictionResultChannel = Channel{
	Name: "ChurnPredictionResultChannel",
}

var ChurnBlockerResultChannel = Channel{
	Name: "ChurnBlockerResultChannel",
}

var ProjectCreationResultChannel = Channel{
	Name: "ProjectCreationResultChannel",
}


var FastHttpModel = &FastHttp{
	Channel: map[string]*Channel{
		"RemoteOfferChannel": &RemoteOfferChannelModel,
		"InterstitialAdChannel": &InterstitialAdChannelModel,
		"ChurnPredictionResultChannel": &ChurnPredictionResultChannel,
		"ChurnBlockerResultChannel": &ChurnBlockerResultChannel,
		"ProjectCreationResultChannel": &ProjectCreationResultChannel,
	},
}

var wgGroup *sync.WaitGroup

func assignAndGetWaitGroup(group *sync.WaitGroup) *sync.WaitGroup{
	wgGroup := group
	return wgGroup
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	id := req.URI().QueryArgs().Peek("clientId")
	projectKey := req.URI().QueryArgs().Peek("projectId")

	wgGroup.Add(5)
	switch string(ctx.Path()) {
	case "/RemoteOfferModel":
		go ClientWorker.StartClientListener(ctx, string(id), string(projectKey),
			FastHttpModel.Channel["RemoteOfferChannel"])
	case "/InterstitialAdModel":
		go ClientWorker.StartClientListener(ctx, string(id), string(projectKey),
			FastHttpModel.Channel["InterstitialAdChannel"])
	case "/ChurnPredictionResultModel":
		go ClientWorker.StartClientListener(ctx, string(id), string(projectKey),
			FastHttpModel.Channel["ChurnPredictionResultChannel"])
	case "/ChurnBlockerResultModel":
		go ClientWorker.StartClientListener(ctx, string(id), string(projectKey),
			FastHttpModel.Channel["ChurnBlockerResultChannel"])
	case "/ProjectCreationResultChannel":
		go CustomerWorker.StartCustomerListener(ctx, string(id), string(projectKey),
			FastHttpModel.Channel["ProjectCreationResultChannel"])
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
	wgGroup.Done()
}

func (f *FastHttp) ListenServer(group *sync.WaitGroup) {
	assignAndGetWaitGroup(group)
	h := requestHandler
	h = fasthttp.CompressHandler(h)
	fastHttpServer := fasthttp.Server{
		Name:    "NotificationWorkerService",
		Handler: h,
	}
	log.Fatal(fastHttpServer.ListenAndServe(f.ConnString))
}

func (f *FastHttp) SendMessageToClient(message *[]byte, clientId string,
	projectId string,
	channelName string){

	for _, v := range f.Channel[channelName].Clients {
		if v.Id == clientId && v.ProjectId == projectId {
			err := v.Connection.WriteMessage(2, *message)
			if err != nil {
				log.Fatal(err)
				return 
			}
		}
	}
}

func (f *FastHttp) SendMessageToAllClient(message *[]byte,
	projectId string,
	channelName string){

	for _, v := range f.Channel[channelName].Clients {
		if v.ProjectId == projectId {
			err := v.Connection.WriteMessage(2, *message)
			if err != nil {
				log.Fatal(err)
				return 
			}
		}
	}
}

func (f *FastHttp) SendMessageToCustomer(message *[]byte, customerId string,
	projectId string,
	channelName string){

	for _, v := range f.Channel[channelName].Customers {
		if v.Id == customerId && v.ProjectId == projectId {
			err := v.Connection.WriteMessage(2, *message)
			if err != nil {
				log.Fatal(err)
				return 
			}
		}
	}
}

func (f *FastHttp) SendMessageToAllCustomer(message *[]byte,
	projectId string,
	channelName string){

	for _, v := range f.Channel[channelName].Customers {
		if v.ProjectId == projectId {
			err := v.Connection.WriteMessage(2, *message)
			if err != nil {
				log.Fatal(err)
				return 
			}
		}
	}
}