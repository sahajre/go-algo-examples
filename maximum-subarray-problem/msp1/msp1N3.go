package msp

func mspN3(n []int) (maxSum, mi, mj int) {
	l := len(n)
	if l == 0 {
		return -1, -1, -1
	}

	maxSum = n[0]
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += n[k]
				if sum > maxSum {
					maxSum = sum
					mi = i
					mj = k
				}
				if sum == maxSum {
					mj = k
				}
			}
		}
	}

	return maxSum, mi, mj
}
