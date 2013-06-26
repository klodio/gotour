package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}
//need to implement Read() for the io.Reader
func (reader rot13Reader) Read(p []byte) (n int, err error){
	n,err=reader.r.Read(p);
    
    for i,v:=range p{
        p[i]=rotConv(v)
    }
    return n,err
}

func rotConv(p byte) byte{
    switch {
    case p < 65:
        return p
    case p >= 65 && p < 78:
        return p+13
    case p < 90:
        return p-13
    case p >= 97 && p < 110:
        return p+13
    case p < 122:
        return p-13
    default:
        return p
    }
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
