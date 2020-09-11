package coralogix

//LoggerManager Logger Manger Process the logs
type LoggerManager struct {
	AppCredentials Credentials
	LogRecords     []Log
	ClassName      string
	MethodName     string
	Category       string
}

//NewCoralogixLoggerManager Returns the instance of Logger Manager
func NewCoralogixLoggerManager(ApplicationName, PrivateKey, SubsystemName string, ClassName, MethodName, Category string) *LoggerManager {

	return &LoggerManager{
		Credentials{
			PrivateKey,
			ApplicationName,
			SubsystemName,
		},
		[]Log{},
		ClassName,
		MethodName,
		Category,
	}
}

//AddLogLine Adds the log to buffer
func (LM *LoggerManager) AddLogLine(LogRecord Log) {
	LM.LogRecords = append(LM.LogRecords, LogRecord)
}

//WriteLogs Writes logs to the coralogix
func (LM *LoggerManager) WriteLogs() error {
	var WindowSize int64
	WindowSize = 10
	LogRecordsLen := int64(len(LM.LogRecords))
	var Index int64
	for Index = 0; Index < LogRecordsLen; Index += WindowSize {
		Start := Index
		End := Index + WindowSize
		if End > LogRecordsLen {
			End = LogRecordsLen
		}
		BulkPayloadObj := NewBulk(LM.AppCredentials)
		BulkPayloadObj.LogEntries = LM.LogRecords[Start:End]

		err := SendRequest(BulkPayloadObj)
		if err != nil {
			return err
		}
	}

	return nil
}
