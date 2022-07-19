package hapc

import "github.com/brutella/hap/characteristic"

const (
	//RouterStatusRead - read(0)
	RouterStatusRead int = 0

	//RouterStatusNotRead - not read(1)
	RouterStatusNotRead int = 1
)

//TypeRouterStatus - 0000020E-0000-1000-8000-0026BB765291
const TypeRouterStatus string = "20E"

//RouterStatus - Formats UINT8
type RouterStatus struct {
	*characteristic.Int
}

//NewRouterStatus return *RouterStatus
func NewRouterStatus() *RouterStatus {
	char := characteristic.NewInt(TypeRouterStatus)
	char.Format = characteristic.FormatUInt8
	char.Permissions = []string{
		characteristic.PermissionRead,
		characteristic.PermissionEvents,
	}

	char.SetMinValue(0)
	char.SetMaxValue(1)
	char.SetStepValue(1)
	char.SetValue(0)

	return &RouterStatus{char}
}
