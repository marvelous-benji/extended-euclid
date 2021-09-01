
package main

import "fmt"

/*
Insight comes from the Bezout lemma:
For nonzero integers a and b, let d be the greatest common divisor d=gcd⁡(a,b)
Then, there exist integers x and y such that
		ax+by=d
a special case is when a and b are coprime (that is relatively prime)
then d = gcd(a,b) = 1 implies ax + by = 1
given a natural number n, if ax = 1 mod n, then we say x is the multiplicative
inverse of 'a' modulo n or the modular multiplicative of 'a' in n.
A necessary and sufficient condition for 'a' to have a multiplicative inverse
in n is that gcd(a,n) = 1. Thus there exist integers x and y such that
ax + yn = 1 ==> ax - 1 = (-y)n.
Using the extended Euclidean algorithm we can get the bezout coefficients
and this is implemented below
*/





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



