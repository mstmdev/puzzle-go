TEXT ·GetAge(SB),$0
    MOVQ age+16(FP),AX
    MOVQ AX,result+40(FP)
    RET

TEXT ·GetPassword(SB),$0
    MOVQ password_data+24(FP),AX
    MOVQ password_len+32(FP),BX
    MOVQ AX,result+40(FP)
    MOVQ BX,result+48(FP)
    RET
