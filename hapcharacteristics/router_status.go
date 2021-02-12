package hapcharacteristics

import "github.com/brutella/hc/characteristic"

const (
	//RouterStatusRead -
	RouterStatusRead int = 0
	//RouterStatusNotRead -
	RouterStatusNotRead int = 1
)

//TypeRouterStatus -
const TypeRouterStatus string = "20E"

//RouterStatus -
type RouterStatus struct {
	*characteristic.Int
}

//NewRouterStatus -
func NewRouterStatus() *RouterStatus {
	char := characteristic.NewInt(TypeRouterStatus)
	char.Format = characteristic.FormatUInt8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetMinValue(0)
	char.SetMaxValue(1)
	char.SetStepValue(1)
	char.SetValue(0)

	return &RouterStatus{char}
}
