package gomips

import "testing"

// TestOpCode tests that opcodes are are extracted correctly
func TestOpCode(t *testing.T) {
	// Given
	anyOpCode := uint8(0x27)
	instruction1, instruction2 := createInstructions(uint8(anyOpCode), 26, 6)

	// When
	opCode1 := instruction1.OpCode()
	opCode2 := instruction2.OpCode()

	// Then
	assertValue(t, uint8(opCode1), anyOpCode)
	assertValue(t, uint8(opCode2), anyOpCode)
}

func TestRs(t *testing.T) {
	anyOpCode := uint8(0x1D)
	instruction1, instruction2 := createInstructions(uint8(anyOpCode), 21, 5)

	// When
	rs1 := RsRtInstruction{instruction1}.Rs()
	rs2 := RsRtInstruction{instruction2}.Rs()

	// Then
	assertValue(t, uint8(rs1), anyOpCode)
	assertValue(t, uint8(rs2), anyOpCode)
}

func TestRt(t *testing.T) {
	anyOpCode := uint8(0x1D)
	instruction1, instruction2 := createInstructions(uint8(anyOpCode), 16, 5)

	// When
	rt1 := RsRtInstruction{instruction1}.Rt()
	rt2 := RsRtInstruction{instruction2}.Rt()

	// Then
	assertValue(t, uint8(rt1), anyOpCode)
	assertValue(t, uint8(rt2), anyOpCode)
}

func TestNewRInstruction(t *testing.T) {
	// Given
	anyInstruction := Instruction(0x821EF021)
	expected := RInstruction{RsRtInstruction{anyInstruction}}

	// When
	actual := NewRInstruction(anyInstruction)

	// Then
	if expected != actual {
		t.Fatalf("Expected: 0x%x, Got: 0x%x", expected.Instruction, actual.Instruction)
	}
}

func TestRd(t *testing.T) {
	anyOpCode := uint8(0x1D)
	instruction1, instruction2 := createInstructions(uint8(anyOpCode), 11, 5)

	// When
	rd1 := RInstruction{RsRtInstruction{instruction1}}.Rd()
	rd2 := RInstruction{RsRtInstruction{instruction2}}.Rd()

	// Then
	assertValue(t, uint8(rd1), anyOpCode)
	assertValue(t, uint8(rd2), anyOpCode)
}
func TestSamt(t *testing.T) {
	anyOpCode := uint8(0x1D)
	instruction1, instruction2 := createInstructions(uint8(anyOpCode), 6, 5)

	// When
	samt1 := RInstruction{RsRtInstruction{instruction1}}.Samt()
	samt2 := RInstruction{RsRtInstruction{instruction2}}.Samt()

	// Then
	assertValue(t, uint8(samt1), anyOpCode)
	assertValue(t, uint8(samt2), anyOpCode)
}

func TestFunction(t *testing.T) {
	anyOpCode := uint8(0x1D)
	instruction1, instruction2 := createInstructions(uint8(anyOpCode), 0, 6)

	// When
	function1 := RInstruction{RsRtInstruction{instruction1}}.Function()
	function2 := RInstruction{RsRtInstruction{instruction2}}.Function()

	// Then
	assertValue(t, uint8(function1), anyOpCode)
	assertValue(t, uint8(function2), anyOpCode)
}

func createInstructions(valuePart uint8, shift uint, size uint) (Instruction, Instruction) {
	mask := uint32(0)

	// Make sure that the implmentation is different than the test
	for i := uint(0); i < size; i++ {
		mask = mask<<1 + 1
	}

	valueAsInt := uint32(valuePart) << shift
	negated := 0xFFFFFFFF ^ mask<<shift | valueAsInt
	return Instruction(valueAsInt), Instruction(negated)
}

func assertValue(t *testing.T, actual uint8, expected uint8) {
	if actual != expected {
		t.Fatalf("Expected: 0x%x, Got: 0x%x", expected, actual)
	}
}
