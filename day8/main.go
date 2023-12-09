package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
  "container/list"

)

type Node struct {
  name string
  left, right *Node
}

func main() {
  file, err := os.Open("./day8/input")

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

  instruction, nodes := parseInput(input)
  
  res1 := task1(instruction, nodes)
  res2 := task2(instruction, nodes)

  fmt.Println(res1, res2)
  
}

func parseInput(input []string) (string, map[string]*Node) {
  if len(input) == 0 {
    fmt.Println("Input string is not correct.")
    return "", map[string]*Node{}
  }

  instruction := input[0] 
  nodes := map[string]*Node{}

  inputPattern := `(\w+)\s+=\s(.+)`
  childNodePattern := `\((\w+),\s+(\w+)\)`

  inputRegex := regexp.MustCompile(inputPattern)
  nodeRegex := regexp.MustCompile(childNodePattern)

  for i := 1; i < len(input); i++ {
    str := input[i]
    tokens := inputRegex.FindStringSubmatch(str)
    if len(tokens) >= 3 {
      children := nodeRegex.FindStringSubmatch(tokens[2])
      if len(children) != 3 {
        fmt.Println("Parse error")
        os.Exit(1)
      }

      nodeName := tokens[1]
      leftName := children[1]
      rightName := children[2]

      var left,right, node *Node

      if nodes[leftName] == nil {
        left = &Node{name: leftName}
        nodes[leftName] = left
      } else {
        left = nodes[leftName]
      }
      if nodes[rightName] == nil {
        right = &Node{name: rightName}
        nodes[rightName] = right
      } else {
        right = nodes[rightName]
      }

      if nodes[nodeName] == nil {
        node = &Node{name: nodeName}
        nodes[nodeName] = node
      } else {
        node = nodes[nodeName]
      }

      node.left = left
      node.right = right
    }
  }

  return instruction, nodes
}


func task1(instruction string, nodes map[string]*Node) int {
  current := "AAA"
  currentNode := nodes[current]
  attempts := 0

  for currentNode.name != "ZZZ" && attempts < 1e8 {
    currentInstruction := instruction[attempts % len(instruction)]

    if currentInstruction == 'R' {
      currentNode = currentNode.right 
    } else {
      currentNode = currentNode.left
    }
    attempts++
  }

  return attempts
}

func task2(instruction string, nodes map[string]*Node) int {
  attempts := 0
  q := list.New()

  minStepsToZ := []int{}

  for _, v := range nodes {
    if v.name[2] == 'A' {
      q.PushBack(v)
    }
  }

  for q.Len() > 0 { 
    n := q.Len()
    direction := instruction[attempts % len(instruction)]
    for n > 0 {
      front := q.Front()
      q.Remove(front)
      v := front.Value.(*Node)
      if v.name[2] == 'Z' {
        minStepsToZ = append(minStepsToZ, attempts)
      } else {
        if direction == 'R' {
          q.PushBack(v.right)
        } else {
          q.PushBack(v.left)
        }
      }
      n--
    }
    attempts++
  }

  res := minStepsToZ[0]

  var gcd func(int, int) int
  gcd = func(a, b int) int {
    if a % b == 0 {
      return b
    }
    return gcd(b, a % b)
  }

  for i := 1; i < len(minStepsToZ); i++ {
    res = (res * minStepsToZ[i]) / gcd(res, minStepsToZ[i])
  }

  return res
}
