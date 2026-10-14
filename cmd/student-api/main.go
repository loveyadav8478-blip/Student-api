package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/loveyadav8478-blip/student-api/internal/config"
	"github.com/loveyadav8478-blip/student-api/internal/http/handlers/student"
	"github.com/loveyadav8478-blip/student-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup

	storage, err := sqlite.New(cfg)
	
	if err!=nil{
		log.Fatal(err)
	}
	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.22"))


	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students",student.New(storage))

	// Setup server

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	slog.Info("server started", slog.String("address",cfg.Addr))

	done := make(chan os.Signal,1)
	
	signal.Notify(done,os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		err := server.ListenAndServe()

		if err!=nil{
			log.Fatal("Failed to start server")
		}
	}()
 	
	<- done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	
	if err := server.Shutdown(ctx); err!=nil{
		slog.Error("failed to shutdown the server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
	
}