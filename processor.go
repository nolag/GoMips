package gomips

import (
	"encoding/binary"

	"github.com/nolag/gocpu/memory"
	"github.com/nolag/gocpu/registers"
)

// Processor represents a MIPS processor, it is meant to be encapsulated by an implementation.
type Processor struct {
	ByteOrder      binary.ByteOrder
	Coprocessors   [4]Coprocessor
	FloatRegisters [32]registers.IFloatRegister32
	Hi             registers.IIntRegister32
	InBranchDelay  bool
	Low            registers.IIntRegister32
	Memory         memory.Memory
	Pc             registers.IIntRegister32
	Registers      [32]registers.IIntRegister32
}
