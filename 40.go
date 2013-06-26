exo 40

package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    words:=strings.Fields(s)
    var result = make(map[string]int)
    for _, w:=range words {
        result[w]+=1
    }
    return result
}

func main() {
    wc.Test(WordCount)
}

