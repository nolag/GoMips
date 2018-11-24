package gomips

import "github.com/nolag/gocpu/instructions"

type verifiable interface {
	Verify() bool
}

// InstructionBuilder builds an instruction
type InstructionBuilder struct {
	instruction Instruction
}

// RsRtInstructionBuilder builds an RsRtInstruction
type RsRtInstructionBuilder struct {
	InstructionBuilder
}

// IInstructionBuilder builds up an IInstruction
type IInstructionBuilder struct {
	RsRtInstructionBuilder
}

// RInstructionBuilder builds up a RInstruction
type RInstructionBuilder struct {
	RsRtInstructionBuilder
}

// JInstructionBuilder builds up a JInstruction
type JInstructionBuilder struct {
	InstructionBuilder
}

// SetOpCode sets the op code portion of an instruction.
// Returns true if successful, false if overflow would occur and the operation was canceled
func (builder *InstructionBuilder) SetOpCode(value instructions.Uint6) bool {
	return builder.setValue(value, uint32(value)<<26, 0x3FFFFFF)
}

// Build returns the value held by this builder
func (builder *InstructionBuilder) Build() Instruction {
	return builder.instruction
}

// SetRs sets the Rs value
func (builder *RsRtInstructionBuilder) SetRs(value instructions.Uint5) bool {
	return builder.setValue(value, uint32(value)<<21, 0xFC1FFFFF)
}

// SetRt sets the Rt value
func (builder *RsRtInstructionBuilder) SetRt(value instructions.Uint5) bool {
	return builder.setValue(value, uint32(value)<<16, 0xFFE0FFFF)
}

// Build returns the value held by this builder
func (builder *RsRtInstructionBuilder) Build() RsRtInstruction {
	return RsRtInstruction{builder.InstructionBuilder.Build()}
}

// SetRd sets the Rd value
func (builder *RInstructionBuilder) SetRd(value instructions.Uint5) bool {
	return builder.setValue(value, uint32(value)<<11, 0xFFFF07FF)
}

// SetShamt sets the SetShamt value
func (builder *RInstructionBuilder) SetShamt(value instructions.Uint5) bool {
	return builder.setValue(value, uint32(value)<<6, 0xFFFFF83F)
}

// SetFunct sets the Function value
func (builder *RInstructionBuilder) SetFunct(value instructions.Uint6) bool {
	return builder.setValue(value, uint32(value), 0xFFFFFFC0)
}

// Build returns the value held by this builder
func (builder *RInstructionBuilder) Build() RInstruction {
	return NewRInstruction(builder.InstructionBuilder.Build())
}

// NewInstructionBuilder returns a builder for Instruction that starts with provided instruction
func NewInstructionBuilder(instruction Instruction) *InstructionBuilder {
	return &InstructionBuilder{instruction}
}

// NewRsRtInstructionBuilder returns a builder for RsRtInstruction that starts with provided instruction
func NewRsRtInstructionBuilder(instruction RsRtInstruction) *RsRtInstructionBuilder {
	return &RsRtInstructionBuilder{InstructionBuilder{instruction.Instruction}}
}

// NewRInstructionBuilder returns a builder for RInstruction that starts with provided instruction
func NewRInstructionBuilder(instruction RInstruction) *RInstructionBuilder {
	return &RInstructionBuilder{RsRtInstructionBuilder{InstructionBuilder{instruction.Instruction}}}
}

func (builder *InstructionBuilder) setValue(value verifiable, valueAsinstructions uint32, mask uint32) bool {
	if !value.Verify() {
		return false
	}

	builder.instruction &= Instruction(mask)
	builder.instruction |= Instruction(valueAsinstructions)
	return true
}
