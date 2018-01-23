package gomips

import "testing"

func TestBasicRegisterGetsValue(t *testing.T) {
	// Given
	anyValue := uint32(0xFEEDBAD1)
	reg := BasicRegister32{anyValue}

	// When - Then
	assertValue(t, uint(reg.Value32()), uint(anyValue))
}

func TestBasicRegisterSetsValue(t *testing.T) {
	// Given
	anyValue := uint32(0xFEEDBAD1)
	reg := BasicRegister32{}

	// When
	wasSet := reg.SetValue32(anyValue)

	// When - Then
	if !wasSet {
		t.Fatalf("Value must be set for basic registers")
	}

	assertValue(t, uint(reg.Value32()), uint(anyValue))
}

func TestZeroRegisterGetsZeroValue(t *testing.T) {
	// Given
	reg := ZeroRegister32{}

	// When - Then
	assertValue(t, uint(reg.Value32()), 0)
}

func TestZeroRegisterDoesNotSetValue(t *testing.T) {
	// Given
	anyValue := uint32(0xFEEDBAD1)
	reg := ZeroRegister32{}

	// When
	wasSet := reg.SetValue32(anyValue)

	// When - Then
	if wasSet {
		t.Fatalf("Value must not be set for zero register")
	}

	assertValue(t, uint(reg.Value32()), 0)
}
