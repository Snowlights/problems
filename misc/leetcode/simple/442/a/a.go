package main

func maxContainers(n, w, maxWeight int) int {
	return min(maxWeight/w, n*n)
}
