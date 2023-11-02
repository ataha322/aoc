package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    nice := 0
    for scanner.Scan() {
        line := scanner.Text()
        if isNice(line) {
            nice++
        }
    }

    fmt.Println(nice)
}

func isNice(line string) bool {
    if len(line) < 4 {
        return false
    }
    if !cond1(line) {
        return false
    }
    if !cond2(line) {
        return false
    }
    return true
}

func cond1(line string) bool {
    m := map[string]int {} //two letter string -> beginning index

    for i := 0; i < len(line)-1; i++ {
        str := line[i:i+2]
        idx, ok := m[str]
        if !ok {
            m[str] = i
        } else if i - idx >= 2 {
            return true
        }
    }

    return false
}

func cond2(line string) bool {
    for i := 0; i < len(line)-2; i++ {
        if line[i] == line[i+2] {
            return true
        }
    }

    return false
}
