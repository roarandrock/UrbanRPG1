package models

/*
Jorbs
*/

//Jorb is base struct
type Jorb struct {
	CompanyName string
	Industry    string
	Title       string
	Schedule    []int
	Minstats    Jorbmin
	Impact      Jorbimpact
}

//Jorbmin minimum stats needed
type Jorbmin struct {
	MinStats  PlayerStats
	MinTraits PlayerTraits
}

//Jorbimpact impact on stats from taking/working jorb
type Jorbimpact struct {
	ImpactStats  PlayerStats
	ImpactTraits PlayerTraits
	ImpactMeters PlayerMeters
}

//FirstJorb first job possible
var FirstJorb Jorb
var SecondJorb Jorb

//GetJorb returns current player struct
func GetJorb() Jorb {
	return FirstJorb
}

//SetJorb sets initial job stats
func SetJorb() {
	FirstJorb.CompanyName = "Fun First Jorb Corp"
	FirstJorb.Industry = "Software Testing"
	FirstJorb.Title = "Entry Level"
	FirstJorb.Schedule = []int{0, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1} //busy morning, working, then 1s

	//Zero for all mins on first jorb
	//ImpactMeters
	FirstJorb.Impact.ImpactMeters.Energy = -20
	FirstJorb.Impact.ImpactMeters.Hygiene = -10
	FirstJorb.Impact.ImpactMeters.Soul = -1
	FirstJorb.Impact.ImpactMeters.Stress = 20
	//ImpactStats
	FirstJorb.Impact.ImpactStats.Connections = 10
	FirstJorb.Impact.ImpactStats.Experience = 10
	FirstJorb.Impact.ImpactStats.Knowledge = 20
	//ImpactTraits
	FirstJorb.Impact.ImpactTraits.Income = 20
	FirstJorb.Impact.ImpactTraits.Independence = 40

	//second one
	SecondJorb.CompanyName = "Not so fun, slightly more serious Jorb Corp"
	SecondJorb.Industry = "Software Testing"
	SecondJorb.Title = "Professional Level"
	SecondJorb.Schedule = []int{0, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1} //busy morning, working, then 1s

	//MinStats
	SecondJorb.Minstats.MinStats.Connections = 10
	SecondJorb.Minstats.MinStats.Experience = 10
	SecondJorb.Minstats.MinStats.Knowledge = 40
	//MinTraits
	SecondJorb.Minstats.MinTraits.Education = 20
	SecondJorb.Minstats.MinTraits.Independence = 30
	SecondJorb.Minstats.MinTraits.Lucky = 10
	SecondJorb.Minstats.MinTraits.Smooth = 10
	//ImpactMeters
	SecondJorb.Impact.ImpactMeters.Energy = -30
	SecondJorb.Impact.ImpactMeters.Hygiene = -20
	SecondJorb.Impact.ImpactMeters.Soul = -5
	SecondJorb.Impact.ImpactMeters.Stress = 40
	//ImpactStats
	SecondJorb.Impact.ImpactStats.Connections = 10
	SecondJorb.Impact.ImpactStats.Experience = 10
	SecondJorb.Impact.ImpactStats.Knowledge = 30
	//ImpactTraits
	SecondJorb.Impact.ImpactTraits.Income = 40
	SecondJorb.Impact.ImpactTraits.Independence = 30
}
