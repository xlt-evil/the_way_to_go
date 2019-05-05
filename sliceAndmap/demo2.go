package main

import "fmt"

func main() {
	m, k := mapinit()
	r := RangeMap(m, k)
	fmt.Println(r)
}

func mapinit() (map[int]int, []int) {
	m := make(map[int]int)
	key := make([]int, 0)
	for i := 0; i < 10; i++ {
		m[i] = i
		key = append(key, i)
	}
	return m, key
}

func RangeMap(m map[int]int, key []int) (result []int) {
	length := len(key)
	for i := 0; i < length; i++ {
		result = append(result, m[key[i]])
	}
	return result
}
