package models

import "fmt"

/*
to track days

Weekday, weekend
Morning, Daytime, Evening, Night, Midnight, Dawn

Both split into arrays
Week = 0 weekday, 1 weekend
Day = 0 Morning- 6 dawn
*/

var currentweek = []int{1, 0}
var currentday = []int{1, 0}

//GetTime returns time
func GetTime() ([]int, []int) {
	return currentday, currentweek
}

//UpdateTimeofDay moves the day along
func UpdateTimeofDay() {
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
}

//SetTime sets the time to the time provided
func SetTime(ct []int, cw []int) {
	currentday = ct
	currentweek = cw
}

//ShowTime displays the current time in human text
func ShowTime() {
	switch currentweek[1] {
	case 0:
		fmt.Println("Weekday")
	case 1:
		fmt.Println("Weekend")
	}
	switch currentday[1] {
	case 0:
		fmt.Println("Morning")
	case 1:
		fmt.Println("Daytime")
	case 2:
		fmt.Println("Evening")
	case 3:
		fmt.Println("Night")
	case 4:
		fmt.Println("Midnight")
	case 5:
		fmt.Println("Dawn")
	}
}
