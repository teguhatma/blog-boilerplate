package http

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
	fe "github.com/teguhatma/blog-boilerplate/errors"
)

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

type ContentType string

const (
	ContentTypeKey  = "content-type"
	ContentTypeJSON = ContentType("application/json")
	defaultMessage  = "Internal Server Error"
	defaultCode     = ""
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

// ErrorResponse - Struct for the error response
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseError struct {
	StatusCode int
	fe.FError
}

type httpResponse struct {
	Data interface{} `json:"data"`
}

func NewHeaders() HTTPHeaders {
	return make(map[string]string)
}

func setHeaders(headers HTTPHeaders, w http.ResponseWriter) {
	if _, ok := headers[ContentTypeKey]; !ok {
		w.Header().Set(ContentTypeKey, string(ContentTypeJSON))
	}

	for key, val := range headers {
		w.Header().Set(key, val)
	}
}

func InitServer(router *mux.Router) (*Server, error) {
	httpConfig := &Config{
		Name:    "test",
		Port:    8030,
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

type AppHandler func(*http.Request) (*Response, error)

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := fn(r)

	if err != nil {
		errResponse := writeErrorResponse(err, w)

		if _, err := w.Write(errResponse); err != nil {
			fmt.Sprintf("http response writing failde: %v", err)
		}

		fmt.Sprintf("error is returned %v", err)
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

func IsValidStatusCode(statusCode int) bool {
	return http.StatusText(statusCode) != ""
}

func writeErrorResponse(err error, w http.ResponseWriter) []byte {
	w.Header().Set(ContentTypeKey, string(ContentTypeJSON))

	switch e := err.(type) {
	case ResponseError: // This for http or sync way
		if IsValidStatusCode(e.StatusCode) {
			w.WriteHeader(e.StatusCode)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return marshalErrorResponse(e.Error(), string(e.Code()))
	case fe.FError: // This for async way or kafka
		w.WriteHeader(http.StatusInternalServerError)
		return marshalErrorResponse(e.Error(), string(e.Code()))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return marshalErrorResponse(defaultMessage, defaultCode)
	}
}

func marshalErrorResponse(message, code string) []byte {
	errResponse := ErrorResponse{
		Message: message,
		Code:    code,
	}

	response, err := json.Marshal(errResponse)
	if err != nil {
		return []byte(http.StatusText(http.StatusInternalServerError))
	}

	return response
}
