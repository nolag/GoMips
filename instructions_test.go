package gomips

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestOpCode tests that opcodes are are extracted correctly
func TestOpCode(t *testing.T) {
	// Given
	anyOpCode := uint(0x27)
	instruction1, instruction2 := createInstructions(anyOpCode, 26, 6)

	// When
	opCode1 := instruction1.OpCode()
	opCode2 := instruction2.OpCode()

	// Then
	assert.Equal(t, anyOpCode, uint(opCode1))
	assert.Equal(t, anyOpCode, uint(opCode2))
}

func TestRs(t *testing.T) {
	anyOpCode := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyOpCode, 21, 5)

	// When
	rs1 := RsRtInstruction{instruction1}.Rs()
	rs2 := RsRtInstruction{instruction2}.Rs()

	// Then
	assert.Equal(t, anyOpCode, uint(rs1))
	assert.Equal(t, anyOpCode, uint(rs2))
}

func TestRt(t *testing.T) {
	anyOpCode := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyOpCode, 16, 5)

	// When
	rt1 := RsRtInstruction{instruction1}.Rt()
	rt2 := RsRtInstruction{instruction2}.Rt()

	// Then
	assert.Equal(t, anyOpCode, uint(rt1))
	assert.Equal(t, anyOpCode, uint(rt2))
}

func TestNewRInstruction(t *testing.T) {
	// Given
	anyInstruction := Instruction(0x821EF021)
	expected := RInstruction{RsRtInstruction{anyInstruction}}

	// When
	actual := NewRInstruction(anyInstruction)

	// Then
	assert.Equal(t, expected, actual)
}

func TestRd(t *testing.T) {
	anyValue := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyValue, 11, 5)

	// When
	rd1 := RInstruction{RsRtInstruction{instruction1}}.Rd()
	rd2 := RInstruction{RsRtInstruction{instruction2}}.Rd()

	// Then
	assert.Equal(t, anyValue, uint(rd1))
	assert.Equal(t, anyValue, uint(rd2))
}
func TestSamt(t *testing.T) {
	anyValue := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyValue, 6, 5)

	// When
	samt1 := RInstruction{RsRtInstruction{instruction1}}.Samt()
	samt2 := RInstruction{RsRtInstruction{instruction2}}.Samt()

	// Then
	assert.Equal(t, anyValue, uint(samt1))
	assert.Equal(t, anyValue, uint(samt2))
}

func TestFunction(t *testing.T) {
	anyFunction := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyFunction, 0, 6)

	// When
	function1 := RInstruction{RsRtInstruction{instruction1}}.Function()
	function2 := RInstruction{RsRtInstruction{instruction2}}.Function()

	// Then
	assert.Equal(t, anyFunction, uint(function1))
	assert.Equal(t, anyFunction, uint(function2))
}

func TestNewIInstruction(t *testing.T) {
	// Given
	anyInstruction := Instruction(0x821EF021)
	expected := IInstruction{RsRtInstruction{anyInstruction}}

	// When
	actual := NewIInstruction(anyInstruction)

	// Then
	assert.Equal(t, expected, actual)
}

func TestAddress(t *testing.T) {
	anyAddress := uint(0xFEED)
	instruction1, instruction2 := createInstructions(anyAddress, 0, 16)

	// When
	function1 := IInstruction{RsRtInstruction{instruction1}}.Address()
	function2 := IInstruction{RsRtInstruction{instruction2}}.Address()

	// Then
	assert.Equal(t, anyAddress, uint(function1))
	assert.Equal(t, anyAddress, uint(function2))
}

func createInstructions(valuePart uint, shift uint, size uint) (Instruction, Instruction) {
	mask := uint32(0)

	// Make sure that the implmentation is different than the test
	for i := uint(0); i < size; i++ {
		mask = mask<<1 + 1
	}

	valueAsInt := uint32(valuePart) << shift
	negated := 0xFFFFFFFF ^ mask<<shift | valueAsInt
	return Instruction(valueAsInt), Instruction(negated)
}
