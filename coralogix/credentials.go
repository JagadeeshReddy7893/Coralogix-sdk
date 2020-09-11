package coralogix

//Credentials Credentials of an application
type Credentials struct {
	PrivateKey      string // Coralogix private key
	ApplicationName string // Your application name
	SubsystemName   string // Subsystem name of your application
}
