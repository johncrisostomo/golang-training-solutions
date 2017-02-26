package main

import (
  "fmt"
  "math/rand"
)

func main() {
  n := []int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
  fmt.Println("The numbers are ", n)
  greatest := findGreatest(n...)
  fmt.Println("The greatest is ", greatest)
}

func findGreatest(n ...int) int {
  greatest := n[0]

  for _, current := range n {
    if current > greatest {
      greatest = current 
    }
  }

  return greatest
}
