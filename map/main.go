package main

import "fmt"

func main(){
    m := make(map[string]int)

    m["k1"]= 10
    m["k2"]=20

    fmt.Println("map: ", m)

    v:= m["k1"]
    fmt.Println("k1 val: ", v)

    fmt.Println("len(m)=",len(m))

    delete(m, "k2")
    fmt.Println("updated map:", m)

    _, prs:= m["k2"]
    fmt.Println("prs: ", prs)

    n := map[string]int {"foo":1, "bar":2}
    fmt.Println("inline map: ",n)


}