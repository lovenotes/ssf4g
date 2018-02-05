package tlog

import (
	"errors"
	"fmt"

	"ssf4g/libs/config"

	"github.com/getsentry/raven-go"
)

var (
	_logger *Logger

	_sentry_switch bool
	_sentry_client *raven.Client
)

func init() {
	tlogLevel := LOG_LEVEL_DEBUG

	runMode := config.GetIniData().String("run_mode")

	if runMode == "prod" {
		tlogLevel = LOG_LEVEL_INFO
	}

	logPath := config.GetIniData().String("log_path")

	_logger = NewLogger(logPath, tlogLevel)

	sentryDsn := config.GetIniData().String("sentry_dsn")
	_sentry_switch = false

	if runMode == "prod" {
		_sentry_switch = true
	}

	if _sentry_switch == true {
		var err error

		_sentry_client, err = raven.New(sentryDsn)

		if err != nil {
			_std_error.Printf("init sentry (%s) err (%v).\n", sentryDsn, err)
		}

		_sentry_client.SetTagsContext(map[string]string{
			"service_name": config.GetIniData().String("app_name"),
		})
	}
}

func Debug(format string, v ...interface{}) string {
	log := fmt.Sprintln(fmt.Sprintf(format, v...))

	_logger.Debug(log)

	return log
}

func Info(format string, v ...interface{}) string {
	log := fmt.Sprintln(fmt.Sprintf(format, v...))

	_logger.Info(log)

	return log
}

func Warn(format string, v ...interface{}) string {
	log := fmt.Sprintln(fmt.Sprintf(format, v...))

	_logger.Warn(log)

	return log
}

func Error(format string, v ...interface{}) string {
	log := fmt.Sprintln(fmt.Sprintf(format, v...))

	_logger.Error(log)

	return log
}

func AsyncSend(errdata *ErrData) {
	if _sentry_switch == true {
		captureError(errdata, raven.ERROR)
	}
}

func captureError(errdata *ErrData, level raven.Severity) *LogEvent {
	stackTrack := errdata._stack_track
	request := errdata._request
	err := errdata.Error()

	var packet *raven.Packet

	if request != nil {
		exceptionMsg := request.URL.Path + ": " + err.Error()
		exception := raven.NewException(errors.New(exceptionMsg), stackTrack)
		exception.Type = exceptionMsg

		packet = raven.NewPacket(errdata._err_msg[0], exception, raven.NewHttp(request))
	} else {
		exceptionMsg := "Error: " + err.Error()
		exception := raven.NewException(errors.New(exceptionMsg), stackTrack)
		exception.Type = exceptionMsg

		packet = raven.NewPacket(errdata._err_msg[0], exception)
	}

	packet.Level = level

	for index, errMsg := range errdata._err_msg {
		packet.Extra[fmt.Sprintf("debug.No%d", index)] = errMsg
	}

	eventID, ch := _sentry_client.Capture(packet, nil)

	return NewLogEvent(eventID, ch)
}
