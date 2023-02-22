package main

import (
	"fmt"
	"net/http"

	"github.com/brunno98/word-analyze/service"
	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Text string `json:"text"`
}

func main() {
	r := gin.Default()
	r.POST("/analyze", func(ctx *gin.Context) {
		var body requestBody
		if err := ctx.BindJSON(&body); err != nil {
			return
		}
		fmt.Println(body)

		analyzeResult, err := service.AnalyzeText(&body.Text)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.IndentedJSON(http.StatusOK, analyzeResult)
	})
	r.Run()
}
