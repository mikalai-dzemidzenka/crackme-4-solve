package main

import (
	"math/bits"
)

func foo(str string) string {
    fin := make([]byte,len(str))
    for i := uint(0); i < uint(len(str)); i++ {
        c := uint(str[i])
        c += i
        c ^= i
        hi := c >> 8
        lo := c
        _, r := bits.Div(hi,lo,26)
        c = r
        c += 0x41
        fin[i] = byte(c)
    }
    return string(fin)
}

func main() {
    print(foo("troner"))
}
