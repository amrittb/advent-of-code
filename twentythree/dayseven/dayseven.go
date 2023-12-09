package dayseven

import (
	"sort"
	"strings"

	"github.com/amrittb/adventofcode/integer"
)

type HandType int

const (
  HIGH_CARD HandType = 0
  ONE_PAIR HandType = 1
  TWO_PAIR HandType = 2
  THREE_OF_A_KIND HandType = 3
  FULL_HOUSE HandType = 4
  FOUR_OF_A_KIND HandType = 5
  FIVE_OF_A_KIND HandType = 6
)

var cardRanks = map[rune]int {
  '2': 2,
  '3': 3,
  '4': 4,
  '5': 5,
  '6': 6,
  '7': 7,
  '8': 8,
  '9': 9,
  'T': 10,
  'J': 11,
  'Q': 12,
  'K': 13,
  'A': 14,
}

var cardRanksWithJoker = map[rune]int {
  'J': 1,
  '2': 2,
  '3': 3,
  '4': 4,
  '5': 5,
  '6': 6,
  '7': 7,
  '8': 8,
  '9': 9,
  'T': 10,
  'Q': 11,
  'K': 12,
  'A': 13,
}

type CamelCardHand struct {
  Hand string
  Bid int
  Type HandType
}

func NewCamelCardHand(hand string, bid int, t HandType) *CamelCardHand {
  return &CamelCardHand{Hand: hand, Bid: bid, Type: t}
}

func TotalWinnings(lines []string) int {
  hands := make([]*CamelCardHand, 0)

  for _, line := range lines {
    handStr, bid := getHandAndBid(line)
    handType := findHandType(handStr)
    hands = append(hands, NewCamelCardHand(handStr, bid, handType))
  }

  sort.Slice(hands, getRankingFunc(hands, cardRanks))

  return calculateTotalBid(hands)
}

func TotalWinningsWithJoker(lines []string) int {
  hands := make([]*CamelCardHand, 0)

  for _, line := range lines {
    handStr, bid := getHandAndBid(line)
    handType := findHandTypeWithJoker(handStr)
    hands = append(hands, NewCamelCardHand(handStr, bid, handType))
  }

  sort.Slice(hands, getRankingFunc(hands, cardRanksWithJoker))

  return calculateTotalBid(hands)
}

func getHandAndBid(line string) (string, int) {
  split := strings.Split(line, " ")
  return strings.Trim(split[0], " "), integer.TrimAndAtoi(split[1])
}

func findHandType(hand string) HandType {
  cardToCounts := map[rune]int{}
  for _, c := range hand {
    _, ok := cardToCounts[c]
    if !ok {
      cardToCounts[c] = 0
    }
    cardToCounts[c] += 1
  }

  countToCards := make([][]rune, 6)
  for k, v := range cardToCounts {
    countToCards[v] = append(countToCards[v], k)
  }

  // 1 card has been seen 5 times
  if len(countToCards[5]) == 1 {
    return FIVE_OF_A_KIND
  }

  // 1 card has been seen 4 times
  if len(countToCards[4]) == 1 {
    return FOUR_OF_A_KIND
  }

  // 1 card has been seen 3 times
  // 1 card has been seen 2 times
  if len(countToCards[3]) == 1 && len(countToCards[2]) == 1 {
    return FULL_HOUSE
  }

  // 1 card has been seen 3 times
  // 1 card has been seen 1 time
  // 1 card has been seen 1 time
  if len(countToCards[3]) == 1 && len(countToCards[1]) == 2 {
    return THREE_OF_A_KIND
  }

  // 2 card has been seen 2 times each
  if len(countToCards[2]) == 2 {
    return TWO_PAIR
  }

  // 1 card has been seen 2 times
  if len(countToCards[2]) == 1 {
    return ONE_PAIR
  }

  // all values are distinct
  return HIGH_CARD
}

func findHandTypeWithJoker(hand string) HandType {
  jocker := rune('J')
  cardToCounts := map[rune]int{}
  for _, c := range hand {
    _, ok := cardToCounts[c]
    if !ok {
      cardToCounts[c] = 0
    }
    cardToCounts[c] += 1
  }

  countToCards := make([][]rune, 6)
  for k, v := range cardToCounts {
    countToCards[v] = append(countToCards[v], k)
  }

  // 1 card has been seen 5 times
  if len(countToCards[5]) == 1 {
    return FIVE_OF_A_KIND
  }

  // 1 card has been seen 4 times
  if len(countToCards[4]) == 1 {
    // If either 4 same cards or 1 card is Jocker, then it's five of a kind
    _, ok := cardToCounts[jocker]; if ok {
      return FIVE_OF_A_KIND
    }
    return FOUR_OF_A_KIND
  }

  // 1 card has been seen 3 times
  // 1 card has been seen 2 times
  if len(countToCards[3]) == 1 && len(countToCards[2]) == 1 {
    // If either 3 same card or 2 same card is Jocker, then it's five of a kind
    _, ok := cardToCounts[jocker]; if ok {
      return FIVE_OF_A_KIND
    }

    return FULL_HOUSE
  }

  // 1 card has been seen 3 times
  // 1 card has been seen 1 time
  // 1 card has been seen 1 time
  if len(countToCards[3]) == 1 && len(countToCards[1]) == 2 {
    // If any one of the cards is jocker, then we can make it four of a kind.
    _, ok := cardToCounts[jocker]; if ok {
      return FOUR_OF_A_KIND
    }

    return THREE_OF_A_KIND
  }

  // 2 card has been seen 2 times each
  if len(countToCards[2]) == 2 {
    // If one of the two pairs is jocker, then we can make it four of a kind.
    if countToCards[2][0] == jocker || countToCards[2][1] == jocker {
      return FOUR_OF_A_KIND
    }
    // If the jocker is found in the different card, then we can make it full house.
    if countToCards[1][0] == jocker {
      return FULL_HOUSE
    }

    return TWO_PAIR
  }

  // 1 card has been seen 2 times
  if len(countToCards[2]) == 1 {
    // If any one of the remaining cards is jocker, then we can make it three of a kind.
    _, ok := cardToCounts[jocker]; if ok {
      return THREE_OF_A_KIND
    }
   
     return ONE_PAIR
  }

  // If any one of the cards is jocker, then we can make it one pair.
  _, ok := cardToCounts[jocker]; if ok {
    return ONE_PAIR
  }

  // all values are distinct
  return HIGH_CARD
}

func getRankingFunc(hands []*CamelCardHand, ranks map[rune]int) func(int, int) bool {
  return func(i, j int) bool {
    a := *hands[i]
    b := *hands[j]

    if a.Type ==  b.Type {
      for i := 0; i < 5; i++ {
        aRank := ranks[rune(a.Hand[i])]
        bRank := ranks[rune(b.Hand[i])]
        if aRank == bRank {
          continue
        }

        return aRank < bRank
      }

      return true
    }

    return a.Type < b.Type
  }
}



func calculateTotalBid(hands []*CamelCardHand) int {
  winning := 0
  for i, hand := range hands {
    winning += (i + 1) * hand.Bid
  }
  return winning
}


