package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reverse(s string) string {
	b := []byte(s)
	fmt.Println(b, len(b))
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func HomepageHandler(c *gin.Context) {
	input := "creating an engine instance with the logger"
	rev := Reverse(input)
	fmt.Println(rev)
	drev := Reverse(rev)
	fmt.Println(drev)
	
	c.JSON(http.StatusOK, gin.H{
		"message" : "Here is HomePage",
		"input" : input,
		"rev input" : rev,
		"double rev input": drev,
	})
}