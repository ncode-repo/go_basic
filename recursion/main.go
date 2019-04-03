package main
import "fmt"

func fact(n int)int{
    if n==0{
        return 1
    }else{
        return n*fact(n-1)
    }
}

func main(){
    i:=3
    fmt.Printf("Factorial of %d is ",i)
    fmt.Println(fact(i))
}