package main

import "fmt"
import "strings"

func Index( vs []string, t string) int {
    for i, s := range vs {
        if t == s{
            return i
        }
    }
    return -1
}

func Include(vs []string, t string) bool {
    return Index(vs, t) > -1
}
func Any (vs []string, f func(string)bool) bool{
    for _, v := range vs {
        if f(v){
            return true
        }
    }
    return false
}
func All(vs []string, f func(string) bool) bool{
    for _, v:= range vs {
        if f(v){
            return false
        }
    }
    return true
}
func Filter(vs []string, f func(string)bool) []string{
    s := make([]string, 0)

    for _,v := range vs {
        if f(v) {
            s = append(s, v)
        }
    }
    return s
}
func Map(vs []string, f func(s string) string) []string {
    vsf := make([]string, len(vs))
    for i, v := range vs {
        vsf[i] = f(v)
    }
    return vsf
}

func main(){
    var strs = []string {"peach","apple","chikoo","grape"}

    fmt.Println(Index(strs,"apple"))
    fmt.Println(Include(strs,"grape"))

    fmt.Println(Any(strs, 
        func(v string) bool{
            return strings.HasPrefix(v, "p")
        }))
    fmt.Println(All(strs,
        func(v string) bool{
            return strings.HasPrefix(v, "p")
        }))
    fmt.Println(Filter(strs,
        func(v string) bool{
            return strings.Contains(v, "e")
        }))
    fmt.Println(Map(strs, strings.ToUpper))
}