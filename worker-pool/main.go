package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, result chan<- int){
    for j:= range jobs {
        fmt.Println("worker ",id," started job",j)
        time.Sleep(time.Second *2)
        fmt.Println("worker ", id, " finished job",j)
        result <- j*2
    }
}

func main(){
    jobs := make(chan int,100)
    result := make(chan int,100)

    for w:=1; w <= 3; w++ {
        go worker(w, jobs, result)
    }
    for j:=1; j<=5; j++ {
        jobs <- j
    }
    close(jobs)

    for k:=1; k<=5; k++ {
        <-result
    }
}