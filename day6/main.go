package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
  time, record int
}

func main() {
  file, err := os.Open("./day6/input")

  if err != nil {
    fmt.Println("Input file not found.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)


  input := []string{}

  for scanner.Scan() {
    input = append(input, scanner.Text())
  }

  if len(input) != 2 {
    fmt.Println("Input format is not right!")
    os.Exit(1)
  }

  raceDatas, mainRace := parseInput(input[0], input[1])

  fmt.Println("running task 1")
  res1 := task1(raceDatas)
  res2 := task1([]Race{mainRace})
  // fmt.Println("running task 2")
  // res2 := task2(seeds, relations, nextEnvRelation)

  fmt.Println(res1, res2)
}

func parseInput(timeStr, recordStr string) ([]Race, Race) {
  timeTokens := strings.Fields(timeStr)
  recordTokens := strings.Fields(recordStr)
  if len(timeTokens) != len(recordTokens) {
    fmt.Println("input cannot be matched!")
  }

  races := []Race{}

  for i := 1;i < len(timeTokens); i++ {
    time, err := strconv.Atoi(timeTokens[i])
    record, err := strconv.Atoi(recordTokens[i])
    if err != nil {
      fmt.Println("Failed to convert input to number")
    }
    races = append(races, Race{ time: time, record: record})
  }

  mainRaceTimeStr := strings.Join(timeTokens[1:], "")
  mainRaceRecordStr := strings.Join(recordTokens[1:], "")

  mainRaceTime, _ := strconv.Atoi(mainRaceTimeStr)
  mainRaceRecord, _ := strconv.Atoi(mainRaceRecordStr)

  mainRace := Race{ time: mainRaceTime, record: mainRaceRecord }

  return races, mainRace
}


func task1(races []Race) int {
  
  res := 1
  for _, race := range races {
    t := float64(race.time)
    r := float64(race.record)
    d := math.Pow(t, 2) - 4 * r
    if d < 0 {
      fmt.Println("No solution!")
      continue
    }
    floor := (t - math.Sqrt(d)) / 2
    ceil := (t + math.Sqrt(d)) / 2
    if floor == math.Ceil(floor) {
      floor += 1
    }
    if ceil == math.Floor(ceil) {
      ceil -= 1
    }
    floor = math.Ceil(floor)
    ceil = math.Floor(ceil)

    res *= (int(ceil) - int(floor)+ 1)
  }
  return res
}
