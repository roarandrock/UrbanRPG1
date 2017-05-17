package models

import "fmt"

/*
to keep and modify schedule

Weekday, weekend
Morning, Daytime, Evening, Night, Midnight, Dawn

Both split into arrays
Week = 0 weekday, 1 weekend
Day = 0 Morning- 6 dawn
Weekday, weekend
Morning, Daytime, Evening, Night, Midnight, Dawn



freetime = 1
worktime = 2
busy = 0

40 hour work week:
0,2,1,1,1,1,1,1,1,1,1,1

*/

var (
	wd0 = 1 //weekday morning
	wd1 = 1
	wd2 = 1
	wd3 = 1
	wd4 = 1
	wd5 = 1
	wd6 = 1
	we0 = 1
	we1 = 1
	we2 = 1
	we3 = 1
	we4 = 1
	we5 = 1
	we6 = 1
)

//SetSchedule takes in new schedule
func SetSchedule(newsc []int) {
	//can make a check then return error
	if len(newsc) != 14 {
		fmt.Println("Schedule error, wrong size", len(newsc))
		return
	}
	wd0 = newsc[0]
	wd1 = newsc[1]
	wd2 = newsc[2]
	wd3 = newsc[3]
	wd4 = newsc[4]
	wd5 = newsc[5]
	wd6 = newsc[6] //separate weekday and weekend schedule?
	we0 = newsc[7]
	we1 = newsc[8]
	we2 = newsc[9]
	we3 = newsc[10]
	we4 = newsc[11]
	we5 = newsc[12]
	we6 = newsc[13]
}

//GetSched returns current schedule
func GetSched() []int {
	exports := []int{wd0, wd1, wd2, wd3, wd4, wd5, wd6,
		we0, we1, we2, we3, we4, we5, we6}
	return exports
}
