package models

import "fmt"

/*
Stats are what you have
Traits are what you are
Base Struct is where you are

need to seperate stats into:
Cumulative (like knowledge and experience, gain over time) stat
Set (like income, always at a level per job, indepedence too) trait
Swing (changing daily like energy, hygience, stress) meters
*/

//Player is base struct
type Player struct {
	Name   string
	Height int
	Loc    int
}

type PlayerStats struct {
	Connections int
	Stuff       int
	Brands      int
	Knowledge   int
	Experience  int
}

type PlayerMeters struct {
	Hygiene int
	Soul    int
	Energy  int
	Stress  int
}

type PlayerTraits struct {
	Income       int
	Independence int
	Education    int
	Alcoholism   int
	Lucky        int
	Beauty       int
	Smooth       int
}

var currentPlayer Player
var currentPlayerStats PlayerStats
var currentPlayerTraits PlayerTraits
var currentPlayerMeters PlayerMeters

//GetCurrentPlayer returns current player struct
func GetCurrentPlayer() Player {
	return currentPlayer
}

//GetPlayerStats returns stats
func GetPlayerStats() PlayerStats {
	return currentPlayerStats
}

//UpdatePlayerStats updates stats
func UpdatePlayerStats(cps PlayerStats) {
	currentPlayerStats = cps
}

func GetPlayerMeters() PlayerMeters {
	return currentPlayerMeters
}

func UpdatePlayerMeters(cps PlayerMeters) {
	currentPlayerMeters = cps
}

func GetPlayerTraits() PlayerTraits {
	return currentPlayerTraits
}

func UpdatePlayerTraits(cps PlayerTraits) {
	currentPlayerTraits = cps
}

//GetPlayerLoc returns the locations as an int
func GetPlayerLoc() int {
	return currentPlayer.Loc
}

//UpdateCurrentPlayer updates the current player struct
func UpdateCurrentPlayer(np Player) {
	currentPlayer = np
}

//SetPlayerAll sets initial stats, assuming no jorb
func SetPlayerAll() {
	currentPlayer.Name = "Ready"
	currentPlayer.Height = 68
	currentPlayer.Loc = 1
	//stats
	currentPlayerStats.Stuff = 0
	currentPlayerStats.Brands = 20
	currentPlayerStats.Knowledge = 20
	currentPlayerStats.Experience = 0
	currentPlayerStats.Connections = 10
	//traits
	currentPlayerTraits.Alcoholism = 40
	currentPlayerTraits.Beauty = 50
	currentPlayerTraits.Education = 50
	currentPlayerTraits.Income = 0
	currentPlayerTraits.Independence = 100
	currentPlayerTraits.Lucky = 20
	currentPlayerTraits.Smooth = 20
	//meters
	currentPlayerMeters.Energy = 100
	currentPlayerMeters.Hygiene = 50
	currentPlayerMeters.Soul = 50
	currentPlayerMeters.Stress = 10
}

//ShowPlayerStats displays all stats
func ShowPlayerStats() {
	fmt.Println("Connections: ", currentPlayerStats.Connections)
	fmt.Println("Stuff: ", currentPlayerStats.Stuff)
	fmt.Println("Brands: ", currentPlayerStats.Brands)
	fmt.Println("Knowledge: ", currentPlayerStats.Knowledge)
	fmt.Println("Experience: ", currentPlayerStats.Experience)
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
	cps := GetPlayerStats()
	cps.Brands = DeltaStat(cps.Brands, nps.Brands)
	cps.Connections = DeltaStat(cps.Connections, nps.Connections)
	cps.Experience = DeltaStat(cps.Experience, nps.Experience)
	cps.Knowledge = DeltaStat(cps.Knowledge, nps.Knowledge)
	cps.Stuff = DeltaStat(cps.Stuff, nps.Stuff)
	if nps.Brands != 0 {
		fmt.Println("Brands:", cps.Brands)
	}
	if nps.Connections != 0 {
		fmt.Println("Connections:", cps.Connections)
	}
	if nps.Experience != 0 {
		fmt.Println("Experience:", cps.Experience)
	}
	if nps.Knowledge != 0 {
		fmt.Println("Knowledge:", cps.Knowledge)
	}
	if nps.Stuff != 0 {
		fmt.Println("Stuff:", cps.Stuff)
	}
	UpdatePlayerStats(cps)
}

//TraitChange implements deltaTraits
//traits are set, not cumulative
func TraitChange(nps PlayerTraits) {
	cps := GetPlayerTraits()

	if nps.Alcoholism != 0 {
		cps.Alcoholism = nps.Alcoholism
		fmt.Println("Alcoholism:", cps.Alcoholism)
	}
	if nps.Beauty != 0 {
		cps.Beauty = nps.Beauty
		fmt.Println("Beauty:", cps.Beauty)
	}
	if nps.Education != 0 {
		cps.Education = nps.Education
		fmt.Println("Education:", cps.Education)
	}
	if nps.Income != 0 {
		cps.Income = nps.Income
		fmt.Println("Income:", cps.Income)
	}
	if nps.Independence != 0 {
		cps.Independence = nps.Independence
		fmt.Println("Independence:", cps.Independence)
	}
	if nps.Lucky != 0 {
		cps.Lucky = nps.Lucky
		fmt.Println("Lucky:", cps.Lucky)
	}
	if nps.Smooth != 0 {
		cps.Smooth = nps.Smooth
		fmt.Println("Smooth:", cps.Smooth)
	}
	UpdatePlayerTraits(cps)
}

//MeterChange implements deltaMeters
func MeterChange(nps PlayerMeters) {
	cps := GetPlayerMeters()
	//fmt.Println("Testing meter change 1:", nps.Energy)
	cps.Energy = DeltaStat(cps.Energy, nps.Energy)
	cps.Hygiene = DeltaStat(cps.Hygiene, nps.Hygiene)
	cps.Soul = DeltaStat(cps.Soul, nps.Soul)
	cps.Stress = DeltaStat(cps.Stress, nps.Stress)

	if nps.Energy != 0 {
		fmt.Println("Energy:", cps.Energy)
	}
	if nps.Hygiene != 0 {
		fmt.Println("Hygiene:", cps.Hygiene)
	}
	if nps.Soul != 0 {
		fmt.Println("Soul:", cps.Soul)
	}
	if nps.Stress != 0 {
		fmt.Println("Stress:", cps.Stress)
	}
	//fmt.Println("Testing meter change 2:", cps.Energy)
	UpdatePlayerMeters(cps)
}

func MinMeterCheck(minm PlayerMeters) bool {
	cpm := GetPlayerMeters()
	cont := true
	if cpm.Energy <= minm.Energy {
		cont = false
		fmt.Println("You don't have the energy.")
		//fmt.Println("Test: action meters,player meters", minm.Energy, cpm.Energy)
	}
	return cont
}

func MinTraitCheck(mint PlayerTraits) bool {
	cpt := GetPlayerTraits()
	cont := true
	if cpt.Income <= mint.Income {
		cont = false
		fmt.Println("You cannot afford this place.")
		//fmt.Println("Test: action meters,player meters", minm.Energy, cpm.Energy)
	}
	return cont
}
