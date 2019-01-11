package  main

import (
	"fmt"
	"github.com/WithLin/ncma/middleware"
	"github.com/WithLin/ncma/model"
	"github.com/WithLin/ncma/router"
	"github.com/gin-gonic/gin"
	"github.com/WithLin/ncma/config"
	"io"
	"os"
)

func main(){
	fmt.Println("gin.Version: ", gin.Version)
	if config.Conf.GoConf.Env != model.DevelopmentMode {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()

	}
	logFile, err := os.OpenFile(config.Conf.GoConf.LogDir, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(-1)
	}
	gin.DefaultWriter = io.MultiWriter(logFile)

	app := gin.New()
	maxSize := int64(32)
	app.MaxMultipartMemory = maxSize << 20 // 3 MiB

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())
	app.Use(middleware.CORSMiddleware())
	router.Route(app)

	app.Run(":" + fmt.Sprintf("%d", config.Conf.GoConf.Port))
}