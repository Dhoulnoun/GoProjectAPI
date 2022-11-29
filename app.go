package GoProjectAPI

import (
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Client *redis.Client
}

func (a *App) Initialize() {
	a.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	a.Router = mux.NewRouter()

}

func (a *App) Run(addr string) {
	// ...
}