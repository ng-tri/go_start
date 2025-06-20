package intermediate

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

func WriteFile() {
	file, err := os.Create("intermediate/donhang.txt")
	if err != nil {
		fmt.Println("❌ Lỗi tạo file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Đơn hàng #12345\n")
	file.WriteString("Khách: Nguyễn Văn A\n")
	file.WriteString("Tổng tiền: 2.000.000 VNĐ\n")

	fmt.Println("✅ Ghi file thành công")
}

func ReadFile() error {
	file, err := os.Open("intermediate/donhang.txt")
	if err != nil {
		return fmt.Errorf("không thể mở file: %w", err) // wrap lại lỗi gốc
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("📄 Nội dung file:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("❌ Lỗi khi đọc file:", err)
		return fmt.Errorf("lỗi khi đọc file: %w", err)
	}

	return nil
}

func CheckFile() {
	err := ReadFile()
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File không tồn tại.")
	} else if err != nil {
		fmt.Println("Lỗi khác:", err)
	}
}

func WriteCsv() {
	file, err := os.Create("intermediate/donhang.csv")
	if err != nil {
		fmt.Println("❌ Không tạo được file CSV:", err)
		return
	}
	defer file.Close()

	file.WriteString("\uFEFF")

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Mã đơn", "Khách hàng", "Tổng tiền"})
	writer.Write([]string{"12345", "Nguyễn Văn A", "2000000"})
	writer.Write([]string{"12346", "Trần Thị B", "1500000"})
	writer.Write([]string{"12347", "Thái Đức C", "1900000"})

	fmt.Println("✅ Ghi file CSV thành công")
}

func logToFile(wg *sync.WaitGroup, id int, file *os.File, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		msg := fmt.Sprintf("Goroutine %d - dòng log %d\n", id, i+1)
		mu.Lock()
		file.WriteString(msg)
		mu.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func LogGoRoutine() {
	file, err := os.Create("intermediate/logs.txt")
	if err != nil {
		fmt.Println("❌ Không thể tạo file log:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go logToFile(&wg, i, file, &mu)
	}

	wg.Wait()
	fmt.Println("✅ Ghi log từ nhiều goroutine hoàn tất")
}
