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

type OPENCONFIGHARDWARECOMPONENT struct {
}

func (id OPENCONFIGHARDWARECOMPONENT) String() string {
	return "openconfig-platform-types:OPENCONFIG_HARDWARE_COMPONENT"
}

type OPENCONFIGSOFTWARECOMPONENT struct {
}

func (id OPENCONFIGSOFTWARECOMPONENT) String() string {
	return "openconfig-platform-types:OPENCONFIG_SOFTWARE_COMPONENT"
}

type CHASSIS struct {
}

func (id CHASSIS) String() string {
	return "openconfig-platform-types:CHASSIS"
}

type BACKPLANE struct {
}

func (id BACKPLANE) String() string {
	return "openconfig-platform-types:BACKPLANE"
}

type POWERSUPPLY struct {
}

func (id POWERSUPPLY) String() string {
	return "openconfig-platform-types:POWER_SUPPLY"
}

type FAN struct {
}

func (id FAN) String() string {
	return "openconfig-platform-types:FAN"
}

type SENSOR struct {
}

func (id SENSOR) String() string {
	return "openconfig-platform-types:SENSOR"
}

type MODULE struct {
}

func (id MODULE) String() string {
	return "openconfig-platform-types:MODULE"
}

type LINECARD struct {
}

func (id LINECARD) String() string {
	return "openconfig-platform-types:LINECARD"
}

type PORT struct {
}

func (id PORT) String() string {
	return "openconfig-platform-types:PORT"
}

type TRANSCEIVER struct {
}

func (id TRANSCEIVER) String() string {
	return "openconfig-platform-types:TRANSCEIVER"
}

type CPU struct {
}

func (id CPU) String() string {
	return "openconfig-platform-types:CPU"
}

type OPERATINGSYSTEM struct {
}

func (id OPERATINGSYSTEM) String() string {
	return "openconfig-platform-types:OPERATING_SYSTEM"
}

