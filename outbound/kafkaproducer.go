package outbound

import (
	"fmt"
	"iespec/protomsg"
)

func Kafka(c chan *protomsg.ZFlow) {
	go func() {
		for {
			msg := <-c
			fmt.Println(msg)
		}
	}()
}
