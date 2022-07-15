#include "textflag.h"

GLOBL ·hellodata<>(SB),NOPTR,$16 // <> suffix makes the name visible only in the current source file
DATA ·hellodata<>+0(SB)/16,$"hello world"

GLOBL ·Hello(SB),NOPTR,$16 // reflect.StringHeader
DATA ·Hello+0(SB)/8,$·hellodata<>(SB) // reflect.StringHeader.Data
DATA ·Hello+8(SB)/8,$11 // reflect.StringHeader.Len
