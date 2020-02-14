package registry

import "time"

type Device struct {
	ID               int       `json:"id"`
	SerialNumber     string    `json:"serial_number" example:"TEST-123"`
	FirmwareVersion  string    `json:"firmware_version" example:"1.0.0-1"`
	RegistrationDate time.Time `json:"registration_date" format:"date-time" example:"2017-07-21T17:32:28Z"`
}
