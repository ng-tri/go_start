package practice

import "fmt"

func helloWorld() {
	fmt.Println("Hello, World!")
}

func variablesDemo() {
	var a string = "Xin chào"
	var b int = 10
	var c float64 = 3.14
	var d bool = true
	e := "Go rất tuyệt"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func operatorsDemo() {
	a := 10
	b := 3
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)
}

func conditionsDemo() {
	x := -1
	if x > 0 {
		fmt.Println("Số dương")
	} else if x == 0 {
		fmt.Println("Bằng 0")
	} else {
		fmt.Println("Số âm")
	}

	day := 3
	switch day {
	case 1:
		fmt.Println("Thứ Hai")
	case 2:
		fmt.Println("Thứ Ba")
	default:
		fmt.Println("Ngày khác")
	}
}

func loopsDemo() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Lần thứ:", i)
	}

	i := 0
	for i < 3 {
		fmt.Println("While style:", i)
		i++
	}
}

func functionsDemo() {
	fmt.Println("Xin chào", "Linh")
	fmt.Println("Tổng:", add(3, 4))
}

func add(a int, b int) int {
	return a + b
}

func arraysSlicesDemo() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Mảng:", arr)

	nums := []int{4, 5, 6}
	nums = append(nums, 7)
	fmt.Println("Slice:", nums)
}

func mapsDemo() {
	ages := map[string]int{
		"Linh": 30,
		"Nam":  25,
	}
	ages["An"] = 28
	fmt.Println("Tuổi của Linh:", ages["Linh"])
}

type Person struct {
	Ten  string
	Tuoi int
}

func structsDemo() {
	a := Person{Ten: "Linh", Tuoi: 30}
	fmt.Println("Tên:", a.Ten)
	fmt.Println("Tuổi:", a.Tuoi)
}

func pointersDemo() {
	a := 10
	increment(&a)
	fmt.Println("Sau khi tăng:", a)
}

func increment(x *int) {
	*x += 1
}

func main() {
	fmt.Println("=== Hello World ===")
	helloWorld()

	fmt.Println("\n=== Biến & Kiểu dữ liệu ===")
	variablesDemo()

	fmt.Println("\n=== Toán tử ===")
	operatorsDemo()

	fmt.Println("\n=== Cấu trúc điều kiện ===")
	conditionsDemo()

	fmt.Println("\n=== Vòng lặp ===")
	loopsDemo()

	fmt.Println("\n=== Hàm ===")
	functionsDemo()

	fmt.Println("\n=== Mảng và Slice ===")
	arraysSlicesDemo()

	fmt.Println("\n=== Map ===")
	mapsDemo()

	fmt.Println("\n=== Struct ===")
	structsDemo()

	fmt.Println("\n=== Pointer ===")
	pointersDemo()
}
