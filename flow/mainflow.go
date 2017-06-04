package flow

import (
	"UrbanRPG1/action"
	"UrbanRPG1/check"
	"UrbanRPG1/inputs"
	"UrbanRPG1/model"
	"UrbanRPG1/narrative"
	"fmt"
	"strconv"
)

//MainFlow runs the game loop
func MainFlow() error {
	fmt.Println("Welcome to the jungle")
	cont, err := starterFlow()
	check.ErrorCheck(err)
	for cont == true {
		//game loop
		_, err = gameLoopFlow()
		//check to continue
		cont = check.GameEndCheck()
	}
	endFlow()
	return err
}

func starterFlow() (bool, error) {
	var cont bool
	_, err := strconv.Atoi("-42") //error cheat
	fmt.Println("Do you want to play a game?")
	r1 := inputs.StringarrayInput([]string{"Yes", "No"})
	if r1 == 1 {
		cont = true
	} else {
		cont = false
		return cont, err
	}
	narrative.Intro()
	models.SetJorb() //setting up job board
	models.SetPlayerAll()
	//setting up city
	models.SetCity()
	//setting up player based on jorb, should be a separate function
	//traits
	cpt := models.GetPlayerTraits()
	cpt.Income = models.GetJorb().Impact.ImpactTraits.Income
	cpt.Independence = models.GetJorb().Impact.ImpactTraits.Independence
	models.TraitChange(cpt)
	//Stats
	models.StatChange(models.GetJorb().Impact.ImpactStats)
	//Meters are on a daily basis
	//update all
	//setting schedule for job, needs to be implemented when changing jobs
	models.SetSchedule(models.GetJorb().Schedule)

	return cont, err
}

func gameLoopFlow() (bool, error) {

	//normal day loop
	dayFlowLoop()
	_, err := strconv.Atoi("-42") //error cheat
	cont := check.GameEndCheck()
	return cont, err
}

func dayFlowLoop() {
	models.ShowTime() //test
	//job check, schedule check, actions at time check
	//cpm := models.GetPlayerMeters()
	//cpt := models.GetPlayerTraits()
	//fmt.Println("Test energy", cpm.Energy)
	//fmt.Println("Test independence", cpt.Independence, cpt)
	//models.GetJorb(), passes
	//new way
	getChoices()
	check.StatZeroCheck()
	models.UpdateTimeofDay() //moves forward fixed amount
	endOfWeek()
}

func getChoices() {
	//check schedule/time events
	cd, cw := models.GetTime()
	csched := models.GetSched()
	actiontype := 1 //default, free time
	options := []string{}
	m1 := 1
	//fmt.Println("Test schedule", sched)
	switch cw[1] {
	case 0: //weekdays
		actiontype = csched[cd[1]]
	case 1: //weekends
		actiontype = csched[cd[1]+7]
	}
	//work or free
	switch actiontype {
	case 0: //booked
		fmt.Println(models.GetenSetSchedEvent("Busy getting ready for work."))
		options = []string{"OK", "Bummer"}
	case 1: //free
		options, m1 = action.FreeTimeActions() //implement basic menu now
	case 2: //work
		options = action.WorkActions()
	}
	//stats events
	//options = append...
	//location events
	if actiontype == 1 {
		cd := models.DistrictGetByLoc(models.GetPlayerLoc())
		fmt.Println("You are in:", cd.Name)
	}
	//display options
	r1 := inputs.StringarrayInput(options)
	switch actiontype {
	case 1:
		action.FreeTimeOutcomes(options[r1-1], m1) //what to do with more options? Need to handle dynamic options
	case 2:
		action.WorkOutcomes(options[r1-1])
	}
}

func schedulePlanner() { //not used currently
	fmt.Println("What would you like to do this week?")
	//daytimeoptions := []string{"Go to work", "Stay home", "Go on holiday"}
	//weeknightoptions := []string{"Stay home", "Go to gym"}
}

func endOfWeek() {
	cd, cw := models.GetTime()
	if cd[1] >= 5 && cw[1] == 1 {
		fmt.Println("How are you feeling about your life choices?")
	}
	//look at traits/stats?
}

func endFlow() {
	fmt.Println("It's all over.")
}
