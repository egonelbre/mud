package negotiator

import "github.com/egonelbre/mud/telnet"

type Initiator func(telnet.Command) telnet.Negotiator

type Automatic struct {
	Initiators  map[byte]Initiator
	Negotiatons map[byte]telnet.Negotiator
}

func NewAutomatic(initiators map[byte]Initiator) telnet.Negotiator {
	return &Automatic{initiators, make(map[byte]telnet.Negotiator)}
}

func NewDefault() telnet.Negotiator {
	return NewAutomatic(DefaultInitiator)
}

func (n *Automatic) Handle(c telnet.Command) telnet.Command {
	cmd, ok := c.(telnet.OptionCommand)
	if !ok {
		return nil
	}
	opt := cmd.OptionCode()

	ongoing, ok := n.Negotiatons[opt]
	if ok {
		return ongoing.Handle(cmd)
	}

	init, ok := n.Initiators[opt]
	if ok {
		neg := init(cmd)
		n.Negotiatons[opt] = neg
		return neg.Handle(cmd)
	}

	return telnet.Wont{opt}
}
