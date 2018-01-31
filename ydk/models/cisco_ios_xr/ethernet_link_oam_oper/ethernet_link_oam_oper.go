// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-link-oam package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ether-link-oam: Ethernet Link OAM operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_link_oam_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_link_oam_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ethernet-link-oam-oper ether-link-oam}", reflect.TypeOf(EtherLinkOam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ethernet-link-oam-oper:ether-link-oam", reflect.TypeOf(EtherLinkOam{}))
}

// Log represents The type of a log entry
type Log string

const (
    // Log entry for an errored symbol event
    Log_log_type_symbol_event Log = "log-type-symbol-event"

    // Log entry for an errored frame period event
    Log_log_type_period_event Log = "log-type-period-event"

    // Log entry for an errored frame event
    Log_log_type_frame_event Log = "log-type-frame-event"

    // Log entry for an errored frame seconds summary
    // event
    Log_log_type_secs_event Log = "log-type-secs-event"

    // Log entry for a link fault
    Log_log_type_link_fault Log = "log-type-link-fault"

    // Log entry for a dying gasp
    Log_log_type_dying_gasp Log = "log-type-dying-gasp"

    // Log entry for a critical event
    Log_log_type_critical_event Log = "log-type-critical-event"
)

// LogLocation represents The location of the event that caused a log entry
type LogLocation string

const (
    // A local event
    LogLocation_log_location_local LogLocation = "log-location-local"

    // A remote event
    LogLocation_log_location_remote LogLocation = "log-location-remote"
)

// LoopbackStatus represents The loopback mode of an OAM interface
type LoopbackStatus string

const (
    // Loopback is not being performed
    LoopbackStatus_none LoopbackStatus = "none"

    // Initiating master loopback
    LoopbackStatus_initiating LoopbackStatus = "initiating"

    // In master loopback mode
    LoopbackStatus_master_loopback LoopbackStatus = "master-loopback"

    // Terminating master loopback mode
    LoopbackStatus_terminating LoopbackStatus = "terminating"

    // In slave loopback mode
    LoopbackStatus_local_loopback LoopbackStatus = "local-loopback"

    // Parser and multiplexer combination unexpected
    LoopbackStatus_unknown LoopbackStatus = "unknown"
)

// OperationalState represents Operational state of an interface
type OperationalState string

const (
    // 802.3 OAM is disabled
    OperationalState_disabled OperationalState = "disabled"

    // 802.3 OAM has encountered a link fault
    OperationalState_link_fault OperationalState = "link-fault"

    // Passive OAM entity waiting to see if peer is
    // OAM capable
    OperationalState_passive_wait OperationalState = "passive-wait"

    // Active OAM entity trying to determine if peer
    // is OAM capable
    OperationalState_active_send_local OperationalState = "active-send-local"

    // OAM discovered peer but still to accept or
    // reject peer config
    OperationalState_send_local_and_remote OperationalState = "send-local-and-remote"

    // OAM peering is allowed by local device
    OperationalState_send_local_and_remote_ok OperationalState = "send-local-and-remote-ok"

    // OAM peering rejected by local device
    OperationalState_peering_locally_rejected OperationalState = "peering-locally-rejected"

    // OAM peering rejected by remote device
    OperationalState_peering_remotely_rejected OperationalState = "peering-remotely-rejected"

    // 802.3 OAM is operational
    OperationalState_operational OperationalState = "operational"

    // 802.3 OAM is operating in half-duplex mode
    OperationalState_operational_half_duplex OperationalState = "operational-half-duplex"
)

// Mode represents Mode of an OAM interface
type Mode string

const (
    // Passive mode
    Mode_passive Mode = "passive"

    // Active mode
    Mode_active Mode = "active"

    // Don't care what the mode is
    Mode_dont_care Mode = "dont-care"
)

// Action represents Actions supported by an OAM interface
type Action string

const (
    // Disabled (do nothing)
    Action_no_action Action = "no-action"

    // Disable the interface
    Action_disable_interface Action = "disable-interface"

    // Log the event and do nothing else
    Action_log Action = "log"

    // EFD the interface
    Action_efd Action = "efd"
)

// ProtocolState represents The state the protocol is in
type ProtocolState string

const (
    // The protocol is in the INACTIVE state
    ProtocolState_protocol_state_inactive ProtocolState = "protocol-state-inactive"

    // The protocol is in the FAULT state
    ProtocolState_protocol_state_fault ProtocolState = "protocol-state-fault"

    // The protocol is in the ACTIVE_SEND_LOCAL state
    ProtocolState_protocol_state_active_send_local ProtocolState = "protocol-state-active-send-local"

    // The protocol is in the SEND_LOCAL_REMOTE state
    ProtocolState_protocol_state_passive_wait ProtocolState = "protocol-state-passive-wait"

    // The protocol is in the LOCAL_REMOTE state
    ProtocolState_protocol_state_send_local_remote ProtocolState = "protocol-state-send-local-remote"

    // The protocol is in the LOCAL_REMOTE_OK state
    ProtocolState_protocol_state_send_local_remote_ok ProtocolState = "protocol-state-send-local-remote-ok"

    // The protocol is in the SEND_ANY state
    ProtocolState_protocol_state_send_any ProtocolState = "protocol-state-send-any"
)

// EtherLinkOam
// Ethernet Link OAM operational data
type EtherLinkOam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Ethernet Link OAM enabled interfaces within Discovery Info
    // container.
    DiscoveryInfoInterfaces EtherLinkOam_DiscoveryInfoInterfaces

    // Table of Ethernet Link OAM enabled interfaces within Interface State
    // container.
    InterfaceStateInterfaces EtherLinkOam_InterfaceStateInterfaces

    // Table of Ethernet Link OAM enabled interfaces within Running Config
    // container.
    RunningConfigInterfaces EtherLinkOam_RunningConfigInterfaces

    // Node table for node-specific operational data.
    Nodes EtherLinkOam_Nodes

    // Table of Ethernet Link OAM enabled interfaces within Event Log Entry
    // container.
    EventLogEntryInterfaces EtherLinkOam_EventLogEntryInterfaces

    // Table of Ethernet Link OAM enabled interfaces within Stats container.
    StatsInterfaces EtherLinkOam_StatsInterfaces
}

func (etherLinkOam *EtherLinkOam) GetFilter() yfilter.YFilter { return etherLinkOam.YFilter }

func (etherLinkOam *EtherLinkOam) SetFilter(yf yfilter.YFilter) { etherLinkOam.YFilter = yf }

func (etherLinkOam *EtherLinkOam) GetGoName(yname string) string {
    if yname == "discovery-info-interfaces" { return "DiscoveryInfoInterfaces" }
    if yname == "interface-state-interfaces" { return "InterfaceStateInterfaces" }
    if yname == "running-config-interfaces" { return "RunningConfigInterfaces" }
    if yname == "nodes" { return "Nodes" }
    if yname == "event-log-entry-interfaces" { return "EventLogEntryInterfaces" }
    if yname == "stats-interfaces" { return "StatsInterfaces" }
    return ""
}

func (etherLinkOam *EtherLinkOam) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-link-oam-oper:ether-link-oam"
}

func (etherLinkOam *EtherLinkOam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discovery-info-interfaces" {
        return &etherLinkOam.DiscoveryInfoInterfaces
    }
    if childYangName == "interface-state-interfaces" {
        return &etherLinkOam.InterfaceStateInterfaces
    }
    if childYangName == "running-config-interfaces" {
        return &etherLinkOam.RunningConfigInterfaces
    }
    if childYangName == "nodes" {
        return &etherLinkOam.Nodes
    }
    if childYangName == "event-log-entry-interfaces" {
        return &etherLinkOam.EventLogEntryInterfaces
    }
    if childYangName == "stats-interfaces" {
        return &etherLinkOam.StatsInterfaces
    }
    return nil
}

func (etherLinkOam *EtherLinkOam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discovery-info-interfaces"] = &etherLinkOam.DiscoveryInfoInterfaces
    children["interface-state-interfaces"] = &etherLinkOam.InterfaceStateInterfaces
    children["running-config-interfaces"] = &etherLinkOam.RunningConfigInterfaces
    children["nodes"] = &etherLinkOam.Nodes
    children["event-log-entry-interfaces"] = &etherLinkOam.EventLogEntryInterfaces
    children["stats-interfaces"] = &etherLinkOam.StatsInterfaces
    return children
}

func (etherLinkOam *EtherLinkOam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherLinkOam *EtherLinkOam) GetBundleName() string { return "cisco_ios_xr" }

func (etherLinkOam *EtherLinkOam) GetYangName() string { return "ether-link-oam" }

func (etherLinkOam *EtherLinkOam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (etherLinkOam *EtherLinkOam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (etherLinkOam *EtherLinkOam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (etherLinkOam *EtherLinkOam) SetParent(parent types.Entity) { etherLinkOam.parent = parent }

func (etherLinkOam *EtherLinkOam) GetParent() types.Entity { return etherLinkOam.parent }

func (etherLinkOam *EtherLinkOam) GetParentYangName() string { return "Cisco-IOS-XR-ethernet-link-oam-oper" }

// EtherLinkOam_DiscoveryInfoInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Discovery Info container
type EtherLinkOam_DiscoveryInfoInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet Link OAM interface to get Discovery Info for. The type is slice of
    // EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface.
    DiscoveryInfoInterface []EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface
}

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetFilter() yfilter.YFilter { return discoveryInfoInterfaces.YFilter }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) SetFilter(yf yfilter.YFilter) { discoveryInfoInterfaces.YFilter = yf }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetGoName(yname string) string {
    if yname == "discovery-info-interface" { return "DiscoveryInfoInterface" }
    return ""
}

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetSegmentPath() string {
    return "discovery-info-interfaces"
}

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discovery-info-interface" {
        for _, c := range discoveryInfoInterfaces.DiscoveryInfoInterface {
            if discoveryInfoInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface{}
        discoveryInfoInterfaces.DiscoveryInfoInterface = append(discoveryInfoInterfaces.DiscoveryInfoInterface, child)
        return &discoveryInfoInterfaces.DiscoveryInfoInterface[len(discoveryInfoInterfaces.DiscoveryInfoInterface)-1]
    }
    return nil
}

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range discoveryInfoInterfaces.DiscoveryInfoInterface {
        children[discoveryInfoInterfaces.DiscoveryInfoInterface[i].GetSegmentPath()] = &discoveryInfoInterfaces.DiscoveryInfoInterface[i]
    }
    return children
}

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetYangName() string { return "discovery-info-interfaces" }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) SetParent(parent types.Entity) { discoveryInfoInterfaces.parent = parent }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetParent() types.Entity { return discoveryInfoInterfaces.parent }

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetParentYangName() string { return "ether-link-oam" }

// EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface
// Ethernet Link OAM interface to get Discovery
// Info for
type EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    MemberInterface interface{}

    // Interface Name. The type is string.
    Name interface{}

    // Operational status. The type is OperationalState.
    OperationalStatus interface{}

    // The loopback mode the interface is in. The type is LoopbackStatus.
    LoopbackMode interface{}

    // Local Mode (passive/active). The type is Mode.
    LocalMode interface{}

    // Has the interface mis-wired?. The type is bool.
    Miswired interface{}

    // Local Mis-wiring Detection key. The type is interface{} with range:
    // 0..4294967295.
    LocalMwdKey interface{}

    // Local Unidirectional support. The type is bool.
    LocalFunctionUnidirectional interface{}

    // Local loopback support. The type is bool.
    LocalFunctionLoopback interface{}

    // Local event support. The type is bool.
    LocalFunctionEvent interface{}

    // Local variable retreival support. The type is bool.
    LocalFunctionvariable interface{}

    // Local revision. The type is interface{} with range: 0..4294967295.
    LocalRevision interface{}

    // Local MTU. The type is interface{} with range: 0..4294967295.
    LocalMtu interface{}

    // Is the local OAM session operational?. The type is bool.
    LocalOperational interface{}

    // Is the local OAM session evaluating?. The type is bool.
    LocalEvaluating interface{}

    // Remote Mode (passive/active). The type is Mode.
    RemoteMode interface{}

    // Remote unidirectional support. The type is bool.
    RemoteUnidirectional interface{}

    // Remote loopback support. The type is bool.
    RemoteLoopback interface{}

    // Remote event support. The type is bool.
    RemoteEvent interface{}

    // Remote variable retreival support. The type is bool.
    RemoteVariable interface{}

    // Remote MTU. The type is interface{} with range: 0..4294967295.
    RemoteMtu interface{}

    // Remote MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RemoteMacAddress interface{}

    // Remote vendor OUI. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RemoteVendorOui interface{}

    // Remote revision. The type is interface{} with range: 0..65535.
    RemoteRevision interface{}

    // Remote vendor info. The type is interface{} with range: 0..4294967295.
    RemoteVendorInfo interface{}

    // Remote Mis-wiring Detection key. The type is interface{} with range:
    // 0..4294967295.
    RemoteMwdKey interface{}

    // Has this value been received successfully?. The type is bool.
    OperationalStatusValid interface{}

    // Has this value been received successfully?. The type is bool.
    LoopbackModeValid interface{}

    // Has this value been received successfully?. The type is bool.
    LocalModeValid interface{}

    // Has this value been received successfully?. The type is bool.
    MiswiredValid interface{}

    // Has this value been received successfully?. The type is bool.
    LocalMwdKeyValid interface{}

    // Has this value been received successfully?. The type is bool.
    LocalFunctionUnidirectionalValid interface{}

    // Has this value been received successfully?. The type is bool.
    LocalFunctionLoopbackValid interface{}

    // Has this value been received successfully?. The type is bool.
    LocalFunctionEventValid interface{}

    // Has this value been received successfully?. The type is bool.
    LocalFunctionvariableValid interface{}

    // Has this value been received successfully?. The type is bool.
    LocalRevisionvalid interface{}

    // Has this value been received successfully?. The type is bool.
    LocalMtuValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteModeValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteUnidirectionalValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteLoopbackValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteEventValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteVariableValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteMtuValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteMacAddressValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteVendorOuiValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteRevisionvalid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteVendorInfoValid interface{}

    // Has this value been received successfully?. The type is bool.
    RemoteMwdKeyValid interface{}

    // Timestamp of when the last At Risk notification was received (in seconds
    // since the UNIX epoch), or 0 if the peer is not currently at risk. The type
    // is interface{} with range: 0..18446744073709551615. Units are second.
    ReceivedAtRiskNotificationTimestamp interface{}

    // Number of seconds remaining that the peer has indicated it will be At Risk.
    // The type is interface{} with range: 0..65535. Units are second.
    ReceivedAtRiskNotificationTimeRemaining interface{}
}

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetFilter() yfilter.YFilter { return discoveryInfoInterface.YFilter }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) SetFilter(yf yfilter.YFilter) { discoveryInfoInterface.YFilter = yf }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetGoName(yname string) string {
    if yname == "member-interface" { return "MemberInterface" }
    if yname == "name" { return "Name" }
    if yname == "operational-status" { return "OperationalStatus" }
    if yname == "loopback-mode" { return "LoopbackMode" }
    if yname == "local-mode" { return "LocalMode" }
    if yname == "miswired" { return "Miswired" }
    if yname == "local-mwd-key" { return "LocalMwdKey" }
    if yname == "local-function-unidirectional" { return "LocalFunctionUnidirectional" }
    if yname == "local-function-loopback" { return "LocalFunctionLoopback" }
    if yname == "local-function-event" { return "LocalFunctionEvent" }
    if yname == "local-functionvariable" { return "LocalFunctionvariable" }
    if yname == "local-revision" { return "LocalRevision" }
    if yname == "local-mtu" { return "LocalMtu" }
    if yname == "local-operational" { return "LocalOperational" }
    if yname == "local-evaluating" { return "LocalEvaluating" }
    if yname == "remote-mode" { return "RemoteMode" }
    if yname == "remote-unidirectional" { return "RemoteUnidirectional" }
    if yname == "remote-loopback" { return "RemoteLoopback" }
    if yname == "remote-event" { return "RemoteEvent" }
    if yname == "remote-variable" { return "RemoteVariable" }
    if yname == "remote-mtu" { return "RemoteMtu" }
    if yname == "remote-mac-address" { return "RemoteMacAddress" }
    if yname == "remote-vendor-oui" { return "RemoteVendorOui" }
    if yname == "remote-revision" { return "RemoteRevision" }
    if yname == "remote-vendor-info" { return "RemoteVendorInfo" }
    if yname == "remote-mwd-key" { return "RemoteMwdKey" }
    if yname == "operational-status-valid" { return "OperationalStatusValid" }
    if yname == "loopback-mode-valid" { return "LoopbackModeValid" }
    if yname == "local-mode-valid" { return "LocalModeValid" }
    if yname == "miswired-valid" { return "MiswiredValid" }
    if yname == "local-mwd-key-valid" { return "LocalMwdKeyValid" }
    if yname == "local-function-unidirectional-valid" { return "LocalFunctionUnidirectionalValid" }
    if yname == "local-function-loopback-valid" { return "LocalFunctionLoopbackValid" }
    if yname == "local-function-event-valid" { return "LocalFunctionEventValid" }
    if yname == "local-functionvariable-valid" { return "LocalFunctionvariableValid" }
    if yname == "local-revisionvalid" { return "LocalRevisionvalid" }
    if yname == "local-mtu-valid" { return "LocalMtuValid" }
    if yname == "remote-mode-valid" { return "RemoteModeValid" }
    if yname == "remote-unidirectional-valid" { return "RemoteUnidirectionalValid" }
    if yname == "remote-loopback-valid" { return "RemoteLoopbackValid" }
    if yname == "remote-event-valid" { return "RemoteEventValid" }
    if yname == "remote-variable-valid" { return "RemoteVariableValid" }
    if yname == "remote-mtu-valid" { return "RemoteMtuValid" }
    if yname == "remote-mac-address-valid" { return "RemoteMacAddressValid" }
    if yname == "remote-vendor-oui-valid" { return "RemoteVendorOuiValid" }
    if yname == "remote-revisionvalid" { return "RemoteRevisionvalid" }
    if yname == "remote-vendor-info-valid" { return "RemoteVendorInfoValid" }
    if yname == "remote-mwd-key-valid" { return "RemoteMwdKeyValid" }
    if yname == "received-at-risk-notification-timestamp" { return "ReceivedAtRiskNotificationTimestamp" }
    if yname == "received-at-risk-notification-time-remaining" { return "ReceivedAtRiskNotificationTimeRemaining" }
    return ""
}

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetSegmentPath() string {
    return "discovery-info-interface" + "[member-interface='" + fmt.Sprintf("%v", discoveryInfoInterface.MemberInterface) + "']"
}

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["member-interface"] = discoveryInfoInterface.MemberInterface
    leafs["name"] = discoveryInfoInterface.Name
    leafs["operational-status"] = discoveryInfoInterface.OperationalStatus
    leafs["loopback-mode"] = discoveryInfoInterface.LoopbackMode
    leafs["local-mode"] = discoveryInfoInterface.LocalMode
    leafs["miswired"] = discoveryInfoInterface.Miswired
    leafs["local-mwd-key"] = discoveryInfoInterface.LocalMwdKey
    leafs["local-function-unidirectional"] = discoveryInfoInterface.LocalFunctionUnidirectional
    leafs["local-function-loopback"] = discoveryInfoInterface.LocalFunctionLoopback
    leafs["local-function-event"] = discoveryInfoInterface.LocalFunctionEvent
    leafs["local-functionvariable"] = discoveryInfoInterface.LocalFunctionvariable
    leafs["local-revision"] = discoveryInfoInterface.LocalRevision
    leafs["local-mtu"] = discoveryInfoInterface.LocalMtu
    leafs["local-operational"] = discoveryInfoInterface.LocalOperational
    leafs["local-evaluating"] = discoveryInfoInterface.LocalEvaluating
    leafs["remote-mode"] = discoveryInfoInterface.RemoteMode
    leafs["remote-unidirectional"] = discoveryInfoInterface.RemoteUnidirectional
    leafs["remote-loopback"] = discoveryInfoInterface.RemoteLoopback
    leafs["remote-event"] = discoveryInfoInterface.RemoteEvent
    leafs["remote-variable"] = discoveryInfoInterface.RemoteVariable
    leafs["remote-mtu"] = discoveryInfoInterface.RemoteMtu
    leafs["remote-mac-address"] = discoveryInfoInterface.RemoteMacAddress
    leafs["remote-vendor-oui"] = discoveryInfoInterface.RemoteVendorOui
    leafs["remote-revision"] = discoveryInfoInterface.RemoteRevision
    leafs["remote-vendor-info"] = discoveryInfoInterface.RemoteVendorInfo
    leafs["remote-mwd-key"] = discoveryInfoInterface.RemoteMwdKey
    leafs["operational-status-valid"] = discoveryInfoInterface.OperationalStatusValid
    leafs["loopback-mode-valid"] = discoveryInfoInterface.LoopbackModeValid
    leafs["local-mode-valid"] = discoveryInfoInterface.LocalModeValid
    leafs["miswired-valid"] = discoveryInfoInterface.MiswiredValid
    leafs["local-mwd-key-valid"] = discoveryInfoInterface.LocalMwdKeyValid
    leafs["local-function-unidirectional-valid"] = discoveryInfoInterface.LocalFunctionUnidirectionalValid
    leafs["local-function-loopback-valid"] = discoveryInfoInterface.LocalFunctionLoopbackValid
    leafs["local-function-event-valid"] = discoveryInfoInterface.LocalFunctionEventValid
    leafs["local-functionvariable-valid"] = discoveryInfoInterface.LocalFunctionvariableValid
    leafs["local-revisionvalid"] = discoveryInfoInterface.LocalRevisionvalid
    leafs["local-mtu-valid"] = discoveryInfoInterface.LocalMtuValid
    leafs["remote-mode-valid"] = discoveryInfoInterface.RemoteModeValid
    leafs["remote-unidirectional-valid"] = discoveryInfoInterface.RemoteUnidirectionalValid
    leafs["remote-loopback-valid"] = discoveryInfoInterface.RemoteLoopbackValid
    leafs["remote-event-valid"] = discoveryInfoInterface.RemoteEventValid
    leafs["remote-variable-valid"] = discoveryInfoInterface.RemoteVariableValid
    leafs["remote-mtu-valid"] = discoveryInfoInterface.RemoteMtuValid
    leafs["remote-mac-address-valid"] = discoveryInfoInterface.RemoteMacAddressValid
    leafs["remote-vendor-oui-valid"] = discoveryInfoInterface.RemoteVendorOuiValid
    leafs["remote-revisionvalid"] = discoveryInfoInterface.RemoteRevisionvalid
    leafs["remote-vendor-info-valid"] = discoveryInfoInterface.RemoteVendorInfoValid
    leafs["remote-mwd-key-valid"] = discoveryInfoInterface.RemoteMwdKeyValid
    leafs["received-at-risk-notification-timestamp"] = discoveryInfoInterface.ReceivedAtRiskNotificationTimestamp
    leafs["received-at-risk-notification-time-remaining"] = discoveryInfoInterface.ReceivedAtRiskNotificationTimeRemaining
    return leafs
}

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetBundleName() string { return "cisco_ios_xr" }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetYangName() string { return "discovery-info-interface" }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) SetParent(parent types.Entity) { discoveryInfoInterface.parent = parent }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetParent() types.Entity { return discoveryInfoInterface.parent }

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetParentYangName() string { return "discovery-info-interfaces" }

// EtherLinkOam_InterfaceStateInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Interface State container
type EtherLinkOam_InterfaceStateInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet Link OAM interface to get Interface State for. The type is slice
    // of EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface.
    InterfaceStateInterface []EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface
}

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetFilter() yfilter.YFilter { return interfaceStateInterfaces.YFilter }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) SetFilter(yf yfilter.YFilter) { interfaceStateInterfaces.YFilter = yf }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetGoName(yname string) string {
    if yname == "interface-state-interface" { return "InterfaceStateInterface" }
    return ""
}

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetSegmentPath() string {
    return "interface-state-interfaces"
}

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-state-interface" {
        for _, c := range interfaceStateInterfaces.InterfaceStateInterface {
            if interfaceStateInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface{}
        interfaceStateInterfaces.InterfaceStateInterface = append(interfaceStateInterfaces.InterfaceStateInterface, child)
        return &interfaceStateInterfaces.InterfaceStateInterface[len(interfaceStateInterfaces.InterfaceStateInterface)-1]
    }
    return nil
}

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceStateInterfaces.InterfaceStateInterface {
        children[interfaceStateInterfaces.InterfaceStateInterface[i].GetSegmentPath()] = &interfaceStateInterfaces.InterfaceStateInterface[i]
    }
    return children
}

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetYangName() string { return "interface-state-interfaces" }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) SetParent(parent types.Entity) { interfaceStateInterfaces.parent = parent }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetParent() types.Entity { return interfaceStateInterfaces.parent }

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetParentYangName() string { return "ether-link-oam" }

// EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface
// Ethernet Link OAM interface to get Interface
// State for
type EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    MemberInterface interface{}

    // The state the protocol is in. The type is ProtocolState.
    ProtocolCode interface{}

    // Has a uni-directional link-fault been detected?. The type is bool.
    RxFault interface{}

    // The local MWD key. The type is interface{} with range: 0..4294967295.
    LocalMwdKey interface{}

    // Does the remote side have an MWD key?. The type is bool.
    RemoteMwdKeyPresent interface{}

    // The remote MWD key. The type is interface{} with range: 0..4294967295.
    RemoteMwdKey interface{}

    // The errors that have occurred on this interface.
    Errors EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors

    // Any present EFD triggers.
    EfdTriggers EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers
}

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetFilter() yfilter.YFilter { return interfaceStateInterface.YFilter }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) SetFilter(yf yfilter.YFilter) { interfaceStateInterface.YFilter = yf }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetGoName(yname string) string {
    if yname == "member-interface" { return "MemberInterface" }
    if yname == "protocol-code" { return "ProtocolCode" }
    if yname == "rx-fault" { return "RxFault" }
    if yname == "local-mwd-key" { return "LocalMwdKey" }
    if yname == "remote-mwd-key-present" { return "RemoteMwdKeyPresent" }
    if yname == "remote-mwd-key" { return "RemoteMwdKey" }
    if yname == "errors" { return "Errors" }
    if yname == "efd-triggers" { return "EfdTriggers" }
    return ""
}

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetSegmentPath() string {
    return "interface-state-interface" + "[member-interface='" + fmt.Sprintf("%v", interfaceStateInterface.MemberInterface) + "']"
}

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "errors" {
        return &interfaceStateInterface.Errors
    }
    if childYangName == "efd-triggers" {
        return &interfaceStateInterface.EfdTriggers
    }
    return nil
}

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["errors"] = &interfaceStateInterface.Errors
    children["efd-triggers"] = &interfaceStateInterface.EfdTriggers
    return children
}

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["member-interface"] = interfaceStateInterface.MemberInterface
    leafs["protocol-code"] = interfaceStateInterface.ProtocolCode
    leafs["rx-fault"] = interfaceStateInterface.RxFault
    leafs["local-mwd-key"] = interfaceStateInterface.LocalMwdKey
    leafs["remote-mwd-key-present"] = interfaceStateInterface.RemoteMwdKeyPresent
    leafs["remote-mwd-key"] = interfaceStateInterface.RemoteMwdKey
    return leafs
}

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetYangName() string { return "interface-state-interface" }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) SetParent(parent types.Entity) { interfaceStateInterface.parent = parent }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetParent() types.Entity { return interfaceStateInterface.parent }

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetParentYangName() string { return "interface-state-interfaces" }

// EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors
// The errors that have occurred on this interface
type EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reason for the Interface Management error (if applicable). The type is
    // string.
    PfiReason interface{}

    // The Interface Management error/success code. The type is interface{} with
    // range: 0..4294967295.
    PfiErrorCode interface{}

    // Reason for the platform error (if applicable). The type is string.
    PlatformReason interface{}

    // The platform error/success code. The type is interface{} with range:
    // 0..4294967295.
    PlatformErrorCode interface{}

    // Reason for the Packet I/O error (if applicable). The type is string.
    SpioReason interface{}

    // The Packet I/O error/success code. The type is interface{} with range:
    // 0..4294967295.
    SpioErrorCode interface{}

    // Reason for the Packet error (if applicable). The type is string.
    EpiReason interface{}

    // The Packet error/success code. The type is interface{} with range:
    // 0..4294967295.
    EpiErrorCode interface{}

    // Reason for the caps add error (if applicable). The type is string.
    CapsAddReason interface{}

    // The caps add error/success code. The type is interface{} with range:
    // 0..4294967295.
    CapsAddErrorCode interface{}
}

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetFilter() yfilter.YFilter { return errors.YFilter }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) SetFilter(yf yfilter.YFilter) { errors.YFilter = yf }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetGoName(yname string) string {
    if yname == "pfi-reason" { return "PfiReason" }
    if yname == "pfi-error-code" { return "PfiErrorCode" }
    if yname == "platform-reason" { return "PlatformReason" }
    if yname == "platform-error-code" { return "PlatformErrorCode" }
    if yname == "spio-reason" { return "SpioReason" }
    if yname == "spio-error-code" { return "SpioErrorCode" }
    if yname == "epi-reason" { return "EpiReason" }
    if yname == "epi-error-code" { return "EpiErrorCode" }
    if yname == "caps-add-reason" { return "CapsAddReason" }
    if yname == "caps-add-error-code" { return "CapsAddErrorCode" }
    return ""
}

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetSegmentPath() string {
    return "errors"
}

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pfi-reason"] = errors.PfiReason
    leafs["pfi-error-code"] = errors.PfiErrorCode
    leafs["platform-reason"] = errors.PlatformReason
    leafs["platform-error-code"] = errors.PlatformErrorCode
    leafs["spio-reason"] = errors.SpioReason
    leafs["spio-error-code"] = errors.SpioErrorCode
    leafs["epi-reason"] = errors.EpiReason
    leafs["epi-error-code"] = errors.EpiErrorCode
    leafs["caps-add-reason"] = errors.CapsAddReason
    leafs["caps-add-error-code"] = errors.CapsAddErrorCode
    return leafs
}

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetBundleName() string { return "cisco_ios_xr" }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetYangName() string { return "errors" }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) SetParent(parent types.Entity) { errors.parent = parent }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetParent() types.Entity { return errors.parent }

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetParentYangName() string { return "interface-state-interface" }

// EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers
// Any present EFD triggers
type EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link-fault messages being received. The type is bool.
    LinkFaultReceived interface{}

    // The discovery process has timed out. The type is bool.
    DiscoveryTimedOut interface{}

    // A capabilities conflict has been detected. The type is bool.
    CapabilitiesConflict interface{}

    // A wiring conflict has been detected. The type is bool.
    WiringConflict interface{}

    // The 802.3 OAM session is down. The type is bool.
    SessionDown interface{}
}

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetFilter() yfilter.YFilter { return efdTriggers.YFilter }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) SetFilter(yf yfilter.YFilter) { efdTriggers.YFilter = yf }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetGoName(yname string) string {
    if yname == "link-fault-received" { return "LinkFaultReceived" }
    if yname == "discovery-timed-out" { return "DiscoveryTimedOut" }
    if yname == "capabilities-conflict" { return "CapabilitiesConflict" }
    if yname == "wiring-conflict" { return "WiringConflict" }
    if yname == "session-down" { return "SessionDown" }
    return ""
}

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetSegmentPath() string {
    return "efd-triggers"
}

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-fault-received"] = efdTriggers.LinkFaultReceived
    leafs["discovery-timed-out"] = efdTriggers.DiscoveryTimedOut
    leafs["capabilities-conflict"] = efdTriggers.CapabilitiesConflict
    leafs["wiring-conflict"] = efdTriggers.WiringConflict
    leafs["session-down"] = efdTriggers.SessionDown
    return leafs
}

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetBundleName() string { return "cisco_ios_xr" }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetYangName() string { return "efd-triggers" }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) SetParent(parent types.Entity) { efdTriggers.parent = parent }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetParent() types.Entity { return efdTriggers.parent }

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetParentYangName() string { return "interface-state-interface" }

// EtherLinkOam_RunningConfigInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Running Config container
type EtherLinkOam_RunningConfigInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet Link OAM interface to get Running Config for. The type is slice of
    // EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface.
    RunningConfigInterface []EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface
}

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetFilter() yfilter.YFilter { return runningConfigInterfaces.YFilter }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) SetFilter(yf yfilter.YFilter) { runningConfigInterfaces.YFilter = yf }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetGoName(yname string) string {
    if yname == "running-config-interface" { return "RunningConfigInterface" }
    return ""
}

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetSegmentPath() string {
    return "running-config-interfaces"
}

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "running-config-interface" {
        for _, c := range runningConfigInterfaces.RunningConfigInterface {
            if runningConfigInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface{}
        runningConfigInterfaces.RunningConfigInterface = append(runningConfigInterfaces.RunningConfigInterface, child)
        return &runningConfigInterfaces.RunningConfigInterface[len(runningConfigInterfaces.RunningConfigInterface)-1]
    }
    return nil
}

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range runningConfigInterfaces.RunningConfigInterface {
        children[runningConfigInterfaces.RunningConfigInterface[i].GetSegmentPath()] = &runningConfigInterfaces.RunningConfigInterface[i]
    }
    return children
}

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetYangName() string { return "running-config-interfaces" }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) SetParent(parent types.Entity) { runningConfigInterfaces.parent = parent }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetParent() types.Entity { return runningConfigInterfaces.parent }

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetParentYangName() string { return "ether-link-oam" }

// EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface
// Ethernet Link OAM interface to get Running
// Config for
type EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    MemberInterface interface{}

    // Is 100ms hello interval time enabled?. The type is bool.
    FastHelloIntervalEnabled interface{}

    // Is link monitoring enabled?. The type is bool.
    LinkMonitorEnabled interface{}

    // Is remote loopback enabled?. The type is bool.
    RemoteLoopbackEnabled interface{}

    // Is MIB retrieval enabled?. The type is bool.
    MibRetrievalEnabled interface{}

    // Is uni-directional link-fault detection enabled?. The type is bool.
    UdlfEnabled interface{}

    // Configured mode. The type is Mode.
    Mode interface{}

    // Connection timeout. The type is interface{} with range: 0..255.
    ConnectionTimeout interface{}

    // Symbol period event window size. The type is interface{} with range:
    // 0..4294967295.
    SymbolPeriodWindow interface{}

    // Symbol period event window units. The type is interface{} with range:
    // 0..255.
    SymbolPeriodWindowUnits interface{}

    // Symbol period event window multiplier. The type is interface{} with range:
    // 0..255.
    SymbolPeriodWindowMultiplier interface{}

    // Symbol period event low threshold. The type is interface{} with range:
    // 0..18446744073709551615.
    SymbolPeriodThresholdLow interface{}

    // Symbol period event high threshold. The type is interface{} with range:
    // 0..18446744073709551615.
    SymbolPeriodThresholdHigh interface{}

    // Symbol period event threshold units. The type is interface{} with range:
    // 0..255.
    SymbolPeriodThresholdUnits interface{}

    // Symbol period event threshold low multiplier. The type is interface{} with
    // range: 0..255.
    SymbolPeriodThresholdLowMultiplier interface{}

    // Symbol period event threshold high multiplier. The type is interface{} with
    // range: 0..255.
    SymbolPeriodThresholdHighMultiplier interface{}

    // Frame event window size. The type is interface{} with range: 0..4294967295.
    FrameWindow interface{}

    // Frame event low threshold. The type is interface{} with range:
    // 0..18446744073709551615.
    FrameThresholdLow interface{}

    // Frame event high threshold. The type is interface{} with range:
    // 0..18446744073709551615.
    FrameThresholdHigh interface{}

    // Frame period event threshold low multiplier. The type is interface{} with
    // range: 0..255.
    FrameThresholdLowMultiplier interface{}

    // Frame event threshold high multiplier. The type is interface{} with range:
    // 0..255.
    FrameThresholdHighMultiplier interface{}

    // Frame period event window size. The type is interface{} with range:
    // 0..4294967295.
    FramePeriodWindow interface{}

    // Frame period event window units. The type is interface{} with range:
    // 0..255.
    FramePeriodWindowUnits interface{}

    // Frame period event window multiplier. The type is interface{} with range:
    // 0..255.
    FramePeriodWindowMultiplier interface{}

    // Frame period event low threshold. The type is interface{} with range:
    // 0..18446744073709551615.
    FramePeriodThresholdLow interface{}

    // Frame period event high threshold. The type is interface{} with range:
    // 0..18446744073709551615.
    FramePeriodThresholdHigh interface{}

    // Frame period event threshold units. The type is interface{} with range:
    // 0..255.
    FramePeriodThresholdUnits interface{}

    // Frame period event threshold low multiplier. The type is interface{} with
    // range: 0..255.
    FramePeriodThresholdLowMultiplier interface{}

    // Frame period event threshold high multiplier. The type is interface{} with
    // range: 0..255.
    FramePeriodThresholdHighMultiplier interface{}

    // Frame seconds event high threshold. The type is interface{} with range:
    // 0..4294967295. Units are second.
    FrameSecondsWindow interface{}

    // Frame seconds event high threshold. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    FrameSecondsThresholdLow interface{}

    // Frame seconds event high threshold. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    FrameSecondsThresholdHigh interface{}

    // Action to perform when a high threshold is breached. The type is Action.
    HighThresholdAction interface{}

    // Action to perform when a link fault occurs. The type is Action.
    LinkFaultAction interface{}

    // Action to perform when a dying gasp occurs. The type is Action.
    DyingGaspAction interface{}

    // Action to perform when a critical event occurs. The type is Action.
    CriticalEventAction interface{}

    // Action to perform when a discovery timeout occurs. The type is Action.
    DiscoveryTimeoutAction interface{}

    // Action to perform when a capabilities conflict occurs. The type is Action.
    CapabilitiesConflictAction interface{}

    // Action to perform when a wiring conflict occurs. The type is Action.
    WiringConflictAction interface{}

    // Action to perform when a session comes up. The type is Action.
    SessionUpAction interface{}

    // Action to perform when a session comes down. The type is Action.
    SessionDownAction interface{}

    // Action to perform when a session enters or exits remote loopback. The type
    // is Action.
    RemoteLoopbackAction interface{}

    // The mode that is required of the remote peer. The type is Mode.
    RequireRemoteMode interface{}

    // Require the remote peer to support MIB retrieval. The type is bool.
    RequireRemoteMibRetrieval interface{}

    // Require the remote peer to support loopback mode. The type is bool.
    RequireLoopback interface{}

    // Require the remote peer to support link monitoring. The type is bool.
    RequireLinkMonitoring interface{}

    // Is this configuration information an interface override?. The type is bool.
    FastHelloIntervalEnabledOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    LinkMonitoringEnabledOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    RemoteLoopbackEnabledOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    MibRetrievalEnabledOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    UdlfEnabledOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    ModeOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    ConnectionTimeoutOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    SymbolPeriodWindowOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    SymbolPeriodThresholdLowOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    SymbolPeriodThresholdHighOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FrameWindowOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FrameThresholdLowOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FrameThresholdHighOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FramePeriodWindowOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FramePeriodThresholdLowOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FramePeriodThresholdHighOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FrameSecondsWindowOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FrameSecondsThresholdLowOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    FrameSecondsThresholdHighOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    HighThresholdActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    LinkFaultActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    DyingGaspActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    CriticalEventActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    DiscoveryTimeoutActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    CapabilitiesConflictActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    WiringConflictActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    SessionDownActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    SessionUpActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    RemoteLoopbackActionOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    RequireModeOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    RequireMibRetrievalOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    RequireLoopbackOverridden interface{}

    // Is this configuration information an interface override?. The type is bool.
    RequireLinkMonitoringOverridden interface{}
}

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetFilter() yfilter.YFilter { return runningConfigInterface.YFilter }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) SetFilter(yf yfilter.YFilter) { runningConfigInterface.YFilter = yf }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetGoName(yname string) string {
    if yname == "member-interface" { return "MemberInterface" }
    if yname == "fast-hello-interval-enabled" { return "FastHelloIntervalEnabled" }
    if yname == "link-monitor-enabled" { return "LinkMonitorEnabled" }
    if yname == "remote-loopback-enabled" { return "RemoteLoopbackEnabled" }
    if yname == "mib-retrieval-enabled" { return "MibRetrievalEnabled" }
    if yname == "udlf-enabled" { return "UdlfEnabled" }
    if yname == "mode" { return "Mode" }
    if yname == "connection-timeout" { return "ConnectionTimeout" }
    if yname == "symbol-period-window" { return "SymbolPeriodWindow" }
    if yname == "symbol-period-window-units" { return "SymbolPeriodWindowUnits" }
    if yname == "symbol-period-window-multiplier" { return "SymbolPeriodWindowMultiplier" }
    if yname == "symbol-period-threshold-low" { return "SymbolPeriodThresholdLow" }
    if yname == "symbol-period-threshold-high" { return "SymbolPeriodThresholdHigh" }
    if yname == "symbol-period-threshold-units" { return "SymbolPeriodThresholdUnits" }
    if yname == "symbol-period-threshold-low-multiplier" { return "SymbolPeriodThresholdLowMultiplier" }
    if yname == "symbol-period-threshold-high-multiplier" { return "SymbolPeriodThresholdHighMultiplier" }
    if yname == "frame-window" { return "FrameWindow" }
    if yname == "frame-threshold-low" { return "FrameThresholdLow" }
    if yname == "frame-threshold-high" { return "FrameThresholdHigh" }
    if yname == "frame-threshold-low-multiplier" { return "FrameThresholdLowMultiplier" }
    if yname == "frame-threshold-high-multiplier" { return "FrameThresholdHighMultiplier" }
    if yname == "frame-period-window" { return "FramePeriodWindow" }
    if yname == "frame-period-window-units" { return "FramePeriodWindowUnits" }
    if yname == "frame-period-window-multiplier" { return "FramePeriodWindowMultiplier" }
    if yname == "frame-period-threshold-low" { return "FramePeriodThresholdLow" }
    if yname == "frame-period-threshold-high" { return "FramePeriodThresholdHigh" }
    if yname == "frame-period-threshold-units" { return "FramePeriodThresholdUnits" }
    if yname == "frame-period-threshold-low-multiplier" { return "FramePeriodThresholdLowMultiplier" }
    if yname == "frame-period-threshold-high-multiplier" { return "FramePeriodThresholdHighMultiplier" }
    if yname == "frame-seconds-window" { return "FrameSecondsWindow" }
    if yname == "frame-seconds-threshold-low" { return "FrameSecondsThresholdLow" }
    if yname == "frame-seconds-threshold-high" { return "FrameSecondsThresholdHigh" }
    if yname == "high-threshold-action" { return "HighThresholdAction" }
    if yname == "link-fault-action" { return "LinkFaultAction" }
    if yname == "dying-gasp-action" { return "DyingGaspAction" }
    if yname == "critical-event-action" { return "CriticalEventAction" }
    if yname == "discovery-timeout-action" { return "DiscoveryTimeoutAction" }
    if yname == "capabilities-conflict-action" { return "CapabilitiesConflictAction" }
    if yname == "wiring-conflict-action" { return "WiringConflictAction" }
    if yname == "session-up-action" { return "SessionUpAction" }
    if yname == "session-down-action" { return "SessionDownAction" }
    if yname == "remote-loopback-action" { return "RemoteLoopbackAction" }
    if yname == "require-remote-mode" { return "RequireRemoteMode" }
    if yname == "require-remote-mib-retrieval" { return "RequireRemoteMibRetrieval" }
    if yname == "require-loopback" { return "RequireLoopback" }
    if yname == "require-link-monitoring" { return "RequireLinkMonitoring" }
    if yname == "fast-hello-interval-enabled-overridden" { return "FastHelloIntervalEnabledOverridden" }
    if yname == "link-monitoring-enabled-overridden" { return "LinkMonitoringEnabledOverridden" }
    if yname == "remote-loopback-enabled-overridden" { return "RemoteLoopbackEnabledOverridden" }
    if yname == "mib-retrieval-enabled-overridden" { return "MibRetrievalEnabledOverridden" }
    if yname == "udlf-enabled-overridden" { return "UdlfEnabledOverridden" }
    if yname == "mode-overridden" { return "ModeOverridden" }
    if yname == "connection-timeout-overridden" { return "ConnectionTimeoutOverridden" }
    if yname == "symbol-period-window-overridden" { return "SymbolPeriodWindowOverridden" }
    if yname == "symbol-period-threshold-low-overridden" { return "SymbolPeriodThresholdLowOverridden" }
    if yname == "symbol-period-threshold-high-overridden" { return "SymbolPeriodThresholdHighOverridden" }
    if yname == "frame-window-overridden" { return "FrameWindowOverridden" }
    if yname == "frame-threshold-low-overridden" { return "FrameThresholdLowOverridden" }
    if yname == "frame-threshold-high-overridden" { return "FrameThresholdHighOverridden" }
    if yname == "frame-period-window-overridden" { return "FramePeriodWindowOverridden" }
    if yname == "frame-period-threshold-low-overridden" { return "FramePeriodThresholdLowOverridden" }
    if yname == "frame-period-threshold-high-overridden" { return "FramePeriodThresholdHighOverridden" }
    if yname == "frame-seconds-window-overridden" { return "FrameSecondsWindowOverridden" }
    if yname == "frame-seconds-threshold-low-overridden" { return "FrameSecondsThresholdLowOverridden" }
    if yname == "frame-seconds-threshold-high-overridden" { return "FrameSecondsThresholdHighOverridden" }
    if yname == "high-threshold-action-overridden" { return "HighThresholdActionOverridden" }
    if yname == "link-fault-action-overridden" { return "LinkFaultActionOverridden" }
    if yname == "dying-gasp-action-overridden" { return "DyingGaspActionOverridden" }
    if yname == "critical-event-action-overridden" { return "CriticalEventActionOverridden" }
    if yname == "discovery-timeout-action-overridden" { return "DiscoveryTimeoutActionOverridden" }
    if yname == "capabilities-conflict-action-overridden" { return "CapabilitiesConflictActionOverridden" }
    if yname == "wiring-conflict-action-overridden" { return "WiringConflictActionOverridden" }
    if yname == "session-down-action-overridden" { return "SessionDownActionOverridden" }
    if yname == "session-up-action-overridden" { return "SessionUpActionOverridden" }
    if yname == "remote-loopback-action-overridden" { return "RemoteLoopbackActionOverridden" }
    if yname == "require-mode-overridden" { return "RequireModeOverridden" }
    if yname == "require-mib-retrieval-overridden" { return "RequireMibRetrievalOverridden" }
    if yname == "require-loopback-overridden" { return "RequireLoopbackOverridden" }
    if yname == "require-link-monitoring-overridden" { return "RequireLinkMonitoringOverridden" }
    return ""
}

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetSegmentPath() string {
    return "running-config-interface" + "[member-interface='" + fmt.Sprintf("%v", runningConfigInterface.MemberInterface) + "']"
}

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["member-interface"] = runningConfigInterface.MemberInterface
    leafs["fast-hello-interval-enabled"] = runningConfigInterface.FastHelloIntervalEnabled
    leafs["link-monitor-enabled"] = runningConfigInterface.LinkMonitorEnabled
    leafs["remote-loopback-enabled"] = runningConfigInterface.RemoteLoopbackEnabled
    leafs["mib-retrieval-enabled"] = runningConfigInterface.MibRetrievalEnabled
    leafs["udlf-enabled"] = runningConfigInterface.UdlfEnabled
    leafs["mode"] = runningConfigInterface.Mode
    leafs["connection-timeout"] = runningConfigInterface.ConnectionTimeout
    leafs["symbol-period-window"] = runningConfigInterface.SymbolPeriodWindow
    leafs["symbol-period-window-units"] = runningConfigInterface.SymbolPeriodWindowUnits
    leafs["symbol-period-window-multiplier"] = runningConfigInterface.SymbolPeriodWindowMultiplier
    leafs["symbol-period-threshold-low"] = runningConfigInterface.SymbolPeriodThresholdLow
    leafs["symbol-period-threshold-high"] = runningConfigInterface.SymbolPeriodThresholdHigh
    leafs["symbol-period-threshold-units"] = runningConfigInterface.SymbolPeriodThresholdUnits
    leafs["symbol-period-threshold-low-multiplier"] = runningConfigInterface.SymbolPeriodThresholdLowMultiplier
    leafs["symbol-period-threshold-high-multiplier"] = runningConfigInterface.SymbolPeriodThresholdHighMultiplier
    leafs["frame-window"] = runningConfigInterface.FrameWindow
    leafs["frame-threshold-low"] = runningConfigInterface.FrameThresholdLow
    leafs["frame-threshold-high"] = runningConfigInterface.FrameThresholdHigh
    leafs["frame-threshold-low-multiplier"] = runningConfigInterface.FrameThresholdLowMultiplier
    leafs["frame-threshold-high-multiplier"] = runningConfigInterface.FrameThresholdHighMultiplier
    leafs["frame-period-window"] = runningConfigInterface.FramePeriodWindow
    leafs["frame-period-window-units"] = runningConfigInterface.FramePeriodWindowUnits
    leafs["frame-period-window-multiplier"] = runningConfigInterface.FramePeriodWindowMultiplier
    leafs["frame-period-threshold-low"] = runningConfigInterface.FramePeriodThresholdLow
    leafs["frame-period-threshold-high"] = runningConfigInterface.FramePeriodThresholdHigh
    leafs["frame-period-threshold-units"] = runningConfigInterface.FramePeriodThresholdUnits
    leafs["frame-period-threshold-low-multiplier"] = runningConfigInterface.FramePeriodThresholdLowMultiplier
    leafs["frame-period-threshold-high-multiplier"] = runningConfigInterface.FramePeriodThresholdHighMultiplier
    leafs["frame-seconds-window"] = runningConfigInterface.FrameSecondsWindow
    leafs["frame-seconds-threshold-low"] = runningConfigInterface.FrameSecondsThresholdLow
    leafs["frame-seconds-threshold-high"] = runningConfigInterface.FrameSecondsThresholdHigh
    leafs["high-threshold-action"] = runningConfigInterface.HighThresholdAction
    leafs["link-fault-action"] = runningConfigInterface.LinkFaultAction
    leafs["dying-gasp-action"] = runningConfigInterface.DyingGaspAction
    leafs["critical-event-action"] = runningConfigInterface.CriticalEventAction
    leafs["discovery-timeout-action"] = runningConfigInterface.DiscoveryTimeoutAction
    leafs["capabilities-conflict-action"] = runningConfigInterface.CapabilitiesConflictAction
    leafs["wiring-conflict-action"] = runningConfigInterface.WiringConflictAction
    leafs["session-up-action"] = runningConfigInterface.SessionUpAction
    leafs["session-down-action"] = runningConfigInterface.SessionDownAction
    leafs["remote-loopback-action"] = runningConfigInterface.RemoteLoopbackAction
    leafs["require-remote-mode"] = runningConfigInterface.RequireRemoteMode
    leafs["require-remote-mib-retrieval"] = runningConfigInterface.RequireRemoteMibRetrieval
    leafs["require-loopback"] = runningConfigInterface.RequireLoopback
    leafs["require-link-monitoring"] = runningConfigInterface.RequireLinkMonitoring
    leafs["fast-hello-interval-enabled-overridden"] = runningConfigInterface.FastHelloIntervalEnabledOverridden
    leafs["link-monitoring-enabled-overridden"] = runningConfigInterface.LinkMonitoringEnabledOverridden
    leafs["remote-loopback-enabled-overridden"] = runningConfigInterface.RemoteLoopbackEnabledOverridden
    leafs["mib-retrieval-enabled-overridden"] = runningConfigInterface.MibRetrievalEnabledOverridden
    leafs["udlf-enabled-overridden"] = runningConfigInterface.UdlfEnabledOverridden
    leafs["mode-overridden"] = runningConfigInterface.ModeOverridden
    leafs["connection-timeout-overridden"] = runningConfigInterface.ConnectionTimeoutOverridden
    leafs["symbol-period-window-overridden"] = runningConfigInterface.SymbolPeriodWindowOverridden
    leafs["symbol-period-threshold-low-overridden"] = runningConfigInterface.SymbolPeriodThresholdLowOverridden
    leafs["symbol-period-threshold-high-overridden"] = runningConfigInterface.SymbolPeriodThresholdHighOverridden
    leafs["frame-window-overridden"] = runningConfigInterface.FrameWindowOverridden
    leafs["frame-threshold-low-overridden"] = runningConfigInterface.FrameThresholdLowOverridden
    leafs["frame-threshold-high-overridden"] = runningConfigInterface.FrameThresholdHighOverridden
    leafs["frame-period-window-overridden"] = runningConfigInterface.FramePeriodWindowOverridden
    leafs["frame-period-threshold-low-overridden"] = runningConfigInterface.FramePeriodThresholdLowOverridden
    leafs["frame-period-threshold-high-overridden"] = runningConfigInterface.FramePeriodThresholdHighOverridden
    leafs["frame-seconds-window-overridden"] = runningConfigInterface.FrameSecondsWindowOverridden
    leafs["frame-seconds-threshold-low-overridden"] = runningConfigInterface.FrameSecondsThresholdLowOverridden
    leafs["frame-seconds-threshold-high-overridden"] = runningConfigInterface.FrameSecondsThresholdHighOverridden
    leafs["high-threshold-action-overridden"] = runningConfigInterface.HighThresholdActionOverridden
    leafs["link-fault-action-overridden"] = runningConfigInterface.LinkFaultActionOverridden
    leafs["dying-gasp-action-overridden"] = runningConfigInterface.DyingGaspActionOverridden
    leafs["critical-event-action-overridden"] = runningConfigInterface.CriticalEventActionOverridden
    leafs["discovery-timeout-action-overridden"] = runningConfigInterface.DiscoveryTimeoutActionOverridden
    leafs["capabilities-conflict-action-overridden"] = runningConfigInterface.CapabilitiesConflictActionOverridden
    leafs["wiring-conflict-action-overridden"] = runningConfigInterface.WiringConflictActionOverridden
    leafs["session-down-action-overridden"] = runningConfigInterface.SessionDownActionOverridden
    leafs["session-up-action-overridden"] = runningConfigInterface.SessionUpActionOverridden
    leafs["remote-loopback-action-overridden"] = runningConfigInterface.RemoteLoopbackActionOverridden
    leafs["require-mode-overridden"] = runningConfigInterface.RequireModeOverridden
    leafs["require-mib-retrieval-overridden"] = runningConfigInterface.RequireMibRetrievalOverridden
    leafs["require-loopback-overridden"] = runningConfigInterface.RequireLoopbackOverridden
    leafs["require-link-monitoring-overridden"] = runningConfigInterface.RequireLinkMonitoringOverridden
    return leafs
}

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetBundleName() string { return "cisco_ios_xr" }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetYangName() string { return "running-config-interface" }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) SetParent(parent types.Entity) { runningConfigInterface.parent = parent }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetParent() types.Entity { return runningConfigInterface.parent }

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetParentYangName() string { return "running-config-interfaces" }

// EtherLinkOam_Nodes
// Node table for node-specific operational data
type EtherLinkOam_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // EtherLinkOam_Nodes_Node.
    Node []EtherLinkOam_Nodes_Node
}

func (nodes *EtherLinkOam_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *EtherLinkOam_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *EtherLinkOam_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *EtherLinkOam_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *EtherLinkOam_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLinkOam_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *EtherLinkOam_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *EtherLinkOam_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *EtherLinkOam_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *EtherLinkOam_Nodes) GetYangName() string { return "nodes" }

func (nodes *EtherLinkOam_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *EtherLinkOam_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *EtherLinkOam_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *EtherLinkOam_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *EtherLinkOam_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *EtherLinkOam_Nodes) GetParentYangName() string { return "ether-link-oam" }

// EtherLinkOam_Nodes_Node
// Node-specific data for a particular node
type EtherLinkOam_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Ethernet Link OAM Summary information for the entire node.
    Summary EtherLinkOam_Nodes_Node_Summary
}

func (node *EtherLinkOam_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *EtherLinkOam_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *EtherLinkOam_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (node *EtherLinkOam_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *EtherLinkOam_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &node.Summary
    }
    return nil
}

func (node *EtherLinkOam_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &node.Summary
    return children
}

func (node *EtherLinkOam_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *EtherLinkOam_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *EtherLinkOam_Nodes_Node) GetYangName() string { return "node" }

func (node *EtherLinkOam_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *EtherLinkOam_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *EtherLinkOam_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *EtherLinkOam_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *EtherLinkOam_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *EtherLinkOam_Nodes_Node) GetParentYangName() string { return "nodes" }

// EtherLinkOam_Nodes_Node_Summary
// Ethernet Link OAM Summary information for the
// entire node
type EtherLinkOam_Nodes_Node_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of interfaces with 802.3 OAM configured. The type is interface{}
    // with range: 0..4294967295.
    Interfaces interface{}

    // The number of interfaces in 'Port Down' state. The type is interface{} with
    // range: 0..4294967295.
    PortDown interface{}

    // The number of interfaces in 'Passive Wait' state. The type is interface{}
    // with range: 0..4294967295.
    PassiveWait interface{}

    // The number of interfaces in 'Active Send' state. The type is interface{}
    // with range: 0..4294967295.
    ActiveSend interface{}

    // The number of interfaces in 'Evaluating' state. The type is interface{}
    // with range: 0..4294967295.
    Evaluating interface{}

    // The number of interfaces in 'Local Accept' state. The type is interface{}
    // with range: 0..4294967295.
    LocalAccept interface{}

    // The number of interfaces in 'Local Reject' state. The type is interface{}
    // with range: 0..4294967295.
    LocalReject interface{}

    // The number of interfaces in 'Remote Reject' state. The type is interface{}
    // with range: 0..4294967295.
    RemoteReject interface{}

    // The number of interfaces in 'Operational' state. The type is interface{}
    // with range: 0..4294967295.
    Operational interface{}

    // The number of interfaces in loopback mode. The type is interface{} with
    // range: 0..4294967295.
    LoopbackMode interface{}

    // The number of miswired connections. The type is interface{} with range:
    // 0..4294967295.
    MiswiredConnections interface{}

    // The number of events recorded. The type is interface{} with range:
    // 0..18446744073709551615.
    Events interface{}

    // The number of local events recorded. The type is interface{} with range:
    // 0..18446744073709551615.
    LocalEvents interface{}

    // The number of local symbol period events recorded. The type is interface{}
    // with range: 0..18446744073709551615.
    LocalSymbolPeriod interface{}

    // The mumber of local frame error events recorded. The type is interface{}
    // with range: 0..18446744073709551615.
    LocalFrame interface{}

    // The number of local frame period events recorded. The type is interface{}
    // with range: 0..18446744073709551615.
    LocalFramePeriod interface{}

    // The number of local frame second events recoded. The type is interface{}
    // with range: 0..18446744073709551615.
    LocalFrameSeconds interface{}

    // The number of remote events recorded. The type is interface{} with range:
    // 0..18446744073709551615.
    RemoteEvents interface{}

    // The number of remote symbol period events recorded. The type is interface{}
    // with range: 0..18446744073709551615.
    RemoteSymbolPeriod interface{}

    // The mumber of remote frame error events recorded. The type is interface{}
    // with range: 0..18446744073709551615.
    RemoteFrame interface{}

    // The number of remote frame period events recorded. The type is interface{}
    // with range: 0..18446744073709551615.
    RemoteFramePeriod interface{}

    // The number of remote frame second events recoded. The type is interface{}
    // with range: 0..18446744073709551615.
    RemoteFrameSeconds interface{}
}

func (summary *EtherLinkOam_Nodes_Node_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *EtherLinkOam_Nodes_Node_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *EtherLinkOam_Nodes_Node_Summary) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "port-down" { return "PortDown" }
    if yname == "passive-wait" { return "PassiveWait" }
    if yname == "active-send" { return "ActiveSend" }
    if yname == "evaluating" { return "Evaluating" }
    if yname == "local-accept" { return "LocalAccept" }
    if yname == "local-reject" { return "LocalReject" }
    if yname == "remote-reject" { return "RemoteReject" }
    if yname == "operational" { return "Operational" }
    if yname == "loopback-mode" { return "LoopbackMode" }
    if yname == "miswired-connections" { return "MiswiredConnections" }
    if yname == "events" { return "Events" }
    if yname == "local-events" { return "LocalEvents" }
    if yname == "local-symbol-period" { return "LocalSymbolPeriod" }
    if yname == "local-frame" { return "LocalFrame" }
    if yname == "local-frame-period" { return "LocalFramePeriod" }
    if yname == "local-frame-seconds" { return "LocalFrameSeconds" }
    if yname == "remote-events" { return "RemoteEvents" }
    if yname == "remote-symbol-period" { return "RemoteSymbolPeriod" }
    if yname == "remote-frame" { return "RemoteFrame" }
    if yname == "remote-frame-period" { return "RemoteFramePeriod" }
    if yname == "remote-frame-seconds" { return "RemoteFrameSeconds" }
    return ""
}

func (summary *EtherLinkOam_Nodes_Node_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *EtherLinkOam_Nodes_Node_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *EtherLinkOam_Nodes_Node_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *EtherLinkOam_Nodes_Node_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interfaces"] = summary.Interfaces
    leafs["port-down"] = summary.PortDown
    leafs["passive-wait"] = summary.PassiveWait
    leafs["active-send"] = summary.ActiveSend
    leafs["evaluating"] = summary.Evaluating
    leafs["local-accept"] = summary.LocalAccept
    leafs["local-reject"] = summary.LocalReject
    leafs["remote-reject"] = summary.RemoteReject
    leafs["operational"] = summary.Operational
    leafs["loopback-mode"] = summary.LoopbackMode
    leafs["miswired-connections"] = summary.MiswiredConnections
    leafs["events"] = summary.Events
    leafs["local-events"] = summary.LocalEvents
    leafs["local-symbol-period"] = summary.LocalSymbolPeriod
    leafs["local-frame"] = summary.LocalFrame
    leafs["local-frame-period"] = summary.LocalFramePeriod
    leafs["local-frame-seconds"] = summary.LocalFrameSeconds
    leafs["remote-events"] = summary.RemoteEvents
    leafs["remote-symbol-period"] = summary.RemoteSymbolPeriod
    leafs["remote-frame"] = summary.RemoteFrame
    leafs["remote-frame-period"] = summary.RemoteFramePeriod
    leafs["remote-frame-seconds"] = summary.RemoteFrameSeconds
    return leafs
}

func (summary *EtherLinkOam_Nodes_Node_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *EtherLinkOam_Nodes_Node_Summary) GetYangName() string { return "summary" }

func (summary *EtherLinkOam_Nodes_Node_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *EtherLinkOam_Nodes_Node_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *EtherLinkOam_Nodes_Node_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *EtherLinkOam_Nodes_Node_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *EtherLinkOam_Nodes_Node_Summary) GetParent() types.Entity { return summary.parent }

func (summary *EtherLinkOam_Nodes_Node_Summary) GetParentYangName() string { return "node" }

// EtherLinkOam_EventLogEntryInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Event Log Entry container
type EtherLinkOam_EventLogEntryInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet Link OAM enabled interface to get Event Log Entry for. The type is
    // slice of EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface.
    EventLogEntryInterface []EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface
}

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetFilter() yfilter.YFilter { return eventLogEntryInterfaces.YFilter }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) SetFilter(yf yfilter.YFilter) { eventLogEntryInterfaces.YFilter = yf }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetGoName(yname string) string {
    if yname == "event-log-entry-interface" { return "EventLogEntryInterface" }
    return ""
}

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetSegmentPath() string {
    return "event-log-entry-interfaces"
}

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "event-log-entry-interface" {
        for _, c := range eventLogEntryInterfaces.EventLogEntryInterface {
            if eventLogEntryInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface{}
        eventLogEntryInterfaces.EventLogEntryInterface = append(eventLogEntryInterfaces.EventLogEntryInterface, child)
        return &eventLogEntryInterfaces.EventLogEntryInterface[len(eventLogEntryInterfaces.EventLogEntryInterface)-1]
    }
    return nil
}

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range eventLogEntryInterfaces.EventLogEntryInterface {
        children[eventLogEntryInterfaces.EventLogEntryInterface[i].GetSegmentPath()] = &eventLogEntryInterfaces.EventLogEntryInterface[i]
    }
    return children
}

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetYangName() string { return "event-log-entry-interfaces" }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) SetParent(parent types.Entity) { eventLogEntryInterfaces.parent = parent }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetParent() types.Entity { return eventLogEntryInterfaces.parent }

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetParentYangName() string { return "ether-link-oam" }

// EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface
// Ethernet Link OAM enabled interface to get
// Event Log Entry for
type EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    MemberInterface interface{}

    // Table of Ethernet Link OAM Event Log entries on the interface.
    EventLogEntryIndexes EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes
}

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetFilter() yfilter.YFilter { return eventLogEntryInterface.YFilter }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) SetFilter(yf yfilter.YFilter) { eventLogEntryInterface.YFilter = yf }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetGoName(yname string) string {
    if yname == "member-interface" { return "MemberInterface" }
    if yname == "event-log-entry-indexes" { return "EventLogEntryIndexes" }
    return ""
}

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetSegmentPath() string {
    return "event-log-entry-interface" + "[member-interface='" + fmt.Sprintf("%v", eventLogEntryInterface.MemberInterface) + "']"
}

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "event-log-entry-indexes" {
        return &eventLogEntryInterface.EventLogEntryIndexes
    }
    return nil
}

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["event-log-entry-indexes"] = &eventLogEntryInterface.EventLogEntryIndexes
    return children
}

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["member-interface"] = eventLogEntryInterface.MemberInterface
    return leafs
}

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetBundleName() string { return "cisco_ios_xr" }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetYangName() string { return "event-log-entry-interface" }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) SetParent(parent types.Entity) { eventLogEntryInterface.parent = parent }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetParent() types.Entity { return eventLogEntryInterface.parent }

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetParentYangName() string { return "event-log-entry-interfaces" }

// EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes
// Table of Ethernet Link OAM Event Log entries
// on the interface
type EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet Link OAM Event Log Entry Index to get data for. The type is slice
    // of
    // EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex.
    EventLogEntryIndex []EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex
}

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetFilter() yfilter.YFilter { return eventLogEntryIndexes.YFilter }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) SetFilter(yf yfilter.YFilter) { eventLogEntryIndexes.YFilter = yf }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetGoName(yname string) string {
    if yname == "event-log-entry-index" { return "EventLogEntryIndex" }
    return ""
}

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetSegmentPath() string {
    return "event-log-entry-indexes"
}

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "event-log-entry-index" {
        for _, c := range eventLogEntryIndexes.EventLogEntryIndex {
            if eventLogEntryIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex{}
        eventLogEntryIndexes.EventLogEntryIndex = append(eventLogEntryIndexes.EventLogEntryIndex, child)
        return &eventLogEntryIndexes.EventLogEntryIndex[len(eventLogEntryIndexes.EventLogEntryIndex)-1]
    }
    return nil
}

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range eventLogEntryIndexes.EventLogEntryIndex {
        children[eventLogEntryIndexes.EventLogEntryIndex[i].GetSegmentPath()] = &eventLogEntryIndexes.EventLogEntryIndex[i]
    }
    return children
}

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetYangName() string { return "event-log-entry-indexes" }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) SetParent(parent types.Entity) { eventLogEntryIndexes.parent = parent }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetParent() types.Entity { return eventLogEntryIndexes.parent }

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetParentYangName() string { return "event-log-entry-interface" }

// EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex
// Ethernet Link OAM Event Log Entry Index to
// get data for
type EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Event Log Entry index. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    EventLogEntryIndex interface{}

    // Index in the log entries table. The type is interface{} with range:
    // 0..4294967295.
    Index interface{}

    // Interface handle for this log entry. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Handle interface{}

    // OUI for the log entry. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Oui interface{}

    // Timestamp in hundredths of a second since unix epoch for when the event
    // occurred. The type is interface{} with range: 0..18446744073709551615.
    Timestamp interface{}

    // Type of event that this entry describes. The type is Log.
    Type interface{}

    // Where the event occurred. The type is LogLocation.
    Location interface{}

    // Total number of times event has occurred. The type is interface{} with
    // range: 0..4294967295.
    EventTotal interface{}

    // Local action taken (If applicable). The type is Action.
    ActionTaken interface{}

    // Size of the window (If applicable). The type is interface{} with range:
    // 0..18446744073709551615.
    Window interface{}

    // Size of the threshold (If applicable). The type is interface{} with range:
    // 0..18446744073709551615.
    Threshold interface{}

    // Size of the local high threshold (If applicable) . For remote threshold
    // events this is scaled for comparison with the Breaching Value. This is to
    // account for different local and remote window sizes. The type is
    // interface{} with range: 0..18446744073709551615.
    LocalHighThreshold interface{}

    // Breaching value (If applicable). The type is interface{} with range:
    // 0..18446744073709551615.
    Value interface{}

    // The running total number of errors seen since OAM was enabled on the
    // interface(If applicable). The type is interface{} with range:
    // 0..18446744073709551615.
    RunningTotal interface{}

    // The window in the units that are currently configured. The type is
    // interface{} with range: 0..18446744073709551615.
    WindowConfigUnits interface{}

    // The units in which the window size is configured . The type is interface{}
    // with range: 0..255.
    WindowUnits interface{}

    // The threshold in the units that are currently configured. The type is
    // interface{} with range: 0..18446744073709551615.
    ThresholdConfigUnits interface{}

    // The units in which the threshold size is configured. The type is
    // interface{} with range: 0..255.
    ThresholdUnits interface{}

    // The local high threshold in the units that are currently configured. The
    // type is interface{} with range: 0..18446744073709551615.
    LocalHighThresholdConfigUnits interface{}

    // The breaching value in the units that are currently configured. The type is
    // interface{} with range: 0..18446744073709551615.
    ValueConfigUnits interface{}
}

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetFilter() yfilter.YFilter { return eventLogEntryIndex.YFilter }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) SetFilter(yf yfilter.YFilter) { eventLogEntryIndex.YFilter = yf }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetGoName(yname string) string {
    if yname == "event-log-entry-index" { return "EventLogEntryIndex" }
    if yname == "index" { return "Index" }
    if yname == "handle" { return "Handle" }
    if yname == "oui" { return "Oui" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "type" { return "Type" }
    if yname == "location" { return "Location" }
    if yname == "event-total" { return "EventTotal" }
    if yname == "action-taken" { return "ActionTaken" }
    if yname == "window" { return "Window" }
    if yname == "threshold" { return "Threshold" }
    if yname == "local-high-threshold" { return "LocalHighThreshold" }
    if yname == "value" { return "Value" }
    if yname == "running-total" { return "RunningTotal" }
    if yname == "window-config-units" { return "WindowConfigUnits" }
    if yname == "window-units" { return "WindowUnits" }
    if yname == "threshold-config-units" { return "ThresholdConfigUnits" }
    if yname == "threshold-units" { return "ThresholdUnits" }
    if yname == "local-high-threshold-config-units" { return "LocalHighThresholdConfigUnits" }
    if yname == "value-config-units" { return "ValueConfigUnits" }
    return ""
}

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetSegmentPath() string {
    return "event-log-entry-index" + "[event-log-entry-index='" + fmt.Sprintf("%v", eventLogEntryIndex.EventLogEntryIndex) + "']"
}

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-log-entry-index"] = eventLogEntryIndex.EventLogEntryIndex
    leafs["index"] = eventLogEntryIndex.Index
    leafs["handle"] = eventLogEntryIndex.Handle
    leafs["oui"] = eventLogEntryIndex.Oui
    leafs["timestamp"] = eventLogEntryIndex.Timestamp
    leafs["type"] = eventLogEntryIndex.Type
    leafs["location"] = eventLogEntryIndex.Location
    leafs["event-total"] = eventLogEntryIndex.EventTotal
    leafs["action-taken"] = eventLogEntryIndex.ActionTaken
    leafs["window"] = eventLogEntryIndex.Window
    leafs["threshold"] = eventLogEntryIndex.Threshold
    leafs["local-high-threshold"] = eventLogEntryIndex.LocalHighThreshold
    leafs["value"] = eventLogEntryIndex.Value
    leafs["running-total"] = eventLogEntryIndex.RunningTotal
    leafs["window-config-units"] = eventLogEntryIndex.WindowConfigUnits
    leafs["window-units"] = eventLogEntryIndex.WindowUnits
    leafs["threshold-config-units"] = eventLogEntryIndex.ThresholdConfigUnits
    leafs["threshold-units"] = eventLogEntryIndex.ThresholdUnits
    leafs["local-high-threshold-config-units"] = eventLogEntryIndex.LocalHighThresholdConfigUnits
    leafs["value-config-units"] = eventLogEntryIndex.ValueConfigUnits
    return leafs
}

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetBundleName() string { return "cisco_ios_xr" }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetYangName() string { return "event-log-entry-index" }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) SetParent(parent types.Entity) { eventLogEntryIndex.parent = parent }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetParent() types.Entity { return eventLogEntryIndex.parent }

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetParentYangName() string { return "event-log-entry-indexes" }

// EtherLinkOam_StatsInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Stats container
type EtherLinkOam_StatsInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet Link OAM interface to get Stats for. The type is slice of
    // EtherLinkOam_StatsInterfaces_StatsInterface.
    StatsInterface []EtherLinkOam_StatsInterfaces_StatsInterface
}

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetFilter() yfilter.YFilter { return statsInterfaces.YFilter }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) SetFilter(yf yfilter.YFilter) { statsInterfaces.YFilter = yf }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetGoName(yname string) string {
    if yname == "stats-interface" { return "StatsInterface" }
    return ""
}

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetSegmentPath() string {
    return "stats-interfaces"
}

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats-interface" {
        for _, c := range statsInterfaces.StatsInterface {
            if statsInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EtherLinkOam_StatsInterfaces_StatsInterface{}
        statsInterfaces.StatsInterface = append(statsInterfaces.StatsInterface, child)
        return &statsInterfaces.StatsInterface[len(statsInterfaces.StatsInterface)-1]
    }
    return nil
}

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statsInterfaces.StatsInterface {
        children[statsInterfaces.StatsInterface[i].GetSegmentPath()] = &statsInterfaces.StatsInterface[i]
    }
    return children
}

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetYangName() string { return "stats-interfaces" }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) SetParent(parent types.Entity) { statsInterfaces.parent = parent }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetParent() types.Entity { return statsInterfaces.parent }

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetParentYangName() string { return "ether-link-oam" }

// EtherLinkOam_StatsInterfaces_StatsInterface
// Ethernet Link OAM interface to get Stats for
type EtherLinkOam_StatsInterfaces_StatsInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    MemberInterface interface{}

    // Number of information OAMPDUs transmitted. The type is interface{} with
    // range: 0..4294967295.
    InformationTx interface{}

    // Number of information OAMPDUs received. The type is interface{} with range:
    // 0..4294967295.
    InformationRx interface{}

    // Number of unique event notification OAMPDUs transmitted. The type is
    // interface{} with range: 0..4294967295.
    UniqueEventNotificationTx interface{}

    // Number of unique event notification OAMPDUs received. The type is
    // interface{} with range: 0..4294967295.
    UniqueEventNotificationRx interface{}

    // Number of duplicate event notification OAMPDUs transmitted. The type is
    // interface{} with range: 0..4294967295.
    DuplicateEventNotificationTx interface{}

    // Number of duplicate event notification OAMPDUs received. The type is
    // interface{} with range: 0..4294967295.
    DuplicateEventNotificationRx interface{}

    // Number of loopback control OAMPDUs transmitted. The type is interface{}
    // with range: 0..4294967295.
    LoopbackControlTx interface{}

    // Number of loopback control OAMPDUs received. The type is interface{} with
    // range: 0..4294967295.
    LoopbackControlRx interface{}

    // Number of variable request OAMPDUs transmitted. The type is interface{}
    // with range: 0..4294967295.
    VariableRequestTx interface{}

    // Number of variable request OAMPDUs received. The type is interface{} with
    // range: 0..4294967295.
    VariableRequestRx interface{}

    // Number of variable response OAMPDUs transmitted. The type is interface{}
    // with range: 0..4294967295.
    VariableResponseTx interface{}

    // Number of variable response OAMPDUs received. The type is interface{} with
    // range: 0..4294967295.
    VariableResponseRx interface{}

    // Number of organization specific OAMPDUs transmitted. The type is
    // interface{} with range: 0..4294967295.
    OrgSpecificTx interface{}

    // Number of organization specific OAMPDUs received. The type is interface{}
    // with range: 0..4294967295.
    OrgSpecificRx interface{}

    // Number of OAMPDUs with unsupported codes transmitted. The type is
    // interface{} with range: 0..4294967295.
    UnsupportedCodesTx interface{}

    // Number of OAMPDUs with unsupported codes received. The type is interface{}
    // with range: 0..4294967295.
    UnsupportedCodesRx interface{}

    // Number of frames lost due to OAM. The type is interface{} with range:
    // 0..4294967295.
    FramesLostDueToOam interface{}

    // Number of RX frames 'fixed' by OAM. The type is interface{} with range:
    // 0..4294967295.
    FixedFramesRx interface{}

    // Number of local error symbol period records. The type is interface{} with
    // range: 0..4294967295.
    LocalErrorSymbolPeriodRecords interface{}

    // Number of local error frame records. The type is interface{} with range:
    // 0..4294967295.
    LocalErrorFrameRecords interface{}

    // Number of local error frame period records. The type is interface{} with
    // range: 0..4294967295.
    LocalErrorFramePeriodRecords interface{}

    // Number of local error frame second records. The type is interface{} with
    // range: 0..4294967295.
    LocalErrorFrameSecondRecords interface{}

    // Number of remote error symbol period records. The type is interface{} with
    // range: 0..4294967295.
    RemoteErrorSymbolPeriodRecords interface{}

    // Number of remote error frame records. The type is interface{} with range:
    // 0..4294967295.
    RemoteErrorFrameRecords interface{}

    // Number of remote error frame period records. The type is interface{} with
    // range: 0..4294967295.
    RemoteErrorFramePeriodRecords interface{}

    // Number of remote error frame second records. The type is interface{} with
    // range: 0..4294967295.
    RemoteErrorFrameSecondRecords interface{}
}

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetFilter() yfilter.YFilter { return statsInterface.YFilter }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) SetFilter(yf yfilter.YFilter) { statsInterface.YFilter = yf }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetGoName(yname string) string {
    if yname == "member-interface" { return "MemberInterface" }
    if yname == "information-tx" { return "InformationTx" }
    if yname == "information-rx" { return "InformationRx" }
    if yname == "unique-event-notification-tx" { return "UniqueEventNotificationTx" }
    if yname == "unique-event-notification-rx" { return "UniqueEventNotificationRx" }
    if yname == "duplicate-event-notification-tx" { return "DuplicateEventNotificationTx" }
    if yname == "duplicate-event-notification-rx" { return "DuplicateEventNotificationRx" }
    if yname == "loopback-control-tx" { return "LoopbackControlTx" }
    if yname == "loopback-control-rx" { return "LoopbackControlRx" }
    if yname == "variable-request-tx" { return "VariableRequestTx" }
    if yname == "variable-request-rx" { return "VariableRequestRx" }
    if yname == "variable-response-tx" { return "VariableResponseTx" }
    if yname == "variable-response-rx" { return "VariableResponseRx" }
    if yname == "org-specific-tx" { return "OrgSpecificTx" }
    if yname == "org-specific-rx" { return "OrgSpecificRx" }
    if yname == "unsupported-codes-tx" { return "UnsupportedCodesTx" }
    if yname == "unsupported-codes-rx" { return "UnsupportedCodesRx" }
    if yname == "frames-lost-due-to-oam" { return "FramesLostDueToOam" }
    if yname == "fixed-frames-rx" { return "FixedFramesRx" }
    if yname == "local-error-symbol-period-records" { return "LocalErrorSymbolPeriodRecords" }
    if yname == "local-error-frame-records" { return "LocalErrorFrameRecords" }
    if yname == "local-error-frame-period-records" { return "LocalErrorFramePeriodRecords" }
    if yname == "local-error-frame-second-records" { return "LocalErrorFrameSecondRecords" }
    if yname == "remote-error-symbol-period-records" { return "RemoteErrorSymbolPeriodRecords" }
    if yname == "remote-error-frame-records" { return "RemoteErrorFrameRecords" }
    if yname == "remote-error-frame-period-records" { return "RemoteErrorFramePeriodRecords" }
    if yname == "remote-error-frame-second-records" { return "RemoteErrorFrameSecondRecords" }
    return ""
}

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetSegmentPath() string {
    return "stats-interface" + "[member-interface='" + fmt.Sprintf("%v", statsInterface.MemberInterface) + "']"
}

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["member-interface"] = statsInterface.MemberInterface
    leafs["information-tx"] = statsInterface.InformationTx
    leafs["information-rx"] = statsInterface.InformationRx
    leafs["unique-event-notification-tx"] = statsInterface.UniqueEventNotificationTx
    leafs["unique-event-notification-rx"] = statsInterface.UniqueEventNotificationRx
    leafs["duplicate-event-notification-tx"] = statsInterface.DuplicateEventNotificationTx
    leafs["duplicate-event-notification-rx"] = statsInterface.DuplicateEventNotificationRx
    leafs["loopback-control-tx"] = statsInterface.LoopbackControlTx
    leafs["loopback-control-rx"] = statsInterface.LoopbackControlRx
    leafs["variable-request-tx"] = statsInterface.VariableRequestTx
    leafs["variable-request-rx"] = statsInterface.VariableRequestRx
    leafs["variable-response-tx"] = statsInterface.VariableResponseTx
    leafs["variable-response-rx"] = statsInterface.VariableResponseRx
    leafs["org-specific-tx"] = statsInterface.OrgSpecificTx
    leafs["org-specific-rx"] = statsInterface.OrgSpecificRx
    leafs["unsupported-codes-tx"] = statsInterface.UnsupportedCodesTx
    leafs["unsupported-codes-rx"] = statsInterface.UnsupportedCodesRx
    leafs["frames-lost-due-to-oam"] = statsInterface.FramesLostDueToOam
    leafs["fixed-frames-rx"] = statsInterface.FixedFramesRx
    leafs["local-error-symbol-period-records"] = statsInterface.LocalErrorSymbolPeriodRecords
    leafs["local-error-frame-records"] = statsInterface.LocalErrorFrameRecords
    leafs["local-error-frame-period-records"] = statsInterface.LocalErrorFramePeriodRecords
    leafs["local-error-frame-second-records"] = statsInterface.LocalErrorFrameSecondRecords
    leafs["remote-error-symbol-period-records"] = statsInterface.RemoteErrorSymbolPeriodRecords
    leafs["remote-error-frame-records"] = statsInterface.RemoteErrorFrameRecords
    leafs["remote-error-frame-period-records"] = statsInterface.RemoteErrorFramePeriodRecords
    leafs["remote-error-frame-second-records"] = statsInterface.RemoteErrorFrameSecondRecords
    return leafs
}

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetBundleName() string { return "cisco_ios_xr" }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetYangName() string { return "stats-interface" }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) SetParent(parent types.Entity) { statsInterface.parent = parent }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetParent() types.Entity { return statsInterface.parent }

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetParentYangName() string { return "stats-interfaces" }

