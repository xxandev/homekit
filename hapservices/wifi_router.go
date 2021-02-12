package hapservices

import (
	"github.com/alpr777/homekit/hapcharacteristics"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//TypeWiFiRouter - 0000020A-0000-1000-8000-0026BB765291
const TypeWiFiRouter string = "20A"

//WiFiRouter (+ConfiguredName, +ManagedNetworkEnable, +NetworkAccessViolationControl, +NetworkClientProfileControl,
//+NetworkClientStatusControl, +RouterStatus, +SupportedRouterConfiguration, +WANConfigurationList, +WANStatusList)
type WiFiRouter struct {
	*service.Service
	ConfiguredName                *characteristic.ConfiguredName
	ManagedNetworkEnable          *hapcharacteristics.ManagedNetworkEnable
	NetworkAccessViolationControl *hapcharacteristics.NetworkAccessViolationControl
	NetworkClientProfileControl   *hapcharacteristics.NetworkClientProfileControl
	NetworkClientStatusControl    *hapcharacteristics.NetworkClientStatusControl
	RouterStatus                  *hapcharacteristics.RouterStatus
	SupportedRouterConfiguration  *hapcharacteristics.SupportedRouterConfiguration
	WANConfigurationList          *hapcharacteristics.WANConfigurationList
	WANStatusList                 *hapcharacteristics.WANStatusList
}

//NewWiFiRouter return *WiFiRouter
func NewWiFiRouter() *WiFiRouter {
	svc := WiFiRouter{}
	svc.Service = service.New(TypeWiFiRouter)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddCharacteristic(svc.ConfiguredName.Characteristic)

	svc.ManagedNetworkEnable = hapcharacteristics.NewManagedNetworkEnable()
	svc.AddCharacteristic(svc.ManagedNetworkEnable.Characteristic)

	svc.NetworkAccessViolationControl = hapcharacteristics.NewNetworkAccessViolationControl()
	svc.AddCharacteristic(svc.NetworkAccessViolationControl.Characteristic)

	svc.NetworkClientProfileControl = hapcharacteristics.NewNetworkClientProfileControl()
	svc.AddCharacteristic(svc.NetworkClientProfileControl.Characteristic)

	svc.NetworkClientStatusControl = hapcharacteristics.NewNetworkClientStatusControl()
	svc.AddCharacteristic(svc.NetworkClientStatusControl.Characteristic)

	svc.RouterStatus = hapcharacteristics.NewRouterStatus()
	svc.AddCharacteristic(svc.RouterStatus.Characteristic)

	svc.SupportedRouterConfiguration = hapcharacteristics.NewSupportedRouterConfiguration()
	svc.AddCharacteristic(svc.SupportedRouterConfiguration.Characteristic)

	svc.WANConfigurationList = hapcharacteristics.NewWANConfigurationList()
	svc.AddCharacteristic(svc.WANConfigurationList.Characteristic)

	svc.WANStatusList = hapcharacteristics.NewWANStatusList()
	svc.AddCharacteristic(svc.WANStatusList.Characteristic)

	return &svc
}
