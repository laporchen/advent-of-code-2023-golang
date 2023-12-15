package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spring struct {
  record string
  groups []int
}

func main() {
  file, err := os.Open("./day12/input")

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

  springs := parseInput(input)
  fmt.Println(task1(springs))
  fmt.Println(task2(springs))
}

func parseInput(s []string) []Spring {
  springs := []Spring{}
  for _, row := range s {
    tokens := strings.Fields(row)
    record := tokens[0]
    groups := []int{}

    for _, v := range strings.Split(tokens[1], ",") {
      n, err := strconv.Atoi(v)
      if err != nil {
        fmt.Print("Failed to parse int")
      }
      groups = append(groups, n)
    }

    springs = append(springs, Spring{ record: record, groups: groups}) 
  }

  return springs
}


func count(spring Spring) int {
  springLen := len(spring.record)
  groupsLen := len(spring.groups) 

  var dp [][][]int // dp[i][j][k], first i records that satisfy j groups rule with k continue springs
  dp = make([][][]int, springLen + 1)
  for i := range dp {
    dp[i] = make([][]int, groupsLen + 1)
    for j := range dp[i] {
      dp[i][j] = make([]int, springLen + 1)
    }
  }

  dp[0][0][0] = 1

  for i := 1; i <= springLen; i++ {
    for j := 0; j < groupsLen; j++ {
      for k := 0; k < i; k++ {
        c := spring.record[i-1]
        if (c == '.' || c == '?') {
          dp[i][j][0] += dp[i-1][j][k]
				} 
        if (c == '#' || c == '?') {
          if spring.groups[j] == k + 1 {
            dp[i][j+1][k+1] += dp[i- k - 1][j][k]
          } else {
            dp[i][j][k+1] += dp[i-1][j][k]
          }
        }
      }
    }
  }

  for i, v := range dp {
    for j, w := range v {
      for o, k := range w {
        fmt.Println("first", i, "record, match first", j , "groups,match", o, "continue springs:", k)
      }
    }
    fmt.Println()
  }
  ret := 0
  for _, v := range dp[springLen][groupsLen] {
    ret += v
  }

  return ret
}
// there's probably a dp solution, I can't figure one rn.
func countPossible(spring Spring) int {
  unknownIndex := []int{}
  record := strings.Split(spring.record, "")
  record = append(record, " ") // add an empty string for easier logic handle
  groups := spring.groups
  for i, c := range spring.record {
    if c == '?' {
      unknownIndex = append(unknownIndex, i)
    }
  }

  ret := 0

  for i := 0; i < (1 << len(unknownIndex)); i++ {
    for j := 0; j < len(unknownIndex); j++ {
      if (i & (1 << j)) > 0 {
        record[unknownIndex[j]]= "#"
      } else {
        record[unknownIndex[j]]= "."
      }
    }

    conti := 0
    currentIdx := 0
    ok := true
    for _, c := range record {

      if c == "#" {
        conti++
      } else {
        if conti != 0 {
          if currentIdx == len(groups) {
            ok = false
            break
          }
          if conti != groups[currentIdx] {
            ok = false
            break
          } else {
            conti = 0
            currentIdx++
          }
        }
      }
    }
    
    if currentIdx != len(groups) {
      ok = false
    }

    if ok {
      ret++
    }
  }

  return ret
}

func task1(springs []Spring) int {
  res := 0
  for _, s := range springs {
    c := count(s)
    fmt.Println(c)
    res += c
  } 

  return res
}

func appendCopyToSlice[T any](original []T, times int) []T {
	for i := 0; i < times; i++ {
		copyOfSlice := make([]T, len(original))
		copy(copyOfSlice, original)
		original = append(original, copyOfSlice...)
	}
	return original
}

func unfold(s Spring) Spring {
  newGroups := appendCopyToSlice(s.groups, 5)
  var strBuilder strings.Builder
  for i := 0; i < 5; i++ {
    strBuilder.WriteString(s.record)
    if i < 5 - 1 {
      strBuilder.WriteString("?")
    }
  }
  newRecord := strBuilder.String()

  return Spring{ record: newRecord, groups: newGroups }
}
func task2(springs []Spring) int {
  res := 0
  for i, s := range springs {
    fmt.Println(i)
    res += countPossible(unfold(s))
  } 

  return res
}
