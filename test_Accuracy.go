package main
import (
"fmt"
)
func main() {
var a float32 = 359.9
fmt.Println(a) // 359.9
var b float64 = float64(a)
fmt.Println(b) // 359.8999938964844

var a1 float32 = 358.99999
var b1 float32 = 359.00001
fmt.Println(b1 - a1) // 0
}