package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    DoWalk(t,ch)
    close(ch)
}
//Always visit left before putting in the channel
//To ensure we have the tree sorted
func DoWalk(t *tree.Tree, ch chan int) {
    if t == nil {
    	return
    }
    DoWalk(t.Left, ch)
    ch <- t.Value
	DoWalk(t.Right, ch)    	
}

func CompareChannels(ch1,ch2 chan int) bool{
    for {
        v1,done1:= <- ch1
        v2,done2:= <- ch2
        
        if v1 != v2 {
            return false
        }
        if done1==true || done2==true {
            return done1 == done2
        }
        
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1:=make(chan int);
    ch2:=make(chan int);
    go Walk(t1, ch1);	
    go Walk(t2, ch2);
    return CompareChannels(ch1, ch2);
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}


