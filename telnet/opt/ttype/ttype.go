package ttype

import (
	. "mud/telnet"
)

// Implements the TTYPE protocol as defined by RFC 1091
// http://www.ietf.org/rfc/rfc1091.txt
//

const Code byte = 24

type Negotiation struct {
	Type string
}

func New() *Negotiation {
	return &Negotiation{""}
}

func InitByClient(c Command) Negotiator {
	return &Negotiation{""}
}

func (n *Negotiation) Request() Command {
	return Do{Code}
}

func (n *Negotiation) Handle(c Command) Command {
	switch cmd := c.(type) {
	case Will:
		// client showed willingness
		return SubNegotiation{Code, []byte{SEND}}
	case SubNegotiation:
		n.Type = string(cmd.Data[1:])
		return nil
	}
	return nil
}
