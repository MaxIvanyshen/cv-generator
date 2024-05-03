package main

import (
	"cv/args"
	"cv/cv"
	"fmt"
)

func main() {
    args := args.ParseCreateArgs()

    cv, err := cv.ParseFromYaml(args.DataFile)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(cv)
}
