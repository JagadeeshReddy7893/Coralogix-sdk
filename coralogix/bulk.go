package coralogix

import (
	"encoding/json"
	"os"
)

//Bulk Batch of logs to feed coralogix API
type Bulk struct {
	PrivateKey      string `json:"privateKey"`      // Coralogix private key
	ApplicationName string `json:"applicationName"` // Your application name
	SubsystemName   string `json:"subsystemName"`   // Subsystem name of your application
	ComputerName    string `json:"computerName"`    // Current machine hostname
	LogEntries      []Log  `json:"logEntries"`      // Log records list
}

//NewBulk Returns the instance of new Bulk
func NewBulk(AppCredentials Credentials) *Bulk {
	Hostname, _ := os.Hostname()
	return &Bulk{
		AppCredentials.PrivateKey,
		AppCredentials.ApplicationName,
		AppCredentials.SubsystemName,
		Hostname,
		[]Log{},
	}
}

//AddLog Adds a new log
func (B *Bulk) AddLog(LogRecord Log) {
	B.LogEntries = append(B.LogEntries, LogRecord)
}

//ToBytes Marshals the Payload
func (B *Bulk) ToBytes() []byte {
	DataBytes, _ := json.Marshal(B)

	return DataBytes
}
