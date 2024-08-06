package main

import(
	"math"
    "mathlib"
)

type mandelbulb struct{
	quat mathlib.Quat
    c mathlib.Quat

}

func createBulb(comp *mathlib.Quat, n1 float64) *mandelbulb{
    q := mathlib.Quat{X: 0, Y: 0, Z: 0, R: n1}
    
    bulb := mandelbulb{	quat: q, c: *comp}
	return &bulb
}

func (man *mandelbulb) phi() float64{
	temp := man.c
	return math.Atan(temp.Y / temp.X) 
}

func (man *mandelbulb) magnitude() float64{
	temp := man.c
	return math.Sqrt(math.Pow(temp.X, 2) + math.Pow(temp.Y, 2) + math.Pow(temp.Z, 2)) 
}

func (man *mandelbulb) theta() float64{
	temp := man.c
	return math.Atan(math.Sqrt(math.Pow(temp.X, 2) + math.Pow(temp.Y, 2)) / temp.Z)
}

func (man *mandelbulb) calcVector() mathlib.Quat{
    vec := man.c
    rn := math.Pow(man.magnitude(), vec.R)

    vec.X = rn * math.Sin(vec.R * man.theta()) * math.Cos(vec.R * man.phi()) + vec.X
    vec.Y = rn * math.Sin(vec.R * man.theta()) * math.Sin(vec.R * man.phi()) + vec.Y 
    vec.Z = rn * math.Cos(vec.R * man.theta()) + vec.Z

    tempBulb := mandelbulb{quat: man.quat, c: vec}
    if (tempBulb.magnitude() <= 1){
        return vec
    }else{
        vec.X += 0.001
        vec.Y += 0.01
        vec.Z -= 0.02
        return vec
    }
}


