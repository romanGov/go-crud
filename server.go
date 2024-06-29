package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/romanGov/go-crud/handler/basic_handler"
	"hanlder/basic_handler"
)

func main() {
	router := gin.Default()
	fmt.Println("handling routes")
	router.GET("/go/work", basic_handler.hello)
	router.Run(":8080")
}
