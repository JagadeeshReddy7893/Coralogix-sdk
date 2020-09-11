package coralogix

import (
	"encoding/json"
	"fmt"
	"time"
)

// Log describe record format for Coralogix API
type Log struct {
	Timestamp  float64 `json:"timestamp"`  // Log record timestamp
	Severity   uint    `json:"severity"`   // Log record severity level
	Text       string  `json:"text"`       // Log record message
	Category   string  `json:"category"`   // Log record category
	ClassName  string  `json:"className"`  // Log record class name
	MethodName string  `json:"methodName"` // Log record method name
	ThreadID   string  `json:"threadId"`   // Thread ID
}

//InitialiseLog Initialises the log
func (LM *LoggerManager) InitialiseLog() Log {
	LogRecord := Log{}
	LogRecord.Category = LM.Category
	LogRecord.ClassName = LM.ClassName
	LogRecord.MethodName = LM.MethodName
	LogRecord.Timestamp = float64(time.Now().UnixNano() / int64(time.Millisecond))

	return LogRecord
}

//Debug Adds debug logs
func (LM *LoggerManager) Debug(TextItems ...interface{}) {
	LogRecord := LM.InitialiseLog()
	LogRecord.Severity = Debug
	LogRecord.Text = LogRecord.ProcessText(TextItems)
	if (LogRecord.Size()) < MaxLogChunkSize {
		LM.LogRecords = append(LM.LogRecords, LogRecord)
	}
}

//Verbose Adds verbose logs
func (LM *LoggerManager) Verbose(TextItems ...interface{}) {
	LogRecord := LM.InitialiseLog()
	LogRecord.Severity = Verbose
	LogRecord.Text = LogRecord.ProcessText(TextItems)
	if LogRecord.Size() < MaxLogChunkSize {
		LM.LogRecords = append(LM.LogRecords, LogRecord)
	}
}

//Info Adds Info level logs
func (LM *LoggerManager) Info(TextItems ...interface{}) {
	LogRecord := LM.InitialiseLog()
	LogRecord.Severity = Info
	LogRecord.Text = LogRecord.ProcessText(TextItems)
	if LogRecord.Size() < MaxLogChunkSize {
		LM.LogRecords = append(LM.LogRecords, LogRecord)
	}
}

//Warn Adds Warn level logs
func (LM *LoggerManager) Warn(TextItems ...interface{}) {
	LogRecord := LM.InitialiseLog()
	LogRecord.Severity = Warn
	LogRecord.Text = LogRecord.ProcessText(TextItems)
	if LogRecord.Size() < MaxLogChunkSize {
		LM.LogRecords = append(LM.LogRecords, LogRecord)
	}
}

//Error Adds Error level logs
func (LM *LoggerManager) Error(TextItems ...interface{}) {
	LogRecord := LM.InitialiseLog()
	LogRecord.Severity = Error
	LogRecord.Text = LogRecord.ProcessText(TextItems)
	if LogRecord.Size() < MaxLogChunkSize {
		LM.LogRecords = append(LM.LogRecords, LogRecord)
	}
}

//Critical Adds Critical level logs
func (LM *LoggerManager) Critical(TextItems ...interface{}) {
	LogRecord := LM.InitialiseLog()
	LogRecord.Severity = Critical
	LogRecord.Text = LogRecord.ProcessText(TextItems)
	if LogRecord.Size() < MaxLogChunkSize {
		LM.LogRecords = append(LM.LogRecords, LogRecord)
	}
}

//ProcessText Process the log text
func (L *Log) ProcessText(TextItems ...interface{}) string {
	var TextString string

	if TextItems != nil {
		var TextItemsInterface interface{}
		TextItemsInterface = TextItems[0]
		TextItemsList := TextItemsInterface.([]interface{})

		for _, TextItemObj := range TextItemsList {
			TextString += fmt.Sprint("", TextItemObj)
		}
	}

	return TextString
}

//Size Calculates and returns the size of log in bytes
func (L *Log) Size() int64 {
	LogRecordBytes, _ := json.Marshal(L)

	return int64(len(string(LogRecordBytes)))
}
