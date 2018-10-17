package main

import (
	"zero-Chan/tcp_state_machine/tcp"
	"zero-Chan/tcp_state_machine/event"
	"fmt"
	"time"
)

func main() {
	cli := tcp.NewTCPClient()
	svr := tcp.NewTCPServer()

	for {
		cli.List()
		svr.List()

		var object string
		fmt.Printf("需要提交信号的对象(cli/svr):")
		fmt.Scanln(&object)

		switch object {
		case "cli":
			var eventNum int
			fmt.Printf("需要提交的事件序号:")
			fmt.Scanln(&eventNum)
			if len(cli.Link.Handlers) < eventNum || eventNum < 0 {
				fmt.Println("无效序号")
				continue
			}

			handler := cli.Link.Handlers[eventNum]
			cli.OfferEvent(handler.Event)

		case "svr":
			var eventNum int
			fmt.Printf("需要提交的事件序号:")
			fmt.Scanln(&eventNum)
			if len(svr.Link.Handlers) < eventNum || eventNum < 0 {
				fmt.Println("无效序号")
				continue
			}

			handler := svr.Link.Handlers[eventNum]
			svr.OfferEvent(handler.Event)
		}

		time.Sleep(1 * time.Second)
	}

	event.WaitExit()
}
