package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
    Row, Col int
}

func main() {
  file, err := os.Open("./day4/input")

  if err != nil {
    fmt.Println("Input file not found.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  copyCount := map[int]int{}
  matches := map[int]int{}
  lastStratchCardId := 0

  sum1 := 0
  sum2 := 0

  for scanner.Scan() {
    match , id := analyzeScratchCard(scanner.Text())
    if match > 0 {
      sum1 += 1 << (match - 1)
    }
    matches[id] = match
    copyCount[id] = 1

    if id > lastStratchCardId {
      lastStratchCardId = id
    }
  }

  for i := 1; i <= lastStratchCardId; i++ {
    match := matches[i]
    for j := 1; j <= match && i + j <= lastStratchCardId; j++ {
      copyCount[i + j] += copyCount[i]
    }
  }

  for i := 1; i <= lastStratchCardId; i++ {
    sum2 += copyCount[i]
  }

  fmt.Println(sum1)
  fmt.Println(sum2)
}

func analyzeScratchCard(s string) (int, int) {
  tokens := strings.Split(s, ": ")
  game := strings.Split(tokens[1], " | ")

  gameId, err := strconv.Atoi(strings.Fields(tokens[0])[1])

  if err != nil {
    fmt.Println(err)
  }

  winningNumberString := game[0]
  yourNumberString := game[1]

  winningNumbers := map[int]bool{}

  for _, v := range strings.Fields(winningNumberString) {
    num, err := strconv.Atoi(v)

    if err != nil {
      fmt.Println(err)
    }

    winningNumbers[num] = true
  }

  match := 0

  for _, v := range strings.Fields(yourNumberString) {
    num, err := strconv.Atoi(v)

    if err != nil {
      fmt.Println(err)
    }

    if winningNumbers[num] {
      match++
    }
  }

  if match == 0 {
    return 0, gameId
  }

  return match, gameId
}

