package abstract


type IInterstitialAdService interface {
	SendMessageToClient(data *[]byte)(success bool, message string)
}