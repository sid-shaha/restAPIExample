package server

import (
	"log"
	"net/http"
	"time"

	"github.com/sid-shaha/restAPIExample/handlers"
	"github.com/sid-shaha/restAPIExample/models"
	"github.com/sid-shaha/restAPIExample/repository"
	"github.com/sid-shaha/restAPIExample/service"
	"github.com/sid-shaha/restAPIExample/web"
)

type App struct {
	httpServer  *http.Server
	PostService service.PostService
}

func NewApp() *App {
	storage := initStorage()

	postRepo := repository.NewPostIMRepository(storage)
	postService := service.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	config := web.RouterConfig{
		PostHandler: postHandler,
	}

	router := web.NewRouter(config)

	return &App{
		httpServer: &http.Server{
			Addr:           ":8080",
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20, // 1 MB max header size
		},
		PostService: postService,
	}
}
func (a *App) Run() {
	log.Println("Starting server on :8080...")
	if err := a.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
func initStorage() *map[int]models.Post {
	myMap := make(map[int]models.Post, 0)
	return &myMap
}
