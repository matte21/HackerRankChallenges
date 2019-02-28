package solution

// Complete the JumpingOnClouds function below.
// Time complexity: θ(len(clouds)).
// Space complexity: θ(1).
func JumpingOnClouds(clouds []int32) int32 {
	var jumps int32
	for i := 0; i < len(clouds)-1; {
		if i+2 < len(clouds) && clouds[i+2] == 0 {
			i += 2
		} else {
			i++
		}
		jumps++
	}
	return jumps
}
