package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type ContentType string

const (
	ContentTypeKey  = "content-type"
	ContentTypeJSON = ContentType("application/json")
)

type HTTPHeaders map[string]string

func (h HTTPHeaders) Add(key, val string) {
	h[key] = val
}

type Response struct {
	Data       interface{}
	StatusCode int
	Headers    HTTPHeaders
}

type httpResponse struct {
	Data interface{} `json:"data"`
}

type serverOptions struct {
	methodNotAllowedHandler http.Handler
	notFoundHandler         http.Handler
	serverStopTimeout       int
}

type Option func(*serverOptions) error

type Server struct {
	server  *http.Server
	options *serverOptions
	exit    chan os.Signal
}

type Config struct {
	Name    string
	Port    int
	Version string
}

func InitServer(router *mux.Router) (*Server, error) {
	httpConfig := &Config{
		Name:    "test",
		Port:    8080,
		Version: "v1",
	}
	server, err := NewHTTPServer(httpConfig, router)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (s *Server) Start(ctx context.Context) (err error) {
	timeout := time.Duration(s.options.serverStopTimeout) * time.Second

	idleConnClosed := make(chan struct{})

	go func() {
		signal.Notify(s.exit, os.Interrupt, syscall.SIGTERM)

		sig := <-s.exit
		ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)

		defer cancel()

		fmt.Printf("Received signal {%v}, shutting down HTTP Server....", sig.String())

		if err = s.server.Shutdown(ctx); err != nil {
			fmt.Printf("Error while shutting down HTTP server: %v", err)
		}

		close(idleConnClosed)
	}()

	if serveErr := s.server.ListenAndServe(); !errors.Is(serveErr, http.ErrServerClosed) {
		err = serveErr
	}

	<-idleConnClosed

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	s.exit <- os.Interrupt
}

func defaultMethodNotAllowedHandler(r *http.Request) (*Response, error) {
	headers := NewHeaders()

	return &Response{
		Data:       fmt.Sprintf("%v Method not allowed", r.Method),
		StatusCode: http.StatusMethodNotAllowed,
		Headers:    headers,
	}, nil
}

func defaultNotFoundHandler(r *http.Request) (*Response, error) {
	headers := NewHeaders()

	return &Response{
		Data:       "Url not found",
		StatusCode: http.StatusNotFound,
		Headers:    headers,
	}, nil
}

func NewHTTPServer(config *Config, router *mux.Router, opts ...Option) (*Server, error) {
	options := &serverOptions{
		methodNotAllowedHandler: AppHandler(defaultMethodNotAllowedHandler),
		notFoundHandler:         AppHandler(defaultNotFoundHandler),
		serverStopTimeout:       2,
	}

	for _, o := range opts {
		if o != nil {
			if err := o(options); err != nil {
				return nil, err
			}
		}
	}

	if config == nil {
		return nil, fmt.Errorf("config is null")
	}

	if config.Port < 1 {
		return nil, fmt.Errorf("port less than 1")
	}

	if router == nil {
		return nil, fmt.Errorf("router should be nil")
	}

	if options.methodNotAllowedHandler != nil {
		router.MethodNotAllowedHandler = options.methodNotAllowedHandler
	}

	if options.notFoundHandler != nil {
		router.NotFoundHandler = options.notFoundHandler
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Port),
		Handler: router,
	}

	return &Server{
		server:  server,
		options: options,
		exit:    make(chan os.Signal),
	}, nil
}

func NewHeaders() HTTPHeaders {
	return make(map[string]string)
}

type AppHandler func(*http.Request) (*Response, error)

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := fn(r)

	if err != nil {
		return
	}

	setHeaders(res.Headers, w)
	w.WriteHeader(res.StatusCode)

	response, err := json.Marshal(
		httpResponse{
			Data: res.Data,
		},
	)
	if err != nil {
		return
	}

	if _, err := w.Write(response); err != nil {
		return
	}
}

func setHeaders(headers HTTPHeaders, w http.ResponseWriter) {
	if _, ok := headers[ContentTypeKey]; !ok {
		w.Header().Set(ContentTypeKey, string(ContentTypeJSON))
	}

	for key, val := range headers {
		w.Header().Set(key, val)
	}
}
