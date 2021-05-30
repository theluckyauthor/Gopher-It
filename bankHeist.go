package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heist() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	openedVault := rand.Intn(100)
	if openedVault >= 70 && isHeistOn {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened. ")
	}
	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("You were discovered")
		case 1:
			isHeistOn = false
			fmt.Println("Bad ending")
		case 2:
			isHeistOn = false
			fmt.Println("wee woo wee woo")
		case 3:
			isHeistOn = false
			fmt.Println("You had to follow the train CJ")
		default:
			fmt.Println("Start the getaway car!")
		}
	}
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("We stole $", amtStolen)
	}
	fmt.Println(isHeistOn)

}
