package gomips_test

import (
	"testing"

	"github.com/nolag/gocpu/instructions"
	. "github.com/nolag/gomips"
	"github.com/stretchr/testify/assert"
)

func TestSetOpCodeSet(t *testing.T) {
	before := Instruction(0xFEEDBEEF)
	builder := NewInstructionBuilder(before)
	anyOpCode := instructions.Uint6(0x27)
	mask := mask(26)

	result := builder.SetOpCode(anyOpCode)
	after := builder.Build()

	assert.True(t, result, "Opcode is in range, value must be set")
	assert.Equal(t, anyOpCode, after.OpCode(), "OpCode does not match after setting")
	assert.Equal(t, before&mask, after&mask, "Other bits must not be overwritten")
}

func TestSetOpCodeOutOfRange(t *testing.T) {
	before := Instruction(0xFEEDBEEF)
	builder := NewInstructionBuilder(before)
	tooLargeValue := instructions.Uint6(0xFF)

	result := builder.SetOpCode(tooLargeValue)
	after := builder.Build()

	assert.False(t, result, "Opcode is not in range, value must not be set")
	assert.Equal(t, before, after, "Value must not be changed when out of range")
}

func TestSetRs(t *testing.T) {
	before := RsRtInstruction{Instruction: Instruction(0xFEEDBEEF)}
	builder := NewRsRtInstructionBuilder(before)
	anyRs := instructions.Uint5(0x15)
	mask := mask(16)

	result := builder.SetRs(anyRs)
	after := builder.Build()

	assert.True(t, result, "Rs is in range, value must be set")
	assert.Equal(t, before.OpCode(), after.OpCode(), "OpCode must not be overwritten")
	assert.Equal(t, anyRs, after.Rs(), "Rs must be overwritten")
	assert.Equal(t, before.Rt(), after.Rt(), "Rt must not be overwritten")
	assert.Equal(t, before.Instruction&mask, after.Instruction&mask, "Other bits must not be overwritten")
}

func TestSetRsOutOfRange(t *testing.T) {
	before := RsRtInstruction{Instruction: Instruction(0xFEEDBEEF)}
	builder := NewRsRtInstructionBuilder(before)
	tooLargeValue := instructions.Uint5(0xFF)

	result := builder.SetRs(tooLargeValue)
	after := builder.Build()

	assert.False(t, result, "Rs is not in range, value must not be set")
	assert.Equal(t, before, after, "Value must not be changed when out of range")
}

func TestSetRt(t *testing.T) {
	before := RsRtInstruction{Instruction: Instruction(0xFEEDBEEF)}
	builder := NewRsRtInstructionBuilder(before)
	anyRt := instructions.Uint5(0x15)
	mask := mask(16)

	result := builder.SetRt(anyRt)
	after := builder.Build()

	assert.True(t, result, "Opcode is in range, value must be set")
	assert.Equal(t, before.OpCode(), after.OpCode(), "OpCode must not be overwritten")
	assert.Equal(t, before.Rs(), after.Rs(), "Rs must not be overwritten")
	assert.Equal(t, anyRt, after.Rt(), "Rt must be overwritten")
	assert.Equal(t, before.Instruction&mask, after.Instruction&mask, "Other bits must not be overwritten")
}

func TestSetRtOutOfRange(t *testing.T) {
	before := RsRtInstruction{Instruction: Instruction(0xFEEDBEEF)}
	builder := NewRsRtInstructionBuilder(before)
	tooLargeValue := instructions.Uint5(0xFF)

	result := builder.SetRt(tooLargeValue)
	after := builder.Build()

	assert.False(t, result, "Rt is not in range, value must not be set")
	assert.Equal(t, before, after, "Value must not be changed when out of range")
}

func TestSetRd(t *testing.T) {
	before := NewRInstruction(0xFEEDBEEF)
	builder := NewRInstructionBuilder(before)
	anyRd := instructions.Uint5(0x15)

	result := builder.SetRd(anyRd)
	after := builder.Build()

	assert.True(t, result, "Rd is in range")
	assert.Equal(t, before.OpCode(), after.OpCode(), "OpCode must not be overwritten")
	assert.Equal(t, before.Rs(), after.Rs(), "Rs must not be overwritten")
	assert.Equal(t, before.Rt(), after.Rt(), "Rs must not be overwritten")
	assert.Equal(t, anyRd, after.Rd(), "Rd must be overwritten")
	assert.Equal(t, before.Shamt(), after.Shamt(), "Shamt must not be overwritten")
	assert.Equal(t, before.Funct(), after.Funct(), "Function must not be overwritten")
}

func TestSetRdOutOfRange(t *testing.T) {
	before := NewRInstruction(0xFEEDBEEF)
	builder := NewRInstructionBuilder(before)
	tooLargeValue := instructions.Uint5(0xFF)

	result := builder.SetRd(tooLargeValue)
	after := builder.Build()

	assert.False(t, result, "Rd is not in range, value must not be set")
	assert.Equal(t, before, after, "Value must not be changed when out of range")
}

func TestSetShamt(t *testing.T) {
	before := NewRInstruction(0xFEEDBEEF)
	builder := NewRInstructionBuilder(before)
	anyShamt := instructions.Uint5(0x15)

	result := builder.SetShamt(anyShamt)
	after := builder.Build()

	assert.True(t, result, "Rd is in range")
	assert.Equal(t, before.OpCode(), after.OpCode(), "OpCode must not be overwritten")
	assert.Equal(t, before.Rs(), after.Rs(), "Rs must not be overwritten")
	assert.Equal(t, before.Rt(), after.Rt(), "Rs must not be overwritten")
	assert.Equal(t, before.Rd(), after.Rd(), "Rd must be overwritten")
	assert.Equal(t, anyShamt, after.Shamt(), "Shamt must not be overwritten")
	assert.Equal(t, before.Funct(), after.Funct(), "Function must not be overwritten")
}

func TestSetShamtOutOfRange(t *testing.T) {
	before := NewRInstruction(0xFEEDBEEF)
	builder := NewRInstructionBuilder(before)
	tooLargeValue := instructions.Uint5(0xFF)

	result := builder.SetShamt(tooLargeValue)
	after := builder.Build()

	assert.False(t, result, "Shamt is not in range, value must not be set")
	assert.Equal(t, before, after, "Value must not be changed when out of range")
}

func TestSetFunct(t *testing.T) {
	before := NewRInstruction(0xFEEDBEEF)
	builder := NewRInstructionBuilder(before)
	anyFunct := instructions.Uint6(0x27)

	result := builder.SetFunct(anyFunct)
	after := builder.Build()

	assert.True(t, result, "Rd is in range")
	assert.Equal(t, before.OpCode(), after.OpCode(), "OpCode must not be overwritten")
	assert.Equal(t, before.Rs(), after.Rs(), "Rs must not be overwritten")
	assert.Equal(t, before.Rt(), after.Rt(), "Rs must not be overwritten")
	assert.Equal(t, before.Rd(), after.Rd(), "Rd must be overwritten")
	assert.Equal(t, before.Shamt(), after.Shamt(), "Shamt must not be overwritten")
	assert.Equal(t, anyFunct, after.Funct(), "Function must not be overwritten")
}

func TestSetFunctOutOfRange(t *testing.T) {
	before := NewRInstruction(0xFEEDBEEF)
	builder := NewRInstructionBuilder(before)
	tooLargeValue := instructions.Uint6(0xFF)

	result := builder.SetFunct(tooLargeValue)
	after := builder.Build()

	assert.False(t, result, "Funct is not in range, value must not be set")
	assert.Equal(t, before, after, "Value must not be changed when out of range")
}

func mask(size int) Instruction {
	if size == 0 {
		return 0
	}

	value := Instruction(1)

	for i := 1; i < size; i++ {
		value <<= 1
		value |= 1
	}

	return value
}
