package main

import (
    "convert/convert"
    "flag"
    "fmt"
)

// initialize flag
var (
    from = flag.String("from", "jpg", "from extension")
    to = flag.String("to", "png", "after extension")
    directory = flag.String("directory", "./images", "directory path")
)

func init() {
    flag.Usage = func() {
        fmt.Printf(`Usage: -from FROM_FORMAT -to TO_FORMAT -dir DIRECTORY
        Use: convert image files.
        Default: from jpg to png.
        `)
        flag.PrintDefaults()
    }
}

func main() {
    flag.Parse()

    conv, err := convert.NewConv(*from, *to, *directory)
    if err != nil {
            fmt.Println(err)
    }

    paths := conv.FileSearch(*directory)
    for _, path := range paths {
        conv.ReplaceExt(path, *from, *to)
    }
    fmt.Printf("finish convert\n")
}