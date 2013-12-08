package telnet

import (
	"bufio"
	"io"
)

type PreprocessLine func(string) string

func write(c Command, to *bufio.Writer) {
	switch cmd := c.(type) {
	case string:
		to.WriteString(cmd)
	case []byte:
		to.Write(cmd)

	case Transaction:
		for _, c := range cmd {
			write(c, to)
		}

	case SubNegotiation:
		to.Write([]byte{IAC, SB, cmd.Option})
		to.Write(cmd.Data)
		to.Write([]byte{IAC, SE})

	case Will:
		to.Write([]byte{IAC, WILL, cmd.Option})
	case Wont:
		to.Write([]byte{IAC, WONT, cmd.Option})
	case Do:
		to.Write([]byte{IAC, DO, cmd.Option})
	case Dont:
		to.Write([]byte{IAC, DONT, cmd.Option})

	case Break:
		to.Write([]byte{IAC, BRK})
	case InterruptProcess:
		to.Write([]byte{IAC, IP})
	case AbortOutput:
		to.Write([]byte{IAC, AO})
	case AreYouThere:
		to.Write([]byte{IAC, AYT})
	case GoAhead:
		to.Write([]byte{IAC, GA})
	}
}

// Serializes Commands into io.Writer
func Serialize(in chan Command, to io.Writer) {
	stream := bufio.NewWriter(to)
	for c := range in {
		write(c, stream)
		stream.Flush()
	}
}
