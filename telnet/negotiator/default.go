package negotiator

import (
	"mud/telnet/opt"
	"mud/telnet/opt/naws"
	"mud/telnet/opt/ttype"
)

var DefaultInitiator = map[byte]Initiator{
	opt.NAWS:  naws.InitByClient,
	opt.TTYPE: ttype.InitByClient,
}
