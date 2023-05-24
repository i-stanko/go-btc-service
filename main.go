package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"

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
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get bitcoin rate"})
		return
	}

	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	rate, ok := result["bitcoin"]["uah"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get bitcoin rate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rate": rate})
}

func subscribeEmail(c *gin.Context) {
	email := c.PostForm("email")

	// Перевірка, чи електронна адреса не є порожньою
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	// Перевірка, чи електронна адреса існує в файлі
	exists, err := isEmailExists(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email existence"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Збереження електронної адреси у файлі
	err = saveEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email subscribed"})
}

func sendEmails(c *gin.Context) {
	subscribers, err := getSubscribers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscribers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subscribers": subscribers})
}

func getSubscribers() ([]string, error) {
	file, err := os.Open(subscribersFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var subscribers []string
	for scanner.Scan() {
		subscribers = append(subscribers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}
