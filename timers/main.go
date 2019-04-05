package main

import "time"
import "fmt"

func main(){
    timer1 := time.NewTimer(time.Second * 2)
    <-timer1.C
    fmt.Println("Timer1 expired")

    timer2 := time.NewTimer(time.Second * 1)
    go func(){
        <-timer2.C
        fmt.Println("Timer2 expired")
    }()

    stop2 := timer2.Stop()
    if stop2{
        fmt.Println("Timer2 stopped")
    }
}