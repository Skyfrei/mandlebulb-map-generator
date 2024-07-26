package main

type Bicomplx struct{
	w Complx
	z Complx
}


func (num *Bicomplx) conjugate() *Bicomplx{
	res := Bicomplx{ w: {pair: [2]float64 {num.w.pair[0], num.w.pair[1]} },
								z: {pair: [2]float64 {-1 * num.z.pair[0], -1 * num.z.pair[1]} }}

	return &res
}

func prod(num1 *Bicomplx, num2 *Bicomplx) *Bicomplx{

}

// wtf is the output here
func (num *Bicomplx) norm(){

}
