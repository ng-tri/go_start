package intermediate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func RunCustomer() {
	r := gin.Default()

	r.GET("/customers", func(c *gin.Context) {
		data := []Customer{
			{ID: 1, Name: "Nguyễn Văn A"},
			{ID: 2, Name: "Trần Thị B"},
		}
		c.JSON(http.StatusOK, data)
	})

	r.POST("/customers", func(c *gin.Context) {
		var newCustomer Customer
		if err := c.ShouldBindJSON(&newCustomer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Thêm khách hàng thành công", "data": newCustomer})
	})

	r.Run(":8080")
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := Product{ID: 1, Name: "Laptop", Price: 15000000}
	json.NewEncoder(w).Encode(p)
}

func RunProduct() {
	http.HandleFunc("/product", productHandler)
	http.ListenAndServe(":8080", nil)
}

type Order struct {
	Code  string `json:"code"`
	Total int    `json:"total"`
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Chỉ hỗ trợ POST", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Lỗi đọc dữ liệu", http.StatusBadRequest)
		return
	}

	var order Order
	if err := json.Unmarshal(body, &order); err != nil {
		http.Error(w, "Dữ liệu không hợp lệ", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Đã nhận đơn hàng: %s - Tổng tiền: %d\n", order.Code, order.Total)
}

func RunOrder() {
	http.HandleFunc("/order", orderHandler)
	http.ListenAndServe(":8080", nil)
}
