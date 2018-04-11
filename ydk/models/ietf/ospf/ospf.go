// This YANG module defines the generic configuration
// data for OSPF, which is common across all of the vendor
// implementations of the protocol. It is intended that the module
// will be extended by vendors to define vendor-specific
// OSPF configuration parameters and policies,
// for example route maps or route policies.
// 
// Terms and Acronyms
// 
// OSPF (ospf): Open Shortest Path First
// 
// IP (ip): Internet Protocol
// 
// IPv4 (ipv4):Internet Protocol Version 4
// 
// IPv6 (ipv6): Internet Protocol Version 6
// 
// MTU (mtu) Maximum Transmission Unit
// 
package ospf

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ospf"))
}

type Ospf struct {
}

func (id Ospf) String() string {
	return "ietf-ospf:ospf"
}

type Ospfv2 struct {
}

func (id Ospfv2) String() string {
	return "ietf-ospf:ospfv2"
}

type Ospfv3 struct {
}

func (id Ospfv3) String() string {
	return "ietf-ospf:ospfv3"
}

type OperationMode struct {
}

func (id OperationMode) String() string {
	return "ietf-ospf:operation-mode"
}

type ShipsInTheNight struct {
}

func (id ShipsInTheNight) String() string {
	return "ietf-ospf:ships-in-the-night"
}

type AreaType struct {
}

func (id AreaType) String() string {
	return "ietf-ospf:area-type"
}

type Normal struct {
}

func (id Normal) String() string {
	return "ietf-ospf:normal"
}

type Stub struct {
}

func (id Stub) String() string {
	return "ietf-ospf:stub"
}

type Nssa struct {
}

func (id Nssa) String() string {
	return "ietf-ospf:nssa"
}

type IfLinkType struct {
}

func (id IfLinkType) String() string {
	return "ietf-ospf:if-link-type"
}

type IfLinkTypeNormal struct {
}

func (id IfLinkTypeNormal) String() string {
	return "ietf-ospf:if-link-type-normal"
}

type IfLinkTypeVirtualLink struct {
}

func (id IfLinkTypeVirtualLink) String() string {
	return "ietf-ospf:if-link-type-virtual-link"
}

type IfLinkTypeShamLink struct {
}

func (id IfLinkTypeShamLink) String() string {
	return "ietf-ospf:if-link-type-sham-link"
}

// IfStateType represents OSPF interface state type.
type IfStateType string

const (
    // Interface down state
    IfStateType_Down IfStateType = "Down"

    // Interface loopback state
    IfStateType_Loopback IfStateType = "Loopback"

    // Interface waiting state
    IfStateType_Waiting IfStateType = "Waiting"

    // Interface point-to-point state
    IfStateType_Point_to_Point IfStateType = "Point-to-Point"

    // Interface Designated Router (DR) state
    IfStateType_DR IfStateType = "DR"

    // Interface Backup Designated Router (BDR) state
    IfStateType_BDR IfStateType = "BDR"

    // Interface Other Designated Router state
    IfStateType_DR_Other IfStateType = "DR-Other"
)

// NbrStateType represents OSPF neighbor state type.
type NbrStateType string

const (
    // Neighbor down state
    NbrStateType_Down NbrStateType = "Down"

    // Neighbor attempt state
    NbrStateType_Attempt NbrStateType = "Attempt"

    // Neighbor init state
    NbrStateType_Init NbrStateType = "Init"

    // Neighbor 2-Way state
    NbrStateType_Y_2_Way NbrStateType = "2-Way"

    // Neighbor exchange start state
    NbrStateType_ExStart NbrStateType = "ExStart"

    // Neighbor exchange state
    NbrStateType_Exchange NbrStateType = "Exchange"

    // Neighbor loading state
    NbrStateType_Loading NbrStateType = "Loading"

    // Neighbor full state
    NbrStateType_Full NbrStateType = "Full"
)

// RestartHelperStatusType represents Restart helper status type.
type RestartHelperStatusType string

const (
    // Restart helper status not helping.
    RestartHelperStatusType_Not_Helping RestartHelperStatusType = "Not-Helping"

    // Restart helper status helping.
    RestartHelperStatusType_Helping RestartHelperStatusType = "Helping"
)

// RestartExitReasonType represents as a helper.
type RestartExitReasonType string

const (
    // Not attempted.
    RestartExitReasonType_None RestartExitReasonType = "None"

    // Restart in progress.
    RestartExitReasonType_InProgress RestartExitReasonType = "InProgress"

    // Successfully completed.
    RestartExitReasonType_Completed RestartExitReasonType = "Completed"

    // Timed out.
    RestartExitReasonType_TimedOut RestartExitReasonType = "TimedOut"

    // Aborted due to topology change.
    RestartExitReasonType_TopologyChanged RestartExitReasonType = "TopologyChanged"
)

// PacketType represents OSPF packet type.
type PacketType string

const (
    // OSPF hello packet.
    PacketType_Hello PacketType = "Hello"

    // OSPF database description packet.
    PacketType_Database_Descripton PacketType = "Database-Descripton"

    // OSPF link state request packet.
    PacketType_Link_State_Request PacketType = "Link-State-Request"

    // OSPF link state update packet.
    PacketType_Link_State_Update PacketType = "Link-State-Update"

    // OSPF link state acknowlegement packet.
    PacketType_Link_State_Ack PacketType = "Link-State-Ack"
)

// NssaTranslatorStateType represents OSPF NSSA translator state type.
type NssaTranslatorStateType string

const (
    // NSSA translator enabled state.
    NssaTranslatorStateType_Enabled NssaTranslatorStateType = "Enabled"

    // NSSA translator elected state.
    NssaTranslatorStateType_Elected NssaTranslatorStateType = "Elected"

    // NSSA translator disabled state.
    NssaTranslatorStateType_Disabled NssaTranslatorStateType = "Disabled"
)

// RestartStatusType represents OSPF graceful restart status type.
type RestartStatusType string

const (
    // Router is not restarting.
    RestartStatusType_Not_Restarting RestartStatusType = "Not-Restarting"

    // Router is going through planned restart.
    RestartStatusType_Planned_Restart RestartStatusType = "Planned-Restart"

    // Router is going through unplanned restart.
    RestartStatusType_Unplanned_Restart RestartStatusType = "Unplanned-Restart"
)

