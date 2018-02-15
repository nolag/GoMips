package gomips

// RInstructionAction takes action based on an RInstructionAction
// Note that InstructionActions are expected to ignore the OpCode, and function and assume the processor called the correct function.
type RInstructionAction func(*Mips32Processor, RInstruction) error

// RunFromRInstruction creates a InstructionAction using the map functions.  If a function is not found, UnknonIntruction32Error is returned
func RunFromRInstruction(functions *[64]RInstructionAction) InstructionAction {
	return func(processor *Mips32Processor, instruction Instruction32) error {
		rinstruction := NewRInstruction(instruction)
		callback := functions[rinstruction.Function()]
		if callback == nil {
			return UnknonIntruction32Error(instruction)
		}

		return callback(processor, rinstruction)
	}
}
