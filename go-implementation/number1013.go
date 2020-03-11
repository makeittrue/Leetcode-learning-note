package main

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A{
		sum += v
	}
	if sum % 3 != 0{
		return false
	}
	sum = sum / 3
	preSum, total := 0,0
	for i := 0; i<len(A); i++{
		preSum += A[i]
		if preSum == sum {
			total++
			if total == 3{
				return true
			}
			preSum = 0
		}
	}
	return false
}