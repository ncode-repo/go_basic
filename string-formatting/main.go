package main

import "fmt"
import "os"

type point struct {
    x,y int
}

func main() {
    p := point{x:2,y:3}
    
    fmt.Printf("%v\n", p)
    fmt.Printf("%+v\n", p)
    fmt.Printf("%#v\n", p)
    
    fmt.Printf("%T\n", p)
    fmt.Printf("%t\n", true)

    fmt.Printf("%d\n", 123)
    fmt.Printf("%b\n", 12)
    fmt.Printf("%c\n", 33)
    fmt.Printf("%x\n", 456)

    fmt.Printf("%f\n", 23.54)
    fmt.Printf("%e\n", 123000000.00)
    fmt.Printf("%E\n", 123000000.00)
    // For basic string printing use `%s`.
    fmt.Printf("%s\n", "\"string\"")

    // To double-quote strings as in Go source, use `%q`.
    fmt.Printf("%q\n", "\"string\"")

    // As with integers seen earlier, `%x` renders
    // the string in base-16, with two output characters
    // per byte of input.
    fmt.Printf("%x\n", "hex this")

    // To print a representation of a pointer, use `%p`.
    fmt.Printf("%p\n", &p)

    // When formatting numbers you will often want to
    // control the width and precision of the resulting
    // figure. To specify the width of an integer, use a
    // number after the `%` in the verb. By default the
    // result will be right-justified and padded with
    // spaces.
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    // You can also specify the width of printed floats,
    // though usually you'll also want to restrict the
    // decimal precision at the same time with the
    // width.precision syntax.
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    // To left-justify, use the `-` flag.
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // You may also want to control width when formatting
    // strings, especially to ensure that they align in
    // table-like output. For basic right-justified width.
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    // To left-justify use the `-` flag as with numbers.
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    // So far we've seen `Printf`, which prints the
    // formatted string to `os.Stdout`. `Sprintf` formats
    // and returns a string without printing it anywhere.
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    // You can format+print to `io.Writers` other than
    // `os.Stdout` using `Fprintf`.
    fmt.Fprintf(os.Stderr, "an %s\n", "error")


}