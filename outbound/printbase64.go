package outbound

import (
	"encoding/base64"
	"fmt"
)

func printEncoded(buf []byte) {
	encoded := base64.StdEncoding.EncodeToString(buf)
	fmt.Println(encoded)
}
