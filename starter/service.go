package starter

import (
	"fmt"
	"net/http"

	"github.com/IMQS/log"
	"github.com/IMQS/nf"
	"github.com/IMQS/serviceauth/permissions"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// Service definition
type Service struct {
	log    *log.Logger
	router *httprouter.Router
	config Config
	db     *gorm.DB
}

// NewService creates a new AssetCore service
func NewService() *Service {
	return &Service{
		log:    log.New(log.Stdout),
		router: httprouter.New(),
	}
}

// ListenAndServe runs the service
func (s *Service) ListenAndServe() {
	// Setup HTTP API
	nf.Handle(s.router, "GET", "/ping", s.ping)
	nf.HandleAuthenticated(s.router, "GET", "/frog/list", s.listFrogs, nil)
	nf.HandleAuthenticated(s.router, "POST", "/frog/add", s.addFrogs, []int{permissions.PermReportCreator}) // User must have ReportCreator permission to call this

	s.log.Infof("Starting HTTP server on port %v", s.config.HttpPort)
	http.ListenAndServe(fmt.Sprintf(":%v", s.config.HttpPort), s.router)
}

// LoadConfig loads our config from the configuration service, and panics upon failure
func (s *Service) LoadConfig() {
	err := s.config.Load()
	if err != nil {
		panic(err)
	}
}

// Initialize connects to our database and performs any other initialization before starting up
func (s *Service) Initialize() {
	var err error
	if s.db, err = openDB(s.log, s.config.DB); err != nil {
		panic(err)
	}
}
