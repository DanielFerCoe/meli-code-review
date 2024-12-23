package service

import (
	"app/internal"
	"errors"
)

var (
	ErrVehiclesNotFound = errors.New("Vehicles not found")
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) GetByTransmissionTypeService(transmissionType string) (map[int]internal.Vehicle, error) {
	vehicles, err := s.rp.FindManyByTransmissionTypeRepository(transmissionType)

	if err != nil {
		return nil, err
	}

	if len(vehicles) == 0 {
		return nil, ErrVehiclesNotFound
	}

	return vehicles, nil
}
