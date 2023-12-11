package main

import (
	"bufio"
	"fmt"
	"os"
)

type Galaxy struct {
  x, y int
}

func main() {
  file, err := os.Open("./day11/input")

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

  galaxies, height, width := parseInput(input)
  colPrefix, rowPrefix := getExpandedPrefix(galaxies, height, width)
  
  res1 := task(galaxies, colPrefix, rowPrefix, 1)
  res2 := task(galaxies, colPrefix, rowPrefix, 1000000)
  
  fmt.Println(res1, res2)
}

func parseInput(s []string) ([]Galaxy, int, int) {
  galaxies := []Galaxy{}
  for i, row := range s {
    for j, c := range row {
      if c == '#' {
        galaxies = append(galaxies, Galaxy{ x: j, y: i })
      }
    }
  }

  return galaxies, len(s), len(s[0])
}

func getExpandedPrefix(galaxies []Galaxy, height, width int) ([]int, []int) {
  row := make([]int, height)
  col := make([]int, width)

  for i := 1; i < width; i++ {
    col[i] = 1
  }

  for i := 1; i < height; i++ {
    row[i] = 1
  }

  for _, g := range galaxies {
    row[g.y] = 0
    col[g.x] = 0
  }

  for i := 1; i < width; i++ {
    col[i] += col[i-1]
  }

  for i := 1; i < height; i++ {
    row[i] += row[i-1]
  }

  return col, row
}

func task(galaxies []Galaxy, colPrefix, rowPrefix []int, multiplier int) int {
  res := 0
  if multiplier != 1 {
    multiplier--
  }
  for i := 0; i < len(galaxies); i++ {
    for j := i + 1; j < len(galaxies); j++ {
      fromX := galaxies[i].x
      fromY := galaxies[i].y
      toX := galaxies[j].x
      toY := galaxies[j].y
      if fromX > toX {
        fromX, toX = toX, fromX
      }
      if fromY > toY {
        fromY, toY = toY, fromY
      }
      path := toX - fromX + toY - fromY + (colPrefix[toX] - colPrefix[fromX] + rowPrefix[toY] - rowPrefix[fromY]) * (multiplier)
      res += path
    }
  }

  return res
}
