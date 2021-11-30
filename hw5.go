package main

import (
    "fmt"
    "strings"
    "sort"
    "encoding/csv"
    "io"
    "log"
    "os"
)
/*
/*  TODO:*/
/*      1. Prune candidates in CandidateGen
/*      2. Fix CandidateGen to work with k > 2*/
/*      3. Glue everything together
*/

func main() {
    transactions := readTransactions()
    f1 := GenerateF1(transactions, 0.50)
    fk := f1
    dict := map[string]int{}

    ck := CandidateGen(fk);

    for _, c := range ck {
        for _, t := range transactions {
            itemsetInTransaction := true
            itemset := strings.Split(c, " ")

            for _, item := range itemset {
                _, found := Find(t, item)
                if found == false {
                    itemsetInTransaction = false
                }
            }

            if itemsetInTransaction {
                if _, ok := dict[c]; ok {
                    dict[c] += 1
                } else {
                    dict[c] = 1
                }
            }
        }
    }

    //fmt.Println(fk)

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
func GenerateF1(t [][]string, supp float64) []string {
    f       := []string{}
    f1      := []string{}
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
            f1 = append(f1, item)
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

// TODO make this generate candidates for the next generation
func CandidateGen(fk []string) []string {
    ck := []string{}

    // Generate the next gen of candidates
    for i := 0; i < len(fk); i++ {
        itemset := strings.Split(fk[i], " ")
        sort.Strings(itemset)
        for j := 0; j < len(fk); j++ {
            itemsetToAdd := strings.Split(fk[j], " ")
            sort.Strings(itemsetToAdd)
            for _, item := range itemsetToAdd {
                _, found := Find(itemset, item)
                if found == false {
                    toAdd := fk[i] + " " + item
                    _, found := Find(ck, toAdd)
                    if found == false {
                        ck = append(ck, fk[i] + " " + item)
                        break
                    }
                }
            }
        }
    }

    // Prune candidates
    // TODO

    return ck
}
