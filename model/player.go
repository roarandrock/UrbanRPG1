package models

import "fmt"

/*
Stats are what you have
Traits are what you are
Base Struct is where you are
*/

//Player is base struct
type Player struct {
	Name   string
	Height int
	Loc    int
}

//PlayerStats keeps track of stats, 1-5 scale or 100?, 11
type PlayerStats struct {
	Income       int
	Connections  int //trait?
	Independence int //should be a trait?
	Stuff        int
	Brands       int
	Knowledge    int
	Experience   int
	Energy       int
	Beauty       int
	Smooth       int
	Stress       int
}

//PlayerTraits tracks traits
type PlayerTraits struct {
	Education  int
	Hygiene    int //should be a stat?
	Soul       int
	Alcoholism int
	Persuasive int
	Lucky      int
	Arrogant   int
}

var currentPlayer Player
var currentPlayerStats PlayerStats
var currentPlayerTraits PlayerTraits

//GetCurrentPlayer returns current player struct
func GetCurrentPlayer() (Player, PlayerStats, PlayerTraits) {
	return currentPlayer, currentPlayerStats, currentPlayerTraits
}

func GetPlayerStats() PlayerStats {
	return currentPlayerStats
}

func UpdatePlayerStats(cps PlayerStats) {
	currentPlayerStats = cps
}

//GetPlayerLoc returns the locations as an int
func GetPlayerLoc() int {
	return currentPlayer.Loc
}

//UpdateCurrentPlayer updates the current player struct
func UpdateCurrentPlayer(np Player) {
	currentPlayer = np
}

//SetPlayerAll sets initial stats
func SetPlayerAll() {
	currentPlayer.Name = "Ready"
	currentPlayer.Height = 68
	currentPlayer.Loc = 1
	currentPlayerStats.Income = 1 //determined by job
	currentPlayerStats.Connections = 10
	currentPlayerStats.Independence = 80
	currentPlayerStats.Stuff = 1
	currentPlayerStats.Brands = 30
	currentPlayerStats.Knowledge = 30
	currentPlayerStats.Experience = 1
	currentPlayerStats.Energy = 50
	currentPlayerStats.Beauty = 50
	currentPlayerStats.Smooth = 20
	currentPlayerStats.Stress = 20
}

//ShowPlayerStats displays all stats
func ShowPlayerStats() {
	fmt.Println("Income:", currentPlayerStats.Income)
	fmt.Println("Connections: ", currentPlayerStats.Connections)
	fmt.Println("Independence: ", currentPlayerStats.Independence)
	fmt.Println("Stuff: ", currentPlayerStats.Stuff)
	fmt.Println("Brands: ", currentPlayerStats.Brands)
	fmt.Println("Knowledge: ", currentPlayerStats.Knowledge)
	fmt.Println("Experience: ", currentPlayerStats.Experience)
	fmt.Println("Energy: ", currentPlayerStats.Energy)
	fmt.Println("Beauty: ", currentPlayerStats.Beauty)
	fmt.Println("Smooth: ", currentPlayerStats.Smooth)
	fmt.Println("Stress: ", currentPlayerStats.Stress)
}

//DeltaStat returns updated stat
func DeltaStat(olds int, news int) int {
	var finals int
	max := 100 //configurable
	min := 0
	finals = olds + news //check first if news is zero?
	if finals < min {
		finals = min
	} else if finals > max {
		finals = max
	}
	return finals
}

//StatChange implements deltastats
func StatChange(nps PlayerStats) {
	//updates stats, prints delta stats (can be a separate function)
	//would it be easier if stats were mapped and tagged? Or just array - 0 is beauty, 10 is stuff
	cps := GetPlayerStats()

	cps.Beauty = DeltaStat(cps.Beauty, nps.Beauty)
	cps.Brands = DeltaStat(cps.Brands, nps.Brands)
	cps.Connections = DeltaStat(cps.Connections, nps.Connections)
	cps.Energy = DeltaStat(cps.Energy, nps.Energy)
	cps.Experience = DeltaStat(cps.Experience, nps.Experience)
	cps.Income = DeltaStat(cps.Income, nps.Income)
	cps.Independence = DeltaStat(cps.Independence, nps.Independence)
	cps.Knowledge = DeltaStat(cps.Knowledge, nps.Knowledge)
	cps.Smooth = DeltaStat(cps.Smooth, nps.Knowledge)
	cps.Stress = DeltaStat(cps.Stress, nps.Stress)
	cps.Stuff = DeltaStat(cps.Stuff, nps.Stuff)

	if nps.Beauty != 0 {
		fmt.Println("Beauty:", cps.Beauty)
	}
	if nps.Brands != 0 {
		fmt.Println("Brands:", cps.Brands)
	}
	if nps.Connections != 0 {
		fmt.Println("Connections:", cps.Connections)
	}
	if nps.Energy != 0 {
		fmt.Println("Energy:", cps.Energy)
	}
	if nps.Experience != 0 {
		fmt.Println("Experience:", cps.Experience)
	}
	if nps.Income != 0 {
		fmt.Println("Income:", cps.Income)
	}
	if nps.Independence != 0 {
		fmt.Println("Independence:", cps.Independence)
	}
	if nps.Knowledge != 0 {
		fmt.Println("Knowledge:", cps.Knowledge)
	}
	if nps.Smooth != 0 {
		fmt.Println("Smooth:", cps.Smooth)
	}
	if nps.Stress != 0 {
		fmt.Println("Stress:", cps.Stress)
	}
	if nps.Stuff != 0 {
		fmt.Println("Stuff:", cps.Stuff)
	}

	UpdatePlayerStats(cps)
}
