package iespec

import (
	"fmt"
	"net/http"
	"strings"
)

func SplunkPOST(config string, jsonString string) {

	URL := "http://localhost:8088/services/collector"
	client := &http.Client{}
	postBody := "{\"event\" : " + jsonString + "}"
	//fmt.Println("post body:" + postBody)
	authHeaderKey := "CFB8BDA8-BF59-487E-96A1-A6452A57F926"
	req, _ := http.NewRequest("POST", URL, strings.NewReader(postBody))
	req.Header.Add("Authorization", "Splunk "+authHeaderKey)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(config)
		fmt.Println(err)
	}
	//fmt.Println(resp)
}
