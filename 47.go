exo 47

package main

import "fmt"

//Using a closure for fun
func newtonCube() func(y complex128) complex128{
    z:=complex128(1);
    return func(x complex128) complex128{
        //Originally z-= (cmplx.Pow(z,3)-x) / (complex128(3)*cmplx.Pow(z,2))
        z-= (z*z*z-x)/(3*z*z)
        
        return z
    }   
}

func Cbrt(x complex128) complex128 {
    var result complex128;
    var method=newtonCube();
    for i:=0;i<10;i++ {
        result=method(x);
    }   
    return result;
}

func main() {
    fmt.Println(Cbrt(2))
}
