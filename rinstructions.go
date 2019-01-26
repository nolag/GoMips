package gomips

import (
	"github.com/nolag/gocpu/instructions"
)

const (
	// SLL shift left logical
	SLL instructions.Uint6 = 0

	// SRL shift right logical
	SRL = 2

	// SAR shift arithmetic right
	SAR = 3

	// SLLV shift left logical variable
	SLLV = 4

	// SRLV shift right logical variable
	SRLV = 6

	// SRAV shift right arithmetic variable
	SRAV = 7

	// JR jump register
	JR = 8

	// JALR jump and link register
	JALR = 9

	// MOVZ move on zero
	MOVZ = 10

	// MOVN move on not zero
	MOVN = 11

	// SYSCALL system call
	SYSCALL = 12

	// BREAK break point
	BREAK = 13

	// SYNC sync memory
	SYNC = 15

	// MFHI move from high
	MFHI = 16

	// MTHI move to high
	MTHI = 17

	// MFLO move from low
	MFLO = 18

	// MTLO move to low
	MTLO = 19

	// DSLLV doubleword shift left logical variable
	DSLLV = 20

	// DSRLV doubleword shift rigth logical variable
	DSRLV = 22

	// DSRAV doubleword shift right arithmetic variable
	DSRAV = 23

	// MULT multiply
	MULT = 24

	// MULTU multiply unsigned
	MULTU = 25

	// DIV divide
	DIV = 26

	// DIVU divide unsigned
	DIVU = 27

	// DMULT doubleword multiply
	DMULT = 28

	// DMULTU doubleword multiply unsigned
	DMULTU = 29

	// DDIV doubleword divide
	DDIV = 30

	// DDIVU doubleword divide unsigned
	DDIVU = 30

	// ADD add
	ADD = 32

	// ADDU add unsigned
	ADDU = 33

	// SUB subtract
	SUB = 34

	// SUBU substract unsigned
	SUBU = 35

	// AND and
	AND = 36

	// OR or
	OR = 37

	// XOR exclusive or
	XOR = 38

	// NOR not or
	NOR = 39

	// SLT set if <
	SLT = 42

	// SLTU set if <=
	SLTU = 43

	// DADD doubleword add
	DADD = 44

	// DADDU doubleword add unsigned
	DADDU = 45

	// DSUB doubleword subtract
	DSUB = 46

	// DSUBU doubleword subtract unsigned
	DSUBU = 47

	// TGE trap if >=
	TGE = 48

	// TGEU trap if >= unsigned
	TGEU = 48

	// TLT trap if <
	TLT = 50

	// TLTU trap if < unsigned
	TLTU = 51

	// TEQ trap if =
	TEQ = 52

	// TNE trap if !=
	TNE = 54

	// DSLL doubleword shift left logical
	DSLL = 56

	// DSRL doubleword shift right logical
	DSRL = 58

	// DSRA doubleword shift right arithmetic
	DSRA = 59

	// DSLL32 doubleword shift left logical plus 32
	DSLL32 = 60

	// DSRL32 doubleword shift right logical plus 32
	DSRL32 = 61

	// DSRA32 doubleword shift right arithmetic plus 32
	DSRA32 = 62
)
