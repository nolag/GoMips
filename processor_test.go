package gomips_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/nolag/gomips"
)

func TestUnknonInstructionErrorReturnsCorrectString(t *testing.T) {
	// Given
	value := uint32(0xFEEDBE)
	err := UnknonInstructionError(value)
	expected := "Unknown instruction 0x00FEEDBE"

	// When
	actual := err.Error()

	// Then
	assert.Equal(t, expected, actual, "Wrong string representation of unknown instruction error")
}

func TestRunUint32MakesCallCorrectly(t *testing.T) {
	// Given
	anyOpCode := uint32(5)
	anyRemainingBits := uint32(1002)
	expectedValue := anyOpCode<<26 + anyRemainingBits
	errExpected := errors.New("Anything")

	processor := Processor{}
	failCall := func(*Processor, Instruction) error {
		t.Fatalf("Unexpected call with op code %v, processor must only call the instruction in the op code", anyOpCode)
		return nil
	}

	for i := 0; i < 64; i++ {
		processor.InstructionActions[i] = failCall
	}

	processor.InstructionActions[anyOpCode] = CreateVerifiedInstructionActionCall(
		t, errExpected, expectedValue, &processor)

	// When
	err := processor.RunUint32(expectedValue)

	// Then
	assert.Equal(t, errExpected, err, "The runner must return the value from the instruction action call")
}

func TestRunUint32ReturnsUnknownInstructionWhenInstructionWithNoCallback(t *testing.T) {
	// Given
	value := uint32(12343)
	errExpected := UnknonInstructionError(value)
	processor := Processor{}

	// When - Then
	RunProcessor(t, &processor, errExpected, value)
}
func TestRunUint32UnknownInstructionCallbackFailWhenNoActionFound(t *testing.T) {
	// Given
	processor := Processor{}
	errExpected := errors.New("Simething cool")
	value := uint32(214)
	processor.ErrorHandler = CreateVerifiedUnknownInstructionCallback(
		t, errExpected, &processor)

	// When - Then
	RunProcessor(t, &processor, errExpected, value)
}

func TestRunUint32UnknownInstructionCallbackFailWhenActionReturnsUnknownInstruction(t *testing.T) {
	// Given
	value := uint32(325)
	processor := Processor{}
	errExpected := errors.New("Some error that should be returned")

	processor.ErrorHandler = CreateVerifiedUnknownInstructionCallback(
		t, errExpected, &processor)

	for i := 0; i < 64; i++ {
		processor.InstructionActions[i] =
			CreateVerifiedInstructionActionCall(t, UnknonInstructionError(value), value, &processor)
	}

	// When - Then
	RunProcessor(t, &processor, errExpected, value)
}

func TestRunUint32ReturnsUnkownInstructionWhenActionReturnsUnknownInstructionAndNoCallbackSet(t *testing.T) {
	// Given
	value := uint32(325)
	processor := Processor{}
	errExpected := UnknonInstructionError(value)

	for i := 0; i < 64; i++ {
		processor.InstructionActions[i] =
			CreateVerifiedInstructionActionCall(t, errExpected, value, &processor)
	}

	// When - Then
	RunProcessor(t, &processor, errExpected, value)
}

func RunProcessor(t *testing.T, processor *Processor, errExpected error, value uint32) {
	// When
	err := processor.RunUint32(value)

	// Then
	assert.Equal(t, errExpected, err, "Wrong error returned form processor")
}

func CreateVerifiedInstructionActionCall(
	t *testing.T,
	errExpected error,
	expectedValue uint32,
	expectedProcessor *Processor) InstructionAction {

	return func(processor *Processor, instruction Instruction) error {
		assert.Equal(t, Instruction(expectedValue), instruction, "Callback made with wrong instruction")
		assert.Equal(t, expectedProcessor, processor, "Callback made for wrong processor")
		return errExpected
	}
}

func CreateVerifiedUnknownInstructionCallback(
	t *testing.T,
	errExpected error,
	expectedProcessor *Processor) ErrorHandler {
	handler := func(processor *Processor, cause ExceptionCause) error {
		assert.Equal(t, RI, cause, "Wrong cause of exception")
		return errExpected
	}

	return handler
}
