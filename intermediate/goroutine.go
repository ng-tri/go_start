package intermediate

import (
	"fmt"
	"time"
)

// Hàm say in ra message 3 lần, mỗi lần cách nhau 0.5 giây
func say(message string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(message, i)
		time.Sleep(1 * time.Second)
	}
}

func RunGoroutine() {
	// Gọi hàm say như một goroutine (chạy song song)
	go say("Goroutine 1")

	// Gọi hàm say bình thường (chạy chính)
	say("Main routine")

	// Giữ chương trình chạy đủ lâu để goroutine có thời gian in hết
	time.Sleep(3 * time.Second)
}
