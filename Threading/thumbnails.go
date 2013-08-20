package main

import (
    //"github.com/nfnt/resize"
    "math/rand"
    "fmt"
)
// Can pick this up tomorrow

// Pick dimensions that we wish to resize our thumbnail to
func pickDimensions(bound int32) (uint32, uint32){
    var (
        height, width uint32 = uint32(rand.Int31n(bound) + 1), uint32(rand.Int31n(bound) + 1)
    )
    return height, width
}

func main() {
    fmt.Println(rand.Uint32())
}
