package main

import (
	"fmt"
	"os"
	"strconv"
)

type coord struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	input := make([]byte, 16384)
	count, err := file.Read(input)
	if err != nil {
		panic(err)
	}

    fmt.Println("houses recieved at least 1 present: " + strconv.Itoa(countHouses(input[:count])))
}

func countHouses(input []byte) int {
    visited := make(map[coord]bool)
    santa := coord{0,0}
    robo := santa
    visited[santa] = true

    for i, c := range input {
        var curr *coord
        if i%2 == 0 {
            curr = &robo 
        } else {
            curr = &santa
        }

        switch c {
        case '>':
            curr.x++
        case '<':
            curr.x--
        case '^':
            curr.y++
        case 'v':
            curr.y--
        }
        
        visited[*curr] = true
    }

    return len(visited)
}
