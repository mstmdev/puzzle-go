TEXT ·Add(SB),$0
    MOVQ num1+0(FP),AX
    MOVQ num2+8(FP),BX
    ADDQ BX,AX
    MOVQ AX,result+16(FP)
    RET

TEXT ·Sub(SB),$0
    MOVQ num1+0(FP),AX
    MOVQ num2+8(FP),BX
    SUBQ BX,AX
    MOVQ AX,result+16(FP)
    RET

TEXT ·Mul(SB),$0
    MOVQ num1+0(FP),AX
    MOVQ num2+8(FP),BX
    IMULQ BX,AX
    MOVQ AX,result+16(FP)
    RET
