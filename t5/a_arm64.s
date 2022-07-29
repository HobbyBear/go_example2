TEXT ·Add(SB),$0-24
    MOVD  a+0(FP), R0
    MOVD  b+8(FP), R1
    ADD  R0,R1, R0
    MOVD R0 ,c+16(FP)
    RET

