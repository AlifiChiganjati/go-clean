package server

import (
	"fmt"

	router "github.com/AlifiChiganjati/go-clean/api/v1"
	"github.com/AlifiChiganjati/go-clean/config"
	"github.com/AlifiChiganjati/go-clean/config/manager"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	control manager.ControlManager
	engine  *gin.Engine
	host    string
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	repo := manager.NewRepoManager(infra)
	uc := manager.NewUseCaseManager(repo)
	control := manager.NewControlManager(uc)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		control: control,
		engine:  engine,
		host:    host,
	}
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	s.engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "server is healthy"})
	})

	router.NewUserRoutes(s.control.UserController(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		logrus.Fatal("Server can't run")
	}
}
