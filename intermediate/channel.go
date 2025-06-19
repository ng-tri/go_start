package intermediate

import (
	"errors"
	"fmt"
	"time"
)

func saveOrder(resultCh chan string, errCh chan error) {
	time.Sleep(1 * time.Second)
	resultCh <- "✅ Đã lưu đơn hàng vào DB"
}

func callShippingAPI(resultCh chan string, errCh chan error) {
	time.Sleep(2 * time.Second)
	errCh <- errors.New("🚫 Lỗi khi gọi API vận chuyển")
}

func sendEmail(resultCh chan string, errCh chan error) {
	time.Sleep(1 * time.Second)
	resultCh <- "📧 Đã gửi email xác nhận"
}

func RunChannel() {
	resultCh := make(chan string)
	errCh := make(chan error)

	go saveOrder(resultCh, errCh)
	go callShippingAPI(resultCh, errCh)
	go sendEmail(resultCh, errCh)

	success := 0
	fail := 0

	for i := 0; i < 3; i++ {
		select {
		case msg := <-resultCh:
			fmt.Println("✅ Thành công:", msg)
			success++
		case err := <-errCh:
			fmt.Println("❌ Thất bại:", err)
			fail++
		}
	}

	fmt.Println("✅ Tổng thành công:", success)
	fmt.Println("❌ Tổng thất bại:", fail)
	fmt.Println("🎉 Hoàn tất xử lý đơn hàng!")
}
