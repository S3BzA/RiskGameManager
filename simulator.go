package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Simulate a single battle round given the number of dice for attacker and defender
func simulateBattleRound(attackerDice, defenderDice int) (attackerLosses, defenderLosses int) {
	// Generate attacker rolls
	attackerRolls := make([]int, attackerDice)
	for i := 0; i < attackerDice; i++ {
		attackerRolls[i] = 1 + rand.Intn(6)
	}
	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(attackerRolls)))

	// Generate defender rolls
	defenderRolls := make([]int, defenderDice)
	for i := 0; i < defenderDice; i++ {
		defenderRolls[i] = 1 + rand.Intn(6)
	}
	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(defenderRolls)))

	// Compare rolls
	n := len(defenderRolls)
	if len(attackerRolls) < n {
		n = len(attackerRolls)
	}

	for i := 0; i < n; i++ {
		if attackerRolls[i] > defenderRolls[i] {
			defenderLosses++
		} else {
			attackerLosses++
		}
	}

	return
}

// Simulate a series of battles and compute win percentages
func simulateBattles(attackerDice, defenderDice, n int) {
	var attackerWins, defenderWins int

	for i := 0; i < n; i++ {
		attackerLosses, defenderLosses := simulateBattleRound(attackerDice, defenderDice)
		if defenderLosses > attackerLosses {
			attackerWins++
		} else {
			defenderWins++
		}
	}

	attackerWinPercentage := float64(attackerWins) / float64(n) * 100
	defenderWinPercentage := float64(defenderWins) / float64(n) * 100

	fmt.Printf("Results after %d battles (Attacker Dice: %d, Defender Dice: %d):\n", n, attackerDice, defenderDice)
	fmt.Printf("Attacker Win Percentage: %.2f%%\n", attackerWinPercentage)
	fmt.Printf("Defender Win Percentage: %.2f%%\n", defenderWinPercentage)
}

func Simulation() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var attackerDice, defenderDice, numBattles int

	fmt.Println("Enter the number of dice the attacker will roll (1-3):")
	fmt.Scan(&attackerDice)
	if attackerDice < 1 || attackerDice > 3 {
		fmt.Println("Invalid number of dice for attacker. Must be between 1 and 3.")
		return
	}

	fmt.Println("Enter the number of dice the defender will roll (1-2):")
	fmt.Scan(&defenderDice)
	if defenderDice < 1 || defenderDice > 2 {
		fmt.Println("Invalid number of dice for defender. Must be between 1 and 2.")
		return
	}

	fmt.Println("Enter the number of battles to simulate:")
	fmt.Scan(&numBattles)
	if numBattles <= 0 {
		fmt.Println("The number of battles must be greater than zero.")
		return
	}

	simulateBattles(attackerDice, defenderDice, numBattles)
}
