package chapter1

// 求最大公约数
// 欧几里得算法，也叫辗转相除法
func gcd(x, y int) int {
	if x == 0 {
		return x
	}
	if y == 0 {
		return y
	}
	k := x % y
	if k == 0 {
		if x < y {
			return x
		}
		return y
	}
	if x < y {
		return gcd(x, k)
	}
	return gcd(y, k)
}
