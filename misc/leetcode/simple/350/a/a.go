package main

func distanceTraveled(mainTank int, additionalTank int) (ans int) {

	for mainTank >= 5 {
		run := mainTank / 5
		ans += 5 * 10 * run
		mainTank -= 5 * run
		add := min(run, additionalTank)
		additionalTank -= add
		mainTank += add
	}

	return ans + mainTank*10
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
