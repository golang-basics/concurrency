// func Sqrt(x float64) float64
TEXT ·Sqrt(SB),$0
	XORPS  X0, X0
	SQRTSD x+0(FP), X0
	MOVSD  X0, ret+8(FP)
	RET
