// Model for managing Ethernet interfaces -- augments the IETF YANG
// model for interfaces described by RFC 7223
package if_ethernet

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package if_ethernet"))
}

type ETHERNETSPEED struct {
}

func (id ETHERNETSPEED) String() string {
	return "openconfig-if-ethernet:ETHERNET_SPEED"
}

type SPEED10MB struct {
}

func (id SPEED10MB) String() string {
	return "openconfig-if-ethernet:SPEED_10MB"
}

type SPEED100MB struct {
}

func (id SPEED100MB) String() string {
	return "openconfig-if-ethernet:SPEED_100MB"
}

type SPEED1GB struct {
}

func (id SPEED1GB) String() string {
	return "openconfig-if-ethernet:SPEED_1GB"
}

type SPEED10GB struct {
}

func (id SPEED10GB) String() string {
	return "openconfig-if-ethernet:SPEED_10GB"
}

type SPEED25GB struct {
}

func (id SPEED25GB) String() string {
	return "openconfig-if-ethernet:SPEED_25GB"
}

type SPEED40GB struct {
}

func (id SPEED40GB) String() string {
	return "openconfig-if-ethernet:SPEED_40GB"
}

type SPEED50GB struct {
}

func (id SPEED50GB) String() string {
	return "openconfig-if-ethernet:SPEED_50GB"
}

type SPEED100GB struct {
}

func (id SPEED100GB) String() string {
	return "openconfig-if-ethernet:SPEED_100GB"
}

type SPEEDUNKNOWN struct {
}

func (id SPEEDUNKNOWN) String() string {
	return "openconfig-if-ethernet:SPEED_UNKNOWN"
}

