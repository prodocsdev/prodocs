/*
 * Copyright 2022 The ProDocs Authors. All rights reserved. Use of this source code
 * is governed by a BSD-style license that can be found in the LICENSE file.
 */

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

// NewProdocsHTTPServer is the main HTTP server that is exposed to the users
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
	p.server.GET("/", Index)
	p.server.GET("/:path", Path)
	p.server.POST("/hook", Hook)
}
