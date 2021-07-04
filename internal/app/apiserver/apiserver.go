package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// STRUCT APIServer
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// FACTORY
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// METHODS
// // Start server // //
func (s *APIServer) Start() error {
	// IF Logger's value is not nil - return loggerHandler
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	// Log that server is running
	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// // CONFIGURE THE LOGGER // //
func (s *APIServer) configureLogger() error {
	// // get log from server
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	// if ok - set level to logger and leave
	s.logger.SetLevel(level)

	return nil
}

// // ROUTES ENTRY
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

// // Wrapper around http-response. // //
// It allows to use middleware before returning the handler
func (s *APIServer) handleHello() http.HandlerFunc {
	//  Middleware here

	// Return router-handler
	return func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Hello")
	}
}