package solution

// Complete the SockMerchant function below.
// Time complexity: O(len(ar)).
// Space complexity: O(len(ar)).
func SockMerchant(ar []int32) int32 {
	counters := make(map[int32]uint8)
	for _, elem := range ar {
		counters[elem]++
	}
	var totalCount uint8
	for _, count := range counters {
		totalCount += count / 2
	}
	return int32(totalCount)
}
