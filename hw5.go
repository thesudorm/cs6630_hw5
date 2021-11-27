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
    f1 := GenerateF1(t)

    //for _, test := range f1 {
    //fmt.Print(test, " ")
    //}
    //fmt.Println()

    for _, item := range f1 {
        for _, trans := range t {
            _, found := Find(trans, item)
            if found {
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

// Find unique items
func GenerateF1(t [][]string) []string {
    f1 := []string{}
    for i := 0; i < len(t); i++ {
        record := t[i]
        for j := 0; j < len(record); j++ {
            item := record[j]
            if len(item) > 0 {
                _, found := Find(f1, item)
                if !found {
                    f1 = append(f1, item)
                }
            }
        }
    }

    return f1
}

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}
