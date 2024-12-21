package main

import (
	// "encoding/json"
	"fmt"
	"math/rand"
)

// Territory represents a game territory.
type Territory struct {
	Name       string
	Continent  string
	Owner      string
	TroopCount int
}

// Card represents a territory card.
type Card struct {
	TerritoryName string
	Type          string // Infantry, Cavalry, Artillery, or Wild
}

// Player represents a player in the game.
type Player struct {
	Name       string
	TroopCount int
	Cards      []Card
	Territories []string
}

// GameState holds the state of the game.
type GameState struct {
	Players    []Player
	Territories map[string]*Territory
	Deck       []Card
}

// InitializeGame sets up the game state.
func InitializeGame(playerNames []string) *GameState {
	// Define territories.
	territoryData := map[string]string{
		"Alaska": "North America",
		"Northwest territory": "North America",
		"Alberta": "North America",
		"Ontario": "North America",
		"Eastern Canada": "North America",
		"Greenland": "North America",
		"Eastern United States": "North America",
		"Western United States": "North America",
		"Central America": "North America",
		"Venezuela": "South America",
		"Brazil": "South America",
		"Peru": "South America",
		"Argentina": "South America",
		"Egypt": "Africa",
		"North Africa": "Africa",
		"East Africa": "Africa",
		"Central Africa": "Africa",
		"South Africa": "Africa",
		"Madagascar": "Africa",
		"Iceland": "Europe",
		"Scandinavia": "Europe",
		"Russia": "Europe",
		"Northern Europe": "Europe",
		"Southern Europe": "Europe",
		"Western Europe": "Europe",
		"Great Britain": "Europe",
		"Middle East": "Asia",
		"Afghanistan": "Asia",
		"Ural": "Asia",
		"Siberia": "Asia",
		"Yakutsk": "Asia",
		"Irkutsk": "Asia",
		"Mongolia": "Asia",
		"China": "Asia",
		"India": "Asia",
		"Southeast Asia": "Asia",
		"Japan": "Asia",
		"Kamchatka": "Asia",
		"New Guinea": "Australia",
		"Indonesia": "Australia",
		"Eastern Australia": "Australia",
		"Western Australia": "Australia",
	}

	// Create territories.
	territories := make(map[string]*Territory)
	for name, continent := range territoryData {
		territories[name] = &Territory{Name: name, Continent: continent}
	}

	// Create players.
	players := make([]Player, len(playerNames))
	initialTroopCount := map[int]int{
		2: 40, 3: 35, 4: 30, 5: 25,
	}[len(playerNames)]

	for i, name := range playerNames {
		players[i] = Player{Name: name, TroopCount: initialTroopCount, Cards: []Card{}}
	}

	return &GameState{Players: players, Territories: territories, Deck: []Card{}}
}

// AddCard manually records a territory card.
func (g *GameState) AddCard(playerName, territoryName, cardType string) error {
	player := g.getPlayer(playerName)
	territory := g.Territories[territoryName]

	if player == nil {
		return fmt.Errorf("player %s not found", playerName)
	}
	if territory == nil {
		return fmt.Errorf("territory %s not found", territoryName)
	}

	player.Cards = append(player.Cards, Card{TerritoryName: territoryName, Type: cardType})
	return nil
}

// AllocateTroop places a troop on a territory for a player.
func (g *GameState) AllocateTroop(playerName, territoryName string) error {
	player := g.getPlayer(playerName)
	territory := g.Territories[territoryName]

	if player == nil {
		return fmt.Errorf("player %s not found", playerName)
	}
	if territory == nil {
		return fmt.Errorf("territory %s not found", territoryName)
	}
	if territory.Owner != "" && territory.Owner != playerName {
		return fmt.Errorf("territory %s is already owned by %s", territoryName, territory.Owner)
	}

	if player.TroopCount <= 0 {
		return fmt.Errorf("player %s has no troops left to allocate", playerName)
	}

	if territory.Owner == "" {
		territory.Owner = playerName
		player.Territories = append(player.Territories, territoryName)
	}

	territory.TroopCount++
	player.TroopCount--
	return nil
}

// getPlayer finds a player by name.
func (g *GameState) getPlayer(name string) *Player {
	for i := range g.Players {
		if g.Players[i].Name == name {
			return &g.Players[i]
		}
	}
	return nil
}

// The attacker can role up to 3 dice, this value get's entered by user
// generate 1-3 random numbers between 1-6 for attacker, and store them in ascending order in a slice
// sort the rolls in descending order
// The defender can role up to 2 dice, this value get's entered by user
// generate 1-2 random numbers between 1-6 for defender, and store them in descending order in a slice
// sort the rolls in descending order
// Now we compare the 2 highest rolls of attacker and defender
// if attacker roll is higher, defender looses 1 troop
// if defender roll is higher or equal, attacker looses 1 troop
func simulateBattle() {
	fmt.Println("Simulating a battle...")

	attackerDice := 0
	for attackerDice < 1 || attackerDice > 3 {
		fmt.Println("Enter the number of dice the attacker will roll (1-3):")
		fmt.Scan(&attackerDice)
		if attackerDice < 1 || attackerDice > 3 {
			fmt.Println("Invalid number of dice. Please enter a number between 1 and 3.")
		}
	}

	attackerRolls := make([]int, attackerDice)
	for i := 0; i < attackerDice; i++ {
		attackerRolls[i] = 1 + rand.Intn(6)
	}

	for i := 0; i < len(attackerRolls); i++ {
		for j := i + 1; j < len(attackerRolls); j++ {
			if attackerRolls[i] < attackerRolls[j] {
				attackerRolls[i], attackerRolls[j] = attackerRolls[j], attackerRolls[i]
			}
		}
	}

	fmt.Println("Attacker rolls:", attackerRolls)

	defenderDice := 0
	for defenderDice < 1 || defenderDice > 2 {
		fmt.Println("Enter the number of dice the defender will roll (1-2):")
		fmt.Scan(&defenderDice)
		if defenderDice < 1 || defenderDice > 2 {
			fmt.Println("Invalid number of dice. Please enter a number between 1 and 2.")
		}
	}

	defenderRolls := make([]int, defenderDice)
	for i := 0; i < defenderDice; i++ {
		defenderRolls[i] = 1 + rand.Intn(6)
	}

	for i := 0; i < len(defenderRolls); i++ {
		for j := i + 1; j < len(defenderRolls); j++ {
			if defenderRolls[i] < defenderRolls[j] {
				defenderRolls[i], defenderRolls[j] = defenderRolls[j], defenderRolls[i]
			}
		}
	}
	fmt.Println("Defender rolls:", defenderRolls)

	for i := 0; i < len(defenderRolls); i++ {
		if attackerRolls[i] > defenderRolls[i] {
			fmt.Println("Defender loses 1 troop.")
		} else {
			fmt.Println("Attacker loses 1 troop.")
		}
	}
}