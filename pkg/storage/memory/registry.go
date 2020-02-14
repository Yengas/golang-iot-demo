package memory

import (
	"iot-demo/pkg/registry"
	"time"
)

// Registry is non threadsafe repository implementation for registry package
type Registry struct {
	id int
	devices map[int]*registry.Device
}

func (r *Registry) Register(serialNumber string, firmwareVersion string, registrationDate time.Time) (*registry.Device, error) {
	r.id += 1
	device := registry.Device{
		ID:               r.id,
		SerialNumber:     serialNumber,
		FirmwareVersion:  firmwareVersion,
		RegistrationDate: registrationDate,
	}
	r.devices[r.id] = &device
	return &device, nil
}

func (r *Registry) Get(id int) (*registry.Device, bool) {
	device, ok := r.devices[id]
	return device, ok
}

func NewRegistry() *Registry {
	devices := make(map[int]*registry.Device)
	return &Registry{id: 0, devices: devices}
}
