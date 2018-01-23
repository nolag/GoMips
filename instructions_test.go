package gomips

import "testing"

// TestOpCode tests that opcodes are are extracted correctly
func TestOpCode(t *testing.T) {
	// Given
	anyOpCode := uint(0x27)
	instruction1, instruction2 := createInstructions(anyOpCode, 26, 6)

	// When
	opCode1 := instruction1.OpCode()
	opCode2 := instruction2.OpCode()

	// Then
	assertValue(t, uint(opCode1), anyOpCode)
	assertValue(t, uint(opCode2), anyOpCode)
}

func TestRs(t *testing.T) {
	anyOpCode := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyOpCode, 21, 5)

	// When
	rs1 := RsRtInstruction{instruction1}.Rs()
	rs2 := RsRtInstruction{instruction2}.Rs()

	// Then
	assertValue(t, uint(rs1), anyOpCode)
	assertValue(t, uint(rs2), anyOpCode)
}

func TestRt(t *testing.T) {
	anyOpCode := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyOpCode, 16, 5)

	// When
	rt1 := RsRtInstruction{instruction1}.Rt()
	rt2 := RsRtInstruction{instruction2}.Rt()

	// Then
	assertValue(t, uint(rt1), anyOpCode)
	assertValue(t, uint(rt2), anyOpCode)
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
	anyOpCode := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyOpCode, 11, 5)

	// When
	rd1 := RInstruction{RsRtInstruction{instruction1}}.Rd()
	rd2 := RInstruction{RsRtInstruction{instruction2}}.Rd()

	// Then
	assertValue(t, uint(rd1), anyOpCode)
	assertValue(t, uint(rd2), anyOpCode)
}
func TestSamt(t *testing.T) {
	anyOpCode := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyOpCode, 6, 5)

	// When
	samt1 := RInstruction{RsRtInstruction{instruction1}}.Samt()
	samt2 := RInstruction{RsRtInstruction{instruction2}}.Samt()

	// Then
	assertValue(t, uint(samt1), anyOpCode)
	assertValue(t, uint(samt2), anyOpCode)
}

func TestFunction(t *testing.T) {
	anyFunction := uint(0x1D)
	instruction1, instruction2 := createInstructions(anyFunction, 0, 6)

	// When
	function1 := RInstruction{RsRtInstruction{instruction1}}.Function()
	function2 := RInstruction{RsRtInstruction{instruction2}}.Function()

	// Then
	assertValue(t, uint(function1), anyFunction)
	assertValue(t, uint(function2), anyFunction)
}

func TestNewIInstruction(t *testing.T) {
	// Given
	anyInstruction := Instruction(0x821EF021)
	expected := IInstruction{RsRtInstruction{anyInstruction}}

	// When
	actual := NewIInstruction(anyInstruction)

	// Then
	if expected != actual {
		t.Fatalf("Expected: 0x%x, Got: 0x%x", expected.Instruction, actual.Instruction)
	}
}

func TestAddress(t *testing.T) {
	anyAddress := uint(0xFEED)
	instruction1, instruction2 := createInstructions(anyAddress, 0, 16)

	// When
	function1 := IInstruction{RsRtInstruction{instruction1}}.Address()
	function2 := IInstruction{RsRtInstruction{instruction2}}.Address()

	// Then
	assertValue(t, uint(function1), anyAddress)
	assertValue(t, uint(function2), anyAddress)
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

func assertValue(t *testing.T, actual uint, expected uint) {
	if actual != expected {
		t.Fatalf("Expected: 0x%x, Got: 0x%x", expected, actual)
	}
}
