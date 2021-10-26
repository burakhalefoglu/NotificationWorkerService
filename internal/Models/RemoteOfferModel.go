package Models

import "time"

type ProductModel struct{

	Name string
	Image []byte
	Count float32
}

type RemoteOfferModel struct{
	ProductModel []*ProductModel
	ClientIdList []string
	ProjectId string
	FirstPrice float32
	LastPrice float32
	OfferId int
	IsGift bool
	GiftTexture []byte
	StartTime time.Time
	FinishTime time.Time
}

type RemoteOfferDto struct{
	ProductModel []*ProductModel
	FirstPrice float32
	LastPrice float32
	OfferId int
	IsGift bool
	GiftTexture []byte
	StartTime time.Time
	FinishTime time.Time
}
