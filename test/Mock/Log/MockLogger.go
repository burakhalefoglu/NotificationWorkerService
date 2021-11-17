package mocklog

import "github.com/stretchr/testify/mock"

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger)SendInfoLog(parentStruct string, methodInfo string, message ...interface{}){}
func (m *MockLogger)SendTraceLog(parentStruct string, methodInfo string, message ...interface{}){}
func (m *MockLogger)SendDebugLog(parentStruct string, methodInfo string, message ...interface{}){}
func (m *MockLogger)SendWarnLog(parentStruct string, methodInfo string, message ...interface{}){}
func (m *MockLogger)SendErrorLog(parentStruct string, methodInfo string, message ...interface{}){}
func (m *MockLogger)SendFatalLog(parentStruct string, methodInfo string, message ...interface{}){}
func (m *MockLogger)SendPanicLog(parentStruct string, methodInfo string, message ...interface{}){}
func (m *MockLogger)SendInfofLog(parentStruct string, methodInfo string, format string, message ...interface{}){}
func (m *MockLogger)SendTracefLog(parentStruct string, methodInfo string, format string, message ...interface{}){}
func (m *MockLogger)SendDebugfLog(parentStruct string, methodInfo string, format string, message ...interface{}){}
func (m *MockLogger)SendWarnfLog(parentStruct string, methodInfo string, format string, message ...interface{}){}
func (m *MockLogger)SendErrorfLog(parentStruct string, methodInfo string, format string, message ...interface{}){}
func (m *MockLogger)SendFatalfLog(parentStruct string, methodInfo string, format string, message ...interface{}){}
func (m *MockLogger)SendPanicfLog(parentStruct string, methodInfo string, format string, message ...interface{}){}