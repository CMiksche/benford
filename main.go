package main

//nolint
type benford struct {
	dist        []float32
	chiSquared  float32
	probability float32
	terms       []float32
}

// firstDigit returns only the first digit of each number but no zeroes
func firstDigit(numbers []int) []int {
	var new []int

	for i := 0; i < len(numbers); i++ {
		n := numbers[i]
		// only the first digit
		for n >= 10 {
			n = n / 10
		}
		// No 0
		if n == 0 {
			continue
		}
		new = append(new, n)
	}

	return new
}

// Count the occurrences of one number
func countOccurrencesOfOne(numbers []int, searched int) int {
	var res int
	res = 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == searched {
			res = res + 1
		}
	}

	return res
}

// Count the occurrences of every number in a array
func countOccurrences(numbers []int) []int {
	var res []int
	var searched = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(searched); i++ {
		found := countOccurrencesOfOne(numbers, searched[i])
		res = append(res, found)
	}

	return res
}

// Get the distribution of 1 to 9 in the number array
//nolint
func countDistribution(occurrences []int) []float32 {
	return nil
}

// Get a struct with information about how well benfords law was matched
//nolint
func calcBenfords(numbers []int) benford {
	benfordNumbers := [9]float32{
		0.301, // 1
		0.176, // 2
		0.125, // 3
		0.097, // 4
		0.079, // 5
		0.067, // 6
		0.058, // 7
		0.051, // 8
		0.046, // 9
	}
	occurrences := countOccurrences(firstDigit(numbers))
	dist := countDistribution(occurrences)
	// TODO: calc actual chi2
	chiSquared := benfordNumbers[0]
	return benford{dist: dist, chiSquared: chiSquared}
}

func main() {

}
