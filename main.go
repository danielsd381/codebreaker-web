package main

import (
  "net/http"
  "os"
  "github.com/gin-gonic/gin"
)

func main() {


  	port := os.Getenv("PORT")

  	if port == "" {
  		port = "8083"
  	}

	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/codebreaker/setup/:number", func(c *gin.Context) {
		 number := c.Param("number")
     setSecret(number)
		c.String(http.StatusOK, "Secret number configured:"+ number)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/codebreaker/guess/:number", func(c *gin.Context) {
		name := c.Param("number")
		result := validate(name)
		c.String(http.StatusOK, "Answer: " + result)
	})

	router.Run(":" +port)
}
