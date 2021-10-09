package fastHttpServer

import (
	models "NotificationWorkerService/internal/Models"
	"NotificationWorkerService/internal/websocket/fasthttp/ChannelWorker"
	"log"

	"github.com/valyala/fasthttp"
)

type FasthttpConn struct {
	ConnString string
}

func (f FasthttpConn) ListenServer() {
	h := requestHandler
	h = fasthttp.CompressHandler(h)

	fastHttpServer := fasthttp.Server{
		Name:    "NotificationWorkerService",
		Handler: h,
	}
	log.Fatal(fastHttpServer.ListenAndServe(f.ConnString))
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	req := &ctx.Request
	id := req.URI().QueryArgs().Peek("clientId")
	projectKey := req.URI().QueryArgs().Peek("projectId")

	// TODO: Add socket channel to listen on case
	switch string(ctx.Path()) {
	case "/RemoteOfferModel":
		ChannelWorker.Work(ctx, string(id), string(projectKey), &models.RemoteOfferChannelModel)
	case "InterstielAdModel":
		ChannelWorker.Work(ctx, string(id), string(projectKey), &models.InterstielAdChannelModel)
	case "DifficultyServerResultModel":
		ChannelWorker.Work(ctx, string(id), string(projectKey), &models.DifficultyServerResultChannelModel)

	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}

}
