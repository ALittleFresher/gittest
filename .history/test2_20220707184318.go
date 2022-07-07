package main

import "fmt"

func main() {
	var num int
	fibo(num)

}

func fibo(res int) int {
	var j int = 1
	for i := 0; i < 10; i++ {
		mid := res
		res = res + j
		j = mid
		fmt.Printf("第%d次叠加的值：%d\n", i, res)
	}
	return 0
}
