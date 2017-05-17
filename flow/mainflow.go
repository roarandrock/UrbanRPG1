package flow

import (
	"UrbanRPG1/action"
	"UrbanRPG1/check"
	"UrbanRPG1/inputs"
	"UrbanRPG1/model"
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
	fmt.Println("Do you want to play a game?")
	r1 := inputs.StringarrayInput([]string{"Yes", "No"})
	if r1 == 1 {
		cont = true
	} else {
		cont = false
	}
	models.SetJorb() //setting up job board
	models.SetPlayerAll()
	//setting up player based on jorb, should be a separate function
	_, cpStats, _ := models.GetCurrentPlayer()
	//fixed
	cpStats.Income = models.GetJorb().Impact.Income
	models.UpdatePlayerStats(cpStats)
//Cumulative //need to finish this 
cpStats.Energy =
cpStats.Connections  int
cpStats.Independence int
cpStats.Knowledge    int
cpStats.Experience   int
cpStats.Smooth       int
cpStats.Stress       int
cpStats.Income       int
//update all
models.StatChange(cpStats)
	//setting schedule for job, needs to be implemented when changing jobs
	s1 := models.GetJorb().Schedule
	models.SetSchedule(s1)
	_, err := strconv.Atoi("-42") //error cheat
	return cont, err
}

func gameLoopFlow() (bool, error) {
	//cont := true // pointless? Redudnant with game end check
	//normal day loop
	dayFlowLoop()
	_, err := strconv.Atoi("-42") //error cheat
	cont := check.GameEndCheck()
	return cont, err
}

func dayFlowLoop() {
	models.ShowTime() //test
	//job check, schedule check, actions at time check
	//models.GetJorb(), passes
	workSchedule() //change this to a time checker - what time, what is happenig?
	//models.ShowPlayerStats() //make a delta stat display?
	check.StatZeroCheck()
	models.UpdateTimeofDay() //based on schedule
	endOfWeek()
}

//WorkSchedule looks at job and time of day
func workSchedule() {
	cd, cw := models.GetTime()
	//fmt.Println("For test:", cd, cw)
	sched := models.GetSched()
	actiontype := 1 //default, free time
	//fmt.Println("Test schedule", sched)
	switch cw[1] {
	case 0: //weekdays
		actiontype = sched[cd[1]]
	case 1: //weekends
		actiontype = sched[cd[1]+7]
	}
	//fmt.Println("Test:", actiontype)
	switch actiontype {
	case 0:
		fmt.Println("You're busy.")
		//need to include the reason why
		fmt.Println("You have to get ready and go to work.")
		inputs.StringarrayInput([]string{"Bummer"})
	case 1:
		action.FreeTimeActions()
	case 2:
		action.WorkActions()
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
