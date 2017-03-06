package main

import (
  "fmt"
)

var count int64

func main() {
  c := incrementor(2)
  var count int

  for s := range c {
    count++  
    fmt.Println(s)
  }
  fmt.Println("Final Counter:", count)
}

func incrementor(n int) chan string {
  c := make(chan string)
  done := make(chan bool)

  for i := 1; i <= n; i++ {
    go func(j int) {
      for k := 0; k < 20; k++ {
        c <-fmt.Sprintf("Process: %d printing: %d", j, k)    
      }
      done <-true
    }(i)
  }

  go func() {
    for i := 0; i < n; i++ {
      <-done
    }
    close(c)
  }()

  return c
}
