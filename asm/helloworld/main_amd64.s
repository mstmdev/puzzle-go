TEXT ·main(SB),$16-0
    MOVQ ·hello+0(SB),AX
    MOVQ AX,0(SP)
    MOVQ ·hello+8(SB),BX
    MOVQ BX,8(SP)
    //CALL runtime·printstring(SB) // compile error
    //CALL runtime·printnl(SB) // compile error
    CALL ·printStr(SB)
    RET
