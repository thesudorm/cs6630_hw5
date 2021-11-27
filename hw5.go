package main

import (
    "fmt"
    "sort"
    "encoding/csv"
    "io"
    "log"
    "os"
)

func main() {
    transactions := readTransactions()
    f1 := GenerateF1(transactions, 0.35)
    fk := f1
    count := 0
    numOfItemsInTrans := len(transactions[0])

    fmt.Println(numOfItemsInTrans)

    for k := 2; len(fk) > 0; k++ {
        ck := CandidateGen(fk, k) // this needs to make the next generation of candidates
        count = 0
        for _, t := range transactions {
            candidatesInTrans := true
            for _, c := range ck {
                for _, ci := range c {
                    _, found := Find(t, ci)
                    if found == false {
                        candidatesInTrans = false
                    }
                }

                if candidatesInTrans {
                    //c.count += 1
                }
            }

        }
    }

    for _, item := range fk {
        fmt.Println(item)
    }
    fmt.Println(count)
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
func GenerateF1(t [][]string, supp float64) [][]string {
    f       := []string{}
    f1      := [][]string{}
    dict    := map[string]int {}

    for i := 0; i < len(t); i++ {
        record := t[i]
        for j := 0; j < len(record); j++ {
            item := record[j]
            if len(item) > 0 {
                _, found := Find(f, item)
                if !found {
                    f = append(f, item)
                }
            }
        }
    }

    sort.Strings(f)

    for _, item := range f {
        for _, trans := range t {
            _, found := Find(trans, item)
            if found {
                if _, ok := dict[item]; ok {
                    dict[item] += 1
                } else {
                    dict[item] = 1
                }
            }
        }
    }

    for _, item := range f {
        if float64(dict[item]) / float64(len(t)) >= supp {
            f1 = append(f1, []string {item})
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

func CandidateGen(fk [][]string, k int) [][]string {
    fmt.Println("Generating..")
    ck := [][]string{}

    for _, q := range fk {
        for i := 0; i < len(q); i++ {
        }
    }

    fmt.Println(ck)
    return fk
}
