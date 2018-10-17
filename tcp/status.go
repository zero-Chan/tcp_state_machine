package tcp

var (
	STATUS_CLOSED      = TCPStatus{Name: "CLOSED", Desc: "没有活跃的连接或者挂起"}
	STATUS_LISTEN      = TCPStatus{Name: "LISTEN", Desc: "服务器等待入境呼叫"}
	STATUS_SYN_RCVD    = TCPStatus{Name: "SYN_RCVD", Desc: "到达一个连接请求，等待ACK"}
	STATUS_SYN_SENT    = TCPStatus{Name: "SYN_SENT", Desc: "应用已经启动了打开一个连接"}
	STATUS_ESTABLISHED = TCPStatus{Name: "ESTABLISHED", Desc: "正常的数据传送状态"}
	STATUS_FIN_WAIT1   = TCPStatus{Name: "FIN_WAIT1", Desc: "应用没有数据要发了"}
	STATUS_FIN_WAIT2   = TCPStatus{Name: "FIN_WAIT2", Desc: "另一端同意释放连接"}
	STATUS_TIME_WAIT   = TCPStatus{Name: "TIME_WAIT", Desc: "等待所有数据包寿终正寝"}
	STATUS_CLOSING     = TCPStatus{Name: "CLOSING", Desc: "两端同时试图关闭连接"}
	STATUS_CLOSE_WAIT  = TCPStatus{Name: "CLOSE_WAIT", Desc: "另一端已经发起关闭连接"}
	STATUS_LAST_ACK    = TCPStatus{Name: "LAST_ACK", Desc: "等待所有数据包寿终正寝"}
)

type TCPStatus struct {
	Name string
	Desc string
}
