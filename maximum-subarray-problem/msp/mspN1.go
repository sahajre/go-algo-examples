package msp

func mspN1(n []int) (maxSum, mi, mj int) {
	l := len(n)
	if l == 0 {
		return -1, -1, -1
	}

	maxSum = n[0]
	sum := n[0]
	i := 0
	for j := 1; j < l; j++ {
		sum += n[j]
		if sum > maxSum {
			maxSum = sum
			mi = i
			mj = j
		} else if sum < 0 {
			i = j + 1
			sum = 0
		}
	}

	return maxSum, mi, mj
}
