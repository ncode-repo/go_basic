package main

import(
    "fmt"
    "errors"
)

func f1(arg int)(int,error){
    if arg==42{
        return -1, errors.New("Error can't work with 42")
    }
    return arg+3,nil
}

type argError struct {
    arg int
    prob string
}
func (e *argError) Error() string{
    return fmt.Sprintf("%d - %s", e.arg,e.prob)
}

func f2(arg int) (int, error){
    if arg == 42{
        return -1, &argError{arg, "Can't work with it"}
    }
    return arg,nil
}

func main(){
    for _,i :=range[] int{3,42}{
        if a,e := f1(i); e!=nil{
            fmt.Println("f1() failed, ",e)
        }else{
            fmt.Println("f1() success, ",a)
        }
    }

    for _, i := range[]int{12,42} {
        if a, e := f2(i); e != nil {
            fmt.Println("f2() failed, ", e)
        }else{
            fmt.Println("f2() success, ", a)
        }
    }
    
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}