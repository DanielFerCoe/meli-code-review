package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	GetByTransmissionTypeService(transmissionType string) (map[int]Vehicle, error)
}
