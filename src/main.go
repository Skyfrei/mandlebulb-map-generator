package main

import (
    "fmt"
    "mathlib"
    // render engind
    "os"

)



func main(){
    fmt.Println("a")
    
    complex := mathlib.Quat{X: 0.2, Y: 0.3, Z: 0.4, R: 9} 
    bomba := createBulb(&complex, 0.1)
   
    f, _ := os.Create("dest.txt")
     
    for i := 0; i < 1000; i++{
        
        bomba.c = bomba.calcVector()
        f.WriteString(fmt.Sprintf("%f ", bomba.c.X))
        f.WriteString(fmt.Sprintf("%f ", bomba.c.Y))
        f.WriteString(fmt.Sprintf("%f ", bomba.c.Z))
        f.WriteString("\n")

    }

}
