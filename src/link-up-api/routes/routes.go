package routes

import (
	"linkupapi/controllers"
	"net/http"
)

func RegisterRoutes(availabilityController *controllers.AvailabilityController) {
	http.HandleFunc("/availability", availabilityController.CheckAvailabilityHandler)
	http.HandleFunc("/proposal", controllers.CreateProposal)
}
