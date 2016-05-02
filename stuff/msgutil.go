package stuff

import (
	"encoding/base64"
)
// can skip hmac for now

func Wrap(msgtype string,data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	msg := "<pb HMAC=\"xxx\" type=\"ZFlow\">" + encoded + "</pb>" 
	return msg
}