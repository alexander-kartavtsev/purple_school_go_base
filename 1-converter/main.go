package main

import "fmt"

func main() {
	const USDEUR = 0.95
	const USDRUB = 110.5
	const EURRUB = USDRUB / USDEUR

	fmt.Println(USDEUR, USDRUB, EURRUB)
}
