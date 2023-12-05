package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"golang.org/x/exp/constraints"
)


func main() {
  file, err := os.Open("./day1/input")

  if err != nil {
    fmt.Println("Input file not found.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  sum1 := 0
  sum2 := 0

  for scanner.Scan() {
    txt := scanner.Text() 
    sum1 += getCalibration(txt)
    sum2 += getCalibration2(txt)
  }

  fmt.Println(sum1)
  fmt.Println(sum2)
}

func getCalibration(s string) int {
  var firstDigit rune = 0
  var lastDigit rune
  for _, c := range s {
    if (unicode.IsDigit(c)) {
      if firstDigit == 0 {
        firstDigit = c
      }
      lastDigit = c
    }
  }

  return (int(firstDigit) - '0') * 10 + int(lastDigit) - '0'
}

func min[T constraints.Ordered](l, r T) T {
  if l < r {
    return l
  }
  return r
}

func getCalibration2(s string) int {
  var firstDigit int = -1
  var lastDigit int
  n := len(s)
  for idx, c := range s {
    if (unicode.IsDigit(c)) {
      if firstDigit == -1 {
        firstDigit = int(c) - '0'
      }
      lastDigit = int(c) - '0'
    } else {
      for l := 3; l <= 5; l++ {
        if res, v := isDigitString(s[idx: min(idx + l, n)]); res {
          if firstDigit == -1 {
            firstDigit = v
          }
          lastDigit = v
        }
      }
    }
  }

  return firstDigit * 10 + lastDigit
}


func isDigitString (s string) (bool, int) {
  numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

  for i := 0; i < 10; i++ {
    if (numbers[i] == s) {
      return true, i
    }
  }
  return false, -1
}
