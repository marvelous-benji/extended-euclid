
package main

import "fmt"




func ModularInverse(a, n int) int{

	if n <= 0{
		panic("n cannot be negative or zero")
	}
	a = a % n
	if a < 0{
		a = n + a
	}

	x := 0; next_x := 1
	r := n; next_r := a

	for next_r != 0{
		quotient := r / next_r
		x, next_x = next_x, x - quotient * next_x
		r, next_r = next_r, r - quotient * next_r
	}
	if r > 1{
		fmt.Printf("%v cannot have a multiplicative inverse modulo %v\n",a,n)
		panic("error")
	}
	if x < 0{
		x += n
	}

	return x

}



