package main

import "fmt"

func main() {
  a, b := half(1)
  fmt.Printf("half(1) returns (%d, %t)\n", a, b)
  a, b = half(2)
  fmt.Printf("half(2) returns (%d, %t)\n", a, b)
}

func half(n int) (int, bool) {
  quotient := n / 2
  isEven := n % 2 == 0

  return quotient, isEven
}
