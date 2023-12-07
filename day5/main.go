package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
  src, dest, length int
}

type Mapping struct {
  src, dest string
}

func main() {
  file, err := os.Open("./day5/input")

  if err != nil {
    fmt.Println("Input file not found.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  var seeds []int

  relations := map[Mapping][]Range{}
  nextEnvRelation := map[string]string{}
  prevEnvRelation := map[string]string{}

  currentSource := ""
  currentDest := ""
  isInputtingMapping := false

  for scanner.Scan() {
    s := scanner.Text()

    if len(s) == 0 {
      isInputtingMapping = true
      continue
    }
    key := strings.Fields(s)[0]

    if strings.Compare(key, "seeds:") == 0 {
      seeds = parseSeed(s)
    } else {
      if isInputtingMapping {
        currentSource, currentDest = parseMapping(s)
        nextEnvRelation[currentSource] = currentDest
        prevEnvRelation[currentDest] = currentSource
        isInputtingMapping = false
      } else {
        dest, src, length := parseRelation(s)
        mapping := Mapping{ dest: currentDest, src: currentSource}
        relations[mapping] = append(relations[mapping], Range{ dest: dest, src: src, length: length })
      }
    }
  }

  fmt.Println("running task 1")
  res1 := task1(seeds, relations, nextEnvRelation)
  fmt.Println("running task 2")
  res2 := task2(seeds, relations, nextEnvRelation)

  fmt.Println(res1, res2)
}

func parseSeed(s string) []int {
  tokens := strings.Split(s, ": ")
  seedStrArr := strings.Fields(tokens[1])
  seeds := []int{}

  for _, seedStr := range seedStrArr {
    val, err := strconv.Atoi(seedStr)
    if err != nil {
      fmt.Println(err)
    }
    seeds = append(seeds, val)
  }
  
  return seeds
}

func parseMapping(s string) (string, string) {
  tokens := strings.Fields(s)
  mapping := strings.Split(tokens[0], "-to-")

  if len(mapping) != 2 {
    fmt.Println("Mapping format failed!")
  }
  
  return mapping[0], mapping[1]
}

func parseRelation(s string) (int, int, int) {
  tokens := strings.Fields(s)
  if len(tokens) != 3 {
    fmt.Println("mapping relation format is wrong!")
  }

  dest, err := strconv.Atoi(tokens[0])
  src, err := strconv.Atoi(tokens[1])
  length, err := strconv.Atoi(tokens[2])

  if err != nil {
    fmt.Println(err)
  }

  return dest, src, length

}

func task1(seeds []int, relations map[Mapping][]Range, mapping map[string]string) int {
  findPosition := func(id int) int {
    current := "seed"
    next := mapping[current]

    for next != "" {
      rangeArr := relations[Mapping{ src:current, dest: next}]
      tmpId := -1
      for _, relation := range rangeArr {
        src := relation.src 
        dest := relation.dest
        length := relation.length 
        if (id >= src && id - src < length) {
            tmpId = dest + id - src
        }
      }
      
      if tmpId != -1 {
        id = tmpId
      }
      current = next
      next = mapping[current]
    }

    return id
  }

  res := -1
  for _, seed := range seeds {
    ret := findPosition(seed)
    if res == -1 || res > ret {
      res = ret
    }
  }
  return res
}

type Interval struct {
  start, end int
}

// brute force for now
func task2(seedPairs []int, relations map[Mapping][]Range, mapping map[string]string) int {

  var solveInterval func([]Interval, string) int

  solveInterval = func(itvs []Interval, env string) int {
    res := -1

    if env == "location" {
      for _, itv := range itvs {
        if res == -1 || res > itv.start {
          res = itv.start
        }
      }
      return res
    }

    next := mapping[env]

    rangeArr := relations[Mapping{ dest: next, src: env }]
    newIntervals := []Interval{}

    count := len(itvs)

    for i := 0; i < count; i++ {
      itv := itvs[i]
      found := false
      for _, rng := range rangeArr {
        src := rng.src
        dest := rng.dest
        length := rng.length
        rngItv := Interval{ start: src, end: src + length - 1 } 
        left := max(rngItv.start, itv.start)
        right := min(rngItv.end, itv.end)
        if left <= right {
          overlapped := Interval{start: left - src + dest, end: right - src + dest }
          newIntervals = append(newIntervals, overlapped)
          if itv.start < left {
            itvs = append(itvs, Interval{start: itv.start, end: left - 1 })
          }
          if itv.end > right {
            itvs = append(itvs, Interval{start: right + 1, end: itv.end })
          }
          count = len(itvs)

          found = true
          break
        }
      }
      
      if !found {
        newIntervals = append(newIntervals, itv)
      }
    }

    return solveInterval(newIntervals, next)
  }

  intervals := []Interval{}

  for i := 0; i < len(seedPairs); i += 2 {
    intervals = append(intervals, Interval{ start: seedPairs[i], end: seedPairs[i] + seedPairs[i+1] - 1 })
  }

  
  return solveInterval(intervals, "seed")
}
