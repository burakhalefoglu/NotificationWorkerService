package logrus_logstash_hook

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

type logrusToLogstashLOG struct {
	Log *logrus.Logger
}

func LogrusToLogstashLOGConstructor() *logrusToLogstashLOG {
	return &logrusToLogstashLOG{Log: connectLogstash()}
}

func connectLogstash() *logrus.Logger{
	log := logrus.New()
	conn, err := net.Dial("tcp", os.Getenv("LOGSTASH_HOST"))
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "StorageWorkerService"}))
	log.Hooks.Add(hook)
	return  log
}


func (l *logrusToLogstashLOG) SendInfoLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Info(message)
}

func (l *logrusToLogstashLOG)SendTraceLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Trace(message)
}

func (l *logrusToLogstashLOG)SendDebugLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Debug(message)
}

func (l *logrusToLogstashLOG)SendWarnLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Warn(message)
}

func (l *logrusToLogstashLOG)SendErrorLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Error(message)
}

func (l *logrusToLogstashLOG)SendFatalLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Fatal(message)
}

func (l *logrusToLogstashLOG)SendPanicLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Panic(message)
}

func (l *logrusToLogstashLOG) SendInfofLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Infof(format,message)
}

func (l *logrusToLogstashLOG)SendTracefLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Tracef(format, message)
}

func (l *logrusToLogstashLOG)SendDebugfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Debugf(format, message)
}

func (l *logrusToLogstashLOG)SendWarnfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Warnf(format, message)
}

func (l *logrusToLogstashLOG)SendErrorfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Errorf(format, message)
}

func (l *logrusToLogstashLOG)SendFatalfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Fatalf(format, message)
}

func (l *logrusToLogstashLOG)SendPanicfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Panicf(format, message)
}

