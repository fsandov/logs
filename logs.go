package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Service is the struct that contains the configuration of the logs. It is used to create a new instance of the service.
// All the fields are optional.
type Service struct {
	// NameApp is the name of the application that will be used in the logs.
	NameApp string
	// URL is the URL of the webhook to send the logs. If it is not provided, the logs will be printed in the console.
	// The URL could be a Discord webhook URL. Example: https://discordapp.com/api/webhooks/1234567890/abcdefghijklmnopqrstuvwxyz
	// The logs will be sent as a Discord message.
	URL string
	// FileLog is a boolean that indicates if the logs should be saved in a file. If it is true, the logs will be saved in a file.
	// The file will be created in the same folder where the application is running. The name of the file will be the name of the application.
	// If the name of the application is not provided, the name of the file will be "logs".
	FileLog bool
	// fileName is the name of the file where the logs will be saved. It is used internally.
	fileName string
}

// NewService returns a new instance of a Service of logs with the configuration provided.
func NewService(config Service) Service {
	if config.NameApp == "" {
		config.NameApp = "LOGS"
	}
	var fileName string
	if config.FileLog {
		fileName = fmt.Sprintf("%s/%s-%s.log", pathLogs, config.NameApp, time.Now().Format("2006-01-02"))
	}

	return Service{
		NameApp:  config.NameApp,
		URL:      config.URL,
		FileLog:  config.FileLog,
		fileName: fileName,
	}
}

var (
	// DefaultService is the default instance of the service.
	// It is used to call the functions of the service without creating a new instance.
	// It is initialized with the default configuration.
	DefaultService = NewService(Service{
		NameApp:  "LOGS",
		FileLog:  true,
		fileName: fmt.Sprintf("%s/%s-%s.log", pathLogs, "LOGS", time.Now().Format("2006-01-02")),
	})
)

const (
	// logTrace is a log level used for tracing the code and executing steps.
	// This is the most verbose level. it should not be used in production.
	logTrace = "TRACE"
	// logDebug is a log level used for debugging the code.
	// It should not be used in production.
	logDebug = "DEBUG"
	// logInfo is a log level used for providing information about the execution of the code.
	// It should be used in production.
	logInfo = "INFO"
	// logNotice is a log level used for all the notable events that are not considered an error.
	// It should be used in production.
	logNotice = "NOTICE"
	// logWarning is a log level used for all the events that can potentially cause application oddities.
	// It should be used in production.
	logWarning = "WARNING"
	// logError is a log level used for all the errors that are not critical and the application can continue running.
	// It should be used in production.
	logError = "ERROR"
	// logFatal is a log level used for all the errors that are critical and may cause the application to stop running.
	// It should be used in production.
	logFatal      = "FATAL"
	caller        = "3"
	callerDefault = "4"
	pathLogs      = "logs"
)

// logBuilder is the function that builds the logs. It is used internally. It receives the content of the log.
// It returns the content of the log built. This is used to orchestrate the type of logs.
func (s Service) logBuilder(mL string, callerLevel string, message string, extraMessage ...string) string {
	switch mL {
	case logInfo:
		extraMessageSTR := ""
		if callerLevel == caller {
			extraMessageSTR = strings.Join(extraMessage, " ")
		}
		msg := fmt.Sprintf(message + " " + extraMessageSTR)
		return fmt.Sprintf(
			"[%s][%s][%s]-[%s] %s",
			time.Now().Format("2006-01-02"),
			time.Now().Format("15:04:05.999"),
			s.NameApp,
			mL,
			msg)
	default:
		return s.logDecorator(mL, callerLevel, message, extraMessage...)
	}
}

// logDecorator is the function that decorates the logs. It is used internally. It receives the content of the log.
// It returns the content of the log decorated.
func (s Service) logDecorator(mL string, callerLevel string, message string, extraMessage ...string) string {
	callerLevelINT, err := strconv.Atoi(callerLevel)
	if err != nil {
		callerLevelINT, _ = strconv.Atoi(caller)
	}
	pc, file, line, _ := runtime.Caller(callerLevelINT)
	files := strings.Split(file, "/")
	file = files[len(files)-1]
	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]
	extraMessageSTR := ""
	if callerLevel == caller {
		extraMessageSTR = strings.Join(extraMessage, " ")
	}
	msg := fmt.Sprintf(message + " " + extraMessageSTR)
	return fmt.Sprintf(
		"[%s][%s][%s]-[%s] %s:%d:%s(): %s",
		time.Now().Format("2006-01-02"),
		time.Now().Format("15:04:05.999"),
		s.NameApp,
		mL,
		file,
		line,
		name,
		msg)
}

// registerOrchestrator is the function that registers the logs in the different services. It is used internally.
// It is called by the functions of the service. It receives the content of the log.
func (s Service) registerOrchestrator(content string) {
	if s.URL != "" {
		s.postLog(content)
	}
	if s.FileLog {
		s.registerFileLog(content)
	}

}

// registerFileLog saves the logs in a file. The function creates a folder called "logs" in the same folder where the application is running.
// The name of the file will be the name of the application. If the name of the application is not provided, the name of the file will be "LOGS".
// The logs will be saved in the file with the date of the day.
func (s Service) registerFileLog(content string) {
	_, err := os.Stat(pathLogs)
	if os.IsNotExist(err) {
		err = os.Mkdir(pathLogs, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	file, err := os.OpenFile(s.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	mw := io.MultiWriter(os.Stdout, file)
	_, err = fmt.Fprintln(mw, content)
	if err != nil {
		return
	}

}

// postLog send the log to the URL specified in the configuration. It is used to send the logs to a server.
func (s Service) postLog(content string) {
	value := map[string]string{"content": content}
	jsonData, err := json.Marshal(value)
	if err != nil {
		fmt.Println("error marshaling SendInfo:", err)
		return
	}
	_, err = http.Post(s.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
}

// Trace is the function that registers the logs with the level Trace. It receives the content of the log.
// It can receive extra information. It is optional.
func (s Service) Trace(message string, extraMessage ...string) {
	s.registerOrchestrator(s.logBuilder(logTrace, getLevelCaller(extraMessage...), message, extraMessage...))
}

// Debug is the function that registers the logs with the level Debug. It receives the content of the log.
// It can receive extra information. It is optional.
func (s Service) Debug(message string, extraMessage ...string) {
	s.registerOrchestrator(s.logBuilder(logDebug, getLevelCaller(extraMessage...), message, extraMessage...))
}

// Info is the function that registers the logs with the level Info. It receives the content of the log.
// It can receive extra information. It is optional.
func (s Service) Info(message string, extraMessage ...string) {
	s.registerOrchestrator(s.logBuilder(logInfo, getLevelCaller(extraMessage...), message, extraMessage...))
}

// Notice is the function that registers the logs with the level Notice. It receives the content of the log.
// It can receive extra information. It is optional.
func (s Service) Notice(message string, extraMessage ...string) {
	s.registerOrchestrator(s.logBuilder(logNotice, getLevelCaller(extraMessage...), message, extraMessage...))
}

// Warning is the function that registers the logs with the level Warning. It receives the content of the log.
// It can receive extra information. It is optional.
func (s Service) Warning(message string, extraMessage ...string) {
	s.registerOrchestrator(s.logBuilder(logWarning, getLevelCaller(extraMessage...), message, extraMessage...))
}

// Error is the function that registers the logs with the level Error. It receives the content of the log.
// It can receive extra information. It is optional.
func (s Service) Error(message string, extraMessage ...string) {
	s.registerOrchestrator(s.logBuilder(logError, getLevelCaller(extraMessage...), message, extraMessage...))
}

// Fatal is the function that registers the logs with the level Fatal. It receives the content of the log.
// It can receive extra information. It is optional.
func (s Service) Fatal(message string, extraMessage ...string) {
	s.registerOrchestrator(s.logBuilder(logFatal, getLevelCaller(extraMessage...), message, extraMessage...))
}

// Trace is the function that registers the logs with the level Trace. It receives the content of the log.
// It can receive extra information. It is optional. The logs will be registered in the DefaultService.
func Trace(message string) {
	DefaultService.Trace(message, callerDefault)
}

// Debug is the function that registers the logs with the level Debug. It receives the content of the log.
// It can receive extra information. It is optional. The logs will be registered in the DefaultService.
func Debug(message string) {
	DefaultService.Debug(message, callerDefault)
}

// Info is the function that registers the logs with the level Info. It receives the content of the log.
// It can receive extra information. It is optional. The logs will be registered in the DefaultService.
func Info(message string) {
	DefaultService.Info(message, callerDefault)
}

// Notice is the function that registers the logs with the level Notice. It receives the content of the log.
// It can receive extra information. It is optional. The logs will be registered in the DefaultService.
func Notice(message string) {
	DefaultService.Notice(message, callerDefault)
}

// Warning is the function that registers the logs with the level Warning. It receives the content of the log.
// It can receive extra information. It is optional. The logs will be registered in the DefaultService.
func Warning(message string) {
	DefaultService.Warning(message, callerDefault)
}

// Error is the function that registers the logs with the level Error. It receives the content of the log.
// It can receive extra information. It is optional. The logs will be registered in the DefaultService.
func Error(message string) {
	DefaultService.Error(message, callerDefault)
}

// Fatal is the function that registers the logs with the level Fatal. It receives the content of the log.
// It can receive extra information. It is optional. The logs will be registered in the DefaultService.
func Fatal(message string) {
	DefaultService.Fatal(message, callerDefault)
}

// getLevelCaller is the function that returns the level of the caller. It is used internally.
func getLevelCaller(extraMessage ...string) string {
	if len(extraMessage) > 0 {
		if extraMessage[0] == callerDefault {
			return callerDefault
		}
	}
	return caller
}
