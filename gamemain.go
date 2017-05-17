/*
Author: roarandrock
Company: Jupiter Engineering

Motto: "Better to produce shit than to be constipated"
*/

package main

import (
	"UrbanRPG1/check"
	"UrbanRPG1/flow"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Game start")
	log.Println(log.Ldate)
	err := flow.MainFlow()
	check.ErrorCheck(err)
}
