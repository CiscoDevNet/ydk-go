// This module defines types related to the LLDP protocol model.
package lldp_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lldp_types"))
}

type LLDPSYSTEMCAPABILITY struct {
}

func (id LLDPSYSTEMCAPABILITY) String() string {
	return "openconfig-lldp-types:LLDP_SYSTEM_CAPABILITY"
}

type OTHER struct {
}

func (id OTHER) String() string {
	return "openconfig-lldp-types:OTHER"
}

type REPEATER struct {
}

func (id REPEATER) String() string {
	return "openconfig-lldp-types:REPEATER"
}

type MACBRIDGE struct {
}

func (id MACBRIDGE) String() string {
	return "openconfig-lldp-types:MAC_BRIDGE"
}

type WLANACCESSPOINT struct {
}

func (id WLANACCESSPOINT) String() string {
	return "openconfig-lldp-types:WLAN_ACCESS_POINT"
}

type ROUTER struct {
}

func (id ROUTER) String() string {
	return "openconfig-lldp-types:ROUTER"
}

type TELEPHONE struct {
}

func (id TELEPHONE) String() string {
	return "openconfig-lldp-types:TELEPHONE"
}

type DOCSISCABLEDEVICE struct {
}

func (id DOCSISCABLEDEVICE) String() string {
	return "openconfig-lldp-types:DOCSIS_CABLE_DEVICE"
}

type STATIONONLY struct {
}

func (id STATIONONLY) String() string {
	return "openconfig-lldp-types:STATION_ONLY"
}

type CVLAN struct {
}

func (id CVLAN) String() string {
	return "openconfig-lldp-types:C_VLAN"
}

type SVLAN struct {
}

func (id SVLAN) String() string {
	return "openconfig-lldp-types:S_VLAN"
}

type TWOPORTMACRELAY struct {
}

func (id TWOPORTMACRELAY) String() string {
	return "openconfig-lldp-types:TWO_PORT_MAC_RELAY"
}

type LLDPTLV struct {
}

func (id LLDPTLV) String() string {
	return "openconfig-lldp-types:LLDP_TLV"
}

type CHASSISID struct {
}

func (id CHASSISID) String() string {
	return "openconfig-lldp-types:CHASSIS_ID"
}

type PORTID struct {
}

func (id PORTID) String() string {
	return "openconfig-lldp-types:PORT_ID"
}

type PORTDESCRIPTION struct {
}

func (id PORTDESCRIPTION) String() string {
	return "openconfig-lldp-types:PORT_DESCRIPTION"
}

type SYSTEMNAME struct {
}

func (id SYSTEMNAME) String() string {
	return "openconfig-lldp-types:SYSTEM_NAME"
}

type SYSTEMDESCRIPTION struct {
}

func (id SYSTEMDESCRIPTION) String() string {
	return "openconfig-lldp-types:SYSTEM_DESCRIPTION"
}

type SYSTEMCAPABILITIES struct {
}

func (id SYSTEMCAPABILITIES) String() string {
	return "openconfig-lldp-types:SYSTEM_CAPABILITIES"
}

type MANAGEMENTADDRESS struct {
}

func (id MANAGEMENTADDRESS) String() string {
	return "openconfig-lldp-types:MANAGEMENT_ADDRESS"
}

// ChassisIdType represents the chassis identifier
type ChassisIdType string

const (
    // Chassis identifier based on the value of entPhysicalAlias
    // object defined in IETF RFC 2737
    ChassisIdType_CHASSIS_COMPONENT ChassisIdType = "CHASSIS_COMPONENT"

    // Chassis identifier based on the value of ifAlias object
    // defined in IETF RFC 2863
    ChassisIdType_INTERFACE_ALIAS ChassisIdType = "INTERFACE_ALIAS"

    // Chassis identifier based on the value of entPhysicalAlias
    // object defined in IETF RFC 2737 for a port or backplane
    // component
    ChassisIdType_PORT_COMPONENT ChassisIdType = "PORT_COMPONENT"

    // Chassis identifier based on the value of a unicast source
    // address (encoded in network byte order and IEEE 802.3
    // canonical bit order), of a port on the containing chassis
    // as defined in IEEE Std 802-2001
    ChassisIdType_MAC_ADDRESS ChassisIdType = "MAC_ADDRESS"

    // Chassis identifier based on a network address,
    // associated with a particular chassis.  The encoded address
    // is composed of two fields.  The first field is a single
    // octet, representing the IANA AddressFamilyNumbers value
    // for the specific address type, and the second field is the
    // network address value
    ChassisIdType_NETWORK_ADDRESS ChassisIdType = "NETWORK_ADDRESS"

    // Chassis identifier based on the name of the interface,
    // e.g., the value of ifName object defined in IETF RFC 2863
    ChassisIdType_INTERFACE_NAME ChassisIdType = "INTERFACE_NAME"

    // Chassis identifier based on a locally defined value
    ChassisIdType_LOCAL ChassisIdType = "LOCAL"
)

// PortIdType represents the port identifier
type PortIdType string

const (
    // Chassis identifier based on the value of ifAlias object
    // defined in IETF RFC 2863
    PortIdType_INTERFACE_ALIAS PortIdType = "INTERFACE_ALIAS"

    // Port identifier based on the value of entPhysicalAlias
    // object defined in IETF RFC 2737 for a port component
    PortIdType_PORT_COMPONENT PortIdType = "PORT_COMPONENT"

    // Port identifier based on the value of a unicast source
    // address (encoded in network byte order and IEEE 802.3
    // canonical bit order) associated with a port
    PortIdType_MAC_ADDRESS PortIdType = "MAC_ADDRESS"

    // Port identifier based on a network address,
    // associated with a particular port
    PortIdType_NETWORK_ADDRESS PortIdType = "NETWORK_ADDRESS"

    // Port identifier based on the name of the interface,
    // e.g., the value of ifName object defined in IETF RFC 2863
    PortIdType_INTERFACE_NAME PortIdType = "INTERFACE_NAME"

    // Port identifer based on the circuit id in the DHCP
    // relay agent information option as defined in IETF
    // RFC 3046
    PortIdType_AGENT_CIRCUIT_ID PortIdType = "AGENT_CIRCUIT_ID"

    // Port identifier based on a locally defined alphanumeric
    // string
    PortIdType_LOCAL PortIdType = "LOCAL"
)

