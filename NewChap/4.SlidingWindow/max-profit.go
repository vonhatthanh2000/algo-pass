package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	l, r := 0, 1
	maxProfit := 0

	for r < len(prices) {

		if prices[l] <= prices[r] {
			profit := prices[r] - prices[l]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			l = r
		}
		r += 1

	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	maxP := 0
	minBuy := math.MaxInt32
	fmt.Println(minBuy)

	for _, sell := range prices {
		if sell-minBuy > maxP {
			maxP = sell - minBuy
		}
		if sell < minBuy {
			minBuy = sell
		}
	}
	return maxP
}

func main() {

	a := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(a))
	fmt.Println(maxProfit2(a))

}
