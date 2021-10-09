package models

import (
	hub "NotificationWorkerService/pkg/FastHttp"
)

var RemoteOfferChannelModel = hub.Channel{
	Name: "RemoteOfferModel",
}

var InterstielAdChannelModel = hub.Channel{
	Name: "InterstielAdModel",
}

var DifficultyServerResultChannelModel = hub.Channel{
	Name: "DifficultyServerResultModel",
}
