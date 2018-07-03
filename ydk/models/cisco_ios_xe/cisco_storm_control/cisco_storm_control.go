// This module defines data model for strom control feature.
// 
// Traffic storm occurs when packets flood a bridge, creating
// excessive traffic and degrading network performance. Traffic
// storm control prevents bridge disruption by suppressing traffic
// when the number of packets reaches configured threshold
// levels.
package cisco_storm_control

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_storm_control"))
}

type ActionShutdown struct {
}

func (id ActionShutdown) String() string {
	return "cisco-storm-control:action-shutdown"
}

type ActionSnmpTrap struct {
}

func (id ActionSnmpTrap) String() string {
	return "cisco-storm-control:action-snmp-trap"
}

type StormControlAction struct {
}

func (id StormControlAction) String() string {
	return "cisco-storm-control:storm-control-action"
}

type ActionDrop struct {
}

func (id ActionDrop) String() string {
	return "cisco-storm-control:action-drop"
}

