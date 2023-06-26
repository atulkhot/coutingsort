package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Customer struct {
	id            string
	num_purchases int
}

func make_random_slice(num_items, max int) []Customer {
	var rand_array []Customer

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num_items; i++ {
		customer := Customer{}
		customer.id = "C" + strconv.Itoa(i)
		customer.num_purchases = rand.Intn(max)
		rand_array = append(rand_array, customer)
	}

	return rand_array
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max_elem(v []Customer) int {
	m := 0
	for i, e := range v {
		if i == 0 || e.num_purchases > m {
			m = e.num_purchases
		}
	}
	return m
}

func print_slice(arr []Customer, num_items int) {
	num_elems := Min(len(arr), num_items)
	fmt.Println(arr[:num_elems])
}

func check_sorted(arr []Customer) bool {
	prev := Customer{}
	for i, cust := range arr {
		if i == 0 {
			prev = cust
		} else {
			if prev.num_purchases > cust.num_purchases {
				return false
			}
		}
	}
	return true
}

func counting_sort(a []Customer, items int) []Customer {
	counts := make([]int, max_elem(a)+1)

	for i := 0; i < len(a); i++ {
		elem := a[i].num_purchases
		counts[elem]++
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	sorted_a := make([]Customer, len(a))

	for k := len(a) - 1; k >= 0; k-- {
		sorted_a[counts[a[k].num_purchases]-1] = a[k]
		counts[a[k].num_purchases]--
	}

	return sorted_a
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	values := make_random_slice(num_items, max)
	print_slice(values, 40)
	fmt.Println()

	//// Sort and display the result.
	sorted := counting_sort(values, max)
	print_slice(sorted, 40)
	//
	//// Verify that it's sorted.
	//check_sorted(sorted)
}
