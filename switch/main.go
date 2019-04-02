package main

import "fmt"
import "time"

func main(){
    i :=2
    fmt.Print("Write ", i, " as ")
    switch i {
        case 1: fmt.Println("One")
        case 2: fmt.Println("Two")
        case 3: fmt.Println("Three")
    }

    switch time.Now().Weekday(){
        case time.Saturday, time.Sunday:
                fmt.Println("This is weekend")
        default:
                fmt.Println("It's weekday.")
    }

    hr := time.Now()
    switch{
    case hr.Hour() < 12 :
        fmt.Println("It's a before noon")
    default:
        fmt.Println("It's after noon")
    }

    whoAmI := func (i interface {}){
        switch t:=i.(type){
            case int: fmt.Println("I'm an int")
            case bool : fmt.Println("I'm bool")
            default: fmt.Printf("Don't know type %T\n", t)
        }
    }
    whoAmI(4)
    whoAmI(true)
    whoAmI("hey")
}