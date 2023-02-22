package main

import (
	"fmt"
	"net/http"

	"github.com/brunno98/word-analyze/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/analyze", func(ctx *gin.Context) {
		var body struct {
			Text string `json:"text"`
		}
		if err := ctx.BindJSON(&body); err != nil {
			return
		}

		analyzeResult, err := service.AnalyzeText(&body.Text)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var response struct {
			Words map[string]int `json:"words"`
		}
		response.Words = analyzeResult
		ctx.IndentedJSON(http.StatusOK, response)
	})
	r.Run()
}
