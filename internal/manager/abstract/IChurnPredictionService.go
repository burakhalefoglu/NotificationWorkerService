package abstract

type IChurnPredictionService interface {
	SendMessageToClient(data *[]byte)(success bool, message string)
}
