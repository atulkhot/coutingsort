package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	id            string
	num_purchases int
}

func Make_random_array(num_items, max int) []int {
	var rand_array []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num_items; i++ {
		rand_array = append(rand_array, rand.Intn(max))
	}

	return rand_array
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func MaxElem(v []int) int {
	m := v[0]
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func Print_array(arr []int, num_items int) {
	num_elems := Min(len(arr), num_items)
	fmt.Println(arr[:num_elems])
}

func main() {
	a := Make_random_array(10, 25)
	fmt.Println(a)

	maxElem := MaxElem(a)
	counts := make([]int, maxElem+1)

	for i := 0; i < len(a); i++ {
		elem := a[i]
		counts[elem]++
	}
	//fmt.Println(counts)

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	fmt.Print("Count = ")
	fmt.Println(counts)

	sorted_a := make([]int, len(a))

	for k := len(a) - 1; k >= 0; k-- {
		sorted_a[counts[a[k]]-1] = a[k]
		counts[a[k]]--
	}

	fmt.Println(sorted_a)
}
