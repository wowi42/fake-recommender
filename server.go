package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {
    router := gin.Default()

    // This handler will match /user/john but will not match neither /user/ or /user
    router.GET("/*prout", func(c *gin.Context) {
		var a [0]string
        c.JSON(http.StatusOK, a)
    })

    router.Run(":8080")
}
