TEXT ·StoreInt64(SB),$0-16
    MOVQ	ptr+0(FP), BX
    MOVQ	val+8(FP), AX
    XCHGQ	AX, 0(BX)
    RET
