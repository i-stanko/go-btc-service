package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/rate", getCurrentBitcoinRate)
	router.POST("/api/subscribe", subscribeEmail)
	router.POST("/api/sendEmails", sendEmails)

	router.Run(":8080")
}

func getCurrentBitcoinRate(c *gin.Context) {
	// Реалізуати логіку для отримання поточного курсу біткоїна (BTC) у гривні (UAH)
	// Використати сторонній сервіс для отримання актуального курсу
}

func subscribeEmail(c *gin.Context) {
	// Реалізувати логіку для підписки електронної адреси на отримання оновлень про зміну курсу
}

func sendEmails(c *gin.Context) {
	// Реалізувати логіку для відправки листів з актуальним курсом на всі підписані електронні адреси
}
