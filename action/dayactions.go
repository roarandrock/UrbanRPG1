package action

import (
	"UrbanRPG1/inputs"
	"UrbanRPG1/model"
	"fmt"
)

//WorkActions chooses actions at work
func WorkActions() []string {
	cj := models.GetJorb()
	fmt.Println("You go to work at ", cj.CompanyName)
	wtlist := []workTimeAct{work, sleep1} //need a function for this
	options := []string{}
	for _, v := range wtlist {
		options = append(options, v.actionname)
	}
	return options
}

//WorkOutcomes is the impact from chosing the action
func WorkOutcomes(choption string) {
	setWorkAct() //move?
	wtlist := []workTimeAct{work, sleep1}
	wtchosen := workTimeAct{}
	for _, v := range wtlist {
		if v.actionname == choption {
			wtchosen = v
		}
	}

	//fmt.Println("Test Work Outcomes 1:", wtchosen)
	//fmt.Println("Test work outcomes 2:", models.GetJorb().Impact.ImpactMeters)

	models.MeterChange(wtchosen.deltameters)
	chance1 := Brng(100) //different function
	if chance1 <= wtchosen.firedchance {
		fmt.Println("You're fired!") //no impact, needs to change jorb to fired
	}
}

type workTimeAct struct {
	actionname  string
	firedchance int
	deltameters models.PlayerMeters //get from jorb definition, except for special actions
	//promotion chances, getting fired chances
}

var (
	work = workTimeAct{"Working for the machine.",
		0,
		models.PlayerMeters{}} //doesn't pull them in at start
	sleep1 = workTimeAct{"Getting paid to nap. The perfect crime.",
		80,
		models.PlayerMeters{Energy: 15}}
)

//need this
func setWorkAct() {
	work.deltameters = models.GetJorb().Impact.ImpactMeters
}

type freeTimeAct struct {
	actionname string
	desc       string //description
	deltatime  int
	loc        int //allowed location
	minstats   models.PlayerStats
	mintraits  models.PlayerTraits
	minmeters  models.PlayerMeters
	impstats   models.PlayerStats
	imptraits  models.PlayerTraits
	impmeters  models.PlayerMeters
}

var (
	nothing = freeTimeAct{"Nothing",
		"For when there's nothing better to do",
		1,
		0,
		models.PlayerStats{},
		models.PlayerTraits{},
		models.PlayerMeters{},
		models.PlayerStats{},
		models.PlayerTraits{},
		models.PlayerMeters{}}
	watchTV = freeTimeAct{"Watch TV",
		"Watching the latest in mass communication.",
		1,
		0, //home?
		models.PlayerStats{Stuff: 20}, //mins
		models.PlayerTraits{Independence: 10},
		models.PlayerMeters{Energy: 10},
		models.PlayerStats{Brands: 10}, //impact
		models.PlayerTraits{},
		models.PlayerMeters{Energy: 5, Stress: -10, Hygiene: -5}}
	sleep = freeTimeAct{"Sleep",
		"ZZZZZZZZZZZZ",
		2,
		0,
		models.PlayerStats{},                                  //mins
		models.PlayerTraits{Alcoholism: 90, Independence: 10}, //need a drink to sleep
		models.PlayerMeters{},
		models.PlayerStats{}, //impact
		models.PlayerTraits{Beauty: 10, Smooth: 5},
		models.PlayerMeters{Energy: 30, Stress: -15, Hygiene: 5, Soul: 1}}
	party = freeTimeAct{"Party", //too vague, and not something that can be done alone. Only good for testing
		"Party!",
		1,
		0,
		models.PlayerStats{}, //mins
		models.PlayerTraits{Independence: 10},
		models.PlayerMeters{Energy: 20, Stress: 80},
		models.PlayerStats{Connections: 5}, //impact
		models.PlayerTraits{Alcoholism: 10, Smooth: 5},
		models.PlayerMeters{Energy: -30, Stress: -10, Hygiene: -15, Soul: 1}}
)

//FreeTimeActions determines actions in freetime
func FreeTimeActions() ([]string, int) {
	//need to implement general menu options for free time
	m1 := freeMenu()
	//this gets the list of individual actions
	options := []string{}
	switch m1 {
	case 1: //meet people
		fpList := funPlaceList()
		for _, v := range fpList {
			options = append(options, v.Name)
		}
		fmt.Println("Where to meet people?")
	case 2: //Maintenance
		ftlist := freeActList()
		for _, v := range ftlist {
			options = append(options, v.actionname)
		}
	case 3: //Movement
		dList := models.TravelOptions(models.GetPlayerLoc())
		for _, v := range dList {
			options = append(options, v.Name)
		}
		fmt.Println("Where would you like to go?")
	}
	return options, m1
}

func freeMenu() int {
	options := []string{"Meet People", "Maintenance", "Movement"}
	r1 := inputs.StringarrayInput(options)
	return r1
}

//FreeTimeOutcomes takes in a string and shows the outcome
func FreeTimeOutcomes(choption string, m1 int) {
	switch m1 {
	case 1:
		cd := models.DistrictGetByLoc(models.GetPlayerLoc())
		//get place bsed on choption
		fp1 := models.FunPlaceGetByName(cd.Name, choption)
		if models.MinTraitCheck(fp1.MinTraits) == true {
			fmt.Println("You go to", fp1.Name, "to find people.")
		} else {
			fmt.Println("Test: You cannot afford", fp1.Name)
		}

	case 2:
		ftlist := freeActList()
		ftchosen := freeTimeAct{}
		for _, v := range ftlist {
			if v.actionname == choption {
				ftchosen = v
			}
		}
		//fmt.Println("Test 1:", ftchosen.imptraits)
		models.StatChange(ftchosen.impstats)
		models.TraitChange(ftchosen.imptraits)
		models.MeterChange(ftchosen.impmeters)
		if ftchosen.deltatime > 1 {
			models.GetenSetSchedEvent("Sleeping") //generic placeholder
			cd, cw := models.GetFutureTime(1)
			nSch := 0 //setting next segment to busy
			models.UpdateSchedPerTime(cd, cw, nSch)
		}
	case 3:
		nd := models.DistrictGetByName(choption)
		cp := models.GetCurrentPlayer()
		cp.Loc = nd.Nummer
		fmt.Println("You go to:", nd.Name)
		models.UpdateCurrentPlayer(cp)
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
	ftlist1 := []freeTimeAct{watchTV, sleep, party} //needs a function
	//need a default action for when nothing else works
	ftlist2 := []freeTimeAct{}
	//check time, location, stats
	for _, v := range ftlist1 {
		if models.MinMeterCheck(v.minmeters) == true {
			ftlist2 = append(ftlist2, v)
		} else {
			fmt.Println("Test: You cannot", v.actionname)
		} //can do the same for traits and stats if this works
	}
	return ftlist2
}

func funPlaceList() []models.FunPlace {
	fpList1 := []models.FunPlace{}
	cd := models.DistrictGetByLoc(models.GetPlayerLoc())
	fpList1 = append(fpList1, cd.Place1) //cheating
	return fpList1
}
