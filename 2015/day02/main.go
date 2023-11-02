package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type box struct {
    l, w, h int
}

func main() {
    ret := 0
    ribbon := 0
    input := parse()

    for _, b := range input {
        ret += oneBox(&b)
        ribbon += computeRibbon(&b)
    }

    fmt.Println("total area:" + strconv.Itoa(ret))
    fmt.Println("ribbon: " + strconv.Itoa(ribbon))
}

func oneBox(b *box) int {
	side1 := b.l * b.w
	side2 := b.l * b.h
	side3 := b.w * b.h

	area := 2*side1 + 2*side2 + 2*side3

	return min(side1, side2, side3) + area
}

func computeRibbon(b *box) int {
    min_perimeter := 2 * (b.l + b.w + b.h - max(b.l, b.w, b.h))
    bow := b.l * b.w * b.h
    return min_perimeter + bow
}

func parse() []box {
    f, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    
    boxes := []box{}

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "x")
        if len (parts) != 3 {
            fmt.Println("WRONG INPUT")
            panic(1)
        }

        l, _ := strconv.Atoi(parts[0])
        w, _ := strconv.Atoi(parts[1])
        h, _ := strconv.Atoi(parts[2])

        boxes = append(boxes, box{l, w, h})
    }

    return boxes
}
