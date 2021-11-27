package main

import (
    "fmt"
    "encoding/csv"
    "io"
    "log"
    "os"
)

func main() {
    t := readTransactions()
    //minSup := 0.20
    dict := map[string]int {}

    for i := 0; i < len(t); i++ {
        record := t[i]
        for j := 0; j < len(record); j++ {
            item := record[j]
            if len(item) > 0 {
                if _, ok := dict[item]; ok {
                    dict[item] += 1
                } else {
                    dict[item] = 0
                }
            }
        }
    }

    for k, v := range dict {
        fmt.Println(k, v)
    }
}

func readTransactions() [][]string {
    csv_file, _ := os.Open("retail_dataset.csv")
    r := csv.NewReader(csv_file)
    t := [][]string{}

    for {
        record, err := r.Read()

        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        t = append(t, record)
    }

    return t
}
