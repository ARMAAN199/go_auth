package router

import (
	controllers "github.com/ARMAAN199/Go_EcomApi/controllers"
	"github.com/ARMAAN199/Go_EcomApi/stores"
	"gorm.io/gorm"
)

func initHandShakeRouter(r *Routes, db *gorm.DB) {

	baseStore := stores.NewDBBaseStore(db)
	handShakeController := controllers.NewHandShakeController(baseStore)

	r.Base.HandleFunc("/test", handShakeController.ShakeHandsHandler()).Methods("GET")
}
