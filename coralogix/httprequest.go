package coralogix

import "fmt"

//SendRequest Makes HTTP request to write logs into the coralogix
func SendRequest(BulkPayload *Bulk) error {
	BulkPayloadBytes := BulkPayload.ToBytes()

	URL := POSTCoralogixLogsEndpoint

	StatusCode, Response, err := HTTPCall("POST", URL, BulkPayloadBytes, nil)
	CoralogixRetryCounter := 0
coralogix_retry:
	if err != nil || StatusCode != 200 {

		fmt.Println("HTTP call failed. Error: ", err, " StatusCode: ", StatusCode)
		fmt.Println("URL: ", URL)
		fmt.Println("Response: ", string(Response))
		CoralogixRetryCounter++
		if CoralogixRetryCounter < 3 {
			fmt.Println("coralogix api retry counter: ", CoralogixRetryCounter)
			goto coralogix_retry
		}
		RequestJSON, _ := RemovePrettyJSON(BulkPayloadBytes)
		fmt.Println("Request: ", RequestJSON)
		return err
	}

	fmt.Println("Pushed the logs successfully to coralogix")

	return nil
}
