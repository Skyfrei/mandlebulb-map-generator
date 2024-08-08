package main

import(
	"math"
    "mathlib"
//    "math/rand"
 //   "fmt"
)

type mandelbulb struct{
	quat mathlib.Quat
    c mathlib.Quat

}

func createBulb(comp, q *mathlib.Quat, ) *mandelbulb{
    bulb := mandelbulb{	quat: *q, c: *comp}
	return &bulb
}


func (man mandelbulb) phi() float64{
	temp := man.quat
	return math.Atan2(temp.Y, temp.X) 
}

func (man mandelbulb) magnitude() float64{
	temp := man.quat
    return math.Sqrt(math.Pow(temp.X, 2) + math.Pow(temp.Y, 2) + math.Pow(temp.Z, 2))
}

func (man mandelbulb) theta() float64{
	temp := man.quat
    return math.Atan2(math.Sqrt(math.Pow(temp.X, 2) + math.Pow(temp.Y, 2)), temp.Z)
}

func (man *mandelbulb) calcVector() mathlib.Quat{
    vec := &man.quat
    theta := man.theta()
    phi := man.phi()

    rn := math.Pow(man.magnitude(), vec.R)
    vec.X = rn * (math.Sin(vec.R * theta) * math.Cos(vec.R * phi))
    vec.Y = rn * (math.Sin(vec.R * theta) * math.Sin(vec.R * phi))
    vec.Z = rn * (math.Cos(vec.R * theta)) 
  
    return *vec
}


