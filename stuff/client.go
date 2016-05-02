package stuff

import (
	"net/http"
	"fmt"
	"strings"
)

func Post(data string) {
	//data := "<pb type="ZFlow" HMAC=\"a82389f7cad3e320710a2731a2b8a1c8d4463f76\">eHh4eAo=</pb>"
	postUrl := "http://ec2-54-210-70-189.compute-1.amazonaws.com:8040/api/datacollection"

	resp,err := http.Post(postUrl, "text/plain;charset=UTF-8", strings.NewReader(data))
	fmt.Println(resp)
	fmt.Println(err)
}
