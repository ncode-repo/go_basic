package main
import "fmt"

func intSeq() func()int{
    i:=1
    return func()int{
        i++
        return i
    }
}

func main(){
    nxtInt:= intSeq()

    fmt.Println("nxtInt:",nxtInt())
    fmt.Println("nxtInt:",nxtInt())
    fmt.Println("nxtInt:",nxtInt())

    newInt:=intSeq()
    fmt.Println("newInt:",newInt())
}