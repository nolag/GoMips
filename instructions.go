package gomips

// Uint6 represents an unsigned 6 bit integer.
type Uint6 uint8

// Uint5 represents an unsigned 5 bit integer.
type Uint5 uint8

// Uint3 represents an unsigned 3 bit integer.
type Uint3 uint8

// Uint26 represents a 26 bit integer.
type Uint26 uint32

// Instruction represents a 32 bit MIPs instruction
type Instruction uint32

// RsRtInstruction represents an instruction that has an rs and rt value, namely an IInstruction or RInstruction.
type RsRtInstruction struct {
	Instruction
}

// RInstruction represents an R instruction
type RInstruction struct {
	RsRtInstruction
}

// IInstruction represents an I instruction
type IInstruction struct {
	RsRtInstruction
}

// JInstruction represents a J instruction
type JInstruction struct {
	Instruction
}

// RInstructionBuilder provides methods to build a RInstruction
type RInstructionBuilder struct {
	RInstruction
}

// OpCode reads the op code from the instruction
func (instruction Instruction) OpCode() Uint6 {
	return Uint6(instruction >> 26)
}

// Rs returns the RS bits of an instruction.
func (instruction RsRtInstruction) Rs() Uint5 {
	return Uint5(instruction.Instruction >> 21 & 0x1F)
}

// Rt returns the Rt bits of an instruction.
func (instruction RsRtInstruction) Rt() Uint5 {
	return Uint5(instruction.Instruction >> 16 & 0x1F)
}

// NewRInstruction is a shortcut to RInstruction{RsRtInstruction{instruction}}
func NewRInstruction(instruction Instruction) RInstruction {
	return RInstruction{RsRtInstruction{instruction}}
}

// Rd returns the Rd bits of an instruction
func (instruction RInstruction) Rd() Uint5 {
	return Uint5(instruction.Instruction >> 11 & 0x1F)
}

// Samt returns the shift amount bits of an instruction
func (instruction RInstruction) Samt() Uint5 {
	return Uint5(instruction.Instruction >> 6 & 0x1F)
}

// Function returns the function from the instruction
func (instruction RInstruction) Function() Uint6 {
	return Uint6(instruction.Instruction & 0x3F)
}

// NewIInstruction is a shortcut to IInstruction{RsRtInstruction{instruction}}
func NewIInstruction(instruction Instruction) IInstruction {
	return IInstruction{RsRtInstruction{instruction}}
}

// Immediate returns the immediate value from the instruction
func (instruction IInstruction) Immediate() uint16 {
	return uint16(instruction.Instruction)
}

// Address returns the address part of the instruction
func (instruction JInstruction) Address() Uint26 {
	return Uint26(instruction.Instruction & 0x3FFFFFF)
}
