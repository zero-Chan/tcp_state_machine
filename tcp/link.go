package tcp

var (
	LINK_CLOSED      = TCPLink{}
	LINK_LISTEN      = TCPLink{}
	LINK_SYN_RCVD    = TCPLink{}
	LINK_SYN_SENT    = TCPLink{}
	LINK_ESTABLISHED = TCPLink{}
	LINK_FIN_WAIT1   = TCPLink{}
	LINK_CLOSING     = TCPLink{}
	LINK_FIN_WAIT2   = TCPLink{}
	LINK_TIME_WAIT   = TCPLink{}
	LINK_CLOSE_WAIT  = TCPLink{}
	LINK_LAST_ACK    = TCPLink{}
)

func init() {
	// link关系
	LINK_CLOSED = TCPLink{
		CurrentStatus: STATUS_CLOSED,
		Handlers: []Handler{
			{
				Event:    EVENT_CONNECT,
				Action:   ACTION_SYN,
				NextLink: &LINK_SYN_SENT,
			},
			{
				Event:    EVENT_LISTEN,
				Action:   ACTION_NOT,
				NextLink: &LINK_LISTEN,
			},
		},
	}

	LINK_LISTEN = TCPLink{
		CurrentStatus: STATUS_LISTEN,
		Handlers: []Handler{
			{
				Event:    EVENT_CLOSE,
				Action:   ACTION_NOT,
				NextLink: &LINK_CLOSED,
			},
			{
				Event:    EVENT_SEND,
				Action:   ACTION_SYN,
				NextLink: &LINK_SYN_SENT,
			},
			{
				Event:    EVENT_SYN,
				Action:   ACTION_SYN_ACK,
				NextLink: &LINK_SYN_RCVD,
			},
		},
	}

	LINK_SYN_RCVD = TCPLink{
		CurrentStatus: STATUS_SYN_RCVD,
		Handlers: []Handler{
			{
				Event:    EVENT_RST,
				Action:   ACTION_NOT,
				NextLink: &LINK_LISTEN,
			},
			{
				Event:    EVENT_ACK,
				Action:   ACTION_NOT,
				NextLink: &LINK_ESTABLISHED,
			},
			{
				Event:    EVENT_CLOSE,
				Action:   ACTION_FIN,
				NextLink: &LINK_FIN_WAIT1,
			},
		},
	}

	LINK_SYN_SENT = TCPLink{
		CurrentStatus: STATUS_SYN_SENT,
		Handlers: []Handler{
			{
				Event:    EVENT_CLOSE,
				Action:   ACTION_NOT,
				NextLink: &LINK_CLOSED,
			},
			{
				Event:    EVENT_SYN,
				Action:   ACTION_SYN_ACK,
				NextLink: &LINK_SYN_RCVD,
			},
			{
				Event:    EVENT_SYN_ACK,
				Action:   ACTION_ACK,
				NextLink: &LINK_ESTABLISHED,
			},
		},
	}

	LINK_ESTABLISHED = TCPLink{
		CurrentStatus: STATUS_ESTABLISHED,
		Handlers: []Handler{
			{
				Event:    EVENT_CLOSE,
				Action:   ACTION_FIN,
				NextLink: &LINK_FIN_WAIT1,
			},
			{
				Event:    EVENT_FIN,
				Action:   ACTION_ACK,
				NextLink: &LINK_CLOSE_WAIT,
			},
		},
	}

	LINK_FIN_WAIT1 = TCPLink{
		CurrentStatus: STATUS_FIN_WAIT1,
		Handlers: []Handler{
			{
				Event:    EVENT_FIN,
				Action:   ACTION_ACK,
				NextLink: &LINK_CLOSING,
			},
			{
				Event:    EVENT_ACK,
				Action:   ACTION_NOT,
				NextLink: &LINK_FIN_WAIT2,
			},
			{
				Event:    EVENT_FIN_ACK,
				Action:   ACTION_ACK,
				NextLink: &LINK_TIME_WAIT,
			},
		},
	}

	LINK_CLOSING = TCPLink{
		CurrentStatus: STATUS_CLOSING,
		Handlers: []Handler{
			{
				Event:    EVENT_ACK,
				Action:   ACTION_NOT,
				NextLink: &LINK_TIME_WAIT,
			},
		},
	}

	LINK_FIN_WAIT2 = TCPLink{
		CurrentStatus: STATUS_FIN_WAIT2,
		Handlers: []Handler{
			{
				Event:    EVENT_FIN,
				Action:   ACTION_ACK,
				NextLink: &LINK_TIME_WAIT,
			},
		},
	}

	LINK_TIME_WAIT = TCPLink{
		CurrentStatus: STATUS_TIME_WAIT,
		Handlers: []Handler{
			{
				Event:    EVENT_TIMEOUT,
				Action:   ACTION_NOT,
				NextLink: &LINK_CLOSED,
			},
		},
	}

	LINK_CLOSE_WAIT = TCPLink{
		CurrentStatus: STATUS_CLOSE_WAIT,
		Handlers: []Handler{
			{
				Event:    EVENT_CLOSE,
				Action:   ACTION_FIN,
				NextLink: &LINK_LAST_ACK,
			},
		},
	}

	LINK_LAST_ACK = TCPLink{
		CurrentStatus: STATUS_LAST_ACK,
		Handlers: []Handler{
			{
				Event:    EVENT_ACK,
				Action:   ACTION_NOT,
				NextLink: &LINK_CLOSED,
			},
		},
	}
}

type TCPLink struct {
	CurrentStatus TCPStatus
	Handlers      []Handler
}

type Handler struct {
	Event    TCPEvent
	Action   TCPAction
	NextLink *TCPLink
}
