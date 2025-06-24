package intermediate

import (
	"encoding/json"
	"fmt"
)

func InfoProduct() {
	p := Product{ID: 1, Name: "Laptop", Price: 20000000}
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Lỗi encode:", err)
		return
	}
	fmt.Println(string(jsonData))
}

// Output:
// {
//   "id": 1,
//   "name": "Laptop",
//   "price": 20000000
// }

func InfoOrder() {
	jsonStr := `{"code": "A123", "total": 500000}`
	var o Order

	err := json.Unmarshal([]byte(jsonStr), &o)
	if err != nil {
		fmt.Println("Lỗi decode:", err)
		return
	}

	fmt.Printf("Mã đơn: %s, Tổng: %d\n", o.Code, o.Total)
}

type CustomerWithOrders struct {
	Customer
	Orders []Order `json:"orders"`
}

func InOrder() {
	jsonStr := `{
        "id": 1,
        "name": "Nguyễn Văn A",
        "orders": [
            {"code": "A123", "total": 100000},
            {"code": "B456", "total": 250000}
        ]
    }`

	var c CustomerWithOrders
	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		fmt.Println("Lỗi:", err)
		return
	}

	fmt.Println("Tên khách:", c.Name)
	for _, o := range c.Orders {
		fmt.Printf("  Đơn %s - Tổng: %d\n", o.Code, o.Total)
	}
}
