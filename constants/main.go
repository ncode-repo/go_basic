package main

import (
    "fmt"
    "math"
)

const constant string = "Const String"
func main(){
    fmt.Println(constant)
    const CONST = 10000000

    const d = 3e20/CONST
    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(CONST))

}