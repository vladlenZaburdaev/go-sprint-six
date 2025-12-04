package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.HandleIndex)
	router.HandleFunc("upload", handlers.HandleUpload)

	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		Server: server,
	}
}
