package main

import (
	"fmt"
	"os"
)

func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    input := data
    
    answer := 0
    for i, c := range input {
        if c == '(' {
            answer++
        } else if c == ')'{
            answer--
        }

        if answer == -1 {
            fmt.Println(i)
            break
        }
    }

    fmt.Println(answer)
}
