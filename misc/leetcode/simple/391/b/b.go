package main

func maxBottlesDrunk(numBottles int, numExchange int) int {
	ans := numBottles

	for numBottles >= numExchange {
		numBottles++
		ans++
		numBottles -= numExchange
		numExchange++
	}

	return ans
}
