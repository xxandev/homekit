package haps

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
	hapc "github.com/xxandev/homekit/hap-characteristic"
)

//TypeWiFiRouter - 0000020A-0000-1000-8000-0026BB765291
const TypeWiFiRouter string = "20A"

//WiFiRouter
//	◈ ConfiguredName
//	◈ ManagedNetworkEnable
//	◈ NetworkAccessViolationControl
//	◈ NetworkClientProfileControl
//	◈ NetworkClientStatusControl
//	◈ RouterStatus
//	◈ SupportedRouterConfiguration
//	◈ WANConfigurationList
//	◈ WANStatusList
type WiFiRouter struct {
	*service.Service
	ConfiguredName                *characteristic.ConfiguredName
	ManagedNetworkEnable          *hapc.ManagedNetworkEnable
	NetworkAccessViolationControl *hapc.NetworkAccessViolationControl
	NetworkClientProfileControl   *hapc.NetworkClientProfileControl
	NetworkClientStatusControl    *hapc.NetworkClientStatusControl
	RouterStatus                  *hapc.RouterStatus
	SupportedRouterConfiguration  *hapc.SupportedRouterConfiguration
	WANConfigurationList          *hapc.WANConfigurationList
	WANStatusList                 *hapc.WANStatusList
}

//NewWiFiRouter return *WiFiRouter
func NewWiFiRouter() *WiFiRouter {
	svc := WiFiRouter{}
	svc.Service = service.New(TypeWiFiRouter)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddCharacteristic(svc.ConfiguredName.Characteristic)

	svc.ManagedNetworkEnable = hapc.NewManagedNetworkEnable()
	svc.AddCharacteristic(svc.ManagedNetworkEnable.Characteristic)

	svc.NetworkAccessViolationControl = hapc.NewNetworkAccessViolationControl()
	svc.AddCharacteristic(svc.NetworkAccessViolationControl.Characteristic)

	svc.NetworkClientProfileControl = hapc.NewNetworkClientProfileControl()
	svc.AddCharacteristic(svc.NetworkClientProfileControl.Characteristic)

	svc.NetworkClientStatusControl = hapc.NewNetworkClientStatusControl()
	svc.AddCharacteristic(svc.NetworkClientStatusControl.Characteristic)

	svc.RouterStatus = hapc.NewRouterStatus()
	svc.AddCharacteristic(svc.RouterStatus.Characteristic)

	svc.SupportedRouterConfiguration = hapc.NewSupportedRouterConfiguration()
	svc.AddCharacteristic(svc.SupportedRouterConfiguration.Characteristic)

	svc.WANConfigurationList = hapc.NewWANConfigurationList()
	svc.AddCharacteristic(svc.WANConfigurationList.Characteristic)

	svc.WANStatusList = hapc.NewWANStatusList()
	svc.AddCharacteristic(svc.WANStatusList.Characteristic)

	return &svc
}
