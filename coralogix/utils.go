package coralogix

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//HTTPCall Generic HTTP Call
func HTTPCall(method, url string, body []byte, Auth *string) (int, []byte, error) {

	payload := strings.NewReader(string(body))
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return -1, nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if Auth != nil {
		req.Header.Add("Authorization", *Auth)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, nil, err
	}

	return res.StatusCode, body, err
}

//RemovePrettyJSON Remove pretty formatting
func RemovePrettyJSON(input []byte) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, input)
	if err != nil {
		return nil, err
	}
	return buffer, err
}
