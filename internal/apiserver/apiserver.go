package apiserver

import (
	"L2.18-calendar/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
}

func New(config *Config) *APIServer {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	return &APIServer{
		config: config,
		logger: logger,
		router: gin.Default(),
	}

}

func (s *APIServer) Run() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) ConfigureRouter(eventHandler *handler.Handler) {
	//POST
	s.router.POST("/create_event", eventHandler.Create)
	//GET
	s.router.GET("/events_for_month", eventHandler.GetForMonth)
	s.router.GET("/events_for_day", eventHandler.GetForDay)
	s.router.GET("/events_for_week", eventHandler.GetForWeek)

}
