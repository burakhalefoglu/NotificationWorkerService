package abstract

type IChurnBlockerService interface {
	SendMessageToClient(data *[]byte)(success bool, message string)
}
