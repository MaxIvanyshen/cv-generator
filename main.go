package main

import (
    "fmt"
    "cv/args"
)

func main() {
    args := args.ParseCreateArgs()
    fmt.Printf("%+v\n", args)
}
