package machine

import (
	"fmt"
	"strconv"
)

func DoSelectingProcess() string {
	fmt.Println("What are you going to buy today :)")
	fmt.Print("Please Enter Item Code > ")

	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
	}
	return input
}

func DoDepositingProcess() float64 {
	fmt.Print("Please Amount of Money > ")

	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
	}
	amount, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("invalid input. expected to get: 1.00, 2.5, ...")
		return 0
	}
	return amount
}
