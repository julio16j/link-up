package controllers

import (
	"encoding/json"
	"linkupapi/models"
	"linkupapi/repositories"
	"linkupapi/services"
	"net/http"
)

type AvailabilityController struct {
	Service *services.AvailabilityService
}

func NewAvailabilityController(service *services.AvailabilityService) *AvailabilityController {
	return &AvailabilityController{
		Service: service,
	}
}

func InitAvailabilityController() *AvailabilityController {
	repo := repositories.NewAvailabilityRepository()
	service := services.NewAvailabilityService(*repo)

	return NewAvailabilityController(service)
}

func (c *AvailabilityController) CheckAvailabilityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var address models.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	available, err := c.Service.IsAddressAvailable(address)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"available": available})
}
