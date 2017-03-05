// Challenge: 100 factorial computations concurrently and in parallel, use the "pipeline" pattern
package main

import (
  "fmt"
)

func main() {
  c := generator()
  
  results := launch(c)

  for n := range results {
    fmt.Println(n)
  }
}

func generator() chan int {
  output := make(chan int)
  go func() {
    for i := 0; i < 10; i++ {
      output<-i+5
    }
    close(output)
  }()
  return output
}

func launch(c chan int) chan int {
  output := make(chan int)
  go func() {
    for n := range c {
      output<-factorial(n)
    }
    close(output)
  }()
  return output
}

func factorial(n int) int {
  total := 1
  for i := n; i > 0; i-- {
    total *= i
  }
  return total 
}
