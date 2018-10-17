package tcp

import (
	"fmt"
)

var (
	clientChannel = make(chan TCPEvent)
)

type TCPClient struct {
	Link *TCPLink
}

func NewTCPClient() *TCPClient {
	cli := &TCPClient{
		Link: &LINK_CLOSED,
	}

	go cli.Recv()

	return cli
}

func (this *TCPClient) Recv() {
	for event := range clientChannel {
		handler, exist := this.getHandler(event)
		if !exist {
			fmt.Printf("TCPClient 接收到无效信号[%s]. 当前状态[%s]\n", event.Name, this.Link.CurrentStatus.Name)
			continue
		}

		this.Handle(handler)
	}
}

func (this *TCPClient) Send(action TCPAction) {
	if action.Name != ACTION_NOT.Name {
		event, exist := Action2EventMap[action.Name]
		if exist {
			serverChannel <- event
		}

	}
}

func (this *TCPClient) getHandler(event TCPEvent) (hdl Handler, exist bool) {
	handlers := this.Link.Handlers
	for _, handler := range handlers {
		if event.Name == handler.Event.Name {
			return handler, true
		}
	}

	return Handler{}, false
}

func (this *TCPClient) Handle(handler Handler) {
	fmt.Println("----------")
	fmt.Println("cli:")
	fmt.Printf("recv:\t%s\t(%s)\n", handler.Event.Name, handler.Event.Desc)
	fmt.Printf("send:\t%s\t(%s)\n", handler.Action.Name, handler.Action.Desc)
	fmt.Printf("nstatus:\t%s\t(%s)\n", handler.NextLink.CurrentStatus.Name, handler.NextLink.CurrentStatus.Desc)
	fmt.Println("----------")

	this.Link = handler.NextLink

	this.Send(handler.Action)
}

func (this *TCPClient) List() {
	fmt.Println("====================")
	fmt.Println("cli:")
	fmt.Printf("status:\t%s\t(%s)\n", this.Link.CurrentStatus.Name, this.Link.CurrentStatus.Desc)
	for idx, handler := range this.Link.Handlers {
		fmt.Printf("%d: eve:%s, act:%s, nstatus:%s\n", idx, handler.Event.Name, handler.Action.Name, handler.NextLink.CurrentStatus.Name)
	}

	fmt.Println("====================")
}

func (this *TCPClient) OfferEvent(event TCPEvent) {
	clientChannel <- event
}
