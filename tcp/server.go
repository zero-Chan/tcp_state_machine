package tcp

import (
	"fmt"
)

var (
	serverChannel = make(chan TCPEvent)
)

type TCPServer struct {
	Link *TCPLink
}

func NewTCPServer() *TCPServer {
	svr := &TCPServer{
		Link: &LINK_LISTEN,
	}

	go svr.Recv()

	return svr
}

func (this *TCPServer) Recv() {
	for event := range serverChannel {
		handler, exist := this.getHandler(event)
		if !exist {
			fmt.Printf("TCPServer 接收到无效信号[%s]. 当前状态[%s]\n", event.Name, this.Link.CurrentStatus.Name)
			continue
		}

		this.Handle(handler)
	}
}

func (this *TCPServer) Send(action TCPAction) {
	if action.Name != ACTION_NOT.Name {
		event, exist := Action2EventMap[action.Name]
		if exist {
			clientChannel <- event
		}
	}
}

func (this *TCPServer) getHandler(event TCPEvent) (hdl Handler, exist bool) {
	handlers := this.Link.Handlers
	for _, handler := range handlers {
		if event.Name == handler.Event.Name {
			return handler, true
		}
	}

	return Handler{}, false
}

func (this *TCPServer) Handle(handler Handler) {
	fmt.Println("----------")
	fmt.Println("svr:")
	fmt.Printf("recv:\t%s\t(%s)\n", handler.Event.Name, handler.Event.Desc)
	fmt.Printf("send:\t%s\t(%s)\n", handler.Action.Name, handler.Action.Desc)
	fmt.Printf("nstatus:\t%s\t(%s)\n", handler.NextLink.CurrentStatus.Name, handler.NextLink.CurrentStatus.Desc)
	fmt.Println("----------")

	this.Link = handler.NextLink

	this.Send(handler.Action)
}

func (this *TCPServer) List() {
	fmt.Println("====================")
	fmt.Println("svr:")
	fmt.Printf("status:\t%s\t(%s)\n", this.Link.CurrentStatus.Name, this.Link.CurrentStatus.Desc)
	for idx, handler := range this.Link.Handlers {
		fmt.Printf("%d: eve:%s, act:%s, nstatus:%s\n", idx, handler.Event.Name, handler.Action.Name, handler.NextLink.CurrentStatus.Name)
	}

	fmt.Println("====================")
}

func (this *TCPServer) OfferEvent(event TCPEvent) {
	serverChannel <- event
}
