package check

import (
	"UrbanRPG1/model"
	"fmt"
)

//ErrorCheck for checking errors
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}
}

//GameEndCheck checks for end conditions
func GameEndCheck() bool {
	cont := true
	_, cw := models.GetTime()
	if cw[0] >= 3 {
		cont = false
	}
	return cont
}

//StatZeroCheck checks if a stat is zero
func StatZeroCheck() bool {
	cont := true
	nps := models.GetPlayerStats()
	if nps.Beauty == 0 {
		fmt.Println("You have no beauty.")
	}
	if nps.Brands == 0 {
		fmt.Println("You have no brands.")
	}
	if nps.Connections == 0 {
		fmt.Println("You have no connections.")
	}
	if nps.Energy == 0 {
		fmt.Println("You have no energy.")
	}
	if nps.Experience == 0 {
		fmt.Println("You have no experience.")
	}
	if nps.Income == 0 {
		fmt.Println("You have no income.")
	}
	if nps.Independence == 0 {
		fmt.Println("You have no independence.")
	}
	if nps.Knowledge == 0 {
		fmt.Println("You have no knowledge")
	}
	if nps.Smooth == 0 {
		fmt.Println("You have no smooth.")
	}
	if nps.Stress == 0 {
		fmt.Println("You have no stress.")
	}
	if nps.Stuff == 0 {
		fmt.Println("You have no stuff.")
	}
	return cont
}
