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
	//impact on energy and soul too, need a variable impact on stats,
	Minstats Jorbmin
	Impact   Jorbimpact
}

//Jorbmin minimum stats needed
type Jorbmin struct {
	Energy       int
	Connections  int
	Independence int
	Knowledge    int
	Experience   int
	Beauty       int
}

//Jorbimpact impact on stats from taking/working jorb
type Jorbimpact struct {
	Energy       int
	Connections  int
	Independence int
	Knowledge    int
	Experience   int
	Smooth       int
	Stress       int
	Income       int
}

//FirstJorb first job possible
var FirstJorb Jorb

//GetJorb returns current player struct
func GetJorb() Jorb {
	return FirstJorb
}

//SetJorb sets initial job stats
func SetJorb() {
	FirstJorb.CompanyName = "Fun First Jorb Corp"
	FirstJorb.Industry = "Software Testing"
	FirstJorb.Title = "Entry Level"
	FirstJorb.Schedule = []int{0, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	FirstJorb.Impact.Energy = -20
	FirstJorb.Impact.Income = 20
	FirstJorb.Impact.Stress = 10
	FirstJorb.Impact.Knowledge = 10
	FirstJorb.Impact.Connections = 10
	FirstJorb.Impact.Independence = -20
}
