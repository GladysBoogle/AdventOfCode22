package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	max := 0
	for _, current_sum := range sums {
		if current_sum > max {
			max = current_sum
		}
	}

	fmt.Print(max)
}
