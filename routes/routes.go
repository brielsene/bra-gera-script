package routes

import (
	"chg-gera-script-brad/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/styles", "./styles")
	r.Static("/assets", "./assets")
	r.Static("/archives", "./archives")
	r.GET("/index", controllers.Index)
	r.POST("/submit", controllers.ArquivoGerado)
	// r.GET("/download", controllers.DownloadArquivo)
	r.GET("/download", controllers.DownloadArquivo)

	r.Run(":8000")

}
