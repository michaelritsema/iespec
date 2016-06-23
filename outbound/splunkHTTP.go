package outbound

import (
	"net/http"
	"strings"
)

func SplunkPOST(config string, jsonString string) {

	URL := "http://localhost:8088/services/collector"
	client := &http.Client{}
	postBody := "{\"event\" : " + jsonString + "}"

	authHeaderKey := "CFB8BDA8-BF59-487E-96A1-A6452A57F926"
	req, _ := http.NewRequest("POST", URL, strings.NewReader(postBody))
	req.Header.Add("Authorization", "Splunk "+authHeaderKey)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
}
