package gomips

import (
	"encoding/binary"
	"errors"
	"testing"
)

func TestStepIncrementsPc(t *testing.T) {
	// Given
	anyPc := uint32(400)
	processor := Mips32Processor{}
	processor.ByteOrder = binary.LittleEndian
	processor.ProgramCounter = anyPc
	expectedPc := anyPc + 4

	pcVerify := func(*Mips32Processor, Instruction) error {
		if expectedPc != processor.ProgramCounter {
			t.Fatalf("Program counter must be changed before the instruction runs")
		}

		return nil
	}

	for i := 0; i < 64; i++ {
		processor.InstructionActions[i] = pcVerify
	}

	processor.Memory = make([]byte, 800)

	// When
	err := processor.Step()

	// Then
	if processor.ProgramCounter != expectedPc {
		t.Fatalf("Expected pc to be %v, but was %v", expectedPc, processor.ProgramCounter)
	}

	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
}

func TestStepChoosesCorrectActionWithCorrectInstruction(t *testing.T) {
	// Given
	anyOpCode := uint32(5)
	numCalls := 0
	anyRemainingBits := uint32(1002)
	expectedValue := anyOpCode<<26 + anyRemainingBits

	processor := Mips32Processor{}
	processor.ByteOrder = binary.LittleEndian
	failCall := func(*Mips32Processor, Instruction) error {
		t.Fatalf("Unexpected call with op code %v, processor must only call the instruction in the op code", anyOpCode)
		return nil
	}

	processor.ProgramCounter = 24

	verifiedCall := func(actualProcessor *Mips32Processor, instruction Instruction) error {
		numCalls++
		actual := uint32(instruction)
		if expectedValue != actual {
			t.Fatalf("Expected instruction %x but got %x", expectedValue, actual)
		}

		if &processor != actualProcessor {
			t.Fatalf("Processor sent to run instrucion must be the same processor running it")
		}

		return nil
	}

	for i := 0; i < 64; i++ {
		processor.InstructionActions[i] = failCall
	}

	processor.InstructionActions[anyOpCode] = verifiedCall

	processor.Memory = make([]byte, 400)
	processor.ByteOrder.PutUint32(processor.Memory[processor.ProgramCounter:], expectedValue)

	// When
	err := processor.Step()

	// Then
	if numCalls != 1 {
		t.Fatalf("Expected call handel instruction once, got it %v times ", numCalls)
	}

	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
}

func TestStepReturnsErrorFromCall(t *testing.T) {
	// Given
	processor := Mips32Processor{}
	processor.ByteOrder = binary.LittleEndian
	anyError := errors.New("Anything")
	errorFn := func(*Mips32Processor, Instruction) error { return anyError }

	for i := 0; i < 64; i++ {
		processor.InstructionActions[i] = errorFn
	}

	processor.Memory = make([]byte, 800)

	// When
	err := processor.Step()

	// Then
	if err != anyError {
		t.Fatalf("Expected error %v but got %v", anyError, err)
	}
}

func TestStepExecutesDelayActionAfterStepIfItIsSet(t *testing.T) {
	// Given
	processor := Mips32Processor{}
	processor.ByteOrder = binary.LittleEndian
	wasInstructionRun := false
	delayActionCallCount := 0

	thirdCall := func(*Mips32Processor, Instruction) error {
		if delayActionCallCount != 1 {
			t.Fatalf("Delayed action must be called after the proceeded instruciton")
		}

		return nil
	}

	secondCall := func(*Mips32Processor, Instruction) error {
		wasInstructionRun = true
		for i := 0; i < 64; i++ {
			processor.InstructionActions[i] = thirdCall
		}

		if processor.DelayAction != nil {
			t.Fatalf("Delay action must be cleared to allow this instruction to set one")
		}

		return nil
	}

	firstCall := func(*Mips32Processor, Instruction) error {
		for i := 0; i < 64; i++ {
			processor.InstructionActions[i] = secondCall
		}

		processor.DelayAction = func(actual *Mips32Processor) {
			if &processor != actual {
				t.Fatalf("Processor sent to run delayed instruction must be the same processor running it")
			}

			if !wasInstructionRun {
				t.Fatalf("Next instruction must be run before the delayed action")
			}

			delayActionCallCount++
		}

		return nil
	}

	for i := 0; i < 64; i++ {
		processor.InstructionActions[i] = firstCall
	}

	processor.Memory = make([]byte, 800)

	// When
	processor.Step()
	processor.Step()
	processor.Step()

	// Then
	if !wasInstructionRun {
		t.Fatalf("Instruction must still be run when there is a delayed action")
	}

	if delayActionCallCount != 1 {
		t.Fatalf("Delayed call must be called exactly once was called %v", delayActionCallCount)
	}
}
