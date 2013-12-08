package opt

import "fmt"

// This is a listing of different Telnet RFC-s

type Info struct {
	Code     byte
	Mnemonic string
	Name     string
}

func ByCode(opt byte) Info {
	if opt, ok := bycode[opt]; ok {
		return opt
	}
	return Info{opt, fmt.Sprintf("U%d"), fmt.Sprintf("Unknown (%d)")}
}

var bycode = map[byte]Info{
	BINARY:         Info{BINARY, "BINARY", "binary"},
	ECHO:           Info{ECHO, "ECHO", "echo"},
	RCP:            Info{RCP, "RCP", "prepare for reconnect"},
	SGA:            Info{SGA, "SGA", "supress go ahead"},
	NAMS:           Info{NAMS, "NAMS", "approximate message size"},
	STATUS:         Info{STATUS, "STATUS", "give status"},
	TIMING_MARK:    Info{TIMING_MARK, "TIMING_MARK", "timing mark"},
	RCTE:           Info{RCTE, "RCTE", "remote controlled transmission and echo"},
	NAOL:           Info{NAOL, "NAOL", "negotiate about ouput line width"},
	NAOP:           Info{NAOP, "NAOP", "negotiate about output page size"},
	NAOCRD:         Info{NAOCRD, "NAOCRD", "negotiate about CR disposition"},
	NAOHTS:         Info{NAOHTS, "NAOHTS", "negotiate about horizontal tabstops"},
	NAOHTD:         Info{NAOHTD, "NAOHTD", "negotiate about horizontal tab disposition"},
	NAOFFD:         Info{NAOFFD, "NAOFFD", "negotiate about formfeed disposition"},
	NAOVTS:         Info{NAOVTS, "NAOVTS", "negotiate about vertical tab stops"},
	NAOVTD:         Info{NAOVTD, "NAOVTD", "negotiate about vertical tab disposition"},
	NAOLFD:         Info{NAOLFD, "NAOLFD", "negotiate about output LF disposition"},
	XASCII:         Info{XASCII, "XASCII", "extended ascii character set"},
	LOGOUT:         Info{LOGOUT, "LOGOUT", "force logout"},
	BM:             Info{BM, "BM", "byte macro"},
	DET:            Info{DET, "DET", "data entry terminal"},
	SUPDUP:         Info{SUPDUP, "SUPDUP", "supdup protocol"},
	SUPDUPOUTPUT:   Info{SUPDUPOUTPUT, "SUPDUPOUTPUT", "supdup output"},
	SNDLOC:         Info{SNDLOC, "SNDLOC", "send location"},
	TTYPE:          Info{TTYPE, "TTYPE", "terminal type"},
	EOR:            Info{EOR, "OPT_EOR", "end or record"},
	TUID:           Info{TUID, "TUID", "TACACS user identification"},
	OUTMRK:         Info{OUTMRK, "OUTMRK", "output marking"},
	TTYLOC:         Info{TTYLOC, "TTYLOC", "terminal location number"},
	VT3270REGIME:   Info{VT3270REGIME, "VT3270REGIME", "3270 regime"},
	X3PAD:          Info{X3PAD, "X3PAD", "X.3 PAD"},
	NAWS:           Info{NAWS, "NAWS", "window size"},
	TERMINAL_SPEED: Info{TERMINAL_SPEED, "TERMINAL_SPEED", "terminal speed"},
	LFLOW:          Info{LFLOW, "LFLOW", "remote flow control"},
	LINE_MODE:      Info{LINE_MODE, "LINE_MODE", "Linemode option"},
	XDISPLOC:       Info{XDISPLOC, "XDISPLOC", "X Display Location"},
	OLD_ENVIRON:    Info{OLD_ENVIRON, "OLD_ENVIRON", "Old - Environment variables"},
	AUTHENTICATION: Info{AUTHENTICATION, "AUTHENTICATION", "Authenticate"},
	ENCRYPT:        Info{ENCRYPT, "ENCRYPT", "Encryption option"},
	NEW_ENVIRON:    Info{NEW_ENVIRON, "NEW_ENVIRON", "Environment variables"},
	CHARSET:        Info{CHARSET, "CHARSET", "Charset"},
	MSDP:           Info{MSDP, "MSDP", "Mud Server Data Protocol"},
	MSSP:           Info{MSSP, "MSSP", "Mud Server Status Protocol"},
	MCCPv1:         Info{MCCPv1, "MCCPv1", "Mud Client Compression Protocol v1"},
	MCCP:           Info{MCCP, "MCCP", "Mud Client Compression Protocol v2"},
	MSP:            Info{MSP, "MSP", "MUD Sound Protocol"},
	MXP:            Info{MXP, "MXP", "Mud eXtension Protocol"},
	ZMP:            Info{ZMP, "ZMP", "Zenith MUD Protocol"},
	GCMP:           Info{GCMP, "GCMP", "Generic Mud Communication Protocol"},
	EXOPL:          Info{EXOPL, "EXOPL", "telnet extended options"},
}

const (
	BINARY         byte = 0  // RFC 856
	ECHO           byte = 1  // RFC 857
	RCP            byte = 2  // RFC 426
	SGA            byte = 3  // RFC 858
	NAMS           byte = 4  //
	STATUS         byte = 5  // RFC 859
	TIMING_MARK    byte = 6  // RFC 860
	RCTE           byte = 7  //
	NAOL           byte = 8  //
	NAOP           byte = 9  //
	NAOCRD         byte = 10 //
	NAOHTS         byte = 11 //
	NAOHTD         byte = 12 //
	NAOFFD         byte = 13 //
	NAOVTS         byte = 14 //
	NAOVTD         byte = 15 //
	NAOLFD         byte = 16 //
	XASCII         byte = 17 //
	LOGOUT         byte = 18 //
	BM             byte = 19 //
	DET            byte = 20 //
	SUPDUP         byte = 21 //
	SUPDUPOUTPUT   byte = 22 //
	SNDLOC         byte = 23 //
	TTYPE          byte = 24 // RFC 930, 1091, http://tintin.sourceforge.net/mtts/
	EOR            byte = 25 // RFC 885
	TUID           byte = 26 //
	OUTMRK         byte = 27 //
	TTYLOC         byte = 28 //
	VT3270REGIME   byte = 29 //
	X3PAD          byte = 30 //
	NAWS           byte = 31
	TERMINAL_SPEED byte = 32  // RFC 1079
	LFLOW          byte = 33  //
	LINE_MODE      byte = 34  // RFC 1184
	XDISPLOC       byte = 35  //
	OLD_ENVIRON    byte = 36  //
	AUTHENTICATION byte = 37  //
	ENCRYPT        byte = 38  //
	NEW_ENVIRON    byte = 39  // RFC 1572
	CHARSET        byte = 42  // RFC 2066
	MSDP           byte = 69  // http://tintin.sourceforge.net/msdp/
	MSSP           byte = 70  // http://tintin.sourceforge.net/mssp/
	MCCPv1         byte = 85  // http://www.zuggsoft.com/zmud/mcp.htm
	MCCP           byte = 86  // http://www.zuggsoft.com/zmud/mcp.htm
	MSP            byte = 90  // http://www.zuggsoft.com/zmud/msp.htm
	MXP            byte = 91  // http://www.zuggsoft.com/zmud/mxp.htm
	ZMP            byte = 93  // http://discworld.starturtle.net/external/protocols/zmp.html
	GCMP           byte = 201 // http://www.ironrealms.com/gmcp-doc
	EXOPL          byte = 255 // RFC 861
)
