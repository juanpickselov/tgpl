// Echo1 prints its command-line arguments arguments arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    var aString, separator string // that's right I don't like the short s variable name, what of it?
    for i:= 1; i < len(os.Args); i++ {
        aString += separator + os.Args[i]
        separator = " "
    }
    fmt.Println(aString)
}
