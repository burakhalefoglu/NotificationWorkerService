package abstract

type IRemoteOfferService interface {
	SendMessageToClient(data *[]byte)(success bool, message string)
}
