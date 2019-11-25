package server

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

const (
	get  = "GET"
	post = "POST"
)

func New() *gin.Engine {
	return gin.New()
}

func InitServer(lc fx.Lifecycle, s *gin.Engine, l *log.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			port := os.Getenv("TODO_PORT")
			if port == "" {
				port = "8080"
			}
			go s.Run(":" + port)
			return nil
		},
	})
}

func RegisterHanders(s *gin.Engine, hss []*handlerSignature) {
	for _, hs := range hss {
		switch hs.method {
		case "GET":
			s.GET(hs.path, hs.handler)
		case "POST":
			s.POST(hs.path, hs.handler)
		default:
			log.Println("Invalid hander signature")
		}
	}
}
