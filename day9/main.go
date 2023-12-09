package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
    nums := parseInput(strings.Fields(history))
    res1 += task(nums, true)

    nums = parseInput(strings.Fields(history))
    slices.Reverse(nums)
    res2 += task(nums, false)
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

func task(nums []int, findNext bool) int {
  res := 0
  
  size := len(nums)

  for i := size - 1; i > 1; i-- {
    nonZeroDiff := 0

    for j := 0; j < i; j++ {
      nums[j] = nums[j + 1] - nums[j]
      if nums[j] != 0 {
        nonZeroDiff++
      }
    }

    sign := 1
    if !findNext && (size - 1) % 2 == 1 {
      sign = -1
    }
    
    res += nums[i] * sign

    if nonZeroDiff == 0 {
      break
    }
  }

  return res
}
