package gomips

// ExceptionCause represents the cause of an exception
type ExceptionCause uint32

const (
	// IRQ Interrupt
	IRQ ExceptionCause = 0

	// MOD TLB read-only fault
	MOD ExceptionCause = 1

	// TLBL TLB miss load
	TLBL ExceptionCause = 2

	// TLBS TLB miss store
	TLBS ExceptionCause = 3

	// ADEL address error load
	ADEL ExceptionCause = 4

	// ADES address error store
	ADES ExceptionCause = 5

	// IBE bus error fetch
	IBE ExceptionCause = 6

	// DBE bus error data access
	DBE ExceptionCause = 7

	// SYS system call
	SYS ExceptionCause = 8

	// BP break point
	BP ExceptionCause = 9

	// RI reserved instruction
	RI ExceptionCause = 10

	// CPU Coprocessor unusable
	CPU ExceptionCause = 11

	// OVF overflow
	OVF ExceptionCause = 12

	// TE Trap excepion
	TE ExceptionCause = 13

	// DBZ Division by zero (floating point)
	DBZ ExceptionCause = 15

	// FOVF floating point overflow
	FOVF ExceptionCause = 16

	// FUNF floating point underflow
	FUNF ExceptionCause = 16
)
