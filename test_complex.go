package main

import "fmt"

func main() {
	x, y, z := complex(1, 2), complex(3, 4), complex(5, 6)
	fmt.Println(x, y)             // (1+2i) (3+4i)
	fmt.Println(real(x), imag(x)) // 1 2
	fmt.Println(x * y)            // (-5+10i)
	fmt.Println(1i)               // (0+1i)
	fmt.Println(1i * 1i)          // (-1+0i)
	fmt.Println(x * z)
	fmt.Println(z * y)
	fmt.Println(x * y * z)
}
