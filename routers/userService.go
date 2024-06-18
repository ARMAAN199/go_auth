package router

import (
	"net/http"

	"github.com/ARMAAN199/Go_EcomApi/controllers"
	"github.com/ARMAAN199/Go_EcomApi/stores"
	"github.com/ARMAAN199/Go_EcomApi/utils"
	"gorm.io/gorm"
)

func initUserRouter(r *Routes, db *gorm.DB) {

	userStore := stores.NewDBUserStore(db)
	userController := controllers.NewUserController(userStore)

	r.Base.HandleFunc("/user/register", userController.RegisterUser()).Methods("POST")
	r.Base.HandleFunc("/user/login", userController.LoginUser()).Methods("POST")
	r.Base.Handle("/user/get", utils.AuthMiddleware(http.HandlerFunc(userController.GetUser()))).Methods("GET")
}
