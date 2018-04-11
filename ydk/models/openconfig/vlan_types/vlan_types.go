// This module defines configuration and state variables for VLANs,
// in addition to VLAN parameters associated with interfaces
package vlan_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package vlan_types"))
}

type TPIDTYPES struct {
}

func (id TPIDTYPES) String() string {
	return "openconfig-vlan-types:TPID_TYPES"
}

type TPID0x8100 struct {
}

func (id TPID0x8100) String() string {
	return "openconfig-vlan-types:TPID_0x8100"
}

type TPID0x8A88 struct {
}

func (id TPID0x8A88) String() string {
	return "openconfig-vlan-types:TPID_0x8A88"
}

type TPID0x9100 struct {
}

func (id TPID0x9100) String() string {
	return "openconfig-vlan-types:TPID_0x9100"
}

type TPID0X9200 struct {
}

func (id TPID0X9200) String() string {
	return "openconfig-vlan-types:TPID_0X9200"
}

// VlanModeType represents VLAN interface mode (trunk or access)
type VlanModeType string

const (
    // Access mode VLAN interface (No 802.1q header)
    VlanModeType_ACCESS VlanModeType = "ACCESS"

    // Trunk mode VLAN interface
    VlanModeType_TRUNK VlanModeType = "TRUNK"
)

