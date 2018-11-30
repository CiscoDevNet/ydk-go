// This module defines data types (e.g., YANG identities)
// to support the OpenConfig component inventory model.
package platform_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package platform_types"))
}

type OPERATINGSYSTEM struct {
}

func (id OPERATINGSYSTEM) String() string {
	return "openconfig-platform-types:OPERATING_SYSTEM"
}

type LINECARD struct {
}

func (id LINECARD) String() string {
	return "openconfig-platform-types:LINECARD"
}

type OPENCONFIGHARDWARECOMPONENT struct {
}

func (id OPENCONFIGHARDWARECOMPONENT) String() string {
	return "openconfig-platform-types:OPENCONFIG_HARDWARE_COMPONENT"
}

type MODULE struct {
}

func (id MODULE) String() string {
	return "openconfig-platform-types:MODULE"
}

type CPU struct {
}

func (id CPU) String() string {
	return "openconfig-platform-types:CPU"
}

type TRANSCEIVER struct {
}

func (id TRANSCEIVER) String() string {
	return "openconfig-platform-types:TRANSCEIVER"
}

type CHASSIS struct {
}

func (id CHASSIS) String() string {
	return "openconfig-platform-types:CHASSIS"
}

type FAN struct {
}

func (id FAN) String() string {
	return "openconfig-platform-types:FAN"
}

type BACKPLANE struct {
}

func (id BACKPLANE) String() string {
	return "openconfig-platform-types:BACKPLANE"
}

type OPENCONFIGSOFTWARECOMPONENT struct {
}

func (id OPENCONFIGSOFTWARECOMPONENT) String() string {
	return "openconfig-platform-types:OPENCONFIG_SOFTWARE_COMPONENT"
}

type SENSOR struct {
}

func (id SENSOR) String() string {
	return "openconfig-platform-types:SENSOR"
}

type PORT struct {
}

func (id PORT) String() string {
	return "openconfig-platform-types:PORT"
}

type POWERSUPPLY struct {
}

func (id POWERSUPPLY) String() string {
	return "openconfig-platform-types:POWER_SUPPLY"
}

