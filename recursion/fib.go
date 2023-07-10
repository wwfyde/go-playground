package recursion

func Fib(n int) (res int) {
	if n <= 1 {
		return n
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

func FibList(n int) (res []int) {
	values := make([]int, n)
	for i := range values {
		res = append(res, Fib(i))
	}
	return
}

func FibList2(n int) []int {
	values := make([]int, n)
	for i := range values {
		values[i] = Fib(i)
	}
	return values
}

// TODO 使用并发
