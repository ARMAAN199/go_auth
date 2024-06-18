package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Routes struct {
	Base *mux.Router
}

func ReturnRouter(db *gorm.DB) *http.Handler {
	routes := Routes{}
	routes.Base = mux.NewRouter().PathPrefix("/api").Subrouter()

	initHandShakeRouter(&routes, db)
	initUserRouter(&routes, db)

	handler := http.Handler(routes.Base)
	return &handler
}
