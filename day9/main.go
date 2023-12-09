package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
  name string
  left, right *Node
}

func main() {
  file, err := os.Open("./day9/input")

  if err != nil {
    fmt.Println("Input file not found.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)


  input := []string{}

  for scanner.Scan() {
    if len(scanner.Text()) == 0 {
      continue
    }
    input = append(input, scanner.Text())
  }

  if len(input) == 0 {
    fmt.Println("Input format is not right!")
    os.Exit(1)
  }

  res1 := 0
  res2 := 0
  for _, history := range input {
    res1 += task(parseInput(strings.Fields(history)), true)
    res2 += task(parseInput(strings.Fields(history)), false)
  }

  fmt.Println(res1)
  fmt.Println(res2)
}

func parseInput(s []string) []int {
  res := []int{}
  for _, v := range s {
    num, err := strconv.Atoi(v)
    if err != nil {
      fmt.Println("Failed to parse int.")
    }

    res = append(res, num)
  }

  return res
}

func task(nums []int, nextHistory bool) int {
  res := 0
  
  size := len(nums)


  for i := size - 1; i > 1; i-- {
    nonZeroDiff := 0

    if nextHistory {
      for j := 0; j < i; j++ {
        nums[j] = nums[j + 1] - nums[j]
        if nums[j] != 0 {
          nonZeroDiff++
        }
      }
      
      res += nums[i]
    } else {
      for j := size - 1; j >= size - i; j-- {
        nums[j] = nums[j] - nums[j - 1]
        if nums[j] != 0 {
          nonZeroDiff++
        }
      }
      diff := nums[size - i - 1] 
      if (size - i) % 2 == 1 {
        diff *= -1
      }
      res -= diff
    }
    if nonZeroDiff == 0 {
      break
    }
  }

  return res
}
