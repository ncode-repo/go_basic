package main
import "fmt"

func f(frm string){
    for i:=0; i<3; i++ {
        fmt.Println(frm,": ",i)
        fmt.Println(frm,": ",i)
        fmt.Println(frm,": ",i)
    }
}

func main(){
    f("direct")
    go f("goroutine")

    go func(msg string){
        fmt.Println(msg)
    }("going")
    fmt.Scanln()
    fmt.Println("Done.")
}