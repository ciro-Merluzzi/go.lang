package main

import (
	"fmt"
)

func main() {
//	fmt.Println("Olá mundo !")

	f := true
//	flag := &f

//	if flag == nil && 2 == 3 {
		fmt.Println("Value is nil")
//	} else if *flag {
		fmt.Println("True")
//	} else {
		fmt.Println("False")
//	}

	arr := []string{"cílios", "merlindo", "vicente"}
	for i, Value := range arr {
		fmt.Println(i)
		fmt.Println(Value)
	}

	mymap := make(map[string]interface{})
	mymap["Nome"] = "Cílios"
	mymap["Idade"] = 21

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v", k, v)
	}

}
