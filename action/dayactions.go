package action

import (
	"UrbanRPG1/inputs"
	"UrbanRPG1/model"
	"fmt"
)

//WorkActions chooses actions at work
func WorkActions() {
	cj := models.GetJorb()
	fmt.Println("You go to work at ", cj.CompanyName)
	//time passes
	//_, cpstats, _ := models.GetCurrentPlayer()

	wtlist := []workTimeAct{work, sleep1}
	options := []string{}
	for _, v := range wtlist {
		options = append(options, v.actionname)
	}
	r1 := inputs.StringarrayInput(options)
	wtchosen := wtlist[r1-1]
	models.StatChange(wtchosen.deltastats)
	chance1 := Brng(100)
	if chance1 <= wtchosen.firedchance {
		fmt.Println("You're fired!") //no impact, needs to change jorb to fired
	}
}

type workTimeAct struct {
	actionname  string
	firedchance int
	deltastats  models.PlayerStats
	//promotion chances, getting fired chances
}

var (
	work = workTimeAct{"Working for the machine.",
		0,
		models.PlayerStats{Energy: -2, Stress: -2, Independence: -1}}
	sleep1 = workTimeAct{"Getting paid to nap. The perfect crime.",
		80,
		models.PlayerStats{Energy: 1, Stress: -1}}
)

type freeTimeAct struct {
	actionname string
	deltastats models.PlayerStats
	deltatime  int
	minstats   models.PlayerStats
	//also need barriers - independence, income, location, could use minstats?
	//description
}

var (
	watchTV = freeTimeAct{"Watch TV",
		models.PlayerStats{Brands: 1, Energy: 1, Stress: -1},
		1,
		models.PlayerStats{}} //11 stats
	sleep = freeTimeAct{"Sleep",
		models.PlayerStats{Energy: 2, Stress: -1, Beauty: 1},
		2,
		models.PlayerStats{}}
	party = freeTimeAct{"Party!",
		models.PlayerStats{Energy: -1, Stress: -1, Connections: 1},
		1,
		models.PlayerStats{Energy: 1}}
)

//FreeTimeActions determines actions in freetime
func FreeTimeActions() {
	//_, cw := models.GetTime()
	//models.ShowTime()
	//cps := models.GetPlayerStats()

	ftlist := freeActList()

	options := []string{}
	for _, v := range ftlist {
		options = append(options, v.actionname)
	}
	r1 := inputs.StringarrayInput(options)
	ftchosen := ftlist[r1-1]
	models.StatChange(ftchosen.deltastats)
	if ftchosen.deltatime > 1 { //assumes it's only 2
		models.UpdateTimeofDay()
	}
	/*
		cps.Beauty = cps.Beauty + ftchosen.deltastats.Beauty
		cps.Brands = cps.Brands + ftchosen.deltastats.Brands
		cps.Connections = cps.Connections + ftchosen.deltastats.Connections
		cps.Energy = cps.Energy + ftchosen.deltastats.Energy
		cps.Experience = cps.Experience + ftchosen.deltastats.Experience
		cps.Income = cps.Income + ftchosen.deltastats.Income
		cps.Independence = cps.Independence + ftchosen.deltastats.Independence
		cps.Knowledge = cps.Knowledge + ftchosen.deltastats.Knowledge
		cps.Smooth = cps.Smooth + ftchosen.deltastats.Knowledge
		cps.Stress = cps.Stress + ftchosen.deltastats.Stress
		cps.Stuff = cps.Stuff + ftchosen.deltastats.Stuff
		models.UpdatePlayerStats(cps) */
}

//FreeActList makes a list of valid actions
func freeActList() []freeTimeAct {
	//get all possible actions
	ftlist1 := []freeTimeAct{watchTV, sleep, party} //needs a function that includes a stat compare
	ftlist2 := []freeTimeAct{}
	//check time, location, stats
	cps := models.GetPlayerStats()
	for _, v := range ftlist1 {
		if v.minstats.Energy <= cps.Energy { //only does this for energy for now
			ftlist2 = append(ftlist2, v)
		}
	}
	return ftlist2
}
