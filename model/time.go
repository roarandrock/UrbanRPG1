package models

import "fmt"

/*
to track days

Weekday, weekend
Morning, Daytime, Evening, Night, Midnight, Dawn

Both split into arrays
Week = 0 weekday, 1 weekend
Day = 0 Morning- 5 dawn
*/

var currentweek = []int{1, 0}
var currentday = []int{1, 0}
var exportweek = []int{0, 0}
var exportday = []int{0, 0}

func localTime() ([]int, []int) {
	exportday = currentday
	exportweek = currentweek
	return exportday, exportweek
}

//GetTime returns time
func GetTime() ([]int, []int) {
	cd, cw := localTime()
	return cd, cw
}

func GetFutureTime(n int) ([]int, []int) {
	cd := []int{0, 0}
	newd, cw := GetTime()
	cd[1] = newd[1] + n
	if cd[1] > 5 {
		cd[0]++
		cd[1] = 0
		cw[1]++
		if cw[1] > 1 {
			cw[0]++
			cw[1] = 0
		}
	}
	return cd, cw
}

//UpdateTimeofDay moves the day along
func UpdateTimeofDay() {
	fmt.Println("Test Time 1: Day, week", currentday, currentweek)
	currentday[1] = currentday[1] + 1
	if currentday[1] > 5 {
		currentday[0]++
		currentday[1] = 0
		currentweek[1]++
		if currentweek[1] > 1 {
			currentweek[0]++
			currentweek[1] = 0
		}
	}
	fmt.Println("Test Time 1: Day, week", currentday, currentweek)
}

//SetTime sets the time to the time provided
func SetTime(ct []int, cw []int) {
	currentday = ct
	currentweek = cw
}

//ShowTime displays the current time in human text
func ShowTime() {
	var s1 string
	switch currentday[1] {
	case 0:
		s1 = "Morning"
	case 1:
		s1 = "Daytime"
	case 2:
		s1 = "Evening"
	case 3:
		s1 = "Night"
	case 4:
		s1 = "Midnight"
	case 5:
		s1 = "Dawn"
	}
	switch currentweek[1] {
	case 0:
		//fmt.Println("Weekday")
		fmt.Printf("It is %s on a weekday.\n", s1)
	case 1:
		//fmt.Println("Weekend")
		fmt.Printf("It is %s on the weekend.\n", s1)
	}
}
