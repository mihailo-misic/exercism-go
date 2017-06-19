package letter

import "fmt"

const testVersion = 1

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(text []string) FreqMap {
	c1 := make(chan FreqMap)
	c2 := make(chan FreqMap)
	c3 := make(chan FreqMap)
	for i := 0; i < len(text); i++ {
		go Frequency(text[i])
	}
	n1 := <-c1
	n2 := <-c2
	n3 := <-c3
	fmt.Println(n1, n2, n3)
	return n1 + n2 + n3
}
