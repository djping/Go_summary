package main
import "fmt"
func main () {
	// var n [5] int 
	var x int = 5
  var n  = [5]int{1,3,4,5,7}
	var y  = [...]string{"second", "article", "in", "juejin"}
	fmt.Println("n: ", n, x)
	fmt.Println("y: ", y)
	fmt.Println("数组的2中遍历方法")
	for i := 0; i < len(n); i++ {
		fmt.Printf("n[%d]=%d\n", i, n[i])
	}
	for index, value := range y {
		fmt.Printf("y[%d]=%s\n", index, value)
	}
	validArr(n)
	fmt.Println("这是原来的数组n: \n", n)

	// var nn = [3]int {1,3,4}
	fmt.Println("下面是切片的讨论\n")
	n0 := n[:]
	fmt.Printf("n0: %d\n", n0)
	n1 := n[1:2]
	validSlice(n1)
	fmt.Printf("通过函数改变切片，现在的n是: %v\n", n)
	fmt.Printf("n0: %d\n", n0)
}
func validArr (arr [5]int) {
	arr[0] = 100
	fmt.Println("这是函数内改变后的数组: \n", arr)
}
func validSlice (arr []int) {
	arr[0] = 100
}