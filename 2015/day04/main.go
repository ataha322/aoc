package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)


func main() {
    var input string = "ckczppom"
    result := bruteforce(input)
    fmt.Println(result)
}

func bruteforce(input string) string {
    for i := 1;; i++ {
        i_str := strconv.Itoa(i)
        h := md5.New()
        io.WriteString(h, input+i_str)
        md5hash := fmt.Sprintf("%x", h.Sum(nil))
        if fiveZeros(md5hash) {
            return i_str
        }
    }
}

func fiveZeros(str string) bool {
    for i := 0; i < 6; i++ {
        if (str[i] != '0') {
            return false
        }
    }
    return true
}
