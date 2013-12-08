package remote

import (
	"io"
	"fmt"
	"log"
	"strings"

	"github.com/egonelbre/mud/telnet"
	"github.com/egonelbre/mud/telnet/negotiator"
	"github.com/egonelbre/mud/telnet/opt"
)

const maxLine = 4096

type Conn struct {
	conn       io.ReadWriter
	Lines      chan string
	send       chan telnet.Command
	negotiator telnet.Negotiator
}

func New(input io.ReadWriter) *Conn {
	return NewWithNegotiator(input, negotiator.NewDefault())
}

func NewWithNegotiator(input io.ReadWriter, n telnet.Negotiator) *Conn {
	r := &Conn{
		conn:       input,
		Lines:      make(chan string),
		send:       make(chan telnet.Command),
		negotiator: n,
	}
	go r.run()
	return r
}

func (r *Conn) formatLine(line string) string {
	return strings.Replace(line, "\n", "\r\n", -1)
}

func (r *Conn) prepare(c telnet.Command) telnet.Command {
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

func (r *Conn) Send(c telnet.Command) { 
	r.send <- r.prepare(c)
}

func (r *Conn) Print(s string){
	r.send <- r.prepare(s)
}

func (r *Conn) Printf(format string, a ...interface{}){
	r.send <- r.prepare(fmt.Sprintf(format, a...))
}

func (r *Conn) run() {
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

func (r *Conn) Terminate() {
	close(r.Lines)
}
