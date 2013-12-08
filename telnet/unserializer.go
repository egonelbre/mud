package telnet

import (
	"bufio"
	"io"
	"log"
)

const maxLineLength = 1024

// Unserializes telnet bytestream into commands
func Unserialize(input io.Reader, out chan Command) {
	stream := bufio.NewReader(input)

	next := func() byte {
		b, err := stream.ReadByte()
		if err != nil {
			panic(err)
		}
		return b
	}

	defer func() {
		err := recover()
		if (err == io.EOF) || (err == nil) {
			return
		}
		log.Printf("Error occurred: %v\n", err)
	}()

	buffer := make([]byte, 0, maxLineLength)
	for {
		b := next()
		switch b {
		case IAC:
			command := next()
			switch command {
			case SB:
				sub := SubNegotiation{next(), make([]byte, 0)}
				for {
					b := next()
					if b == IAC {
						suboption := next()
						if suboption == SB {
							sub.Data = append(sub.Data, IAC)
						} else if suboption == SE {
							break
						} else {
							panic("Protocol error!")
						}
					} else {
						sub.Data = append(sub.Data, b)
					}
				}
				out <- sub
			case WILL:
				out <- Will{next()}
			case WONT:
				out <- Wont{next()}
			case DO:
				out <- Do{next()}
			case DONT:
				out <- Dont{next()}

			case BRK:
				out <- Break{}
			case IP:
				out <- InterruptProcess{}
			case AO:
				out <- AbortOutput{}
			case AYT:
				out <- AreYouThere{}

			case EC:
				if len(buffer) > 0 {
					buffer = buffer[:len(buffer)-1]
				}
			case EL:
				buffer = buffer[:0]

			case DM:
				panic("Don't know how to handle data mark!")

			case NOP, GA: // do nothing
			default:
				buffer = append(buffer, command)
			}
		case CarriageReturn, LineFeed:
			if len(buffer) > 0 {
				out <- string(buffer)
				buffer = buffer[:0]
			}
		default:
			buffer = append(buffer, b)
			if len(buffer) >= maxLineLength {
				panic("Line too long!")
			}
		}
	}
}
