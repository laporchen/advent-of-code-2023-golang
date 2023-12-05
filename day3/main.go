package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
    Row, Col int
}

func main() {
  file, err := os.Open("./day3/input")

  if err != nil {
    fmt.Println("Input file not found.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  var schematic []string

  for scanner.Scan() {
    schematic = append(schematic, scanner.Text())
  }

  fmt.Println(checkGear(schematic))
}

func checkGear(schematic []string) (int, int) {
  height := len(schematic)

  if height == 0 {
    return 0, 0
  }

  width := len(schematic[0])

  visit := make([][]bool, height)
  adjs := make([][][]int, height)

  for i := range visit {
    visit[i] = make([]bool , width)
    adjs[i] = make([][]int, width)
  }

  updateAndGetGearInfo := func(col, row int) (int, bool) {
    currentNumber := 0
    isGear := false

    index := row

    for index < width {
      c := schematic[col][index]
      if (c < '0' || c > '9') {
        break
      }

      currentNumber = currentNumber * 10 + int(c - '0')
      visit[col][index] = true

      index++
    }

    updated := map[Position]bool{}

    for row < index {
      top := max(0, col - 1)
      bottom := min(height, col + 2)
      left := max(0, row - 1)
      right := min(width, row + 2)
      for i := top; i < bottom; i++ {
        for j := left; j < right; j++ {
          pos := Position{ Col: i, Row: j}
          adj := schematic[i][j]
          
          if (adj < '0' || adj > '9') && adj != '.' && !updated[pos] {
            adjs[i][j] = append(adjs[i][j] , currentNumber)
            isGear = true
            updated[pos] = true
          }
        }
      }

      row++
    }

    return currentNumber, isGear
  }

  ans1 := 0
  ans2 := 0

  for i := range schematic {
    for j := range schematic[i] {
      if (!visit[i][j] && schematic[i][j] >= '0' && schematic[i][j] <= '9') {
        id, isGear := updateAndGetGearInfo(i, j)

        if isGear {
          ans1 += id
        }
      }
    }
  }

  for i := range schematic {
    for j := range schematic[i] {
      if len(adjs[i][j]) == 2 {
        ans2 += adjs[i][j][0] * adjs[i][j][1]
      }
    }
  }


  return ans1, ans2
}
