// Package gomips provides a MIPS simulators and the barebones to create your own MIPs like processor
// NOTE: This package is not yet officially released and may change at any time.
package gomips

import (
	"encoding/binary"
	"fmt"
)

// Processor is an interface for a simulated processor
type Processor interface {
	// Step runs the next instruction, returns error to indicate a trap or unhandeled exception
	Step() error
}

// InstructionAction takes action based on an Instruction
// Note that InstructionActions are expected to ignore the OpCode and assume the processor called the correct function.
type InstructionAction func(*Mips32Processor, Instruction32) error

// ProcessorHook is a hook to allow a callback for later.
// This is used to allow delayed actions (like branch and load delays) to be simulated
type ProcessorHook func(*Mips32Processor)

// UnknonIntruction32Error represents an unknown instruction
type UnknonIntruction32Error Instruction32

func (err UnknonIntruction32Error) Error() string {
	return fmt.Sprintf("Unknown instruction %8x", uint32(err))
}

// Mips32Processor is a simulated MIPs processor.
type Mips32Processor struct {
	Registers          [32]Register32
	Memory             []byte
	InstructionActions [64]InstructionAction
	ProgramCounter     uint32
	ByteOrder          binary.ByteOrder
	DelayAction        ProcessorHook
	UnknownInstruction InstructionAction
	Hi                 uint32
	Low                uint32
}

// Step runs the next instruction, returns if there are more instructions to run.
func (processor *Mips32Processor) Step() error {
	instruction := Instruction32(processor.ByteOrder.Uint32(processor.Memory[processor.ProgramCounter:]))
	processor.ProgramCounter += 4
	priorDelay := processor.DelayAction
	processor.DelayAction = nil
	decodedInstruction := processor.InstructionActions[instruction.OpCode()]
	ran := false
	var result error
	if decodedInstruction != nil {
		result = decodedInstruction(processor, instruction)
		if result != nil {
			_, ok := result.(UnknonIntruction32Error)
			ran = !ok
		} else {
			ran = true
		}
	}

	if !ran {
		result = processor.handelUnknownInstruction(instruction)
	}

	if priorDelay != nil {
		priorDelay(processor)
	}

	return result
}

func (processor *Mips32Processor) handelUnknownInstruction(instruction Instruction32) error {
	if processor.UnknownInstruction == nil {
		return UnknonIntruction32Error(instruction)
	}

	return processor.UnknownInstruction(processor, instruction)
}
