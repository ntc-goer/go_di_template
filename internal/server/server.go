package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HTTPServer struct {
	Engine *gin.Engine
	Server *http.Server
}

func NewHTTPServer(engine *gin.Engine) *HTTPServer {
	server := &http.Server{Addr: ":8080", Handler: engine}
	return &HTTPServer{
		Engine: engine,
		Server: server,
	}
}

func (h *HTTPServer) Start() error {
	if err := h.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Errorf("ListenAndServe fail %s", err)
		return err
	}
	return nil
}
