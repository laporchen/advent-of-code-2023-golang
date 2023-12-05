package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)


func main() {
  file, err := os.Open("./day2/input")

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
    resId, ok := checkGame1(txt)
    if ok {
      sum1 += resId
    }
    sum2 += checkGame2(txt)
  }

  fmt.Println(sum1)
  fmt.Println(sum2)
}

func max[T constraints.Ordered](l, r T) T {
  if (l < r) {
    return r
  }

  return l
}

func parseGame(s string) (int, map[string]int) {
  tokens := strings.Split(s, ":") 
  gameIdTokens := strings.Split(tokens[0], " ")
  gameInfoTokens := strings.Split(tokens[1], ";")

  id, _ := strconv.ParseInt(gameIdTokens[1], 10, 8)

  minNums := map[string]int{ "green": 0, "red": 0, "blue": 0 }

  for _, round := range gameInfoTokens {
    cubes := strings.Split(round, ",")
    for _, cube := range cubes {
      data := strings.Split(cube[1:], " ")
      color := data[1]
      num, _ := strconv.ParseInt(data[0], 10, 8)
      minNums[color] = max(minNums[color], int(num))
    }
  }

  return int(id), minNums
}

func checkGame1(s string) (int, bool) {
  possible := true
  
  id, minNums := parseGame(s)

  if minNums["red"] > 12 || minNums["blue"] > 14 || minNums["green"] > 13 {
    possible = false
  }

  return int(id), possible
}

func checkGame2(s string) int {
  _, minNums := parseGame(s)
  return minNums["red"] * minNums["blue"] * minNums["green"]
}
