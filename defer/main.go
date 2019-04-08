package main

import "fmt"
import "os"

func main (){
    f := createFile("defer/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(fname string) *os.File {
    fmt.Println("creating file...")
    f, err := os.Create(fname)
    if err != nil {
        panic (err)
    }
    return f
}
func writeFile(file *os.File) {
    fmt.Println("writing file...")
    fmt.Fprintln(file, "data")
}
func closeFile(file *os.File) {
    fmt.Println("closing file.")
    file.Close()
}