package main

import (
    "bufio"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "time"

    "github.com/schollz/documentsimilarity"
)

func GetFileToStr(p string) string {
    f, err := ioutil.ReadFile(p)
    if err != nil {
        fmt.Println(err)
    }
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
    args := flag.Args()

    if len(args) < 1 {
        fmt.Println("Please provide a file path")
        os.Exit(1)
    }

    startTime := time.Now()

    doc1, _ := readLines(args[0])
    doc1 = append(doc1, GetFileToStr(args[1]))
    documents := doc1

    ds, error := documentsimilarity.New(documents)
    fmt.Println("ds: ", ds.Bags)
    if error != nil {
        fmt.Println(error)
    }
    if len(args) == 2 {
        diffDocuments := GetFileToStr(args[1])
        var score float64
        cSimilarity, error := ds.CosineSimilarity(diffDocuments)
        if error != nil {
            fmt.Println("cSimilarity: ", cSimilarity)
        }
        for _, cs := range cSimilarity {
            score += cs.Similarity
        }

        if score != 0 {
            fmt.Println("== schollz documentsimilarity result ==")
            fmt.Printf("%s matches %s (%f)\n", args[0], args[1], score)
        } else if error != nil {
            fmt.Println(error)
        } else {
            fmt.Println("The files doesn't match")
        }
    } else {
           fmt.Printf("please put two docs")
    }

    endTime := time.Since(startTime)
    fmt.Printf("Elipsed Time: %f", endTime.Seconds())

}
