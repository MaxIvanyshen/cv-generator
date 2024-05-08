package main

import (
	"cv/args"
	"cv/cv"
	"fmt"
	"strings"
)

func main() {
    args := args.ParseCreateArgs()

    var resume cv.CV 
    var err error

    fileType := strings.Split(args.DataFile, ".")[1]
    if fileType == "json" {
        resume, err = cv.ParseFromJson(args.DataFile)
    } else if fileType == "yml" {
        resume, err = cv.ParseFromYaml(args.DataFile)
    } else {
        err = fmt.Errorf("unsupported file type: '.%s'. Use json or yaml", fileType)
    }

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(resume)
}
