package main

import "fmt"

func main(){
    nums := [] int {1,2,3,4}
    sum :=0
    for _, v := range nums{
        sum+= v
    }
    fmt.Println("Sum= ", sum)

    for i, v:= range nums{
        if v==3{
            fmt.Println("Idx of 3 is ", i)
        }
    }

    kvs := map[string]string{"a":"apple", "b":"bannana"}
    for k,v:=range kvs{
        fmt.Printf("%s-> %s\n",k,v)
    }
    for k:=range kvs{
        fmt.Println("key:",k)
    }

    for i,a := range "go"{
        fmt.Println(i,a)
    }
}