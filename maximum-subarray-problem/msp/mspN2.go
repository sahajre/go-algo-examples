package msp

func mspN2(n []int) (maxSum, mi, mj int) {
	l := len(n)
	if l == 0 {
		return -1, -1, -1
	}

	maxSum = n[0]
	for i := 0; i < l; i++ {
		sum := 0
		for k := i; k < l; k++ {
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

	return maxSum, mi, mj
}
