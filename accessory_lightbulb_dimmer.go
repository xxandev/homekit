package homekit

import (
	"github.com/alpr777/homekit/hapservices"
	"github.com/brutella/hc/accessory"
)

//AccessoryLightbulbDimmer struct
type AccessoryLightbulbDimmer struct {
	*accessory.Accessory
	LightbulbDimmer *hapservices.LightbulbDimmer
}

//NewAccessoryLightbulbDimmer return AccessoryLightbulbDimmer (args... are not used)
func NewAccessoryLightbulbDimmer(info accessory.Info, args ...interface{}) *AccessoryLightbulbDimmer {
	acc := AccessoryLightbulbDimmer{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbDimmer = hapservices.NewLightbulbDimmer()
	acc.AddService(acc.LightbulbDimmer.Service)
	return &acc
}
