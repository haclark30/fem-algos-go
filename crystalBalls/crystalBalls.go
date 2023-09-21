package main

import (
	"math"
)

func CrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jumpAmount
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j <= jumpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}
	return -1
}

// func main() {
// 	ans := CrystalBalls([]bool{false, false, false, false, false, true, true, true})
// 	log.Println(ans)
// }
