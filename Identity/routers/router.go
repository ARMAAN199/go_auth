package router

import (
	"net/http"

	"github.com/ARMAAN199/Go_EcomApi/redis"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Routes struct {
	Base *mux.Router
}

func ReturnRouter(db *gorm.DB, redisClient *redis.RedisStore) *http.Handler {
	routes := Routes{}
	routes.Base = mux.NewRouter().PathPrefix("/api").Subrouter()

	initHandShakeRouter(&routes, db)
	initUserRouter(&routes, db, redisClient)

	handler := http.Handler(routes.Base)
	return &handler
}
