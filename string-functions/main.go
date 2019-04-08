package main

import s "strings"
import "fmt"

var p = fmt.Println
func main(){
    p("Contains:\t",s.Contains("test","es"))
    p("Count:\t\t", s.Count("test","t"))
    p("HasPrefix:\t", s.HasPrefix("test","te"))
    p("HasSuffix:\t", s.HasSuffix("test","st"))
    p("Index:\t\t", s.Index("test","s"))
    p("Join:\t\t", s.Join([]string{"a","b"},"-" ))
    p("Repeat:\t\t", s.Repeat("a",5))
    p("Replace:\t", s.Replace("foo","o","0",-1))
    p("Replace:\t", s.Replace("foo","o","0",1))
    p("Split:\t\t", s.Split("a-b-c-d","-"))
    p("ToLower:\t", s.ToLower("TEST"))
    p("ToUpper:\t", s.ToUpper("test"))
    p()
    p("Len:\t\t",len("hellooo"))
    p("Char:\t\t","hellooo"[3])


}