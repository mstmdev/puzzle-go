#include "funcdata.h"

TEXT ·Factorial(SB),$16-16
    NO_LOCAL_POINTERS
    MOVQ num+0(FP),AX
    CMPQ AX,$1
    JLE return_1
    JMP recursion

return_1:
    MOVQ $1,result+8(FP)
    RET

recursion:
    SUBQ $1,AX
    MOVQ AX,0(SP)
    CALL ·Factorial(SB)
    MOVQ 8(SP),CX
    MOVQ num+0(FP),AX
    IMULQ AX,CX
    MOVQ CX,result+8(FP)
    RET
