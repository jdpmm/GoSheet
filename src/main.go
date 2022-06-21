package main

import (
    "fmt"
    "os"
    "bufio"
)

func main () {
    if len(os.Args) != 2 {
        fmt.Println("No argument given.")
        os.Exit(1)
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("(%s file): Does not exist.", os.Args[1])
        os.Exit(1)
    }

    scner := bufio.NewScanner(file)
    var content string
    var rowvalue int = 0
    for scner.Scan () {
        content = scner.Text()
        Tbl_getrow(content, rowvalue)
        rowvalue++
    }
    Tbl_maketable()
}
