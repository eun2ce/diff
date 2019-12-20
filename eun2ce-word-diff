package main

import (
    "bufio"
    "flag"
    "fmt"
    "io/ioutil"
    "os"

    "time"

    "github.com/eun2ce/textsimilarity"
)

func GetFileToStr(p string) string {
    f, err := ioutil.ReadFile(p)
    if err != nil { fmt.Println(err) }

    return string(f)
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func main() {
    flag.Parse()
    args := flag.Args() //파일 패스를 받는다.

    if len(args) < 1 {
        fmt.Println("Please provide a file path")
        os.Exit(1)
    }

    startTime := time.Now()

    tsdoc, error := readLines(args[0])
    ts := textsimilarity.New(tsdoc)

    fmt.Println("ts : ", ts)
    if error != nil {
        fmt.Println(error)
        os.Exit(1)
    }

    if len(args) == 2 { 
        docA := GetFileToStr(args[0])
        docB := GetFileToStr(args[1])

        score, err := ts.Similarity(docA, docB)

        if score != 0 {
            fmt.Println("== eun2ce textsimilarity result ==")
            fmt.Printf("%s matches %s (%f)\n", args[0], args[1], score)
        } else if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("The files doesn't match")
        }
    } else {
           fmt.Printf("please put two docs")
    }

    endTime := time.Since(startTime)
    fmt.Printf("Elipsed Time: %f", endTime.Seconds())
}
