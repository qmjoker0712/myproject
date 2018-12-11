package ginrpc

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

var (
	UnsupportWSError = errors.New("Unsupport websocket now!")
)

//==========================
var loger Logger

type Service struct {
	router *gin.Engine
	config *Config
}

// NewService create new rpc service
// - input log instance
func NewService(l Logger, conf *Config) *Service {
	loger = l
	s := &Service{}
	s.config = conf
	s.router = gin.Default()
	return s
}

// RegisterPostHandler register an http handler to handle request by given router path
func (s *Service) RegisterPostHandler(h http.Handler, routerPath string) {
	path := "/" + routerPath
	s.router.POST(path, func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
}

// RegisterGetHandler register an http handler to handle request by given router path
func (s *Service) RegisterGetHandler(h http.Handler, routerPath string) {
	path := "/" + routerPath
	s.router.GET(path, func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
}

// RegisterAPI create a service for given interface{}
// - register all public methods in given interface
// - all params and returns must public, otherwise will failed
// - methods return values no more then 2, and the second return value must be 'error'
// - not support websocket now
func (s *Service) RegisterAPI(api *API, routerPath string) error {
	err := registerService(api.Namespace, api.Service)
	if err != nil {
		return err
	}
	s.registerAPIRouter(routerPath)
	return nil
}

// Start start run rpc service
func (s *Service) Start() error {
	port := ":" + s.config.GinPort
	go func() {
		if err := s.router.Run(port); err != nil {
			msg := fmt.Errorf("[ginrpc] could not start server: %v", err)
			loger.Error(msg)
		}
	}()

	return nil
}

// Stop stop run rpc service
func (s *Service) Stop() error {
	return nil
}

//==================================================
func (s *Service) registerAPIRouter(routerPath string) {
	path := "/" + routerPath
	s.router.POST(path, func(c *gin.Context) {
		var rawJson jsonRequest
		if err := c.ShouldBindJSON(&rawJson); err != nil {
			loger.Error("[ginrpc] Request body unpack to json rpc model failed! ", err)
			c.JSON(http.StatusBadRequest, CreateErrorResponse(rawJson.Id, err))
			return
		}

		serverReq, err := createServerRequest(&rawJson)
		if err != nil {
			loger.Error("[ginrpc] createServerRequest error:", err)
			c.JSON(http.StatusBadRequest, CreateErrorResponse(rawJson.Id, err))
			return
		}

		ret, err := handle(serverReq)
		if err != nil {
			loger.Error("[ginrpc] handle error:", err)
			c.JSON(http.StatusBadRequest, CreateErrorResponse(rawJson.Id, err))
			return
		}

		// success
		c.JSON(http.StatusOK, CreateResponse(rawJson.Id, ret))
	})
}

func createServerRequest(jsonReq *jsonRequest) (*serverRequest, error) {
	rpcReq, err := parseRequest(jsonReq)
	if err != nil {
		return nil, fmt.Errorf("parse method failed:", err)
	}

	if rpcReq.isPubSub {
		return nil, UnsupportWSError
	}

	svc, ok := apiService[rpcReq.service]
	if !ok {
		return nil, fmt.Errorf("Unsupport service:", rpcReq.service)
	}

	callfunc, exist := svc.callbacks[rpcReq.method]
	if !exist {
		return nil, fmt.Errorf("Cannot find method:", rpcReq.method)
	}

	serverReq := &serverRequest{
		id:      rpcReq.id,
		svcname: svc.name,
		callb:   callfunc,
	}
	if len(callfunc.argTypes) == 0 {
		return serverReq, nil
	}
	if rpcReq.params != nil && len(callfunc.argTypes) > 0 {
		var parseArgErr error
		serverReq.args, parseArgErr = parseRequestArguments(callfunc.argTypes, rpcReq.params)
		if parseArgErr != nil {
			return nil, fmt.Errorf("Method:%s Parse arguments failed! %s", rpcReq.method, parseArgErr.Error())
		}
	} else {
		return nil, fmt.Errorf("Input param mismatch:", rpcReq.method)
	}

	return serverReq, nil
}

func handle(req *serverRequest) (interface{}, error) {
	if req.isUnsubscribe {
		return nil, UnsupportWSError
	}

	if req.isUnsubscribe {
		return nil, UnsupportWSError
	}

	arguments := []reflect.Value{req.callb.rcvr}
	if req.callb.hasCtx {
		//arguments = append(arguments, reflect.ValueOf(ctx))
		// todo balabala
	}
	if len(req.args) > 0 {
		arguments = append(arguments, req.args...)
	}

	reply := req.callb.method.Func.Call(arguments)
	if len(reply) == 0 {
		return nil, nil
	}
	if req.callb.errPos >= 0 {
		if !reply[req.callb.errPos].IsNil() {
			e := reply[req.callb.errPos].Interface().(error)
			return nil, e
		}
	}
	return reply[0].Interface(), nil
}
