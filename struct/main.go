package main
import "fmt"

type person struct{
    age int
    name string
}

func main(){
    fmt.Println(person{30,"sam"})
    fmt.Println(person{age:30,name:"ammya"})
    fmt.Println(person{name:"dhanya"})
    p:=person{35,"neel"}
    fmt.Println(p)
    pp:=&p
    fmt.Println(pp.name)
    pp.age=22
    fmt.Println(pp)
}