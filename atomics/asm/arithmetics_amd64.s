// func Abs(x float64) float64
TEXT ·Abs(SB), $0
    MOVQ   $(1<<63), BX
    MOVQ   BX, X0
    MOVSD  x+0(FP), X1
    ANDNPD X1, X0
    MOVSD  X0, ret+8(FP)
    RET

// func Sqrt(x float64) float64
TEXT ·Sqrt(SB), $0
	XORPS  X0, X0
	SQRTSD x+0(FP), X0
	MOVSD  X0, ret+8(FP)
	RET
