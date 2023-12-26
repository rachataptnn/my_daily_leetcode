package main

import "fmt"

func main() {
	result := numRollsToTarget(2, 6, 7)
	fmt.Println(result)
}

func numRollsToTarget(diceAmt int, faceAmt int, target int) int {
	curWays := make([]int, target+1)
	prevWays := make([]int, target+1)
	prevWays[0] = 1 // if target is 0, there is 1 way to get it
	mod := int(1e9 + 7)

	for dc := 1; dc <= diceAmt; dc++ {
		for tg := 0; tg <= target; tg++ {
			curWays[tg] = 0
			for fc := 1; (fc <= faceAmt) && (fc <= tg); fc++ {
				curWaysForTarget := curWays[tg]
				prevWaysForReducedTarget := prevWays[tg-fc]
				curWays[tg] = (curWaysForTarget + prevWaysForReducedTarget) % mod

				fmt.Println("breakpoint placing")
			}
		}
		curWays, prevWays = prevWays, curWays

		fmt.Println("breakpoint placing")
	}

	return prevWays[target]
}
