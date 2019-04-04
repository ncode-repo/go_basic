package main

import "fmt"

func main(){
    message := make(chan string)
    signal := make(chan bool)

    select{
    case msg := <- message:
        fmt.Println("Recieved message",msg)
    default:
        fmt.Println("No message received")
    }
    msg:= "hi"
    select {
    case message <- msg :
        fmt.Println("Message sent")
    default:
        fmt.Println("Message not sent")
    }

    select{
    case msg:= <- message:
        fmt.Println("message recieved", msg)
    case s:= <- signal:
        fmt.Println("signal recieved", s)
    default:
        fmt.Println("no activity")
    }
}