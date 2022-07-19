package hapc

import "github.com/brutella/hap/characteristic"

//TypeCurrentTransport - 0000021E-0000-1000-8000-0000022B
const TypeCurrentTransport = "22B"

//CurrentTransport - Formats BOOL
type CurrentTransport struct {
	*characteristic.Bool
}

//NewCurrentTransport return *CurrentTransport
func NewCurrentTransport() *CurrentTransport {
	char := characteristic.NewBool(TypeCurrentTransport)
	char.Format = characteristic.FormatBool
	char.Permissions = []string{characteristic.PermissionRead}

	char.SetValue(false)

	return &CurrentTransport{char}
}
