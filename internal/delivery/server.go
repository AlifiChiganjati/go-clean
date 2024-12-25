package delivery

import (
	"fmt"
	"log"

	"github.com/AlifiChiganjati/go-clean/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host   string
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		engine: engine,
		host:   host,
	}
}

func (s *Server) setupRoutes() {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (s *Server) Run() {
	s.setupRoutes()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
