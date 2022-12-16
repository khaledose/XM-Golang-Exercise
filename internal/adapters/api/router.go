package api

import (
	"net/http"
	"strconv"
	"time"
	"xm/companies/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	Config config.ServerConfigurations
	Router *chi.Mux
	log    *zap.SugaredLogger
}

func NewHTTPServer(logger *zap.SugaredLogger, config config.ServerConfigurations) *HTTPServer {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.Timeout(60 * time.Second))

	return &HTTPServer{
		Config: config,
		Router: router,
		log:    logger,
	}
}

func (r *HTTPServer) Start() {
	listeningAddr := ":" + strconv.Itoa(r.Config.Port)
	r.log.Infof("Server listening on port %s", listeningAddr)

	err := http.ListenAndServe(listeningAddr, r.Router)
	if err != nil {
		r.log.Fatalf("Failed to start http server. %v", err)
	}
}
