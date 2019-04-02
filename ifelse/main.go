package main

import "fmt"

func main(){
    if 7%2==0 {
        fmt.Println("EVEN")
    }else{
        fmt.Println("ODD")
    }

    if i:=9; i<0 {
        fmt.Println(i, " is -ve")
    } else if(i<10){
        fmt.Println(i, " is single digit")
    }else{
        fmt.Println(i, " has multiple digits")
    }
}