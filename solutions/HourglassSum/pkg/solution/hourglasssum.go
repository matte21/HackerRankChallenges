package solution

// Complete the HourglassSum function below.
// No point in giving asymptotic space and time complexities because the problem
// statement contstraints the size of the input to a constant value.
// Time complexity: 144 iterations.
// Space complexity: 2 int32 variables and 4 int indices.
// There's a more efficient solution wrt the provided one: not using the two
// innermost for loops and explicitly summing only the elements of the
// hourglass.
func HourglassSum(arr [][]int32) int32 {
	maxSum := int32(-64)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := int32(0)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if k != 1 || l == 1 {
						sum += arr[i+k][j+l]
					}
				}
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}
