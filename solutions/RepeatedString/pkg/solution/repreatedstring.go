package solution

// Complete the RepeatedString function below.
// Time complexity: o(1), O(len(s)).
// Space complexity: θ(1).
func RepeatedString(s string, n int64) int64 {
	if len(s) == 0 || n == 0 {
		return 0
	}

	var aCounterInS, aCounterInStump int64
	for i, c := range s {
		if c == 'a' {
			aCounterInS++
			if int64(i) < (n % int64(len(s))) {
				aCounterInStump++
			}
		}
	}
	return aCounterInS*(n/int64(len(s))) + aCounterInStump
}
