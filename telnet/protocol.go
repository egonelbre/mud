package telnet

// Defines the base codes for commands as defined in RFC 855
const (
	IAC byte = 255 // Marks the start of a negotiation sequence.

	WILL byte = 251 // Confirm willingness to negotiate.
	WONT byte = 252 // Confirm unwillingness to negotiate.
	DO   byte = 253 // Indicate willingness to negotiate.
	DONT byte = 254 // Indicate unwillingness to negotiate.

	SB byte = 250 // The start of sub-negotiation options.
	SE byte = 240 // The end of sub-negotiation options.

	NOP byte = 241 // No operation.
	DM  byte = 242 // Data Mark

	BRK byte = 243 // Break
	IP  byte = 244 // Interupt Process
	AO  byte = 245 // Abort Output
	AYT byte = 246 // Are You There
	EC  byte = 247 // Erase Character
	EL  byte = 248 // Erase Line
	GA  byte = 249 // Go Ahead

	EOR   byte = 239 // End of record (transparent mode)
	ABORT byte = 238 // Abort process
	SUSP  byte = 237 // Suspend process
	EOF   byte = 236 // End of file
)

const (
	IS   byte = 0 // Sub-negotiation IS command.
	SEND byte = 1 // Sub-negotiation SEND command.
	INFO byte = 2 // Sub-negotiation INFO command.
)
