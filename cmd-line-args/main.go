package main

import "fmt"
import "os"

func main(){
    p1 := os.Args
    p2 := os.Args[1:]
    p3 := os.Args[3]

    fmt.Println(p1)
    fmt.Println(p2)
    fmt.Println(p3)
}