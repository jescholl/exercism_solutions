package diffsquares

func SquareOfSums(firstN int) int {
	sums := 0
	for i := 1; i <= firstN; i++ {
		sums += i
	}
	return (sums * sums)
}

func SumOfSquares(firstN int) (sumOfSquares int) {
	for i := 1; i <= firstN; i++ {
		sumOfSquares += i * i
	}
	return
}

func Difference(firstN int) int {
	return SquareOfSums(firstN) - SumOfSquares(firstN)
}
