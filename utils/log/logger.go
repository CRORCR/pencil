package log

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger *logrus.Logger

type Log struct {
}

func init() {
	Logger = logrus.New()
	Logger.SetReportCaller(false)
	Logger.SetFormatter(&FluentdFormatter{})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
	//Logger.AddHook(&ContextHook{})

}

/////////////////////////////////////////////////////////////////////////////////////
// -------------------------------- FluentdFormatter --------------------------------
/////////////////////////////////////////////////////////////////////////////////////

// Default logrus to FluentD severity map
var SeverityMap = map[string]string{
	"panic":   "fatal",
	"fatal":   "fatal",
	"warning": "warn",
	"debug":   "debug",
	"error":   "error",
	"trace":   "trace",
	"info":    "info",
}

// logrus to stackdriver severity map
func UseStackdriverSeverity() {
	SeverityMap = map[string]string{
		"panic":   "CRITICAL",
		"fatal":   "CRITICAL",
		"warning": "WARNING",
		"debug":   "DEBUG",
		"error":   "ERROR",
		"trace":   "DEBUG",
		"info":    "INFO",
	}
}

// FluentdFormatter is similar to logrus.JSONFormatter but with log level that are recongnized
// by kubernetes fluentd.
type FluentdFormatter struct {
	TimestampFormat string
	SeverityMap     map[string]string
}

// Format the log entry. Implements logrus.Formatter.
func (f *FluentdFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/Sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	prefixFieldClashes(data)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339Nano
	}

	data["timestamp"] = entry.Time.Format(timestampFormat)

	msg := ""

	if entry.HasCaller() && (entry.Level == logrus.ErrorLevel || entry.Level == logrus.WarnLevel || entry.Level == logrus.FatalLevel) {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		msg = fmt.Sprintf(`[file]%s,[func]%s,[msg]%s`, fileVal, funcVal, entry.Message)
	} else {
		msg = entry.Message
	}
	data["message"] = msg

	if ms, ok := SeverityMap[entry.Level.String()]; ok {
		data["severity"] = ms
	} else {
		data["severity"] = SeverityMap["debug"]
	}

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

func prefixFieldClashes(data logrus.Fields) {
	if t, ok := data["time"]; ok {
		data["fields.time"] = t
	}

	if m, ok := data["msg"]; ok {
		data["fields.msg"] = m
	}

	if l, ok := data["level"]; ok {
		data["fields.level"] = l
	}
}
