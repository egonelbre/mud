package negotiator

import (
	"github.com/egonelbre/mud/telnet/opt"
	"github.com/egonelbre/mud/telnet/opt/naws"
	"github.com/egonelbre/mud/telnet/opt/ttype"
)

var DefaultInitiator = map[byte]Initiator{
	opt.NAWS:  naws.InitByClient,
	opt.TTYPE: ttype.InitByClient,
}
