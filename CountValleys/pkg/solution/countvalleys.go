package solution

// Complete the countingValleys function below.
// Time Complexity: θ(len(steps)).
// Space Complexity: θ(1).
func CountingValleys(steps string) int32 {
	var depth, valleys int32
	for _, aStep := range steps {
		if aStep == 'D' {
			depth--
		} else {
			depth++
			if depth == 0 {
				valleys++
			}
		}
	}
	return valleys
}
