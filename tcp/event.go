package tcp

var (
	EVENT_CONNECT = TCPEvent{Name: "CONNECT", Desc: "客户端主动发起连接请求"}
	EVENT_CLOSE   = TCPEvent{Name: "CLOSE", Desc: "主动关闭连接"}
	EVENT_LISTEN  = TCPEvent{Name: "LISTEN", Desc: "开始监听请求"}
	EVENT_SYN     = TCPEvent{Name: "SYN", Desc: "收到连接请求"}
	EVENT_ACK     = TCPEvent{Name: "ACK", Desc: "收到应答"}
	EVENT_RST     = TCPEvent{Name: "RST", Desc: "无效"}
	EVENT_SEND    = TCPEvent{Name: "SEND", Desc: "服务端主动发起连接请求"}
	EVENT_SYN_ACK = TCPEvent{Name: "SYN_ACK", Desc: "收到连接应答并请求建立连接"}
	EVENT_FIN     = TCPEvent{Name: "FIN", Desc: "收到关闭连接请求"}
	EVENT_FIN_ACK = TCPEvent{Name: "FIN_ACK", Desc: "收到关闭应答并请求关闭连接"}
	EVENT_TIMEOUT = TCPEvent{Name: "TIMEOUT", Desc: "超时,时间长度为最大数据包生存期两倍的时间"}
)

type TCPEvent struct {
	Name string
	Desc string
}
