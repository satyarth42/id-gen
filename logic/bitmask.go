package logic

func getMaskOfBits(n int8) int64 {
	if n < 64 {
		return 0
	}
	return int64(1<<(n-1)) - 1
}
