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

// Processor represents a MIPS processor, it is meant to be used with a FixedInstructionLenRunnerUint32
// Use one of the New... functions to assure you set all needed fields
type Processor struct {
	ByteOrder          binary.ByteOrder
	Coprocessors       [4]Coprocessor
	ErrorHandler       ErrorHandler
	FloatRegisters     [32]registers.IFloatRegister32
	Hi                 uint32
	InBranchDelay      bool
	InstructionActions [64]InstructionAction
	Low                uint32
	Memory             memory.Memory
	Pc                 registers.IRegister32
	Registers          [32]registers.IIntRegister32
}

// RunUint32 runs a single instrution (without incrementing the PC for its own read)
func (processor *Processor) RunUint32(instruction uint32) error {
	mipsInstruciton := Instruction(instruction)
	action := processor.InstructionActions[mipsInstruciton.OpCode()]
	if action != nil {
		err := action(processor, mipsInstruciton)
		_, ok := err.(UnknonInstructionError)
		if ok {
			return processor.runUnknownInstructionAction(mipsInstruciton)
		}

		return err
	}

	return processor.runUnknownInstructionAction(mipsInstruciton)
}

func (processor *Processor) runUnknownInstructionAction(instruction Instruction) error {
	if processor.ErrorHandler == nil {
		return UnknonInstructionError(instruction)
	}

	return processor.ErrorHandler(processor, RI)
}
