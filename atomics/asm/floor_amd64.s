#define Big		0x4330000000000000

// func Floor(x float64) float64
TEXT ·Floor(SB),$0
	MOVQ	x+0(FP), AX
	MOVQ	$~(1<<63), DX
	ANDQ	AX,DX
	SUBQ	$1,DX
	MOVQ    $(Big - 1), CX
	CMPQ	DX,CX
	JAE     isBig_floor
	MOVQ	AX, X0
	CVTTSD2SQ	X0, AX
	CVTSQ2SD	AX, X1
	CMPSD	X1, X0, 1
	MOVSD	$(-1.0), X2
	ANDPD	X2, X0
	ADDSD	X1, X0
	MOVSD	X0, ret+8(FP)
	RET
isBig_floor:
	MOVQ    AX, ret+8(FP)
	RET
