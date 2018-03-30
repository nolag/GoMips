package gomips

// RInstructionAction takes action based on an RInstruction
// Note that RInstructionAction are expected to ignore the OpCode, and function and assume the processor called the correct function.
type RInstructionAction func(*Processor, RInstruction) error

// RunFromRInstruction creates a InstructionAct using the map functions.  If a function is not found, UnknonInstructionError is returned
func RunFromRInstruction(functions *[64]RInstructionAction) InstructionAction {
	return func(processor *Processor, instruction Instruction) error {
		rinstruction := NewRInstruction(instruction)
		callback := functions[rinstruction.Function()]
		if callback == nil {
			return UnknonInstructionError(instruction)
		}

		return callback(processor, rinstruction)
	}
}
