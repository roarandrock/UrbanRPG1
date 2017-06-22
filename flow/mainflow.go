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
	models.SetPlayerAll()
	models.SetJorb()          //setting up job board
	models.PlayerJorbUpdate() //setting first job impact and schedule
	//setting up city
	models.SetCity()
	models.SetEvents()
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
	//models.GetCurJorb(), passes
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
	fmt.Println("Test schedule", csched)
	switch cw[1] {
	case 0: //weekdays
		actiontype = csched[cd[1]]
	case 1: //weekends
		actiontype = csched[cd[1]+6]
	}
	//work or free
	switch actiontype {
	case 0: //booked
		fmt.Println(models.GetenSetSchedEvent("Busy getting ready for work."))
		options = []string{"OK", "Bummer"}
	case 1: //free
		//Unemployed
		if check.UnemployedCheck() == true {
			action.UnemployedActions()
		}
		options, m1 = action.FreeTimeActions() //implement basic menu now
	case 2: //work
		options = action.WorkActions()
	}
	//stats events
	//options = append...
	//display options
	r1 := inputs.StringarrayInput(options)
	switch actiontype {
	case 1:
		action.FreeTimeOutcomes(options[r1-1], m1) //what to do with more options? Need to handle dynamic options
	case 2:
		action.WorkOutcomes(options[r1-1])
	}
	//location events
	if check.EventLocationCheck() == true {
		action.EventLocationAction()
	}
}

func schedulePlanner() { //not used currently
	fmt.Println("What would you like to do this week?")
	//daytimeoptions := []string{"Go to work", "Stay home", "Go on holiday"}
	//weeknightoptions := []string{"Stay home", "Go to gym"}
}

func endOfWeek() {
	cd, cw := models.GetTime()
	if cd[1] == 5 && cw[1] == 1 {
		fmt.Println("How are you feeling about your life choices?")
		//resets schedule to Job schedule. cheat, need an exception for Weekday morning?
		models.SetSchedule(models.GetCurrentPlayer().CurJorb.Schedule)
	}
	//look at traits/stats?
}

func endFlow() {
	fmt.Println("It's all over.")
}
