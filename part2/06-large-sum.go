package main

import (
  "fmt"
  "os"
  "bufio"
  "math"
  "strconv"
)

/*
  Solution to Project Euler #13

  Problem Description: 

  Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.  

  Problem URL : https://projecteuler.net/problem=13
*/

func main() {
  file, _ := os.Open("largenumber.txt") 
  defer file.Close()

  reader := bufio.NewReader(file)
  scanner := bufio.NewScanner(reader)

  scanner.Split(bufio.ScanLines)
  numsInString := []string{}

  for scanner.Scan() {
    numsInString = append(numsInString, scanner.Text())
  }
  rawResult := sum(numsInString)
  fmt.Println("The first 10 digits are ", rawResult[:10])
}

func sum(nums []string) string {
  sum := int64(0)

  for _, current := range nums {
    currentLen := len(current)
    offset := int(math.Ceil(math.Log10(float64(currentLen))))
    length := 10 + offset
    parsedCurrent, _ := strconv.ParseInt(current[:length], 10, 64)
    sum += parsedCurrent 
  }

  return strconv.FormatInt(sum, 10)
}
