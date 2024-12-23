package handler

import (
	"app/internal"
	"app/internal/service"
	responseAPI "app/pkg/response"
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

func vehicleToVehicleJSON(v internal.Vehicle) VehicleJSON {
	return VehicleJSON{
		ID:              v.Id,
		Brand:           v.Brand,
		Model:           v.Model,
		Registration:    v.Registration,
		Color:           v.Color,
		FabricationYear: v.FabricationYear,
		Capacity:        v.Capacity,
		MaxSpeed:        v.MaxSpeed,
		FuelType:        v.FuelType,
		Transmission:    v.Transmission,
		Weight:          v.Weight,
		Height:          v.Height,
		Length:          v.Length,
		Width:           v.Width,
	}
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, vehicle := range v {
			data[key] = vehicleToVehicleJSON(vehicle)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmissionType := chi.URLParam(r, "type")

		vehicles, err := h.sv.GetByTransmissionTypeService(transmissionType)

		if err != nil {
			switch {
			case errors.Is(err, service.ErrVehiclesNotFound):
				resBody := responseAPI.ResponseBody{
					Error:   true,
					Message: err.Error(),
				}

				response.JSON(w, http.StatusNotFound, resBody)
			default:
				resBody := responseAPI.ResponseBody{
					Error:   true,
					Message: "Internal error server",
				}
				response.JSON(w, http.StatusInternalServerError, resBody)
			}

			return
		}

		data := make(map[int]VehicleJSON, len(vehicles))

		for key, vehicle := range vehicles {
			data[key] = vehicleToVehicleJSON(vehicle)
		}

		resBody := responseAPI.ResponseBody{
			Data: data,
		}

		response.JSON(w, http.StatusOK, resBody)
	}
}
