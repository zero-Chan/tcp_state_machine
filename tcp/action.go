package tcp

var (
	ACTION_SYN     = TCPAction{Name: "SYN", Desc: "发起建立连接请求"}
	ACTION_SYN_ACK = TCPAction{Name: "SYN_ACK", Desc: "响应建立连接请求并请求建立连接"}
	ACTION_ACK     = TCPAction{Name: "ACK", Desc: "响应"}
	ACTION_FIN     = TCPAction{Name: "FIN", Desc: "发起关闭连接请求"}
	ACTION_FIN_ACK = TCPAction{Name: "FIN_ACK", Desc: "响应关闭连接请求并发起关闭连接请求"}
	ACTION_NOT     = TCPAction{Name: "NOT", Desc: "什么都不做"}
)

type TCPAction struct {
	Name string
	Desc string
}

var (
	// map[action.name]
	Action2EventMap = map[string]TCPEvent{
		ACTION_SYN.Name:     EVENT_SYN,
		ACTION_SYN_ACK.Name: EVENT_SYN_ACK,
		ACTION_ACK.Name:     EVENT_ACK,
		ACTION_FIN.Name:     EVENT_FIN,
		ACTION_FIN_ACK.Name: EVENT_FIN_ACK,
	}
)
