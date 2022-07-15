#include "textflag.h"

GLOBL ·intslicedata<>(SB),NOPTR,$40

DATA ·intslicedata<>+0(SB)/8,$101
DATA ·intslicedata<>+8(SB)/8,$102
DATA ·intslicedata<>+16(SB)/8,$0
DATA ·intslicedata<>+24(SB)/8,$0
DATA ·intslicedata<>+32(SB)/8,$0

GLOBL ·IntSlice(SB),NOPTR,$24 // reflect.SliceHeader

DATA ·IntSlice+0(SB)/8,$·intslicedata<>(SB) // reflect.SliceHeader.Data
DATA ·IntSlice+8(SB)/8,$3 // reflect.SliceHeader.Len
DATA ·IntSlice+16(SB)/8,$5 // reflect.SliceHeader.Cap
