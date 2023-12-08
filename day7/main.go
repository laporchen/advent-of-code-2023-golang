package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HandType int

const (
  HIGH HandType= iota
  ONEPAIR
  TWOPAIR
  THREEOFAKIND
  FULLHOUSE
  FOUROFAKIND
  FIVEOFAKIND
)

type Hand struct {
  cards [5]int
  originalCardOrder [5]int
  str string
  normalRank HandType
  jokerRank HandType
  bet int
}

func parseCard(s string) ([5]int, [5]int) {
  cardMapping := map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9}
  
  cards := [5]int{}
  originalCards := [5]int{}
  for i, c := range s {
    if c >= '2' && c <= '9' {
      originalCards[i] = int(c - '1')
    } else {
      originalCards[i] = cardMapping[c]
    }
  }

  copy(cards[:], originalCards[:])
  slices.Sort(cards[:])

  return cards, originalCards
}

func getHandRank(cards []int) HandType {
  handType := HIGH
  cnt := len(cards)

  if cnt >= 5 && cards[0] == cards[4] {
    handType = FIVEOFAKIND
  } else if cnt >= 4 && cards[0] == cards[3] || (cnt > 4 && cards[1] == cards[4]) {
    handType = FOUROFAKIND
  } else {
    isFullHouse := cnt == 5 && ((cards[0] == cards[2] && cards[3] == cards[4]) || (cards[0] == cards[1] && cards[2] == cards[4]))
    if isFullHouse {
      handType = FULLHOUSE 
    }
  }

  if handType == HIGH {
    for i := 0; i < cnt - 2; i++ {
      if cards[i] == cards[i + 2] {
        handType = THREEOFAKIND
        break
      }
    }
    if handType != THREEOFAKIND {
      pairCount := 0 
      for i := 0; i < cnt - 1; i++ {
        if cards[i] == cards[i+1] {
          pairCount++
        }
      }
      switch pairCount {
      case 1:
        handType = ONEPAIR
      case 2:
        handType = TWOPAIR
      case 0:
      default:
        fmt.Println("Unexpected pair count encountered.")
      }
    }
  }

  return handType
}

func makeHand(handStr string, bet int) Hand {
  cards, originalCards := parseCard(handStr) 
  cardNoJoker := []int{}
  for _, v := range cards {
    if v != 10 {
      cardNoJoker = append(cardNoJoker, v)
    }
  }

  normalRank := getHandRank(cards[:])
  noJokerRank := getHandRank(cardNoJoker)

  jokerRank := noJokerRank
  jokerCnt := 5 - len(cardNoJoker)

  if jokerCnt == 5 {
    jokerRank = FIVEOFAKIND
  } else if jokerCnt == 4 {
    jokerRank = FIVEOFAKIND
  } else if jokerCnt == 3 {
    if noJokerRank == HIGH {
      jokerRank = FOUROFAKIND
    } else if noJokerRank == ONEPAIR {
      jokerRank = FIVEOFAKIND
    }
  } else if jokerCnt == 2 {
    if noJokerRank == HIGH {
      jokerRank = THREEOFAKIND
    } else if noJokerRank == ONEPAIR {
      jokerRank = FOUROFAKIND
    } else if noJokerRank == THREEOFAKIND {
      jokerRank = FIVEOFAKIND
    }
  } else if jokerCnt == 1 {
    switch noJokerRank {
    case FOUROFAKIND:
      jokerRank = FIVEOFAKIND
    case THREEOFAKIND:
      jokerRank = FOUROFAKIND
    case TWOPAIR:
      jokerRank = FULLHOUSE
    case ONEPAIR:
      jokerRank = THREEOFAKIND
    case HIGH:
      jokerRank = ONEPAIR
    }
  }   

  return Hand{cards: cards, originalCardOrder: originalCards, normalRank: normalRank, jokerRank: jokerRank, bet: bet, str: handStr}
}

func main() {
  file, err := os.Open("./day7/input")

  if err != nil {
    fmt.Println("Input file not found.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)


  hands := []Hand{}

  for scanner.Scan() {
    hands = append(hands, parseInput(scanner.Text()))
  }

  res1 := task1(hands)
  res2 := task2(hands)

  fmt.Println(res1, res2)
}

func parseInput(s string) Hand{
  tokens := strings.Fields(s)
  cards := tokens[0]
  bet, _ := strconv.Atoi(tokens[1])

  return makeHand(cards, bet)
}

func task1(hands []Hand) int {
  cmp := func(l, r Hand) int {
    if l.normalRank != r.normalRank {
      if int(l.normalRank) < int(r.normalRank) {
        return -1
      }

      return 1
    }

    for i := 0; i < 5; i++ {
      if l.originalCardOrder[i] != r.originalCardOrder[i]  {
        if l.originalCardOrder[i] < r.originalCardOrder[i] {
          return -1
        } 
        
        return 1
      }
    }

    return 0
  }
  slices.SortFunc(hands[:], cmp)

  res := 0

  for i, h := range hands {
    res += (i + 1) * h.bet
  }
  return res
}

func task2(hands []Hand) int {
  cmp := func(l, r Hand) int {
    if l.jokerRank != r.jokerRank {
      if int(l.jokerRank) < int(r.jokerRank) {
        return -1
      }

      return 1
    }

    for i := 0; i < 5; i++ {
      if l.originalCardOrder[i] != r.originalCardOrder[i]  {
        if l.originalCardOrder[i] == 10 {
          return -1
        } else if r.originalCardOrder[i] == 10 {
          return 1
        }

        if l.originalCardOrder[i] < r.originalCardOrder[i] {
          return -1
        } 
        
        return 1
      }
    }

    return 0
  }
  slices.SortFunc(hands[:], cmp)

  res := 0

  for i, h := range hands {
    fmt.Println(h.str, h.jokerRank)
    res += (i + 1) * h.bet
  }
  return res
}
