package apiserver

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/spf13/viper"
)

type APIServer struct {
	server        *http.Server
	listenIP      string
	listenPort    int64
	enableSwagger bool
	enablePprof   bool
}

func (s *APIServer) Initialize() error {
	s.listenIP = viper.GetString("APIServer.listenIP")
	s.listenPort = viper.GetInt64("APIServer.listenPort")
	s.enableSwagger = viper.GetBool("APIServer.enableSwagger")
	s.enablePprof = viper.GetBool("APIServer.enablePprof")
	if s.listenIP == "" {
		return errors.New("listenIP is empty")
	}
	if s.listenPort == 0 {
		return errors.New("listenPort is empty")
	}
	return nil
}

func (s *APIServer) Stop() error {
	return nil
}

func (s *APIServer) Wait() error {
	return nil
}

func (s *APIServer) Run() error {
	container, err := s.createRestfulContainer()
	if err != nil {
		return err
	}
	s.server = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", s.listenIP, s.listenPort),
		Handler:        container,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	listener, err := net.Listen("tcp", s.server.Addr)
	if err != nil {
		return err
	}
	fmt.Printf("Listening on %s", s.server.Addr)
	return s.server.Serve(listener)
}

func (s *APIServer) createRestfulContainer() (*restful.Container, error) {
	return nil, nil
}
