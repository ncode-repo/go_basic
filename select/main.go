package main
import "fmt"
import "time"

func main(){
    c1 := make(chan string)
    c2 := make(chan string)

    go func(c1 chan string){
        time.Sleep(1 * time.Second)
        c1 <- "One"
    }(c1)
    go func(c2 chan string)  {
        time.Sleep(2 * time.Second)
        c2 <- "Two"
    }(c2)

    for i:=0;i<2;i++ {
        select{
        case msg1 := <- c1:
            fmt.Println("recieved..",msg1)
        case msg2 := <- c2:
            fmt.Println("recieved..",msg2)
        }

    }
}