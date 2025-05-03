package main

import (
	"fmt"
	"time"
)

func main() {
	if today := time.Now().Weekday(); today == time.Saturday || today == time.Saturday {
		fmt.Println("це вихідний")
	} else if today == time.Friday {
		fmt.Print("п'ятниця завтро суббота а после воскресенье")
	} else {
		fmt.Println("WORK DAY")
	}
}
