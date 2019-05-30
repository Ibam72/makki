package main

import (
	"os/exec"
    "github.com/gin-gonic/gin"
)

func str() string {
	data, _ := exec.Command("ojichat", "-e10", "まっきー").Output()
	return string(data)
}

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "index.html", gin.H{"data": str()})
    })

    router.Run()
}
