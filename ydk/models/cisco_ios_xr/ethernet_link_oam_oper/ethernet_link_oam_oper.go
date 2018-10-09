// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-link-oam package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ether-link-oam: Ethernet Link OAM operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
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

func (etherLinkOam *EtherLinkOam) GetEntityData() *types.CommonEntityData {
    etherLinkOam.EntityData.YFilter = etherLinkOam.YFilter
    etherLinkOam.EntityData.YangName = "ether-link-oam"
    etherLinkOam.EntityData.BundleName = "cisco_ios_xr"
    etherLinkOam.EntityData.ParentYangName = "Cisco-IOS-XR-ethernet-link-oam-oper"
    etherLinkOam.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-link-oam-oper:ether-link-oam"
    etherLinkOam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    etherLinkOam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    etherLinkOam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    etherLinkOam.EntityData.Children = types.NewOrderedMap()
    etherLinkOam.EntityData.Children.Append("discovery-info-interfaces", types.YChild{"DiscoveryInfoInterfaces", &etherLinkOam.DiscoveryInfoInterfaces})
    etherLinkOam.EntityData.Children.Append("interface-state-interfaces", types.YChild{"InterfaceStateInterfaces", &etherLinkOam.InterfaceStateInterfaces})
    etherLinkOam.EntityData.Children.Append("running-config-interfaces", types.YChild{"RunningConfigInterfaces", &etherLinkOam.RunningConfigInterfaces})
    etherLinkOam.EntityData.Children.Append("nodes", types.YChild{"Nodes", &etherLinkOam.Nodes})
    etherLinkOam.EntityData.Children.Append("event-log-entry-interfaces", types.YChild{"EventLogEntryInterfaces", &etherLinkOam.EventLogEntryInterfaces})
    etherLinkOam.EntityData.Children.Append("stats-interfaces", types.YChild{"StatsInterfaces", &etherLinkOam.StatsInterfaces})
    etherLinkOam.EntityData.Leafs = types.NewOrderedMap()

    etherLinkOam.EntityData.YListKeys = []string {}

    return &(etherLinkOam.EntityData)
}

// EtherLinkOam_DiscoveryInfoInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Discovery Info container
type EtherLinkOam_DiscoveryInfoInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet Link OAM interface to get Discovery Info for. The type is slice of
    // EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface.
    DiscoveryInfoInterface []*EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface
}

func (discoveryInfoInterfaces *EtherLinkOam_DiscoveryInfoInterfaces) GetEntityData() *types.CommonEntityData {
    discoveryInfoInterfaces.EntityData.YFilter = discoveryInfoInterfaces.YFilter
    discoveryInfoInterfaces.EntityData.YangName = "discovery-info-interfaces"
    discoveryInfoInterfaces.EntityData.BundleName = "cisco_ios_xr"
    discoveryInfoInterfaces.EntityData.ParentYangName = "ether-link-oam"
    discoveryInfoInterfaces.EntityData.SegmentPath = "discovery-info-interfaces"
    discoveryInfoInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discoveryInfoInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discoveryInfoInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discoveryInfoInterfaces.EntityData.Children = types.NewOrderedMap()
    discoveryInfoInterfaces.EntityData.Children.Append("discovery-info-interface", types.YChild{"DiscoveryInfoInterface", nil})
    for i := range discoveryInfoInterfaces.DiscoveryInfoInterface {
        discoveryInfoInterfaces.EntityData.Children.Append(types.GetSegmentPath(discoveryInfoInterfaces.DiscoveryInfoInterface[i]), types.YChild{"DiscoveryInfoInterface", discoveryInfoInterfaces.DiscoveryInfoInterface[i]})
    }
    discoveryInfoInterfaces.EntityData.Leafs = types.NewOrderedMap()

    discoveryInfoInterfaces.EntityData.YListKeys = []string {}

    return &(discoveryInfoInterfaces.EntityData)
}

// EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface
// Ethernet Link OAM interface to get Discovery
// Info for
type EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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

func (discoveryInfoInterface *EtherLinkOam_DiscoveryInfoInterfaces_DiscoveryInfoInterface) GetEntityData() *types.CommonEntityData {
    discoveryInfoInterface.EntityData.YFilter = discoveryInfoInterface.YFilter
    discoveryInfoInterface.EntityData.YangName = "discovery-info-interface"
    discoveryInfoInterface.EntityData.BundleName = "cisco_ios_xr"
    discoveryInfoInterface.EntityData.ParentYangName = "discovery-info-interfaces"
    discoveryInfoInterface.EntityData.SegmentPath = "discovery-info-interface" + types.AddKeyToken(discoveryInfoInterface.MemberInterface, "member-interface")
    discoveryInfoInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discoveryInfoInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discoveryInfoInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discoveryInfoInterface.EntityData.Children = types.NewOrderedMap()
    discoveryInfoInterface.EntityData.Leafs = types.NewOrderedMap()
    discoveryInfoInterface.EntityData.Leafs.Append("member-interface", types.YLeaf{"MemberInterface", discoveryInfoInterface.MemberInterface})
    discoveryInfoInterface.EntityData.Leafs.Append("name", types.YLeaf{"Name", discoveryInfoInterface.Name})
    discoveryInfoInterface.EntityData.Leafs.Append("operational-status", types.YLeaf{"OperationalStatus", discoveryInfoInterface.OperationalStatus})
    discoveryInfoInterface.EntityData.Leafs.Append("loopback-mode", types.YLeaf{"LoopbackMode", discoveryInfoInterface.LoopbackMode})
    discoveryInfoInterface.EntityData.Leafs.Append("local-mode", types.YLeaf{"LocalMode", discoveryInfoInterface.LocalMode})
    discoveryInfoInterface.EntityData.Leafs.Append("miswired", types.YLeaf{"Miswired", discoveryInfoInterface.Miswired})
    discoveryInfoInterface.EntityData.Leafs.Append("local-mwd-key", types.YLeaf{"LocalMwdKey", discoveryInfoInterface.LocalMwdKey})
    discoveryInfoInterface.EntityData.Leafs.Append("local-function-unidirectional", types.YLeaf{"LocalFunctionUnidirectional", discoveryInfoInterface.LocalFunctionUnidirectional})
    discoveryInfoInterface.EntityData.Leafs.Append("local-function-loopback", types.YLeaf{"LocalFunctionLoopback", discoveryInfoInterface.LocalFunctionLoopback})
    discoveryInfoInterface.EntityData.Leafs.Append("local-function-event", types.YLeaf{"LocalFunctionEvent", discoveryInfoInterface.LocalFunctionEvent})
    discoveryInfoInterface.EntityData.Leafs.Append("local-functionvariable", types.YLeaf{"LocalFunctionvariable", discoveryInfoInterface.LocalFunctionvariable})
    discoveryInfoInterface.EntityData.Leafs.Append("local-revision", types.YLeaf{"LocalRevision", discoveryInfoInterface.LocalRevision})
    discoveryInfoInterface.EntityData.Leafs.Append("local-mtu", types.YLeaf{"LocalMtu", discoveryInfoInterface.LocalMtu})
    discoveryInfoInterface.EntityData.Leafs.Append("local-operational", types.YLeaf{"LocalOperational", discoveryInfoInterface.LocalOperational})
    discoveryInfoInterface.EntityData.Leafs.Append("local-evaluating", types.YLeaf{"LocalEvaluating", discoveryInfoInterface.LocalEvaluating})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-mode", types.YLeaf{"RemoteMode", discoveryInfoInterface.RemoteMode})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-unidirectional", types.YLeaf{"RemoteUnidirectional", discoveryInfoInterface.RemoteUnidirectional})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-loopback", types.YLeaf{"RemoteLoopback", discoveryInfoInterface.RemoteLoopback})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-event", types.YLeaf{"RemoteEvent", discoveryInfoInterface.RemoteEvent})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-variable", types.YLeaf{"RemoteVariable", discoveryInfoInterface.RemoteVariable})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-mtu", types.YLeaf{"RemoteMtu", discoveryInfoInterface.RemoteMtu})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-mac-address", types.YLeaf{"RemoteMacAddress", discoveryInfoInterface.RemoteMacAddress})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-vendor-oui", types.YLeaf{"RemoteVendorOui", discoveryInfoInterface.RemoteVendorOui})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-revision", types.YLeaf{"RemoteRevision", discoveryInfoInterface.RemoteRevision})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-vendor-info", types.YLeaf{"RemoteVendorInfo", discoveryInfoInterface.RemoteVendorInfo})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-mwd-key", types.YLeaf{"RemoteMwdKey", discoveryInfoInterface.RemoteMwdKey})
    discoveryInfoInterface.EntityData.Leafs.Append("operational-status-valid", types.YLeaf{"OperationalStatusValid", discoveryInfoInterface.OperationalStatusValid})
    discoveryInfoInterface.EntityData.Leafs.Append("loopback-mode-valid", types.YLeaf{"LoopbackModeValid", discoveryInfoInterface.LoopbackModeValid})
    discoveryInfoInterface.EntityData.Leafs.Append("local-mode-valid", types.YLeaf{"LocalModeValid", discoveryInfoInterface.LocalModeValid})
    discoveryInfoInterface.EntityData.Leafs.Append("miswired-valid", types.YLeaf{"MiswiredValid", discoveryInfoInterface.MiswiredValid})
    discoveryInfoInterface.EntityData.Leafs.Append("local-mwd-key-valid", types.YLeaf{"LocalMwdKeyValid", discoveryInfoInterface.LocalMwdKeyValid})
    discoveryInfoInterface.EntityData.Leafs.Append("local-function-unidirectional-valid", types.YLeaf{"LocalFunctionUnidirectionalValid", discoveryInfoInterface.LocalFunctionUnidirectionalValid})
    discoveryInfoInterface.EntityData.Leafs.Append("local-function-loopback-valid", types.YLeaf{"LocalFunctionLoopbackValid", discoveryInfoInterface.LocalFunctionLoopbackValid})
    discoveryInfoInterface.EntityData.Leafs.Append("local-function-event-valid", types.YLeaf{"LocalFunctionEventValid", discoveryInfoInterface.LocalFunctionEventValid})
    discoveryInfoInterface.EntityData.Leafs.Append("local-functionvariable-valid", types.YLeaf{"LocalFunctionvariableValid", discoveryInfoInterface.LocalFunctionvariableValid})
    discoveryInfoInterface.EntityData.Leafs.Append("local-revisionvalid", types.YLeaf{"LocalRevisionvalid", discoveryInfoInterface.LocalRevisionvalid})
    discoveryInfoInterface.EntityData.Leafs.Append("local-mtu-valid", types.YLeaf{"LocalMtuValid", discoveryInfoInterface.LocalMtuValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-mode-valid", types.YLeaf{"RemoteModeValid", discoveryInfoInterface.RemoteModeValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-unidirectional-valid", types.YLeaf{"RemoteUnidirectionalValid", discoveryInfoInterface.RemoteUnidirectionalValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-loopback-valid", types.YLeaf{"RemoteLoopbackValid", discoveryInfoInterface.RemoteLoopbackValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-event-valid", types.YLeaf{"RemoteEventValid", discoveryInfoInterface.RemoteEventValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-variable-valid", types.YLeaf{"RemoteVariableValid", discoveryInfoInterface.RemoteVariableValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-mtu-valid", types.YLeaf{"RemoteMtuValid", discoveryInfoInterface.RemoteMtuValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-mac-address-valid", types.YLeaf{"RemoteMacAddressValid", discoveryInfoInterface.RemoteMacAddressValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-vendor-oui-valid", types.YLeaf{"RemoteVendorOuiValid", discoveryInfoInterface.RemoteVendorOuiValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-revisionvalid", types.YLeaf{"RemoteRevisionvalid", discoveryInfoInterface.RemoteRevisionvalid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-vendor-info-valid", types.YLeaf{"RemoteVendorInfoValid", discoveryInfoInterface.RemoteVendorInfoValid})
    discoveryInfoInterface.EntityData.Leafs.Append("remote-mwd-key-valid", types.YLeaf{"RemoteMwdKeyValid", discoveryInfoInterface.RemoteMwdKeyValid})
    discoveryInfoInterface.EntityData.Leafs.Append("received-at-risk-notification-timestamp", types.YLeaf{"ReceivedAtRiskNotificationTimestamp", discoveryInfoInterface.ReceivedAtRiskNotificationTimestamp})
    discoveryInfoInterface.EntityData.Leafs.Append("received-at-risk-notification-time-remaining", types.YLeaf{"ReceivedAtRiskNotificationTimeRemaining", discoveryInfoInterface.ReceivedAtRiskNotificationTimeRemaining})

    discoveryInfoInterface.EntityData.YListKeys = []string {"MemberInterface"}

    return &(discoveryInfoInterface.EntityData)
}

// EtherLinkOam_InterfaceStateInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Interface State container
type EtherLinkOam_InterfaceStateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet Link OAM interface to get Interface State for. The type is slice
    // of EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface.
    InterfaceStateInterface []*EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface
}

func (interfaceStateInterfaces *EtherLinkOam_InterfaceStateInterfaces) GetEntityData() *types.CommonEntityData {
    interfaceStateInterfaces.EntityData.YFilter = interfaceStateInterfaces.YFilter
    interfaceStateInterfaces.EntityData.YangName = "interface-state-interfaces"
    interfaceStateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaceStateInterfaces.EntityData.ParentYangName = "ether-link-oam"
    interfaceStateInterfaces.EntityData.SegmentPath = "interface-state-interfaces"
    interfaceStateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStateInterfaces.EntityData.Children = types.NewOrderedMap()
    interfaceStateInterfaces.EntityData.Children.Append("interface-state-interface", types.YChild{"InterfaceStateInterface", nil})
    for i := range interfaceStateInterfaces.InterfaceStateInterface {
        interfaceStateInterfaces.EntityData.Children.Append(types.GetSegmentPath(interfaceStateInterfaces.InterfaceStateInterface[i]), types.YChild{"InterfaceStateInterface", interfaceStateInterfaces.InterfaceStateInterface[i]})
    }
    interfaceStateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaceStateInterfaces.EntityData.YListKeys = []string {}

    return &(interfaceStateInterfaces.EntityData)
}

// EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface
// Ethernet Link OAM interface to get Interface
// State for
type EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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

func (interfaceStateInterface *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface) GetEntityData() *types.CommonEntityData {
    interfaceStateInterface.EntityData.YFilter = interfaceStateInterface.YFilter
    interfaceStateInterface.EntityData.YangName = "interface-state-interface"
    interfaceStateInterface.EntityData.BundleName = "cisco_ios_xr"
    interfaceStateInterface.EntityData.ParentYangName = "interface-state-interfaces"
    interfaceStateInterface.EntityData.SegmentPath = "interface-state-interface" + types.AddKeyToken(interfaceStateInterface.MemberInterface, "member-interface")
    interfaceStateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStateInterface.EntityData.Children = types.NewOrderedMap()
    interfaceStateInterface.EntityData.Children.Append("errors", types.YChild{"Errors", &interfaceStateInterface.Errors})
    interfaceStateInterface.EntityData.Children.Append("efd-triggers", types.YChild{"EfdTriggers", &interfaceStateInterface.EfdTriggers})
    interfaceStateInterface.EntityData.Leafs = types.NewOrderedMap()
    interfaceStateInterface.EntityData.Leafs.Append("member-interface", types.YLeaf{"MemberInterface", interfaceStateInterface.MemberInterface})
    interfaceStateInterface.EntityData.Leafs.Append("protocol-code", types.YLeaf{"ProtocolCode", interfaceStateInterface.ProtocolCode})
    interfaceStateInterface.EntityData.Leafs.Append("rx-fault", types.YLeaf{"RxFault", interfaceStateInterface.RxFault})
    interfaceStateInterface.EntityData.Leafs.Append("local-mwd-key", types.YLeaf{"LocalMwdKey", interfaceStateInterface.LocalMwdKey})
    interfaceStateInterface.EntityData.Leafs.Append("remote-mwd-key-present", types.YLeaf{"RemoteMwdKeyPresent", interfaceStateInterface.RemoteMwdKeyPresent})
    interfaceStateInterface.EntityData.Leafs.Append("remote-mwd-key", types.YLeaf{"RemoteMwdKey", interfaceStateInterface.RemoteMwdKey})

    interfaceStateInterface.EntityData.YListKeys = []string {"MemberInterface"}

    return &(interfaceStateInterface.EntityData)
}

// EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors
// The errors that have occurred on this interface
type EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors struct {
    EntityData types.CommonEntityData
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

func (errors *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_Errors) GetEntityData() *types.CommonEntityData {
    errors.EntityData.YFilter = errors.YFilter
    errors.EntityData.YangName = "errors"
    errors.EntityData.BundleName = "cisco_ios_xr"
    errors.EntityData.ParentYangName = "interface-state-interface"
    errors.EntityData.SegmentPath = "errors"
    errors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errors.EntityData.Children = types.NewOrderedMap()
    errors.EntityData.Leafs = types.NewOrderedMap()
    errors.EntityData.Leafs.Append("pfi-reason", types.YLeaf{"PfiReason", errors.PfiReason})
    errors.EntityData.Leafs.Append("pfi-error-code", types.YLeaf{"PfiErrorCode", errors.PfiErrorCode})
    errors.EntityData.Leafs.Append("platform-reason", types.YLeaf{"PlatformReason", errors.PlatformReason})
    errors.EntityData.Leafs.Append("platform-error-code", types.YLeaf{"PlatformErrorCode", errors.PlatformErrorCode})
    errors.EntityData.Leafs.Append("spio-reason", types.YLeaf{"SpioReason", errors.SpioReason})
    errors.EntityData.Leafs.Append("spio-error-code", types.YLeaf{"SpioErrorCode", errors.SpioErrorCode})
    errors.EntityData.Leafs.Append("epi-reason", types.YLeaf{"EpiReason", errors.EpiReason})
    errors.EntityData.Leafs.Append("epi-error-code", types.YLeaf{"EpiErrorCode", errors.EpiErrorCode})
    errors.EntityData.Leafs.Append("caps-add-reason", types.YLeaf{"CapsAddReason", errors.CapsAddReason})
    errors.EntityData.Leafs.Append("caps-add-error-code", types.YLeaf{"CapsAddErrorCode", errors.CapsAddErrorCode})

    errors.EntityData.YListKeys = []string {}

    return &(errors.EntityData)
}

// EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers
// Any present EFD triggers
type EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers struct {
    EntityData types.CommonEntityData
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

func (efdTriggers *EtherLinkOam_InterfaceStateInterfaces_InterfaceStateInterface_EfdTriggers) GetEntityData() *types.CommonEntityData {
    efdTriggers.EntityData.YFilter = efdTriggers.YFilter
    efdTriggers.EntityData.YangName = "efd-triggers"
    efdTriggers.EntityData.BundleName = "cisco_ios_xr"
    efdTriggers.EntityData.ParentYangName = "interface-state-interface"
    efdTriggers.EntityData.SegmentPath = "efd-triggers"
    efdTriggers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    efdTriggers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    efdTriggers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    efdTriggers.EntityData.Children = types.NewOrderedMap()
    efdTriggers.EntityData.Leafs = types.NewOrderedMap()
    efdTriggers.EntityData.Leafs.Append("link-fault-received", types.YLeaf{"LinkFaultReceived", efdTriggers.LinkFaultReceived})
    efdTriggers.EntityData.Leafs.Append("discovery-timed-out", types.YLeaf{"DiscoveryTimedOut", efdTriggers.DiscoveryTimedOut})
    efdTriggers.EntityData.Leafs.Append("capabilities-conflict", types.YLeaf{"CapabilitiesConflict", efdTriggers.CapabilitiesConflict})
    efdTriggers.EntityData.Leafs.Append("wiring-conflict", types.YLeaf{"WiringConflict", efdTriggers.WiringConflict})
    efdTriggers.EntityData.Leafs.Append("session-down", types.YLeaf{"SessionDown", efdTriggers.SessionDown})

    efdTriggers.EntityData.YListKeys = []string {}

    return &(efdTriggers.EntityData)
}

// EtherLinkOam_RunningConfigInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Running Config container
type EtherLinkOam_RunningConfigInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet Link OAM interface to get Running Config for. The type is slice of
    // EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface.
    RunningConfigInterface []*EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface
}

func (runningConfigInterfaces *EtherLinkOam_RunningConfigInterfaces) GetEntityData() *types.CommonEntityData {
    runningConfigInterfaces.EntityData.YFilter = runningConfigInterfaces.YFilter
    runningConfigInterfaces.EntityData.YangName = "running-config-interfaces"
    runningConfigInterfaces.EntityData.BundleName = "cisco_ios_xr"
    runningConfigInterfaces.EntityData.ParentYangName = "ether-link-oam"
    runningConfigInterfaces.EntityData.SegmentPath = "running-config-interfaces"
    runningConfigInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    runningConfigInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    runningConfigInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    runningConfigInterfaces.EntityData.Children = types.NewOrderedMap()
    runningConfigInterfaces.EntityData.Children.Append("running-config-interface", types.YChild{"RunningConfigInterface", nil})
    for i := range runningConfigInterfaces.RunningConfigInterface {
        runningConfigInterfaces.EntityData.Children.Append(types.GetSegmentPath(runningConfigInterfaces.RunningConfigInterface[i]), types.YChild{"RunningConfigInterface", runningConfigInterfaces.RunningConfigInterface[i]})
    }
    runningConfigInterfaces.EntityData.Leafs = types.NewOrderedMap()

    runningConfigInterfaces.EntityData.YListKeys = []string {}

    return &(runningConfigInterfaces.EntityData)
}

// EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface
// Ethernet Link OAM interface to get Running
// Config for
type EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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

func (runningConfigInterface *EtherLinkOam_RunningConfigInterfaces_RunningConfigInterface) GetEntityData() *types.CommonEntityData {
    runningConfigInterface.EntityData.YFilter = runningConfigInterface.YFilter
    runningConfigInterface.EntityData.YangName = "running-config-interface"
    runningConfigInterface.EntityData.BundleName = "cisco_ios_xr"
    runningConfigInterface.EntityData.ParentYangName = "running-config-interfaces"
    runningConfigInterface.EntityData.SegmentPath = "running-config-interface" + types.AddKeyToken(runningConfigInterface.MemberInterface, "member-interface")
    runningConfigInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    runningConfigInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    runningConfigInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    runningConfigInterface.EntityData.Children = types.NewOrderedMap()
    runningConfigInterface.EntityData.Leafs = types.NewOrderedMap()
    runningConfigInterface.EntityData.Leafs.Append("member-interface", types.YLeaf{"MemberInterface", runningConfigInterface.MemberInterface})
    runningConfigInterface.EntityData.Leafs.Append("fast-hello-interval-enabled", types.YLeaf{"FastHelloIntervalEnabled", runningConfigInterface.FastHelloIntervalEnabled})
    runningConfigInterface.EntityData.Leafs.Append("link-monitor-enabled", types.YLeaf{"LinkMonitorEnabled", runningConfigInterface.LinkMonitorEnabled})
    runningConfigInterface.EntityData.Leafs.Append("remote-loopback-enabled", types.YLeaf{"RemoteLoopbackEnabled", runningConfigInterface.RemoteLoopbackEnabled})
    runningConfigInterface.EntityData.Leafs.Append("mib-retrieval-enabled", types.YLeaf{"MibRetrievalEnabled", runningConfigInterface.MibRetrievalEnabled})
    runningConfigInterface.EntityData.Leafs.Append("udlf-enabled", types.YLeaf{"UdlfEnabled", runningConfigInterface.UdlfEnabled})
    runningConfigInterface.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", runningConfigInterface.Mode})
    runningConfigInterface.EntityData.Leafs.Append("connection-timeout", types.YLeaf{"ConnectionTimeout", runningConfigInterface.ConnectionTimeout})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-window", types.YLeaf{"SymbolPeriodWindow", runningConfigInterface.SymbolPeriodWindow})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-window-units", types.YLeaf{"SymbolPeriodWindowUnits", runningConfigInterface.SymbolPeriodWindowUnits})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-window-multiplier", types.YLeaf{"SymbolPeriodWindowMultiplier", runningConfigInterface.SymbolPeriodWindowMultiplier})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-threshold-low", types.YLeaf{"SymbolPeriodThresholdLow", runningConfigInterface.SymbolPeriodThresholdLow})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-threshold-high", types.YLeaf{"SymbolPeriodThresholdHigh", runningConfigInterface.SymbolPeriodThresholdHigh})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-threshold-units", types.YLeaf{"SymbolPeriodThresholdUnits", runningConfigInterface.SymbolPeriodThresholdUnits})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-threshold-low-multiplier", types.YLeaf{"SymbolPeriodThresholdLowMultiplier", runningConfigInterface.SymbolPeriodThresholdLowMultiplier})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-threshold-high-multiplier", types.YLeaf{"SymbolPeriodThresholdHighMultiplier", runningConfigInterface.SymbolPeriodThresholdHighMultiplier})
    runningConfigInterface.EntityData.Leafs.Append("frame-window", types.YLeaf{"FrameWindow", runningConfigInterface.FrameWindow})
    runningConfigInterface.EntityData.Leafs.Append("frame-threshold-low", types.YLeaf{"FrameThresholdLow", runningConfigInterface.FrameThresholdLow})
    runningConfigInterface.EntityData.Leafs.Append("frame-threshold-high", types.YLeaf{"FrameThresholdHigh", runningConfigInterface.FrameThresholdHigh})
    runningConfigInterface.EntityData.Leafs.Append("frame-threshold-low-multiplier", types.YLeaf{"FrameThresholdLowMultiplier", runningConfigInterface.FrameThresholdLowMultiplier})
    runningConfigInterface.EntityData.Leafs.Append("frame-threshold-high-multiplier", types.YLeaf{"FrameThresholdHighMultiplier", runningConfigInterface.FrameThresholdHighMultiplier})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-window", types.YLeaf{"FramePeriodWindow", runningConfigInterface.FramePeriodWindow})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-window-units", types.YLeaf{"FramePeriodWindowUnits", runningConfigInterface.FramePeriodWindowUnits})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-window-multiplier", types.YLeaf{"FramePeriodWindowMultiplier", runningConfigInterface.FramePeriodWindowMultiplier})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-threshold-low", types.YLeaf{"FramePeriodThresholdLow", runningConfigInterface.FramePeriodThresholdLow})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-threshold-high", types.YLeaf{"FramePeriodThresholdHigh", runningConfigInterface.FramePeriodThresholdHigh})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-threshold-units", types.YLeaf{"FramePeriodThresholdUnits", runningConfigInterface.FramePeriodThresholdUnits})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-threshold-low-multiplier", types.YLeaf{"FramePeriodThresholdLowMultiplier", runningConfigInterface.FramePeriodThresholdLowMultiplier})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-threshold-high-multiplier", types.YLeaf{"FramePeriodThresholdHighMultiplier", runningConfigInterface.FramePeriodThresholdHighMultiplier})
    runningConfigInterface.EntityData.Leafs.Append("frame-seconds-window", types.YLeaf{"FrameSecondsWindow", runningConfigInterface.FrameSecondsWindow})
    runningConfigInterface.EntityData.Leafs.Append("frame-seconds-threshold-low", types.YLeaf{"FrameSecondsThresholdLow", runningConfigInterface.FrameSecondsThresholdLow})
    runningConfigInterface.EntityData.Leafs.Append("frame-seconds-threshold-high", types.YLeaf{"FrameSecondsThresholdHigh", runningConfigInterface.FrameSecondsThresholdHigh})
    runningConfigInterface.EntityData.Leafs.Append("high-threshold-action", types.YLeaf{"HighThresholdAction", runningConfigInterface.HighThresholdAction})
    runningConfigInterface.EntityData.Leafs.Append("link-fault-action", types.YLeaf{"LinkFaultAction", runningConfigInterface.LinkFaultAction})
    runningConfigInterface.EntityData.Leafs.Append("dying-gasp-action", types.YLeaf{"DyingGaspAction", runningConfigInterface.DyingGaspAction})
    runningConfigInterface.EntityData.Leafs.Append("critical-event-action", types.YLeaf{"CriticalEventAction", runningConfigInterface.CriticalEventAction})
    runningConfigInterface.EntityData.Leafs.Append("discovery-timeout-action", types.YLeaf{"DiscoveryTimeoutAction", runningConfigInterface.DiscoveryTimeoutAction})
    runningConfigInterface.EntityData.Leafs.Append("capabilities-conflict-action", types.YLeaf{"CapabilitiesConflictAction", runningConfigInterface.CapabilitiesConflictAction})
    runningConfigInterface.EntityData.Leafs.Append("wiring-conflict-action", types.YLeaf{"WiringConflictAction", runningConfigInterface.WiringConflictAction})
    runningConfigInterface.EntityData.Leafs.Append("session-up-action", types.YLeaf{"SessionUpAction", runningConfigInterface.SessionUpAction})
    runningConfigInterface.EntityData.Leafs.Append("session-down-action", types.YLeaf{"SessionDownAction", runningConfigInterface.SessionDownAction})
    runningConfigInterface.EntityData.Leafs.Append("remote-loopback-action", types.YLeaf{"RemoteLoopbackAction", runningConfigInterface.RemoteLoopbackAction})
    runningConfigInterface.EntityData.Leafs.Append("require-remote-mode", types.YLeaf{"RequireRemoteMode", runningConfigInterface.RequireRemoteMode})
    runningConfigInterface.EntityData.Leafs.Append("require-remote-mib-retrieval", types.YLeaf{"RequireRemoteMibRetrieval", runningConfigInterface.RequireRemoteMibRetrieval})
    runningConfigInterface.EntityData.Leafs.Append("require-loopback", types.YLeaf{"RequireLoopback", runningConfigInterface.RequireLoopback})
    runningConfigInterface.EntityData.Leafs.Append("require-link-monitoring", types.YLeaf{"RequireLinkMonitoring", runningConfigInterface.RequireLinkMonitoring})
    runningConfigInterface.EntityData.Leafs.Append("fast-hello-interval-enabled-overridden", types.YLeaf{"FastHelloIntervalEnabledOverridden", runningConfigInterface.FastHelloIntervalEnabledOverridden})
    runningConfigInterface.EntityData.Leafs.Append("link-monitoring-enabled-overridden", types.YLeaf{"LinkMonitoringEnabledOverridden", runningConfigInterface.LinkMonitoringEnabledOverridden})
    runningConfigInterface.EntityData.Leafs.Append("remote-loopback-enabled-overridden", types.YLeaf{"RemoteLoopbackEnabledOverridden", runningConfigInterface.RemoteLoopbackEnabledOverridden})
    runningConfigInterface.EntityData.Leafs.Append("mib-retrieval-enabled-overridden", types.YLeaf{"MibRetrievalEnabledOverridden", runningConfigInterface.MibRetrievalEnabledOverridden})
    runningConfigInterface.EntityData.Leafs.Append("udlf-enabled-overridden", types.YLeaf{"UdlfEnabledOverridden", runningConfigInterface.UdlfEnabledOverridden})
    runningConfigInterface.EntityData.Leafs.Append("mode-overridden", types.YLeaf{"ModeOverridden", runningConfigInterface.ModeOverridden})
    runningConfigInterface.EntityData.Leafs.Append("connection-timeout-overridden", types.YLeaf{"ConnectionTimeoutOverridden", runningConfigInterface.ConnectionTimeoutOverridden})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-window-overridden", types.YLeaf{"SymbolPeriodWindowOverridden", runningConfigInterface.SymbolPeriodWindowOverridden})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-threshold-low-overridden", types.YLeaf{"SymbolPeriodThresholdLowOverridden", runningConfigInterface.SymbolPeriodThresholdLowOverridden})
    runningConfigInterface.EntityData.Leafs.Append("symbol-period-threshold-high-overridden", types.YLeaf{"SymbolPeriodThresholdHighOverridden", runningConfigInterface.SymbolPeriodThresholdHighOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-window-overridden", types.YLeaf{"FrameWindowOverridden", runningConfigInterface.FrameWindowOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-threshold-low-overridden", types.YLeaf{"FrameThresholdLowOverridden", runningConfigInterface.FrameThresholdLowOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-threshold-high-overridden", types.YLeaf{"FrameThresholdHighOverridden", runningConfigInterface.FrameThresholdHighOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-window-overridden", types.YLeaf{"FramePeriodWindowOverridden", runningConfigInterface.FramePeriodWindowOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-threshold-low-overridden", types.YLeaf{"FramePeriodThresholdLowOverridden", runningConfigInterface.FramePeriodThresholdLowOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-period-threshold-high-overridden", types.YLeaf{"FramePeriodThresholdHighOverridden", runningConfigInterface.FramePeriodThresholdHighOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-seconds-window-overridden", types.YLeaf{"FrameSecondsWindowOverridden", runningConfigInterface.FrameSecondsWindowOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-seconds-threshold-low-overridden", types.YLeaf{"FrameSecondsThresholdLowOverridden", runningConfigInterface.FrameSecondsThresholdLowOverridden})
    runningConfigInterface.EntityData.Leafs.Append("frame-seconds-threshold-high-overridden", types.YLeaf{"FrameSecondsThresholdHighOverridden", runningConfigInterface.FrameSecondsThresholdHighOverridden})
    runningConfigInterface.EntityData.Leafs.Append("high-threshold-action-overridden", types.YLeaf{"HighThresholdActionOverridden", runningConfigInterface.HighThresholdActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("link-fault-action-overridden", types.YLeaf{"LinkFaultActionOverridden", runningConfigInterface.LinkFaultActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("dying-gasp-action-overridden", types.YLeaf{"DyingGaspActionOverridden", runningConfigInterface.DyingGaspActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("critical-event-action-overridden", types.YLeaf{"CriticalEventActionOverridden", runningConfigInterface.CriticalEventActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("discovery-timeout-action-overridden", types.YLeaf{"DiscoveryTimeoutActionOverridden", runningConfigInterface.DiscoveryTimeoutActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("capabilities-conflict-action-overridden", types.YLeaf{"CapabilitiesConflictActionOverridden", runningConfigInterface.CapabilitiesConflictActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("wiring-conflict-action-overridden", types.YLeaf{"WiringConflictActionOverridden", runningConfigInterface.WiringConflictActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("session-down-action-overridden", types.YLeaf{"SessionDownActionOverridden", runningConfigInterface.SessionDownActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("session-up-action-overridden", types.YLeaf{"SessionUpActionOverridden", runningConfigInterface.SessionUpActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("remote-loopback-action-overridden", types.YLeaf{"RemoteLoopbackActionOverridden", runningConfigInterface.RemoteLoopbackActionOverridden})
    runningConfigInterface.EntityData.Leafs.Append("require-mode-overridden", types.YLeaf{"RequireModeOverridden", runningConfigInterface.RequireModeOverridden})
    runningConfigInterface.EntityData.Leafs.Append("require-mib-retrieval-overridden", types.YLeaf{"RequireMibRetrievalOverridden", runningConfigInterface.RequireMibRetrievalOverridden})
    runningConfigInterface.EntityData.Leafs.Append("require-loopback-overridden", types.YLeaf{"RequireLoopbackOverridden", runningConfigInterface.RequireLoopbackOverridden})
    runningConfigInterface.EntityData.Leafs.Append("require-link-monitoring-overridden", types.YLeaf{"RequireLinkMonitoringOverridden", runningConfigInterface.RequireLinkMonitoringOverridden})

    runningConfigInterface.EntityData.YListKeys = []string {"MemberInterface"}

    return &(runningConfigInterface.EntityData)
}

// EtherLinkOam_Nodes
// Node table for node-specific operational data
type EtherLinkOam_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // EtherLinkOam_Nodes_Node.
    Node []*EtherLinkOam_Nodes_Node
}

func (nodes *EtherLinkOam_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ether-link-oam"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// EtherLinkOam_Nodes_Node
// Node-specific data for a particular node
type EtherLinkOam_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Ethernet Link OAM Summary information for the entire node.
    Summary EtherLinkOam_Nodes_Node_Summary
}

func (node *EtherLinkOam_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// EtherLinkOam_Nodes_Node_Summary
// Ethernet Link OAM Summary information for the
// entire node
type EtherLinkOam_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *EtherLinkOam_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("interfaces", types.YLeaf{"Interfaces", summary.Interfaces})
    summary.EntityData.Leafs.Append("port-down", types.YLeaf{"PortDown", summary.PortDown})
    summary.EntityData.Leafs.Append("passive-wait", types.YLeaf{"PassiveWait", summary.PassiveWait})
    summary.EntityData.Leafs.Append("active-send", types.YLeaf{"ActiveSend", summary.ActiveSend})
    summary.EntityData.Leafs.Append("evaluating", types.YLeaf{"Evaluating", summary.Evaluating})
    summary.EntityData.Leafs.Append("local-accept", types.YLeaf{"LocalAccept", summary.LocalAccept})
    summary.EntityData.Leafs.Append("local-reject", types.YLeaf{"LocalReject", summary.LocalReject})
    summary.EntityData.Leafs.Append("remote-reject", types.YLeaf{"RemoteReject", summary.RemoteReject})
    summary.EntityData.Leafs.Append("operational", types.YLeaf{"Operational", summary.Operational})
    summary.EntityData.Leafs.Append("loopback-mode", types.YLeaf{"LoopbackMode", summary.LoopbackMode})
    summary.EntityData.Leafs.Append("miswired-connections", types.YLeaf{"MiswiredConnections", summary.MiswiredConnections})
    summary.EntityData.Leafs.Append("events", types.YLeaf{"Events", summary.Events})
    summary.EntityData.Leafs.Append("local-events", types.YLeaf{"LocalEvents", summary.LocalEvents})
    summary.EntityData.Leafs.Append("local-symbol-period", types.YLeaf{"LocalSymbolPeriod", summary.LocalSymbolPeriod})
    summary.EntityData.Leafs.Append("local-frame", types.YLeaf{"LocalFrame", summary.LocalFrame})
    summary.EntityData.Leafs.Append("local-frame-period", types.YLeaf{"LocalFramePeriod", summary.LocalFramePeriod})
    summary.EntityData.Leafs.Append("local-frame-seconds", types.YLeaf{"LocalFrameSeconds", summary.LocalFrameSeconds})
    summary.EntityData.Leafs.Append("remote-events", types.YLeaf{"RemoteEvents", summary.RemoteEvents})
    summary.EntityData.Leafs.Append("remote-symbol-period", types.YLeaf{"RemoteSymbolPeriod", summary.RemoteSymbolPeriod})
    summary.EntityData.Leafs.Append("remote-frame", types.YLeaf{"RemoteFrame", summary.RemoteFrame})
    summary.EntityData.Leafs.Append("remote-frame-period", types.YLeaf{"RemoteFramePeriod", summary.RemoteFramePeriod})
    summary.EntityData.Leafs.Append("remote-frame-seconds", types.YLeaf{"RemoteFrameSeconds", summary.RemoteFrameSeconds})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// EtherLinkOam_EventLogEntryInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Event Log Entry container
type EtherLinkOam_EventLogEntryInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet Link OAM enabled interface to get Event Log Entry for. The type is
    // slice of EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface.
    EventLogEntryInterface []*EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface
}

func (eventLogEntryInterfaces *EtherLinkOam_EventLogEntryInterfaces) GetEntityData() *types.CommonEntityData {
    eventLogEntryInterfaces.EntityData.YFilter = eventLogEntryInterfaces.YFilter
    eventLogEntryInterfaces.EntityData.YangName = "event-log-entry-interfaces"
    eventLogEntryInterfaces.EntityData.BundleName = "cisco_ios_xr"
    eventLogEntryInterfaces.EntityData.ParentYangName = "ether-link-oam"
    eventLogEntryInterfaces.EntityData.SegmentPath = "event-log-entry-interfaces"
    eventLogEntryInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventLogEntryInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventLogEntryInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventLogEntryInterfaces.EntityData.Children = types.NewOrderedMap()
    eventLogEntryInterfaces.EntityData.Children.Append("event-log-entry-interface", types.YChild{"EventLogEntryInterface", nil})
    for i := range eventLogEntryInterfaces.EventLogEntryInterface {
        eventLogEntryInterfaces.EntityData.Children.Append(types.GetSegmentPath(eventLogEntryInterfaces.EventLogEntryInterface[i]), types.YChild{"EventLogEntryInterface", eventLogEntryInterfaces.EventLogEntryInterface[i]})
    }
    eventLogEntryInterfaces.EntityData.Leafs = types.NewOrderedMap()

    eventLogEntryInterfaces.EntityData.YListKeys = []string {}

    return &(eventLogEntryInterfaces.EntityData)
}

// EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface
// Ethernet Link OAM enabled interface to get
// Event Log Entry for
type EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    MemberInterface interface{}

    // Table of Ethernet Link OAM Event Log entries on the interface.
    EventLogEntryIndexes EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes
}

func (eventLogEntryInterface *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface) GetEntityData() *types.CommonEntityData {
    eventLogEntryInterface.EntityData.YFilter = eventLogEntryInterface.YFilter
    eventLogEntryInterface.EntityData.YangName = "event-log-entry-interface"
    eventLogEntryInterface.EntityData.BundleName = "cisco_ios_xr"
    eventLogEntryInterface.EntityData.ParentYangName = "event-log-entry-interfaces"
    eventLogEntryInterface.EntityData.SegmentPath = "event-log-entry-interface" + types.AddKeyToken(eventLogEntryInterface.MemberInterface, "member-interface")
    eventLogEntryInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventLogEntryInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventLogEntryInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventLogEntryInterface.EntityData.Children = types.NewOrderedMap()
    eventLogEntryInterface.EntityData.Children.Append("event-log-entry-indexes", types.YChild{"EventLogEntryIndexes", &eventLogEntryInterface.EventLogEntryIndexes})
    eventLogEntryInterface.EntityData.Leafs = types.NewOrderedMap()
    eventLogEntryInterface.EntityData.Leafs.Append("member-interface", types.YLeaf{"MemberInterface", eventLogEntryInterface.MemberInterface})

    eventLogEntryInterface.EntityData.YListKeys = []string {"MemberInterface"}

    return &(eventLogEntryInterface.EntityData)
}

// EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes
// Table of Ethernet Link OAM Event Log entries
// on the interface
type EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet Link OAM Event Log Entry Index to get data for. The type is slice
    // of
    // EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex.
    EventLogEntryIndex []*EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex
}

func (eventLogEntryIndexes *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes) GetEntityData() *types.CommonEntityData {
    eventLogEntryIndexes.EntityData.YFilter = eventLogEntryIndexes.YFilter
    eventLogEntryIndexes.EntityData.YangName = "event-log-entry-indexes"
    eventLogEntryIndexes.EntityData.BundleName = "cisco_ios_xr"
    eventLogEntryIndexes.EntityData.ParentYangName = "event-log-entry-interface"
    eventLogEntryIndexes.EntityData.SegmentPath = "event-log-entry-indexes"
    eventLogEntryIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventLogEntryIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventLogEntryIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventLogEntryIndexes.EntityData.Children = types.NewOrderedMap()
    eventLogEntryIndexes.EntityData.Children.Append("event-log-entry-index", types.YChild{"EventLogEntryIndex", nil})
    for i := range eventLogEntryIndexes.EventLogEntryIndex {
        eventLogEntryIndexes.EntityData.Children.Append(types.GetSegmentPath(eventLogEntryIndexes.EventLogEntryIndex[i]), types.YChild{"EventLogEntryIndex", eventLogEntryIndexes.EventLogEntryIndex[i]})
    }
    eventLogEntryIndexes.EntityData.Leafs = types.NewOrderedMap()

    eventLogEntryIndexes.EntityData.YListKeys = []string {}

    return &(eventLogEntryIndexes.EntityData)
}

// EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex
// Ethernet Link OAM Event Log Entry Index to
// get data for
type EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Event Log Entry index. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    EventLogEntryIndex interface{}

    // Index in the log entries table. The type is interface{} with range:
    // 0..4294967295.
    Index interface{}

    // Interface handle for this log entry. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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

func (eventLogEntryIndex *EtherLinkOam_EventLogEntryInterfaces_EventLogEntryInterface_EventLogEntryIndexes_EventLogEntryIndex) GetEntityData() *types.CommonEntityData {
    eventLogEntryIndex.EntityData.YFilter = eventLogEntryIndex.YFilter
    eventLogEntryIndex.EntityData.YangName = "event-log-entry-index"
    eventLogEntryIndex.EntityData.BundleName = "cisco_ios_xr"
    eventLogEntryIndex.EntityData.ParentYangName = "event-log-entry-indexes"
    eventLogEntryIndex.EntityData.SegmentPath = "event-log-entry-index" + types.AddKeyToken(eventLogEntryIndex.EventLogEntryIndex, "event-log-entry-index")
    eventLogEntryIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventLogEntryIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventLogEntryIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventLogEntryIndex.EntityData.Children = types.NewOrderedMap()
    eventLogEntryIndex.EntityData.Leafs = types.NewOrderedMap()
    eventLogEntryIndex.EntityData.Leafs.Append("event-log-entry-index", types.YLeaf{"EventLogEntryIndex", eventLogEntryIndex.EventLogEntryIndex})
    eventLogEntryIndex.EntityData.Leafs.Append("index", types.YLeaf{"Index", eventLogEntryIndex.Index})
    eventLogEntryIndex.EntityData.Leafs.Append("handle", types.YLeaf{"Handle", eventLogEntryIndex.Handle})
    eventLogEntryIndex.EntityData.Leafs.Append("oui", types.YLeaf{"Oui", eventLogEntryIndex.Oui})
    eventLogEntryIndex.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", eventLogEntryIndex.Timestamp})
    eventLogEntryIndex.EntityData.Leafs.Append("type", types.YLeaf{"Type", eventLogEntryIndex.Type})
    eventLogEntryIndex.EntityData.Leafs.Append("location", types.YLeaf{"Location", eventLogEntryIndex.Location})
    eventLogEntryIndex.EntityData.Leafs.Append("event-total", types.YLeaf{"EventTotal", eventLogEntryIndex.EventTotal})
    eventLogEntryIndex.EntityData.Leafs.Append("action-taken", types.YLeaf{"ActionTaken", eventLogEntryIndex.ActionTaken})
    eventLogEntryIndex.EntityData.Leafs.Append("window", types.YLeaf{"Window", eventLogEntryIndex.Window})
    eventLogEntryIndex.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", eventLogEntryIndex.Threshold})
    eventLogEntryIndex.EntityData.Leafs.Append("local-high-threshold", types.YLeaf{"LocalHighThreshold", eventLogEntryIndex.LocalHighThreshold})
    eventLogEntryIndex.EntityData.Leafs.Append("value", types.YLeaf{"Value", eventLogEntryIndex.Value})
    eventLogEntryIndex.EntityData.Leafs.Append("running-total", types.YLeaf{"RunningTotal", eventLogEntryIndex.RunningTotal})
    eventLogEntryIndex.EntityData.Leafs.Append("window-config-units", types.YLeaf{"WindowConfigUnits", eventLogEntryIndex.WindowConfigUnits})
    eventLogEntryIndex.EntityData.Leafs.Append("window-units", types.YLeaf{"WindowUnits", eventLogEntryIndex.WindowUnits})
    eventLogEntryIndex.EntityData.Leafs.Append("threshold-config-units", types.YLeaf{"ThresholdConfigUnits", eventLogEntryIndex.ThresholdConfigUnits})
    eventLogEntryIndex.EntityData.Leafs.Append("threshold-units", types.YLeaf{"ThresholdUnits", eventLogEntryIndex.ThresholdUnits})
    eventLogEntryIndex.EntityData.Leafs.Append("local-high-threshold-config-units", types.YLeaf{"LocalHighThresholdConfigUnits", eventLogEntryIndex.LocalHighThresholdConfigUnits})
    eventLogEntryIndex.EntityData.Leafs.Append("value-config-units", types.YLeaf{"ValueConfigUnits", eventLogEntryIndex.ValueConfigUnits})

    eventLogEntryIndex.EntityData.YListKeys = []string {"EventLogEntryIndex"}

    return &(eventLogEntryIndex.EntityData)
}

// EtherLinkOam_StatsInterfaces
// Table of Ethernet Link OAM enabled interfaces
// within Stats container
type EtherLinkOam_StatsInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet Link OAM interface to get Stats for. The type is slice of
    // EtherLinkOam_StatsInterfaces_StatsInterface.
    StatsInterface []*EtherLinkOam_StatsInterfaces_StatsInterface
}

func (statsInterfaces *EtherLinkOam_StatsInterfaces) GetEntityData() *types.CommonEntityData {
    statsInterfaces.EntityData.YFilter = statsInterfaces.YFilter
    statsInterfaces.EntityData.YangName = "stats-interfaces"
    statsInterfaces.EntityData.BundleName = "cisco_ios_xr"
    statsInterfaces.EntityData.ParentYangName = "ether-link-oam"
    statsInterfaces.EntityData.SegmentPath = "stats-interfaces"
    statsInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsInterfaces.EntityData.Children = types.NewOrderedMap()
    statsInterfaces.EntityData.Children.Append("stats-interface", types.YChild{"StatsInterface", nil})
    for i := range statsInterfaces.StatsInterface {
        statsInterfaces.EntityData.Children.Append(types.GetSegmentPath(statsInterfaces.StatsInterface[i]), types.YChild{"StatsInterface", statsInterfaces.StatsInterface[i]})
    }
    statsInterfaces.EntityData.Leafs = types.NewOrderedMap()

    statsInterfaces.EntityData.YListKeys = []string {}

    return &(statsInterfaces.EntityData)
}

// EtherLinkOam_StatsInterfaces_StatsInterface
// Ethernet Link OAM interface to get Stats for
type EtherLinkOam_StatsInterfaces_StatsInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Member Interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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

func (statsInterface *EtherLinkOam_StatsInterfaces_StatsInterface) GetEntityData() *types.CommonEntityData {
    statsInterface.EntityData.YFilter = statsInterface.YFilter
    statsInterface.EntityData.YangName = "stats-interface"
    statsInterface.EntityData.BundleName = "cisco_ios_xr"
    statsInterface.EntityData.ParentYangName = "stats-interfaces"
    statsInterface.EntityData.SegmentPath = "stats-interface" + types.AddKeyToken(statsInterface.MemberInterface, "member-interface")
    statsInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsInterface.EntityData.Children = types.NewOrderedMap()
    statsInterface.EntityData.Leafs = types.NewOrderedMap()
    statsInterface.EntityData.Leafs.Append("member-interface", types.YLeaf{"MemberInterface", statsInterface.MemberInterface})
    statsInterface.EntityData.Leafs.Append("information-tx", types.YLeaf{"InformationTx", statsInterface.InformationTx})
    statsInterface.EntityData.Leafs.Append("information-rx", types.YLeaf{"InformationRx", statsInterface.InformationRx})
    statsInterface.EntityData.Leafs.Append("unique-event-notification-tx", types.YLeaf{"UniqueEventNotificationTx", statsInterface.UniqueEventNotificationTx})
    statsInterface.EntityData.Leafs.Append("unique-event-notification-rx", types.YLeaf{"UniqueEventNotificationRx", statsInterface.UniqueEventNotificationRx})
    statsInterface.EntityData.Leafs.Append("duplicate-event-notification-tx", types.YLeaf{"DuplicateEventNotificationTx", statsInterface.DuplicateEventNotificationTx})
    statsInterface.EntityData.Leafs.Append("duplicate-event-notification-rx", types.YLeaf{"DuplicateEventNotificationRx", statsInterface.DuplicateEventNotificationRx})
    statsInterface.EntityData.Leafs.Append("loopback-control-tx", types.YLeaf{"LoopbackControlTx", statsInterface.LoopbackControlTx})
    statsInterface.EntityData.Leafs.Append("loopback-control-rx", types.YLeaf{"LoopbackControlRx", statsInterface.LoopbackControlRx})
    statsInterface.EntityData.Leafs.Append("variable-request-tx", types.YLeaf{"VariableRequestTx", statsInterface.VariableRequestTx})
    statsInterface.EntityData.Leafs.Append("variable-request-rx", types.YLeaf{"VariableRequestRx", statsInterface.VariableRequestRx})
    statsInterface.EntityData.Leafs.Append("variable-response-tx", types.YLeaf{"VariableResponseTx", statsInterface.VariableResponseTx})
    statsInterface.EntityData.Leafs.Append("variable-response-rx", types.YLeaf{"VariableResponseRx", statsInterface.VariableResponseRx})
    statsInterface.EntityData.Leafs.Append("org-specific-tx", types.YLeaf{"OrgSpecificTx", statsInterface.OrgSpecificTx})
    statsInterface.EntityData.Leafs.Append("org-specific-rx", types.YLeaf{"OrgSpecificRx", statsInterface.OrgSpecificRx})
    statsInterface.EntityData.Leafs.Append("unsupported-codes-tx", types.YLeaf{"UnsupportedCodesTx", statsInterface.UnsupportedCodesTx})
    statsInterface.EntityData.Leafs.Append("unsupported-codes-rx", types.YLeaf{"UnsupportedCodesRx", statsInterface.UnsupportedCodesRx})
    statsInterface.EntityData.Leafs.Append("frames-lost-due-to-oam", types.YLeaf{"FramesLostDueToOam", statsInterface.FramesLostDueToOam})
    statsInterface.EntityData.Leafs.Append("fixed-frames-rx", types.YLeaf{"FixedFramesRx", statsInterface.FixedFramesRx})
    statsInterface.EntityData.Leafs.Append("local-error-symbol-period-records", types.YLeaf{"LocalErrorSymbolPeriodRecords", statsInterface.LocalErrorSymbolPeriodRecords})
    statsInterface.EntityData.Leafs.Append("local-error-frame-records", types.YLeaf{"LocalErrorFrameRecords", statsInterface.LocalErrorFrameRecords})
    statsInterface.EntityData.Leafs.Append("local-error-frame-period-records", types.YLeaf{"LocalErrorFramePeriodRecords", statsInterface.LocalErrorFramePeriodRecords})
    statsInterface.EntityData.Leafs.Append("local-error-frame-second-records", types.YLeaf{"LocalErrorFrameSecondRecords", statsInterface.LocalErrorFrameSecondRecords})
    statsInterface.EntityData.Leafs.Append("remote-error-symbol-period-records", types.YLeaf{"RemoteErrorSymbolPeriodRecords", statsInterface.RemoteErrorSymbolPeriodRecords})
    statsInterface.EntityData.Leafs.Append("remote-error-frame-records", types.YLeaf{"RemoteErrorFrameRecords", statsInterface.RemoteErrorFrameRecords})
    statsInterface.EntityData.Leafs.Append("remote-error-frame-period-records", types.YLeaf{"RemoteErrorFramePeriodRecords", statsInterface.RemoteErrorFramePeriodRecords})
    statsInterface.EntityData.Leafs.Append("remote-error-frame-second-records", types.YLeaf{"RemoteErrorFrameSecondRecords", statsInterface.RemoteErrorFrameSecondRecords})

    statsInterface.EntityData.YListKeys = []string {"MemberInterface"}

    return &(statsInterface.EntityData)
}

