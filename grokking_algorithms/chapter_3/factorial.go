package main

func fact(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * fact(x-1)
	}
}

func main() {
	println(fact(4)) // 24
}
