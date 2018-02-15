package gomips

import (
	"errors"
	"testing"
)

func TestRunFromRInstructionMapsToCorrectInstruction(t *testing.T) {
	// Given
	instruction := Instruction32(0xFEEDBABE)
	rinstruction := NewRInstruction(instruction)
	expectedRun := int(rinstruction.Function())
	processor := Mips32Processor{}

	expectedReturn := errors.New("Any error")
	var actions [64]RInstructionAction
	for i := 0; i < 64; i++ {
		x := i
		actions[i] = func(*Mips32Processor, RInstruction) error {
			t.Fatalf("Must only run action for %v, ran for %v", expectedRun, x)
			return nil
		}
	}

	actions[expectedRun] = func(actualProc *Mips32Processor, actualInstr RInstruction) error {
		if actualInstr != rinstruction {
			t.Fatalf("Expected to run on instruction %v got %v", rinstruction, actualInstr)
		}

		if actualProc != &processor {
			t.Fatalf("Wrong processor passed to run the RInstructions")
		}

		return expectedReturn
	}

	wrapped := RunFromRInstruction(&actions)

	// When
	err := wrapped(&processor, Instruction32(instruction))

	// Then
	if err != expectedReturn {
		t.Fatalf("Expected return of %v got %v", expectedReturn, err)
	}
}

func TestRunFromRInstructionReturnsCorrectErrorIfInstructionisNil(t *testing.T) {
	// Given
	instruction := Instruction32(0xFEEDBABE)
	processor := Mips32Processor{}
	var actions [64]RInstructionAction
	wrapped := RunFromRInstruction(&actions)

	// When
	err := wrapped(&processor, Instruction32(instruction))
	err32, ok := err.(UnknonIntruction32Error)

	// Then
	if !ok || err32 != UnknonIntruction32Error(instruction) {
		t.Fatalf("Expected unknown instruction error for %v error but got %v", instruction, err)
	}
}
