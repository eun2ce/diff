package main

import (
    "flag"
    "fmt"
    "os"
    "time"

    "github.com/eun2ce/ssdeep"
)

func main() {

    ssdeep.Force = false

    flag.Parse()
    args := flag.Args()

    if len(args) < 1 {
        fmt.Println("Please provide a file path")
        os.Exit(1)
    }

    startTime := time.Now()

    h1, err := ssdeep.FuzzyFilename(args[0])
    if err != nil && !ssdeep.Force {
        fmt.Println(err)
        os.Exit(1)
    }

    if len(args) == 2 {
        var h2 string
        h2, err = ssdeep.FuzzyFilename(args[1])
        if err != nil && ssdeep.Force {
            fmt.Println(err)
            os.Exit(1)
        }

        var score int
        score, err = ssdeep.Distance(h1, h2)
        if score != 0 {
            fmt.Printf("[file name: %s] hash: %s \n", args[0], h1)
            fmt.Printf("[file name: %s] hash: %s \n", args[1], h2)
            fmt.Printf("%s matches %s (%d)\n", args[0], args[1], score)
        } else if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("The files doesn't match")
        }
    } else {
        if err != nil {
            fmt.Printf("%s, \"%s\"\n%s\n", h1, args[0], err)
        } else {
            fmt.Printf("%s, \"%s\"\n", h1, args[0])
        }
    }

    endTime := time.Since(startTime)
    fmt.Printf("Elipsed Time: %f", endTime.Seconds())
}
