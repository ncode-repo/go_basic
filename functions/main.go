package main

import "fmt"

func add(a int, b int) int{
    return a+b
}
func plusplus(a,b,c int)int{
    return a+b+c
}
func main(){
    sum1:=add(4,5)
    fmt.Println("sum1: ", sum1)
    sum2:=plusplus(3,4,5)
    fmt.Println("sum2: ", sum2)
}