// This module contains a collection of Cisco specific YANG type
// definitions for Layer 2 Bridging.
// 
// Terms and Acronyms
//   BD : Bridge Domain
// 
//   DAI : Dynamic ARP Inspection
// 
//   DHCP : Dynamic Host Configuration Protocol
// 
//   IGMP :  Internet Group Management Protocol
// 
//   IPSG : IP Source Guard
// 
//   MLD : Multicast Listener Discovery
// 
package cisco_bridge_common

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_bridge_common"))
}

type NotifSyslog struct {
}

func (id NotifSyslog) String() string {
	return "cisco-bridge-common:notif-syslog"
}

type NotifSnmpTrap struct {
}

func (id NotifSnmpTrap) String() string {
	return "cisco-bridge-common:notif-snmp-trap"
}

type NotifNone struct {
}

func (id NotifNone) String() string {
	return "cisco-bridge-common:notif-none"
}

type MacLimitNotificationType struct {
}

func (id MacLimitNotificationType) String() string {
	return "cisco-bridge-common:mac-limit-notification-type"
}

type NotifSyslogAndSnmpTrap struct {
}

func (id NotifSyslogAndSnmpTrap) String() string {
	return "cisco-bridge-common:notif-syslog-and-snmp-trap"
}

// MacLimitAction represents Actions to be taken once mac limit threshold is exceeded.
type MacLimitAction string

const (
    // No action
    MacLimitAction_none MacLimitAction = "none"

    // Stop mac learning and flood unknown unicast traffic.
    MacLimitAction_flood MacLimitAction = "flood"

    // Stop mac learning and drop unknown unicast traffic.
    MacLimitAction_drop MacLimitAction = "drop"

    // Bring down operational status of the interface.
    MacLimitAction_shutdown MacLimitAction = "shutdown"
)

// EthTrafficClass represents Traffic class for layer 2 ethernet transport
type EthTrafficClass string

const (
    // Ethernet frames with destination mac-address 
    // eqaul to FFFF.FFFF.FFFF
    EthTrafficClass_broadcast EthTrafficClass = "broadcast"

    // Ethernet frame with destination MAC address not equal
    // to the broadcast address, but with the multicast bit set
    // to 1.
    EthTrafficClass_multicast EthTrafficClass = "multicast"

    // Ethernet frames with with a packet destination MAC
    // address not yet learned.
    EthTrafficClass_unknown_unicast EthTrafficClass = "unknown-unicast"
)

// MacAgingType represents MAC aging mechanism.
type MacAgingType string

const (
    // Dynamically learnt MAC entries are aged out after
    // configured aging time only if no data traffic is 
    // detected during aging period.
    MacAgingType_inactivity MacAgingType = "inactivity"

    // Dynamically learnt MAC entries are aged out after 
    // configured aging time.
    MacAgingType_absolute MacAgingType = "absolute"
)

// MacSecureAction represents Actions to be taken upon mac secure violation.
type MacSecureAction string

const (
    // Forward the violating packet and allow the MAC to be
    // relearned.
    MacSecureAction_none MacSecureAction = "none"

    // Drop violating packet.
    MacSecureAction_restrict MacSecureAction = "restrict"

    // Force shutdown the violating bridge port.
    MacSecureAction_shutdown MacSecureAction = "shutdown"
)

