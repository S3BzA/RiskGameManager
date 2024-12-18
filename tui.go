package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "gopkg.in/yaml.v3"
)

// DisplayMenu displays the main menu and returns the selected option.
func DisplayMenu() int {
    fmt.Println("Risk Game State Manager")
    fmt.Println("1. Create a Save")
    fmt.Println("2. Load a Save")
    fmt.Println("3. Delete a Save")
    fmt.Println("4. Update a Save")
    fmt.Println("Enter your choice (1-4):")

    var choice int
    _, err := fmt.Scan(&choice)
    if err != nil {
        fmt.Println("Invalid input. Please enter a number between 1 and 4.")
        return DisplayMenu() // Retry on invalid input.
    }

    if choice < 1 || choice > 4 {
        fmt.Println("Invalid choice. Please select a valid option.")
        return DisplayMenu() // Retry on out-of-range input.
    }

    return choice
}

// PromptUser prompts the user for input with the given message.
func PromptUser(message string) string {
    fmt.Println(message)
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        return strings.TrimSpace(scanner.Text())
    }

    return ""
}

// SaveGame saves the game state to a YAML file.
func SaveGame(filename string, gameState interface{}) error {
    data, err := yaml.Marshal(gameState)
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

// LoadGame loads the game state from a YAML file.
func LoadGame(filename string, gameState interface{}) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    return yaml.Unmarshal(data, gameState)
}

// HandleMenuOption handles the selected menu option.
func HandleMenuOption(choice int) {
    switch choice {
    case 1:
        fmt.Println("Creating a new save...")
        filename := PromptUser("Enter the save file name (e.g., my_conquest):")
        
        // Prompt for the number of players
        numPlayers := 0
        for numPlayers < 2 || numPlayers > 5 {
            fmt.Println("Enter the number of players (2-5):")
            fmt.Scan(&numPlayers)
            if numPlayers < 2 || numPlayers > 5 {
                fmt.Println("Invalid number of players. Please enter a number between 2 and 5.")
            }
        }

        // Prompt for player names
        playerNames := make([]string, numPlayers)
        for i := 0; i < numPlayers; i++ {
            playerNames[i] = PromptUser(fmt.Sprintf("Enter name for player %d:", i+1))
        }

		if !strings.HasSuffix(filename, ".yaml") {
			filename += ".yaml"
		}
		gameState := InitializeGame(playerNames)
		if err := SaveGame(filename, gameState); err != nil {
			fmt.Println("Error saving game:", err)
		} else {
			fmt.Println("Game successfully saved to", filename)
		}

    case 2:
        fmt.Println("Loading a save...")
        filename := PromptUser("Enter the save file name to load:")
		if !strings.HasSuffix(filename, ".yaml") {
			filename += ".yaml"
		}
        var gameState GameState
        if err := LoadGame(filename, &gameState); err != nil {
            fmt.Println("Error loading game:", err)
        } else {
            fmt.Printf("Game successfully loaded. Players: %+v\n", gameState.Players)
        }

    case 3:
        fmt.Println("Deleting a save...")
        filename := PromptUser("Enter the save file name to delete:")
		if !strings.HasSuffix(filename, ".yaml") {
			filename += ".yaml"
		}
        if err := os.Remove(filename); err != nil {
            fmt.Println("Error deleting file:", err)
        } else {
            fmt.Println("Save successfully deleted.")
        }

    case 4:
		fmt.Println("Updating a save...")
		filename := PromptUser("Enter the save file name to update:")
		if !strings.HasSuffix(filename, ".yaml") {
			filename += ".yaml"
		}
		var gameState GameState
		if err := LoadGame(filename, &gameState); err != nil {
			fmt.Println("Error loading game for update:", err)
			return
		}
	
		for i := range gameState.Players {
			player := &gameState.Players[i]
			fmt.Printf("Updating player: %s\n", player.Name)
	
			// Update troop count
			troops := PromptUser(fmt.Sprintf("Enter new troop count for %s:", player.Name))
			fmt.Sscanf(troops, "%d", &player.TroopCount)
	
			// Update territories
			for {
				territoryName := PromptUser("Enter territory name to update (or 'done' to finish):")
				if territoryName == "done" {
					break
				}
				territory, exists := gameState.Territories[territoryName]
				if !exists {
					fmt.Println("Territory does not exist.")
					continue
				}
				troops := PromptUser(fmt.Sprintf("Enter troop count for %s:", territoryName))
				var troopCount int
				fmt.Sscanf(troops, "%d", &troopCount)
				territory.TroopCount = troopCount
				territory.Owner = player.Name // Set the current player as the owner of the territory

				// Remove territory from other players' lists
				for j := range gameState.Players {
					if gameState.Players[j].Name != player.Name {
						for k, t := range gameState.Players[j].Territories {
							if t == territoryName {
								gameState.Players[j].Territories = append(gameState.Players[j].Territories[:k], gameState.Players[j].Territories[k+1:]...)
								break
							}
						}
					}
				}
	
				// Add territory to player's list of territories
				player.Territories = append(player.Territories, territoryName)
			}
	
			// Update cards
			for {
				action := PromptUser("Do you want to add or remove a card? (add/remove/done):")
				if action == "done" {
					break
				}
				if action == "add" {
					territory := PromptUser("Enter territory name for the card:")
					troopType := PromptUser("Enter troop type for the card:")
					player.Cards = append(player.Cards, Card{TerritoryName: territory, Type: troopType})
				} else if action == "remove" {
					territory := PromptUser("Enter territory name of the card to remove:")
					for j, card := range player.Cards {
						if card.TerritoryName == territory {
							player.Cards = append(player.Cards[:j], player.Cards[j+1:]...)
							break
						}
					}
				}
			}
		}
	
		if err := SaveGame(filename, gameState); err != nil {
			fmt.Println("Error updating game:", err)
		} else {
			fmt.Println("Game successfully updated.")
		}

    default:
        fmt.Println("Unknown option. This should never happen.")
    }
}