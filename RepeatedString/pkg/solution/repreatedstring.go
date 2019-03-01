package solution

// edge cases: len(s) > n, len(s) = 0, n = 0

// Complete the repeatedString function below.
// Time complexity: o(1), O(len(s)).
// Space complexity: θ(1).
func RepeatedString(s string, n int64) int64 {
	if len(s) == 0 || n == 0 {
		return 0
	}

	var aCounterInWholeRepetitions, aCounterInLeftover int64
	for i, c := range s {
		if c == 'a' {
			aCounterInWholeRepetitions++
			if int64(i) < (n % int64(len(s))) {
				aCounterInLeftover++
			}
		}
	}
	return aCounterInWholeRepetitions*(n/int64(len(s))) + aCounterInLeftover
}
