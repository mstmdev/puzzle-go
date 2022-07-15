TEXT ·Max(SB),$0
    MOVQ num1+0(FP),AX
    MOVQ num2+8(FP),BX
    CMPQ AX,BX
    JGE return_num1
    MOVQ BX,result+16(FP)
    RET

return_num1:
    MOVQ AX,result+16(FP)
    RET

TEXT ·Factorial(SB),$0
    MOVQ $1,DX // result
    MOVQ num+0(FP),AX
    CMPQ AX,$0
    JE return
    MOVQ $1,CX // i

loop_condition:
    CMPQ CX,AX
    JLE loop_calc
    JMP return

loop_calc:
    IMULQ CX,DX
    ADDQ $1,CX // i++
    JMP loop_condition

return:
    MOVQ DX,result+8(FP)
    RET
