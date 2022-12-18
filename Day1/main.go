package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	inventory := [][]int{}
	current_inventory := []int{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			inventory = append(inventory, current_inventory)
			current_inventory = []int{}
			continue
		}
		calorie, err := strconv.Atoi(line)
		if err != nil {
			log.Panic(err)
		}
		current_inventory = append(current_inventory, calorie)
	}

	sums := []int{}
	sum := 0
	for _, current_inventory := range inventory {
		for _, calory := range current_inventory {
			sum += calory
		}
		sums = append(sums, sum)
		sum = 0
	}

	slices.Sort(sums)
	slices.SortFunc(sums, func(a, b int) bool { return a > b })

	fmt.Printf("Max calories: %d\n", sums[0])
	fmt.Printf("Top 3 calories: %d\n", Sum(sums[0:3]))
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[N Number](integers []N) N {
	var sum N
	for i := range integers {
		sum += integers[i]
	}
	return sum
}
