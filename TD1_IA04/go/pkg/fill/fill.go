package fill

import (
	"math/rand"
	"sort"
)

// Fill the slice with the random numbers
func Fill(sl []int) []int {
	for i := range sl {
		sl[i] = rand.Intn(100)
	}
	return sl
}

func Moyenne(sl []int) float64 {
	var sum int
	for _, v := range sl {
		sum += v
	}
	return float64(sum) / float64(len(sl))
}

func ValeursCentrales(sl []int) []int {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	if len(sl)%2 == 0 {
		return []int{sl[len(sl)/2], sl[len(sl)/2-1]}
	} else {
		return []int{sl[len(sl)/2]}
	}
}

func Plus1(sl []int) []int {
	for i := range sl {
		sl[i]++
	}
	return sl
}
