package main

import "fmt"
import "sort"

func main(){
    strs := []string{"b","d","a","c"}
    fmt.Println("strs:",strs)
    sort.Strings(strs)
    fmt.Println("sorted strs:", strs)

    ints := [] int{2,5,6,3,1,4}
    fmt.Println("ints:",ints)
    isSorted := sort.IntsAreSorted(ints)
    fmt.Println("Is ints sorted:",isSorted)
    sort.Ints(ints)
    fmt.Println("sorted ints:",ints)
}