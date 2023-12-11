package main

import (
	"bufio"
	"fmt"
	"os"
)

type PipeType int

type Pipe struct {
  x, y int
  t PipeType
}

const (
  V PipeType = iota
  H
  TR
  TL
  DR
  DL
  START
  NONE
)

type Direction int
type DirectionOffset [2]int

var (
  dUP DirectionOffset = [2]int{ -1, 0 }
  dDOWN = [2]int{ 1, 0 }
  dRIGHT = [2]int{ 0, 1 }
  dLEFT = [2]int{ 0, -1 }
)


const (
  NODIR Direction = iota
  UP 
  DOWN
  LEFT
  RIGHT
  MARK
)

func main() {
  file, err := os.Open("./day10/input")

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
    fmt.Println("No input")
    os.Exit(1)
  }

  maze, startPipe := parseInput(input)

  res1, res2 := task(maze, startPipe)
  
  fmt.Println(res1, res2)
}

func parseInput(s []string) ([][]Pipe, Pipe)  {
  startPos := Pipe{}
  maze := [][]Pipe{}

  for i, row := range s {
    rowPipe := []Pipe{}
    for j, c := range row {
      t := NONE
      switch c {
      case 'L':
        t = TR
      case 'J':
        t = TL
      case 'F':
        t = DR
      case '7':
        t = DL
      case '|':
        t = V
      case '-':
        t = H
      case 'S':
        startPos.x = j
        startPos.y = i
        startPos.t = START
        t = START
      }

      rowPipe = append(rowPipe, Pipe{x: j, y: i, t: t})
    }
    maze = append(maze, rowPipe)
  }
  return maze, startPos
}

var validDirections = map[PipeType][]Direction {
  H: { LEFT, RIGHT },
  V: { UP, DOWN },
  DL: { DOWN, LEFT },
  DR: { DOWN, RIGHT },
  TL: { UP, LEFT },
  TR: { UP, RIGHT },
}

var reverseDirections = map[Direction]Direction {
  UP: DOWN,
  DOWN: UP,
  LEFT: RIGHT,
  RIGHT: LEFT,
}

var offset = map[Direction]DirectionOffset {
  UP: dUP,
  DOWN: dDOWN,
  LEFT: dLEFT,
  RIGHT: dRIGHT,
}

var verticalDirection = map[PipeType]PipeType {
  H: H,
  V: V,
}

var validConnections = map[PipeType]map[Direction][]PipeType {
    H: {
        LEFT:  {H, DR, TR},
        RIGHT: {H, DL, TL},
    },
    V: {
        UP:   {V, DL, DR},
        DOWN: {V, TL, TR},
    },
    DL: {
        DOWN: {V, TL, TR},
        LEFT:  {H, DR, TR},
    },
    DR: {
        DOWN: {V, TL, TR},
        RIGHT: {H, DL, TL},
    },
    TL: {
        UP:   {V, DL, DR},
        LEFT:  {H, DR, TR},
    },
    TR: {
        UP:   {V, DL, DR},
        RIGHT: {H, DL, TL},
    },
}

func canConnect(from, to PipeType, dir Direction) bool {
    validConnectionsForType, ok := validConnections[from]
    if !ok {
        return false
    }

    validToTypes, ok := validConnectionsForType[dir]
    if !ok {
        return false
    }
    //fmt.Println(validToTypes, to)
    for _, validToType := range validToTypes {
        if to == validToType {
            return true
        }
    }

    return false
}

func task(maze [][]Pipe, startPipe Pipe) (int,int) {
  possiblePipes := []PipeType{ V,H,DL,DR,TL,TR }

  height := len(maze)
  width := len(maze[0])

  res := 0

  var vis [][]Direction

  var traverse func(Pipe, int, Direction, []Pipe) []Pipe
  traverse = func(cur Pipe, length int, prevDir Direction, pathPipes []Pipe) []Pipe {
    if vis[cur.y][cur.x] != NODIR {
      if cur.x == startPipe.x && cur.y == startPipe.y {
        return pathPipes
      }

      return []Pipe{}
    }
    vis[cur.y][cur.x] = prevDir
    for _, dir := range validDirections[cur.t] {
      if reverseDirections[dir] == prevDir {
        continue
      }

      d := offset[dir]
      nx := cur.x + d[1]
      ny := cur.y + d[0]
      if (ny < height && ny >= 0) && (nx < width && nx >= 0) {
        next := maze[ny][nx]
        if canConnect(cur.t, next.t, dir) {
          if res:= traverse(next, length + 1, dir, append(pathPipes, next)); len(res) > 0 {
            return res
          }
        }
      }
    }

    return []Pipe{}
  }


  var ansPath []Pipe

  for _, p := range possiblePipes {
    vis = make([][]Direction, height)
    for i := range vis {
      vis[i] = make([]Direction, width)
    }
    
    maze[startPipe.y][startPipe.x].t = p
    if ret := traverse(maze[startPipe.y][startPipe.x], 0, MARK, []Pipe{}); len(ret) > 0 {
      if len(ret) > res {
        res = len(ret)
        ansPath = ret
      }
    }
  }


  cornerPipes := []Pipe{}


  for _, p := range ansPath {
    if p.t != H && p.t != V {
      cornerPipes = append(cornerPipes, p)
    }
  }

  // showlace formula, add start point to the end
  area := 0

  cornerPipes = append(cornerPipes, cornerPipes[0])
  for i := 0; i < len(cornerPipes) - 1; i++ {
    area += (cornerPipes[i].x * cornerPipes[i+1].y) - (cornerPipes[i].y * cornerPipes[i+1].x)
  }
  
  if area < 0 {
    area *= -1
  }
  area /= 2

  // pick's theorem
  insideTiles := area - (len(ansPath)) / 2 + 1

  return res / 2, insideTiles
} 
