package main

func numberOfEmployeesWhoMetTarget(hours []int, tar int) (ans int) {
	for _, v := range hours {
		if v >= tar {
			ans++
		}
	}
	return
}
