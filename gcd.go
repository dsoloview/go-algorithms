package main

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	if a >= b {
		return gcd(a%b, b)
	}

	if b >= a {
		return gcd(a, b%a)
	}

	return a
}
