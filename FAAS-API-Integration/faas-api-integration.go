package main

import (
	coralogixlogger "coralogix"
	"encoding/json"
	"fmt"
)

func main() {

	GETAllActiveFacilities()
}

//GETAllActiveFacilities Returns all faas active facilities
func GETAllActiveFacilities() error {

	logger := coralogixlogger.NewCoralogixLoggerManager("mts-v2-dev", "6998c1d8-ef61-5073-1a1a-58f226e46a7f", "lambda:mts-v2-Internal-API-endpoint-Invocation-Test-Developer", "FAAS", "FAAS-Integration", "FAAS-Call")

	defer logger.WriteLogs()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic: ", err)
			logger.Error("Panic Error Occured: ", err)
		}
	}()

	URL := "https://faas-api.delhivery.com/v3/facilities/?page=1&status=active&fields=property_lat,property_long"

	//Add valid token here to test
	Token := "Bearer {token}"
	FAASAPIResponseObj := FAASAPIResponse{}

	StatusCode, Response, err := coralogixlogger.HTTPCall("GET", URL, nil, &Token)

	if StatusCode != 200 || err != nil {
		logger.Error("FAAS API failed to return facilities. Error: ", err)
		logger.Error("URL: ", URL, " StatusCode: ", StatusCode)
		ErrorResponseHandler := make(map[string]interface{})
		ErrorResponseHandler["faas_api_response"] = Response
		ErrorResponseHandlerBytes, _ := json.Marshal(ErrorResponseHandler)
		logger.Error(string(ErrorResponseHandlerBytes))

		return err
	}

	err = json.Unmarshal(Response, &FAASAPIResponseObj)

	if err != nil {
		logger.Error("Failed to unmarshal. Error: ", err)
		ErrorResponseHandler := make(map[string]interface{})
		ErrorResponseHandler["aws_request_id"] = "local"
		ErrorResponseHandler["fms_api_response"] = string(Response)
		ErrorResponseHandlerBytes, _ := json.Marshal(ErrorResponseHandler)
		logger.Error(string(ErrorResponseHandlerBytes))
		return err
	}

	logger.Info("API worked successfully")

	var Tmp interface{}
	Tmp = nil
	MapElement := Tmp.(map[string]interface{})
	fmt.Println(MapElement)

	APIResponseHandler := make(map[string]interface{})
	APIResponseHandler["api_response"] = FAASAPIResponseObj
	APIResponseHandlerBytes, _ := json.Marshal(APIResponseHandler)
	APIResponseHandlerJSON, _ := coralogixlogger.RemovePrettyJSON(APIResponseHandlerBytes)
	logger.Info("API Response: ", APIResponseHandlerJSON)
	logger.Info(string(APIResponseHandlerBytes))
	return nil
}

//FacilityDetailsResponse Facility details response from dynamodb
type FacilityDetailsResponse struct {
	PropertyLat  interface{} `json:"property_lat"`
	FacilityCode string      `json:"facility_code"`
	Name         string      `json:"name"`
	PropertyLong interface{} `json:"property_long"`
	PropertyID   string      `json:"property_id"`
	PropertyName string      `json:"property_name"`
	FacilityType []string    `json:"facility_type"`
}

//FAASAPIResponse FAAS API Response structure
type FAASAPIResponse struct {
	Result struct {
		Data []FacilityDetailsResponse `json:"data"`
	} `json:"result"`
}
