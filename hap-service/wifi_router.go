package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
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
	*service.S
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
	svc.S = service.New(TypeWiFiRouter)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddC(svc.ConfiguredName.C)

	svc.ManagedNetworkEnable = hapc.NewManagedNetworkEnable()
	svc.AddC(svc.ManagedNetworkEnable.C)

	svc.NetworkAccessViolationControl = hapc.NewNetworkAccessViolationControl()
	svc.AddC(svc.NetworkAccessViolationControl.C)

	svc.NetworkClientProfileControl = hapc.NewNetworkClientProfileControl()
	svc.AddC(svc.NetworkClientProfileControl.C)

	svc.NetworkClientStatusControl = hapc.NewNetworkClientStatusControl()
	svc.AddC(svc.NetworkClientStatusControl.C)

	svc.RouterStatus = hapc.NewRouterStatus()
	svc.AddC(svc.RouterStatus.C)

	svc.SupportedRouterConfiguration = hapc.NewSupportedRouterConfiguration()
	svc.AddC(svc.SupportedRouterConfiguration.C)

	svc.WANConfigurationList = hapc.NewWANConfigurationList()
	svc.AddC(svc.WANConfigurationList.C)

	svc.WANStatusList = hapc.NewWANStatusList()
	svc.AddC(svc.WANStatusList.C)

	return &svc
}
