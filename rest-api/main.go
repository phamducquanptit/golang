package main

import (
	"errors"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"rest-api/config"
	"rest-api/models"
	"rest-api/routers"
	"rest-api/routers/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

// @title REST API TO DO LIST
// @version 1.0
// @description This is rest api demo
// @termsOfService http://swagger.io/terms/

// @contact.name quanpham
// @contact.url https://www.facebook.com/quanphamptit
// @contact.email phamducquanptithcm@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3001
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization


func main() {
	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init DB
	models.DB.Init()
	defer models.DB.Close()

	// set gin mode run
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	// Routes
	routers.Load(
		g,
		middleware.RequestId(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("Các routers không phản hồi hoặc có thể mất quá nhiều thời gian để phản hồi")
		}
		log.Println("Đã deployed thành công !!")
	}()

	log.Printf("Start listening request on port %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url"))
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Println("Đang đợi server..., thử lại sau 1s")
		time.Sleep(time.Second)
	}
	return errors.New("Không thể kết nối đến server")
}
