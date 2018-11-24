package gomips

import "github.com/nolag/gocpu/instructions"

// Coprocessor is a MIPs co-processor
type Coprocessor interface {
	LoadWord(register instructions.Uint5, bank instructions.Uint3) uint32
	StoreWord(register instructions.Uint5, bank instructions.Uint3, value uint32)
}
