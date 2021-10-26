package abstract

type IProjectCreationService interface {
	SendMessageToCustomer(data *[]byte)(success bool, message string)
}
