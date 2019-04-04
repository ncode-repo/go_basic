package main
import "fmt"

func main(){
    msg := make(chan string,2)

    //go func(){ 
        msg <- "ping1"
        msg <- "ping2"
        //msg <- "ping3"
    //}()

    fmt.Println(<-msg)
    fmt.Println(<-msg)
    //fmt.Println(<-msg)
}