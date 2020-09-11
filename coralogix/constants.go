package coralogix

const (
	//Debug Debug level
	Debug = 1

	//Verbose Verbose level
	Verbose = 2

	//Info Info level
	Info = 3

	//Warn Warn level
	Warn = 4

	//Error Error level
	Error = 5

	//Critical Critical level
	Critical = 6

	//POSTCoralogixLogsEndpoint Endpoint to post logs into the coralogix
	POSTCoralogixLogsEndpoint = "https://api.app.coralogix.in/api/v1/logs"

	//MaxLogChunkSize Max size of log should be 1.5MB
	MaxLogChunkSize int64 = 1.5 * (1024 * 1024)
)
