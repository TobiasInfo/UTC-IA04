package main

import (
	"fmt"
	"sync"
)

func Fill(tab []int, v int) {
	for i := range tab {
		tab[i] = v
	}
}

func FillParallel(tab []int, v int, numGoroutines int) {
	// Divide in smaller sections called chunks
    chunkSize := (len(tab) + numGoroutines - 1) / numGoroutines
    var wg sync.WaitGroup

    for i := 0; i < numGoroutines; i++ {
        start := i * chunkSize
        end := (i + 1) * chunkSize
        if end > len(tab) {
            end = len(tab)
        }
        wg.Add(1)
        go func(start, end int) {
            defer wg.Done()
            for j := start; j < end; j++ {
                tab[j] = v
            }
        }(start, end)
    }

    wg.Wait()
}




// applique f sur chaque élément de tab et remplace la valeur initiale par le résultat de f
func ForEach(tab []int, f func(int) int) {
	for i := range tab {
		tab[i] = f(tab[i])
	}
}

func ForEachParallel(tab []int, f func(int) int, numGoroutines int) {
    chunkSize := (len(tab) + numGoroutines - 1) / numGoroutines
    var wg sync.WaitGroup

    for i := 0; i < numGoroutines; i++ {
        start := i * chunkSize
        end := (i + 1) * chunkSize
        if end > len(tab) {
            end = len(tab)
        }
        wg.Add(1)
        go func(start, end int) {
            defer wg.Done()
            for j := start; j < end; j++ {
                tab[j] = f(tab[j])
            }
        }(start, end)
    }

    wg.Wait()
}


// copy le tableau src dans dest
func Copy(src []int, dest []int) {
	if len(src) != len(dest){
		panic("Les deux tableaux n'ont pas la même taille")
	}
	for i,val := range src{
		dest[i] = val
	}
}

func CopyParallel(src []int, dest []int, numGoroutines int) {
    if len(src) != len(dest) {
        panic("Source and destination slices must be of the same length")
    }
    chunkSize := (len(src) + numGoroutines - 1) / numGoroutines
    var wg sync.WaitGroup

    for i := 0; i < numGoroutines; i++ {
        start := i * chunkSize
        end := (i + 1) * chunkSize
        if end > len(src) {
            end = len(src)
        }
        wg.Add(1)
        go func(start, end int) {
            defer wg.Done()
            for j := start; j < end; j++ {
                dest[j] = src[j]
            }
        }(start, end)
    }

    wg.Wait()
}


// vérifie que tab1 et tab2 sont identiques
func Equal(tab1 []int, tab2 []int) bool {
	if len(tab1) != len(tab2){
		return false
	}
	for i := range tab1{
		if tab1[i] != tab2[i]{
			return false
		}
	}
	return true
}

func EqualParallel(tab1 []int, tab2 []int, numGoroutines int) bool {
    if len(tab1) != len(tab2) {
        return false
    }
    chunkSize := (len(tab1) + numGoroutines - 1) / numGoroutines
    var wg sync.WaitGroup
    result := true
    var mu sync.Mutex

    for i := 0; i < numGoroutines; i++ {
        start := i * chunkSize
        end := (i + 1) * chunkSize
        if end > len(tab1) {
            end = len(tab1)
        }
        wg.Add(1)
        go func(start, end int) {
            defer wg.Done()
            localResult := true
            for j := start; j < end; j++ {
                if tab1[j] != tab2[j] {
                    localResult = false
                }
            }
            mu.Lock()
            result = result && localResult
            mu.Unlock()
        }(start, end)
    }

    wg.Wait()
    return result
}


func add1(i int) int {
	return i + 1
}

func main() {

	table := [...]int{0, 0}
	var table2 [20]int
	Fill(table[:], 5)
	Fill(table2[:], 33)
	fmt.Println(table)
	fmt.Println(table2)
	ForEach(table2[:], add1)
	fmt.Println(table2)
	//	fmt.Println(ForEach(
	//		table2[:],
	//		func (i int) int{
	//			return i + 1
	//		},
	//	),
	//
	// )
}
