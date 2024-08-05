package main

import (
    "fmt"
    "math/rand"
)


func buildBuilb(){

}

func main(){
    fmt.Println("a")
    
    bomba := createBulb(0.1, 9, 0, 0, 0)

    for i := 0; i < 10; i++{
        bomba = bomba.calcVector()
        bomba.quat.X += x
        bomba.quat.Y += y
        bomba.quat.Z += z
    }
        
}
