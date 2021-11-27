package main

import (
    "fmt"
    "encoding/csv"
    "io"
    "log"
    "os"
)

func main() {
    csv_file, _ := os.Open("retail_dataset.csv")
    r := csv.NewReader(csv_file)

    for {
        record, err := r.Read()

        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(record)
    }
}
