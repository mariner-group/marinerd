package apiserver

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
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
	container := restful.NewContainer()
	//container.Filter(s.logRequest)
	if s.enableSwagger {
		container.Add(s.createSwaggerService(container))
	}
	return container, nil
}

func (s *APIServer) logRequest(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	chain.ProcessFilter(req, resp)
}

func (s *APIServer) createSwaggerService(container *restful.Container) *restful.WebService {
	config := restfulspec.Config{
		WebServices:                   container.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject,
	}
	return restfulspec.NewOpenAPIService(config)
}
