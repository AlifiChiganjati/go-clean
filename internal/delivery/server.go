package delivery

import (
	"fmt"
	"log"

	"github.com/AlifiChiganjati/go-clean/config"
	"github.com/AlifiChiganjati/go-clean/internal/auth/handler"
	"github.com/AlifiChiganjati/go-clean/internal/auth/usecase"
	router "github.com/AlifiChiganjati/go-clean/internal/delivery/http"
	"github.com/AlifiChiganjati/go-clean/internal/delivery/middleware"
	"github.com/AlifiChiganjati/go-clean/internal/manager"
	"github.com/AlifiChiganjati/go-clean/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	hm         manager.HandlerManager
	auth       *handler.AuthHandler
	engine     *gin.Engine
	host       string
	jwtService jwt.JwtToken
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}
	jwtService := jwt.NewJwtToken(cfg.TokenConfig)
	repo := manager.NewRepoManager(infra)
	uc := manager.NewUseCaseManager(repo)
	engine := gin.Default()
	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	hm := manager.NewHandlerManager(uc, engine.Group("/api/v1"), authMiddleware)
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		hm:         hm,
		engine:     engine,
		host:       host,
		auth:       handler.NewAuthHandler(usecase.NewAuthUseCase(uc.UserUseCase(), jwtService)),
		jwtService: jwtService,
	}
}

func (s *Server) setupRoutes() {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	rg := s.engine.Group("/api/v1")
	router.NewAuthRouter(*s.auth, rg).Route()
	router.NewUserRouter(s.hm.UserHandler(), rg, authMiddleware).Route()
	router.NewServiceRouter(s.hm.ServiceHandler(), rg, authMiddleware).Route()
}

func (s *Server) Run() {
	s.setupRoutes()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
