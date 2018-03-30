package gomips_test

import (
	"errors"
	"testing"

	. "github.com/nolag/gomips"
)

func TestRunFromRInstructionMapsToCorrectInstruction(t *testing.T) {
	// Given
	instruction := Instruction(0xFEEDBABE)
	rinstruction := NewRInstruction(instruction)
	expectedRun := int(rinstruction.Function())
	processor := Processor{}

	expectedReturn := errors.New("Any error")
	var actions [64]RInstructionAction
	for i := 0; i < 64; i++ {
		x := i
		actions[i] = func(*Processor, RInstruction) error {
			t.Fatalf("Must only run action for %v, ran for %v", expectedRun, x)
			return nil
		}
	}

	actions[expectedRun] = func(actualProc *Processor, actualInstr RInstruction) error {
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
	err := wrapped(&processor, Instruction(instruction))

	// Then
	if err != expectedReturn {
		t.Fatalf("Expected return of %v got %v", expectedReturn, err)
	}
}

func TestRunFromRInstructionReturnsCorrectErrorIfInstructionisNil(t *testing.T) {
	// Given
	instruction := Instruction(0xFEEDBABE)
	processor := Processor{}
	var actions [64]RInstructionAction
	wrapped := RunFromRInstruction(&actions)

	// When
	err := wrapped(&processor, Instruction(instruction))
	err32, ok := err.(UnknonInstructionError)

	// Then
	if !ok || err32 != UnknonInstructionError(instruction) {
		t.Fatalf("Expected unknown instruction error for %v error but got %v", instruction, err)
	}
}
