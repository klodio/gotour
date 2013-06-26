exo 43

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    current,before2,before:=0,0,0;
    return func() int {
        if before2 == 0 && before == 0 && current == 0 {
     		current=1
            return 0
        }
        before2=before;
        before=current;
        current=before+before2;
    	return before;
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}


