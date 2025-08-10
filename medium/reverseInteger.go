package medium

func Reverse(x int) int {

	flag := 1
	if x < 0 {
		flag = -1
		x *= -1
	}
	newX := x
	cnt := 1
	for x > 9 {
		cnt *= 10
		x /= 10
	}
	res := 0
	for newX > 0 {
		res += cnt * (newX % 10)
		newX /= 10
		cnt /= 10
	}
	res *= flag
	if res > (1<<31)-1 || res < -(1<<31) {
		return 0
	}
	return res
}
