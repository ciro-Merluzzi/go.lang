package main

import (
	"fmt"
)

func main() {

	flag := true
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			flag = false
			break
		} else if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(flag)

	day := "Sex"
	switch day {
	case "Sex", "KKK":
		fmt.Println("É HOOJEE!\n BORAAAA")
	case "Seg", "Ter", "XXX":
		fmt.Println("meeee")
	default:
		fmt.Println("padrão")
	}

	switch {
	case day == "Sex":
		fmt.Println("É HOOJEE!\n BORAAAA")
		break
	}

}
