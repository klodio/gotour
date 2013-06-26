package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    current,before:=0,0;
    return func() int {
        if before == 0 && current == 0 { 
        	current=1
            return 0
        }
        current,before = current+before,current
      return before;
    }   
}

func main() {
    f := fibonacci()
    for i := 0; i < 11; i++ {
        fmt.Println(f())
    }   
}
