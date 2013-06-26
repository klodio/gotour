package main

import "code.google.com/p/go-tour/pic"

//This feels rather primitive to me,
//I hope I'm going to discover maps as I go through this turorial!
func Pic(dx, dy int) [][]uint8 {
    
    picture:=make([][]uint8,dy)
    
    for i:=range picture {
        row:=make([]uint8,dx)
        for j:=range row {
            row[j]=uint8(i)^uint8(j)
        }
        picture[i]=row;
    }
    return picture
}

func main() {
    pic.Show(Pic)
}
