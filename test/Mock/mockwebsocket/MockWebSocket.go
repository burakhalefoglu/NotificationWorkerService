package mockwebsocket

import (
	"github.com/stretchr/testify/mock"
	"sync"
)

type MockWebSocket struct {
	mock.Mock
}
func (m *MockWebSocket) ListenServer(group *sync.WaitGroup) {
}

func (m *MockWebSocket) SendMessageToClient(message *[]byte, clientId string,
	projectId string,
	channelName string) error{
	args := m.Called(message, clientId, projectId, channelName)
	return  args.Error(0)
}

func (m *MockWebSocket) SendMessageToAllClient(message *[]byte,
	projectId string,
	channelName string) error{
	args := m.Called(message, projectId, channelName)
	return  args.Error(0)
}

func (m *MockWebSocket) SendMessageToCustomer(message *[]byte, customerId string,
	projectId string,
	channelName string) error{
	args := m.Called(message, customerId, projectId, channelName)
	return  args.Error(0)
}

func (m *MockWebSocket) SendMessageToAllCustomer(message *[]byte,
	projectId string,
	channelName string) error{
	args := m.Called(message, projectId, channelName)
	return  args.Error(0)
}


