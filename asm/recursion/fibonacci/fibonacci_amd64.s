#include "funcdata.h"

TEXT ·Fibonacci(SB),$16-16
    NO_LOCAL_POINTERS
    MOVQ n+0(FP),AX
    CMPQ AX,$0
    JE return_0
    CMPQ AX,$2
    JLE return_1
    JMP recursion

return_0:
    MOVQ $0,result+8(FP)
    RET

return_1:
    MOVQ $1,result+8(FP)
    RET

recursion:
    SUBQ $1,AX
    MOVQ AX,0(SP)
    CALL ·Fibonacci(SB)
    MOVQ 8(SP),DX
    MOVQ DX,result+8(FP) // storage Fibonacci(n-1)

    MOVQ n+0(FP),AX
    SUBQ $2,AX
    MOVQ AX,0(SP)
    CALL ·Fibonacci(SB)
    MOVQ result+8(FP),DX // fetch Fibonacci(n-1)
    ADDQ 8(SP),DX

    MOVQ DX,result+8(FP)
    RET
