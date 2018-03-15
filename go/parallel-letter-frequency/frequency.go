package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(input []string) FreqMap {
	ch := make(chan FreqMap)

	for _, s := range input {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}
	m := <-ch
	for i := 1; i < len(input); i++ {
		for k, v := range <-ch {
			m[k] += v
		}
	}

	return m
}
