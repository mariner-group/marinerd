package apiserver

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
)

type APIServer struct {
	listenIP      string       `yaml:"listenIP"`
	listenPort    uint32       `yaml:"listenPort"`
	server        *http.Server `yaml:"-"`
	enablePprof   bool         `yaml:"enablePprof"`
	enableSwagger bool         `yaml:"enableSwagger"`
}

func (s *APIServer) Run() {
	fmt.Printf("start API server.\n")
	var err error
	address := fmt.Sprintf("%v:%v", s.listenIP, s.listenPort)
	var wsContainer *restful.Container
	wsContainer, err = s.createRestfulContainer()
	if err != nil {
		return
	}
	server := http.Server{Addr: address, Handler: wsContainer, WriteTimeout: 1 * time.Minute}
	var ln net.Listener
	ln, err = net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("net listen(%s) err: %s\n", address, err.Error())
		return
	}
	s.server = &server
	err = server.Serve(ln)
	fmt.Printf("API server stop.")
}

func (s *APIServer) createRestfulContainer() (*restful.Container, error) {
	return nil, nil
}
