package gomips

const (
	// BLTZ branch if < 0
	BLTZ = 0

	// BGEZ branch if >= 0
	BGEZ = 1

	// BLTZL branch < 0 likely
	BLTZL = 2

	// BGEZL branch <= 0 likely
	BGEZL = 3

	// TGEI trap >= immeditate
	TGEI = 8

	// TGEIU trap >= immediate unsigned
	TGEIU = 9

	// TLTI trap < immediate
	TLTI = 10

	// TLTIU trap < immediate unsigned
	TLTIU = 11

	// TEQI trap = immediate
	TEQI = 12

	// TNEI trap != immediate
	TNEI = 14

	// BLTZAL branch if < 0 and link
	BLTZAL = 16

	// BGEZAL branch if >= 0 and link
	BGEZAL = 17

	// BLTZALL branch if < 0 and link likely
	BLTZALL = 18
)
