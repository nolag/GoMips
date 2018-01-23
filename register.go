package gomips

// Register32 represents a 32 bit register
type Register32 interface {
	// Value32 gets the 32 bit value held by this register
	Value32() uint32

	// SetValue32 sets the 32 bit value held by this register.  Returns true if it is actually set
	SetValue32(uint32) bool
}

// BasicRegister32 is a basic way to get or set 32 bit values
type BasicRegister32 struct {
	value uint32
}

// ZeroRegister32 is a register that will always return zero, but allows writes
type ZeroRegister32 struct{}

// Value32 gets the 32 bit value held by this register
func (register *BasicRegister32) Value32() uint32 {
	return register.value
}

// SetValue32 sets the 32 bit value held by this register.  Returns true if it is actually set
func (register *BasicRegister32) SetValue32(value uint32) bool {
	register.value = value
	return true
}

// Value32 gets the 32 bit value held by this register
func (*ZeroRegister32) Value32() uint32 {
	return 0
}

// SetValue32 sets the 32 bit value held by this register.  Returns true if it is actually set
func (*ZeroRegister32) SetValue32(uint32) bool {
	return false
}
