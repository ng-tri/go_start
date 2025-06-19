package basic

import (
	"errors"
	"fmt"
)

func RunTest() {
	var name string
	var age int

	// 📌 Scan đọc lần lượt các giá trị (cách nhau bởi khoảng trắng)
	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age)
	fmt.Println("Name:", name, "| Age:", age)

	// 📌 Scanln giống Scan, nhưng dừng khi gặp dấu xuống dòng
	fmt.Print("Enter your city: ")
	var city string
	fmt.Scanln(&city)
	fmt.Println("City:", city)

	// 📌 Scanf đọc theo định dạng cụ thể
	var day int
	fmt.Print("Enter your birthday (dd): ")
	fmt.Scanf("%d\n", &day)
	fmt.Println("Birthday (day):", day)

	username := "Alice"
	score := 95.75

	// 📌 Sprintf: định dạng dữ liệu thành chuỗi
	s1 := fmt.Sprintf("Student: %s - Score: %.2f", username, score)
	fmt.Println(s1)

	// 📌 Sprint: giống Print nhưng trả về chuỗi
	s2 := fmt.Sprint("Hello", ",", username)
	fmt.Println(s2)

	// 📌 Sprintln: tự động thêm dấu cách giữa các tham số + dấu xuống dòng
	s3 := fmt.Sprintln("Hi", username)
	fmt.Print(s3) // Sử dụng Print để không in 2 lần dòng mới

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	// 📌 Errorf cũng có thể gắn vào errors.Wrap-style
	err2 := fmt.Errorf("file not found: %w", errors.New("config.json"))
	fmt.Println(err2)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

func RunBasic() {
	fmt.Println("Xin chào từ Go trên WSL!")

	// 1. Biến & kiểu dữ liệu
	var a string = "Xin chào"
	var b int = 10
	var c float64 = 3.14
	var d bool = true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// Khai báo ngắn gọn
	e := "Go rất tuyệt"
	fmt.Println(e)

	// 2. Toán tử
	g := 10
	h := 3

	fmt.Println(g + h) // Cộng
	fmt.Println(g - h) // Trừ
	fmt.Println(g * h) // Nhân
	fmt.Println(g / h) // Chia
	fmt.Println(g % h) // Chia dư

	// 3. Cấu trúc điều kiện
	x := 5

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

	// 4. Vòng lặp
	for i := 1; i <= 5; i++ {
		fmt.Println("Lần thứ:", i)
	}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	chaoTen("Linh")

	tong := cong(11, 22)
	fmt.Println(tong)

	thuong, err := chia(22, 11)
	if err != nil {
		fmt.Println("❌ Lỗi:", err)
		return
	}
	fmt.Println("✅ Kết quả phép chia:", thuong)
	array()
}

// 5. Hàm
func chaoTen(ten string) {
	fmt.Println("Xin chào", ten)
}

// Hàm có return
func cong(a int, b int) int {
	return a + b
}

func chia(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("không thể chia cho 0")
	}
	return a / b, nil
}

func array() {
	// 6. Mảng & Slice

	// Mảng
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// Slice (linh hoạt hơn)
	numbers := []int{4, 5, 6}
	numbers = append(numbers, 7)
	fmt.Println(numbers[1]) // 5

	nums := []int{1, 2, 3, 4, 5}
	i := 2 // xóa phần tử tại index 2 (giá trị 3)
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println("Sau khi xóa :", nums)
	// nums[:2] chính là [1, 2]
	// nums[3:] là [4, 5]

	// Hàm append(a, b...) nối slice b vào cuối slice a.
	// ... là toán tử "unpack" (mở rộng slice b thành từng phần tử).
	// Kết quả: append([1, 2], 4, 5) ⇒ [1, 2, 4, 5]

	//  Xóa phần tử theo giá trị (không biết index)
	val := 3
	for i, v := range nums {
		if v == val {
			nums = append(nums[:i], nums[i+1:]...)
			fmt.Println("Sau khi xóa :", nums)
			break // chỉ xóa lần đầu tiên
		}
	}

	// 7. Map (dictionary)
	ages := map[string]int{
		"Linh": 30,
		"Nam":  25,
	}

	fmt.Println(ages["Linh"])
	ages["An"] = 28
}

// // 8. Struct (như class)
// type Nguoi struct {
//     Ten  string
//     Tuoi int
// }

// func main() {
//     a := Nguoi{Ten: "Linh", Tuoi: 30}
//     fmt.Println(a.Ten)
// }

// // 9. Pointer
// func tang(x *int) {
//     *x += 1
// }

// func main() {
//     a := 10
//     tang(&a)
//     fmt.Println(a) // 11
// }
