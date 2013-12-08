package naws

import (
	. "github.com/egonelbre/mud/telnet"
)

// Implements the NAWS protocol as defined by RFC 1073
// http://www.ietf.org/rfc/rfc1073.txt

const Code byte = 31

type Negotiation struct {
	ServerAsked bool
	Width       int
	Height      int
}

func New() *Negotiation {
	return &Negotiation{true, -1, -1}
}

func InitByClient(c Command) Negotiator {
	return &Negotiation{false, -1, -1}
}

func (n *Negotiation) Request() Command {
	return Do{Code}
}

func (n *Negotiation) Handle(c Command) Command {
	switch cmd := c.(type) {
	case Will:
		// if the server didn't ask, we need to respond
		if !n.ServerAsked {
			return Do{Code}
		}
	case SubNegotiation:
		n.Width = int(cmd.Data[0])<<4 + int(cmd.Data[1])
		n.Height = int(cmd.Data[2])<<4 + int(cmd.Data[3])
		return nil
	}
	return nil
}
