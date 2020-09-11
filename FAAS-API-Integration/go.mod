module faas-api-integration

replace coralogix => ../coralogix

go 1.12

require (
	coralogix v0.0.0-00010101000000-000000000000
	github.com/gocarina/gocsv v0.0.0-20190919154618-09be0c8175b6
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
)
