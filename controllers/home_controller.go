package controllers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func HomeIndex(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello from Go + Gin",
// 	})
// }

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🏠 Trang chủ")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ℹ️ Giới thiệu về chúng tôi")
}

func Http() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("🌐 Server đang chạy tại http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Lỗi khi chạy server:", err)
	}
}
