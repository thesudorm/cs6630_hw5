package main

import (
    "fmt"
    "strings"
    "sort"
    "encoding/csv"
    "io"
    "log"
    "strconv"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please include only 1 argument")
    } else {
        if n, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
            minSupp := n
            transactions := readTransactions()
            f1, dict := GenerateF1(transactions, minSupp)
            fk := f1
            fkPrev := f1

            for k := 0; len(fkPrev) > 0; k++ {
                ck := CandidateGen(fkPrev);
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

                // Once you are done looking at candidates, keep the frequent ones in fk
                fk = fkPrev
                fkPrev = []string {}
                for _, c := range ck {
                    supp := float64(dict[c]) / float64(len(transactions))
                    _, found := Find(fk, c)
                    if supp >= minSupp && found == false{
                        fk = append(fk, c)
                        fkPrev = append(fkPrev, c)
                    }
                }
            }

            keys := []string{}
            for _, k := range fk {
                keys = append(keys, k)
            }

            sort.Strings(keys)

            // Write all freq itemsets to console
            for _, k := range keys {
                fmt.Println(k, fmt.Sprintf("%.2f", ((float64(dict[k]) / float64(len(transactions))))))
            }
        } else {
            fmt.Println("Please include a support value.")
        }
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
func GenerateF1(t [][]string, supp float64) ([]string, map[string]int) {
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

    return f1, dict
}

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func CandidateGen(fk []string) []string {
    ck := []string{}

    // Generate the next gen of candidates
    // Maybe I should use arrays of arrays instead? This is kinda dumb.. :(
    for i := 0; i < len(fk); i++ {
        itemset := strings.Split(fk[i], " ")
        sort.Strings(itemset)
        for j := 0; j < len(fk); j++ {
            itemsetToAdd := strings.Split(fk[j], " ")
            sort.Strings(itemsetToAdd)
            for _, item := range itemsetToAdd {
                _, found := Find(itemset, item)
                if found == false { // prevents dupe items in the same itemset
                    if found == false {
                        toAdd := fk[i] + " " + item
                        sortedToAdd := strings.Split(toAdd, " ")
                        sort.Strings(sortedToAdd)
                        toAdd = ""
                        for _, str := range sortedToAdd {
                            toAdd += str + " "
                        }
                        toAdd = toAdd[:len(toAdd)-1]
                        _, found := Find(ck, toAdd)
                        if found == false {
                            ck = append(ck, toAdd)
                        }
                        break
                    }
                }
            }
        }
    }

    return ck
}
