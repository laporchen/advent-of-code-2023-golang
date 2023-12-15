package main

import (
	"bufio"
	"fmt"
	"os"
)

type MirrorDirection int
const (
  ROW MirrorDirection = iota
  COL
  NOTFOUND
)

type Note struct {
  graph []string // true is rock
  height, width int
  mirrorPos int // between mirrorPos and mirrorPos + 1
  mirrorDir MirrorDirection // ROW -, COL |
}

func (n *Note) findMirrorPosition(smudge int) {
  foundMirror := false
  for i := 0; i < n.height - 1 && !foundMirror; i++ {
    isMirror := true
    error := 0
    
    for offset := 0; offset < n.height; offset++ {
      top := i - offset
      bottom := i + 1 + offset
      if top < 0 || bottom >= n.height {
        break
      }

      for j := 0; j < n.width; j++ {
        if n.graph[top][j] != n.graph[bottom][j] {
          if error >= smudge {
            isMirror = false
            break
          }
          error++
        }
      }
    }

    if isMirror && error == smudge {
      foundMirror = true
      n.mirrorDir = ROW
      n.mirrorPos = i
      break
    }
  }
  
  if foundMirror {
    return
  }

  for i := 0; i < n.width - 1 && !foundMirror; i++ {
    isMirror := true
    error := 0 
    for offset := 0; offset < n.width; offset++ {
      left := i - offset
      right := i + 1 + offset
      if left < 0 || right >= n.width {
        break
      }

      for j := 0; j < n.height; j++ {
        if n.graph[j][left] != n.graph[j][right] {
          if error >= smudge {
            isMirror = false
            break
          }
          error++
        }
      }
    }

    if isMirror && error == smudge {
      foundMirror = true
      n.mirrorDir = COL
      n.mirrorPos = i
      break
    }
  }

  if !foundMirror {
    fmt.Println("Mirror not found!")
    for _, row := range n.graph {
      fmt.Println(row)
    }
    fmt.Println()
  }
}

func main() {
  file, err := os.Open("./day13/input")

  if err != nil {
    fmt.Println("Input file not found.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)


  input := []string{}
  notes := []Note{}

  for scanner.Scan() {
    if len(scanner.Text()) == 0 {
      notes = append(notes, makeNote(input))
      input = []string{}
      continue
    }
    input = append(input, scanner.Text())
  }

  if len(input) > 0 {
    notes = append(notes, makeNote(input))
  }

  fmt.Println(task1(notes))
  fmt.Println(task2(notes))
}

func makeNote(s []string) Note {
  graph := [][]bool{}
  for _,row := range s {
    mappedRow := []bool{}
    for _,v := range row {
      if v == '.' {
        mappedRow = append(mappedRow, false)
      } else {
        mappedRow = append(mappedRow, true)
      }
    }
    graph = append(graph, mappedRow)
  }

  return Note{
    graph: s,
    height: len(graph),
    width: len(graph[0]),
    mirrorPos: 0,
    mirrorDir: NOTFOUND,
  }
}

func task1(notes []Note) int {
  res := 0
  for _, n := range notes {
    n.findMirrorPosition(0)
    if n.mirrorDir == ROW {
      res += (n.mirrorPos + 1)* 100
    } else if n.mirrorDir == COL {
      res += n.mirrorPos + 1
    }
  }

  return res
}

func task2(notes []Note) int {
  res := 0
  for _, n := range notes {
    n.findMirrorPosition(1)
    if n.mirrorDir == ROW {
      res += (n.mirrorPos + 1)* 100
    } else if n.mirrorDir == COL {
      res += n.mirrorPos + 1
    }
  }

  return res
}
