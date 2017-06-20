package letter

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
	res := FreqMap{}
	semChan := make(chan FreqMap, len(text))

	for i := 0; i < len(text); i++ {
		go func(i int) {
			semChan <- Frequency(text[i])
		}(i)
	}

	for i := 0; i < len(text); i++ {
		for k, v := range <-semChan {
			res[k] += v
		}
	}

	return res
}
