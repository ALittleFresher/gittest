package main
import "fmt"
func main() {
	var num int
	fibo(num)

}

func fibo(res int) int {
	//var res int
	for i := 0; i < 10; i++ {
		res = i + res
		fmt.Printf("第%d次叠加的值：%d", i, res)
	}
	return 0
}
