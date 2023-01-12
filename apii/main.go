package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var num int = 2

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, num)
}
func main() {
	router := gin.Default()
	router.GET("2", getBooks)
	router.Run("https://reqres.in/api/users/")
}
