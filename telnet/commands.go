package telnet

type Negotiator interface {
	Handle(c Command) (response Command)
}

type (
	Command interface{}
	// commands that have are targeted to a telnet option
	OptionCommand interface {
		OptionCode() byte
	}

	Break            struct{}
	InterruptProcess struct{}
	AbortOutput      struct{}
	AreYouThere      struct{}
	GoAhead          struct{}

	Will struct{ Option byte }
	Wont struct{ Option byte }
	Do   struct{ Option byte }
	Dont struct{ Option byte }

	SubNegotiation struct {
		Option byte
		Data   []byte
	}

	// for sending only
	Transaction []Command
)

func (c Will) OptionCode() byte           { return c.Option }
func (c Wont) OptionCode() byte           { return c.Option }
func (c Do) OptionCode() byte             { return c.Option }
func (c Dont) OptionCode() byte           { return c.Option }
func (c SubNegotiation) OptionCode() byte { return c.Option }
