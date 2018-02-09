package gomips

import (
	"encoding/binary"
)

// Processor is an interface for a simulated processor
type Processor interface {
	// Step runs the next instruction, returns error to indicate a trap or unhandeled exception
	Step() error
}

// MemoryAction simulates a special memory location when written to or read from
type MemoryAction interface {
	// Writes 32 bits to the special memory location, returning an error if on a failure
	Wrtie32(value uint32) error

	// Reads 32 bits from the special memory location, returning an error on a failure
	Read32() (uint32, error)
}

// InstructionAction takes action based on an Instruction
// Note that InstructionActions are expected to ignore the OpCode and assume the processor called the correct function.
type InstructionAction func(*Mips32Processor, Instruction) error

// ProcessorHook is a hook to allow a callback for later.
// This is used to allow delayed actions (like branch and load delays) to be simulated
type ProcessorHook func(*Mips32Processor)

// Mips32Processor is a simulated MIPs processor.
type Mips32Processor struct {
	Registers          [32]Register32
	Memory             []byte
	InstructionActions [64]InstructionAction
	ProgramCounter     uint32
	ByteOrder          binary.ByteOrder
	DelayAction        ProcessorHook
}

// Step runs the next instruction, returns if there are more instructions to run.
func (processor *Mips32Processor) Step() error {
	instruction := Instruction(processor.ByteOrder.Uint32(processor.Memory[processor.ProgramCounter:]))
	processor.ProgramCounter += 4
	result := processor.InstructionActions[instruction.OpCode()](processor, instruction)
	if processor.DelayAction != nil {
		processor.DelayAction(processor)
	}

	return result
}
