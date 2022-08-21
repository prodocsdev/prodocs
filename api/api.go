package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/prodocsdev/prodocs/config"
)

type ProdocsHTTPServer struct {
	server *echo.Echo
	port   string
}

// NewProdocsHTTPServer is
func NewProdocsHTTPServer(config *config.ProdocsConfig) (*ProdocsHTTPServer, error) {
	s := ProdocsHTTPServer{
		server: echo.New(),
		port:   config.Port,
	}
	return &s, nil
}

func (p *ProdocsHTTPServer) Start() {
	p.setupRoutes()
	p.server.Logger.Fatal(p.server.Start(fmt.Sprintf(":%v", p.port)))
}

func (p *ProdocsHTTPServer) setupRoutes() {

}
