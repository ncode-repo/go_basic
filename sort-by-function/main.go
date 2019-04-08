package main

import "fmt"
import "sort"

type byLength []string

func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[j], s[i] = s[i] , s[j]
}
func (s byLength) Less (i, j int) bool{
    return len(s[i]) < len(s[j])
}

func main(){
    s := [] string{"bannana","apple","kiwi","papaya"}
    fmt.Println("fruits=", s)
    sort.Sort(byLength(s))
    fmt.Println("sorted fruits=",s)
}