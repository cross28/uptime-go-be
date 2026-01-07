package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"crosssystems.co/uptime-go-be/middleware"
	"crosssystems.co/uptime-go-be/internal/auth"
	"crosssystems.co/uptime-go-be/internal/health"
	"crosssystems.co/uptime-go-be/internal/users"

	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	port int
	env string
}

type App struct {
	config Configuration
}

func main() {
	var cfg Configuration

	flag.IntVar(&cfg.port, "port", 8000, "api server port")
	flag.StringVar(&cfg.env, "env", "dev", "server environment(development|qa|prod)")
	flag.Parse()

	// app := &App {
	// 	config: cfg,
	// }

	log.SetReportCaller(true)

	r := chi.NewRouter();
	
	cors := cors.New(cors.Options{
		AllowedOrigins: []string { "https://*", "http://*" },
		AllowedMethods: []string { "GET", "POST", "PUT", "DELETE", "OPTION" },
		AllowedHeaders: []string { "*" },
	})

	r.Use(chi_middleware.Logger)
	r.Use(chi_middleware.StripSlashes)
	r.Use(chi_middleware.RequestID)
	r.Use(middleware.JsonContentType)
	r.Use(cors.Handler)

	r.Route("/health", health.RegisterRoutes)
	r.Route("/user", users.RegisterRoutes)
	r.Route("/", auth.RegisterRoutes)

	log.Infof("server starting: port %d | env %s", cfg.port, cfg.env)

	srvr := &http.Server{
		Addr: 	fmt.Sprintf(":%d", cfg.port),
		Handler: r,
		IdleTimeout: time.Minute,
		WriteTimeout: time.Second * 10,
		ReadTimeout: time.Second * 5,
	}

	err := srvr.ListenAndServe()
	log.Error(err.Error())
	os.Exit(1)
}