package main
import "fmt"

func sum(nums ...int){
    fmt.Println("nums: ",nums)
    sum:=0
    for _,v:=range nums{
        sum+=v
    }
    fmt.Println("sum=",sum)
}
func main(){
    sum(1,2)
    sum(3,5,7,8)
    sls:=[]int{11,22,33,44}
    sum(sls...)
}