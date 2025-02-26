package controllers

// controllers are basically repositories i guess.

import (
	"net/http"

	"github.com/ARMAAN199/Go_EcomApi/stores"
)

type HandShakeController struct {
	store stores.BaseStore
}

func NewHandShakeController(store stores.BaseStore) *HandShakeController {
	controller := HandShakeController{
		store: store,
	}
	return &controller
}

func (ctrl *HandShakeController) ShakeHandsHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "GET")

		// car, err := ctrl.store.GetCar()

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println(car)

		// stringifiedResponse, err := json.Marshal(car)
		// if err != nil {
		// 	log.Fatal()
		// }

		// w.Write([]byte(fmt.Sprintf("Hello World! I used to drive a %s %s made in the year %d. \nHere are the details %s", car.Brand, car.Model, car.Year, stringifiedResponse)))

		w.Write([]byte("Hello Armaan!"))
	}
}
