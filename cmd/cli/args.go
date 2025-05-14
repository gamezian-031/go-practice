package args

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) == 2 {
        fmt.Printf("Hello to %s!\n", os.Args[1])
    }   else {
        fmt.Println("Hello World")
    }

}
