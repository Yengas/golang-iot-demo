package auth

type Token string

type DeviceCredential struct {
	DeviceID int `json:"id"`
}
