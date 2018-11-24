package gomips

import (
	"encoding/binary"
	"fmt"

	"github.com/nolag/gocpu/memory"
	"github.com/nolag/gocpu/registers"
)

// InstructionAction takes action based on an Instruction
// Note that InstructionAction are expected to ignore the OpCode and assume the processor called the correct function.
type InstructionAction func(*Processor, Instruction) error

// UnknonInstructionError represents an unknown instruction
type UnknonInstructionError Instruction

func (err UnknonInstructionError) Error() string {
	return fmt.Sprintf("Unknown instruction 0x%08X", uint32(err))
}

// ErrorHandler is used to handel errors from Processor
// Note that the coprocessor 0 will be set up correctly before the call
type ErrorHandler func(*Processor, ExceptionCause) error

// Processor represents a MIPS processor, it is meant to be encapsulated by an implementation.
type Processor struct {
	ByteOrder      binary.ByteOrder
	Coprocessors   [4]Coprocessor
	FloatRegisters [32]registers.IFloatRegister32
	Hi             uint32
	InBranchDelay  bool
	Low            uint32
	Memory         memory.Memory
	Pc             registers.IRegister32
	Registers      [32]registers.IIntRegister32
}
