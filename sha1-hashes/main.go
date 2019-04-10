package main

import "fmt"
import "crypto/sha1"
import "crypto/md5"

func main() {
	s1 := "sh1 this string"
	s2 := "md5 this string"

	h1 := sha1.New()
	h1.Write([]byte(s1))
	bs1 := h1.Sum(nil)

	h2 := md5.New()
	h2.Write([]byte(s2))
	bs2 := h2.Sum(nil)

	fmt.Println("s1= ", s1)
	fmt.Println("sha1 = ", bs1)
	fmt.Printf("sha1 hex= %x\n", bs1)

	fmt.Println("s2= ", s2)
	fmt.Println("md5 = ", bs2)
	fmt.Printf("md5 hex= %x\n", bs2)
}
