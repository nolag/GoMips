package gomips

import (
	"github.com/nolag/gocpu/instructions"
)

const (
	// RInst RInstruction
	RInst instructions.Uint6 = 0

	// BCond branch conditionally
	BCond = 1

	// J jump
	J = 2

	// JAL jump and link
	JAL = 3

	// BEQ branch =
	BEQ = 4

	// BNE branch not =
	BNE = 5

	// BLEZ branch if <= 0
	BLEZ = 6

	// BGTZ branch if > 0
	BGTZ = 7

	// ADDI add immediate
	ADDI = 8

	// ADDIU add immediate unsigned
	ADDIU = 9

	// SLTI set if < immediate
	SLTI = 10

	// SLTIU set if < immediate unsigned
	SLTIU = 11

	// ANDI and immediate
	ANDI = 12

	// ORI or immediate
	ORI = 13

	// XORI exclusive or immediate
	XORI = 14

	// LUI load upper immediate
	LUI = 15

	// COProc0 is the opcode for coprocessor 0
	COProc0 = 16

	// COProc1 is the opcode for coprocessor 1
	COProc1 = 17

	// COProc2 is the opcode for coprocessor 2
	COProc2 = 18

	// COProc3 is the opcode for coprocessor 3
	COProc3 = 19

	// BEQL branch = likely
	BEQL = 20

	// BNEL branch != likely
	BNEL = 21

	// BLEZL branch <= 0 likely
	BLEZL = 22

	// BGTZL branch > 0 likely
	BGTZL = 23

	// DADDI doubleword add immediate
	DADDI = 24

	// DADDIU doubleword add immediate unsigned
	DADDIU = 25

	// LDL load doubleword left
	LDL = 26

	// LDR load doubleword right
	LDR = 27

	// LB load byte
	LB = 32

	// LH load halfword
	LH = 33

	// LWL load word left
	LWL = 34

	// LW load word
	LW = 35

	// LBU load byte unsigned
	LBU = 36

	// LHU load halfword unsigned
	LHU = 37

	// LWR load word right
	LWR = 38

	// LWU load word unsigned
	LWU = 39

	// SB store byte
	SB = 40

	// SH store halfword
	SH = 41

	// SWL store word left
	SWL = 42

	// SW store word
	SW = 43

	// SDL store doubleword left
	SDL = 44

	// SDR store doubleword right
	SDR = 45

	// SWR store word right
	SWR = 46

	// CACHE cache
	CACHE = 47

	// LL load linked
	LL = 48

	// LWCProc1 load word from coprocessor 1
	LWCProc1 = 49

	// LWCProc2 load word from coprocessor 2
	LWCProc2 = 50

	// LWCProc3 load word from coprocessor 3
	LWCProc3 = 51

	// PREF prefetch indexed
	PREF = LWCProc3 // Coprocessor 3 was removed in MIP III and the instruciton was reused in IV

	// LLD load linked doubleword
	LLD = 52

	// LDC1 load doubleword to coprocessor 1
	LDC1 = 53

	// LDC2 load doubleword to coprocessor 1
	LDC2 = 54

	// LDC3 load doubleword to coprocessor 1
	LDC3 = 55

	// LD load doubleword
	LD = LDC3 // Coprocessor 3 was removed in MIP III and the instruciton was reused

	// SC store condional
	SC = 56

	// SWCProc1 store word from coprocessor 1
	SWCProc1 = 57

	// SWCProc2 store word from coprocessor 2
	SWCProc2 = 58

	// SWCProc3 store word from coprocessor 3
	SWCProc3 = 59

	// SCD store conditional doubleword
	SCD = 60

	// SCD1 store doubleword to coprocessor 1
	SCD1 = 61

	// SCD2 store doubleword to coprocessor 2
	SCD2 = 62

	// SCD3 store doubleword to coprocessor 3
	SCD3 = 63
)
