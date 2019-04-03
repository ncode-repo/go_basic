package main
import "fmt"

func zeroVal(i int){
    i=0
}
func zeroPtr(i *int){
    *i=0
}

func main(){
    i:=1
    fmt.Println("Intital:",i)
    zeroVal(i)
    fmt.Println("After zeroVal():",i)
    zeroPtr(&i)
    fmt.Println("After zeroPtr():",i)
    fmt.Println("address=",&i)
}