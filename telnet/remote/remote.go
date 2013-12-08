package remote

import (
	"io"
	"log"
	"mud/telnet"
	"mud/telnet/negotiator"
	"mud/telnet/opt"
	"strings"
)

const maxLine = 4096

type Remote struct {
	conn       io.ReadWriter
	Lines      chan string
	send       chan telnet.Command
	negotiator telnet.Negotiator
}

func New(input io.ReadWriter) *Remote {
	r := &Remote{
		conn:       input,
		Lines:      make(chan string),
		send:       make(chan telnet.Command),
		negotiator: negotiator.NewDefault(),
	}
	go r.run()
	return r
}

func NewWithNegotiator(input io.ReadWriter, n telnet.Negotiator) *Remote {
	r := &Remote{
		conn:       input,
		Lines:      make(chan string),
		send:       make(chan telnet.Command),
		negotiator: n,
	}
	go r.run()
	return r
}

func (r *Remote) formatLine(line string) string {
	return strings.Replace(line, "\n", "\r\n", -1)
}

func (r *Remote) prepare(c telnet.Command) telnet.Command {
	switch cmd := c.(type) {
	case string:
		return r.formatLine(cmd)
	case telnet.Transaction:
		for i, c := range cmd {
			cmd[i] = r.prepare(c)
		}
		return cmd
	default:
		return cmd
	}
}

func (r *Remote) Send(c telnet.Command) {
	r.send <- r.prepare(c)
}

func (r *Remote) run() {
	defer r.Terminate()

	cmds := make(chan telnet.Command)
	go telnet.Unserialize(r.conn, cmds)
	go telnet.Serialize(r.send, r.conn)

	for c := range cmds {
		switch cmd := c.(type) {
		default:
			log.Printf("Unexpected type %T\n", cmd)
		case string:
			log.Printf("User action '%s'\n", cmd)
			r.Lines <- cmd
		case telnet.OptionCommand:
			info := opt.ByCode(cmd.OptionCode())
			log.Printf("Command [%v] %v\n", info.Name, cmd)
			if r.negotiator != nil {
				r.negotiator.Handle(cmd)
			} else {
				r.Send(telnet.Wont{cmd.OptionCode()})
			}
		}
	}
}

func (r *Remote) Terminate() {
	close(r.Lines)
}
