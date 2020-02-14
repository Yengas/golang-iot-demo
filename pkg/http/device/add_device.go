package http_device

import (
	"github.com/gin-gonic/gin"
	add_device "iot-demo/pkg/device/add-device"
	"iot-demo/pkg/device/registry"
	"time"
)

type deviceRegisterRequestDTO struct {
	SerialNumber    string `json:"serial_number" example:"TEST-123"`
	FirmwareVersion string `json:"firmware_version" example:"1.0.0-1"`
}

type deviceRegisterResponseDTO struct {
	Device *registry.Device `json:"device"`
	Token  string           `json:"token"`
}

const (
	invalidRegisterRequestMessage = "please supply serial_number and firmware_version"
)

// DeviceRegistry godoc
// @Summary register a new device
// @Description register a new device with the given parameters
// @Tags device
// @Accept json
// @Produce json
// @Param device body http_device.deviceRegisterRequestDTO true "info of the device to register"
// @Success 201 {object} http_device.deviceRegisterResponseDTO "created new device"
// @Failure 400 {string} string "invalid request parameters"
// @Failure 500 {string} string "unexpected error occurred"
// @Router /device [post]
func makeAddDeviceHandler(addDevice *add_device.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestDTO deviceRegisterRequestDTO
		if err := c.BindJSON(&requestDTO); err != nil {
			c.String(400, invalidRegisterRequestMessage)
			return
		}

		device, token, err := addDevice.Register(requestDTO.SerialNumber, requestDTO.FirmwareVersion, time.Now())
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, deviceRegisterResponseDTO{
			Device: device,
			Token:  string(token),
		})
		return
	}
}

type Handlers struct {
	Service *add_device.Service
}

func (drh Handlers) Register(engine *gin.Engine) {
	engine.POST("/device", makeAddDeviceHandler(drh.Service))
}
