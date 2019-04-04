package main
import "fmt"

func main(){
    message := make(chan string)
    
    go func(){
        message<- "ping1"
        message<- "ping2"
        message<- "ping3"
        message<- "ping4"
    }()

    str := <-message
    fmt.Println(str)
    str =<-message
    fmt.Println(str)
    fmt.Println(<-message)
    fmt.Println(<-message)
}