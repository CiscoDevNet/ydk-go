// This module contains a collection of YANG definitions
// for Cisco IOS-XR terminal-device package configuration.
// 
// This module contains definitions
// for the following management objects:
//   logical-channels: Logical channel in mxp
//   optical-channels: optical channels
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package terminal_device_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package terminal_device_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-terminal-device-cfg logical-channels}", reflect.TypeOf(LogicalChannels{}))
    ydk.RegisterEntity("Cisco-IOS-XR-terminal-device-cfg:logical-channels", reflect.TypeOf(LogicalChannels{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-terminal-device-cfg optical-channels}", reflect.TypeOf(OpticalChannels{}))
    ydk.RegisterEntity("Cisco-IOS-XR-terminal-device-cfg:optical-channels", reflect.TypeOf(OpticalChannels{}))
}

// LogicalTribRate represents Logical trib rate
type LogicalTribRate string

const (
    // TribRate1G
    LogicalTribRate_trib_rate1g LogicalTribRate = "trib-rate1g"

    // TribRate25G
    LogicalTribRate_trib_rate2_5g LogicalTribRate = "trib-rate2-5g"

    // TribRate10G
    LogicalTribRate_trib_rate10g LogicalTribRate = "trib-rate10g"

    // TribRate40G
    LogicalTribRate_trib_rate40g LogicalTribRate = "trib-rate40g"

    // TribRate100G
    LogicalTribRate_trib_rate100g LogicalTribRate = "trib-rate100g"
)

// LogicalLoopbackMode represents Logical loopback mode
type LogicalLoopbackMode string

const (
    // None
    LogicalLoopbackMode_none LogicalLoopbackMode = "none"

    // Facility
    LogicalLoopbackMode_facility LogicalLoopbackMode = "facility"

    // Terminal
    LogicalLoopbackMode_terminal LogicalLoopbackMode = "terminal"
)

// LogicalChannelOtnTtiAuto represents Logical channel otn tti auto
type LogicalChannelOtnTtiAuto string

const (
    // Otn tti auto mode false
    LogicalChannelOtnTtiAuto_false LogicalChannelOtnTtiAuto = "false"

    // Otn tti auto mode true
    LogicalChannelOtnTtiAuto_true LogicalChannelOtnTtiAuto = "true"
)

// LogicalAdminState represents Logical admin state
type LogicalAdminState string

const (
    // Enable
    LogicalAdminState_enable LogicalAdminState = "enable"

    // Disable
    LogicalAdminState_disable LogicalAdminState = "disable"

    // Maintenance
    LogicalAdminState_maintenance LogicalAdminState = "maintenance"
)

// LogicalChannelAssignment represents Logical channel assignment
type LogicalChannelAssignment string

const (
    // Type Logical channel
    LogicalChannelAssignment_type_logical_channel LogicalChannelAssignment = "type-logical-channel"

    // Type Optical channel
    LogicalChannelAssignment_type_optical_channel LogicalChannelAssignment = "type-optical-channel"
)

// LogicalTribProtocol represents Logical trib protocol
type LogicalTribProtocol string

const (
    // 1G Ethernet protocol
    LogicalTribProtocol_trib_proto_type1ge LogicalTribProtocol = "trib-proto-type1ge"

    // OC48 protocol
    LogicalTribProtocol_trib_proto_type_oc48 LogicalTribProtocol = "trib-proto-type-oc48"

    // STM 16 protocol
    LogicalTribProtocol_trib_proto_type_stm16 LogicalTribProtocol = "trib-proto-type-stm16"

    // 10G Ethernet LAN protocol
    LogicalTribProtocol_trib_proto_type10gelan LogicalTribProtocol = "trib-proto-type10gelan"

    // 10G Ethernet WAN protocol
    LogicalTribProtocol_trib_proto_type10gewan LogicalTribProtocol = "trib-proto-type10gewan"

    // OC 192 (9.6GB) port protocol
    LogicalTribProtocol_trib_proto_type_oc192 LogicalTribProtocol = "trib-proto-type-oc192"

    // STM 64 protocol
    LogicalTribProtocol_trib_proto_type_stm64 LogicalTribProtocol = "trib-proto-type-stm64"

    // OTU 2 protocol
    LogicalTribProtocol_trib_proto_type_otu2 LogicalTribProtocol = "trib-proto-type-otu2"

    // OTU 2e protocol
    LogicalTribProtocol_trib_proto_type_otu2e LogicalTribProtocol = "trib-proto-type-otu2e"

    // OTU 1e protocol
    LogicalTribProtocol_trib_proto_type_otu1e LogicalTribProtocol = "trib-proto-type-otu1e"

    // ODU 2 protocol
    LogicalTribProtocol_trib_proto_type_odu2 LogicalTribProtocol = "trib-proto-type-odu2"

    // ODU 2e protocol
    LogicalTribProtocol_trib_proto_type_odu2e LogicalTribProtocol = "trib-proto-type-odu2e"

    // 40G Ethernet port protocol
    LogicalTribProtocol_trib_proto_type40ge LogicalTribProtocol = "trib-proto-type40ge"

    // OC 768 protocol
    LogicalTribProtocol_trib_proto_type_oc768 LogicalTribProtocol = "trib-proto-type-oc768"

    // STM 256 protocol
    LogicalTribProtocol_trib_proto_type_stm256 LogicalTribProtocol = "trib-proto-type-stm256"

    // OTU 3 protocol
    LogicalTribProtocol_trib_proto_type_otu3 LogicalTribProtocol = "trib-proto-type-otu3"

    // ODU 3 protocol
    LogicalTribProtocol_trib_proto_type_odu3 LogicalTribProtocol = "trib-proto-type-odu3"

    // 100G Ethernet protocol
    LogicalTribProtocol_trib_proto_type100ge LogicalTribProtocol = "trib-proto-type100ge"

    // 100G MLG protocol
    LogicalTribProtocol_trib_proto_type100g_mlg LogicalTribProtocol = "trib-proto-type100g-mlg"

    // OTU4 signal protocol (112G) for transporting
    // 100GE signal
    LogicalTribProtocol_trib_proto_type_otu4 LogicalTribProtocol = "trib-proto-type-otu4"

    // OTU Cn protocol
    LogicalTribProtocol_trib_proto_type_otu_cn LogicalTribProtocol = "trib-proto-type-otu-cn"

    // ODU 4 protocol
    LogicalTribProtocol_trib_proto_type_odu4 LogicalTribProtocol = "trib-proto-type-odu4"
)

// LogicalProtocol represents Logical protocol
type LogicalProtocol string

const (
    // Type Ethernet
    LogicalProtocol_type_ethernet LogicalProtocol = "type-ethernet"

    // Type OTN
    LogicalProtocol_type_otn LogicalProtocol = "type-otn"
)

// LogicalChannels
// Logical channel in mxp
type LogicalChannels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logical channel index. The type is slice of LogicalChannels_Channel.
    Channel []LogicalChannels_Channel
}

func (logicalChannels *LogicalChannels) GetFilter() yfilter.YFilter { return logicalChannels.YFilter }

func (logicalChannels *LogicalChannels) SetFilter(yf yfilter.YFilter) { logicalChannels.YFilter = yf }

func (logicalChannels *LogicalChannels) GetGoName(yname string) string {
    if yname == "channel" { return "Channel" }
    return ""
}

func (logicalChannels *LogicalChannels) GetSegmentPath() string {
    return "Cisco-IOS-XR-terminal-device-cfg:logical-channels"
}

func (logicalChannels *LogicalChannels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "channel" {
        for _, c := range logicalChannels.Channel {
            if logicalChannels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LogicalChannels_Channel{}
        logicalChannels.Channel = append(logicalChannels.Channel, child)
        return &logicalChannels.Channel[len(logicalChannels.Channel)-1]
    }
    return nil
}

func (logicalChannels *LogicalChannels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range logicalChannels.Channel {
        children[logicalChannels.Channel[i].GetSegmentPath()] = &logicalChannels.Channel[i]
    }
    return children
}

func (logicalChannels *LogicalChannels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logicalChannels *LogicalChannels) GetBundleName() string { return "cisco_ios_xr" }

func (logicalChannels *LogicalChannels) GetYangName() string { return "logical-channels" }

func (logicalChannels *LogicalChannels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logicalChannels *LogicalChannels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logicalChannels *LogicalChannels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logicalChannels *LogicalChannels) SetParent(parent types.Entity) { logicalChannels.parent = parent }

func (logicalChannels *LogicalChannels) GetParent() types.Entity { return logicalChannels.parent }

func (logicalChannels *LogicalChannels) GetParentYangName() string { return "Cisco-IOS-XR-terminal-device-cfg" }

// LogicalChannels_Channel
// Logical channel index
type LogicalChannels_Channel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Logical Channel Index. The type is interface{}
    // with range: -2147483648..2147483647.
    ChannelIndex interface{}

    // Protocol framing of the tributary signal. The type is LogicalTribProtocol.
    TribProtocol interface{}

    // Description (Max 255 characters). The type is string with length: 1..255.
    Description interface{}

    // Configure ingress client port for this logical channel. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    IngressClientPort interface{}

    // Configure ingress physical channel for this logical channel. The type is
    // interface{} with range: 1..4.
    IngressPhysicalChannel interface{}

    // Configure the admin-state . The type is LogicalAdminState.
    AdminState interface{}

    // Configure the loopback mode . The type is LogicalLoopbackMode.
    LoopbackMode interface{}

    // Configure the logical-channel-type . The type is LogicalProtocol.
    LogicalChannelType interface{}

    // Rounded bit rate of the tributary signal. The type is LogicalTribRate.
    RateClass interface{}

    // Logical channel assignment for logical channel.
    LogicalChannelAssignments LogicalChannels_Channel_LogicalChannelAssignments

    // Otn Related configs for Logical channel.
    Otn LogicalChannels_Channel_Otn
}

func (channel *LogicalChannels_Channel) GetFilter() yfilter.YFilter { return channel.YFilter }

func (channel *LogicalChannels_Channel) SetFilter(yf yfilter.YFilter) { channel.YFilter = yf }

func (channel *LogicalChannels_Channel) GetGoName(yname string) string {
    if yname == "channel-index" { return "ChannelIndex" }
    if yname == "trib-protocol" { return "TribProtocol" }
    if yname == "description" { return "Description" }
    if yname == "ingress-client-port" { return "IngressClientPort" }
    if yname == "ingress-physical-channel" { return "IngressPhysicalChannel" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "loopback-mode" { return "LoopbackMode" }
    if yname == "logical-channel-type" { return "LogicalChannelType" }
    if yname == "rate-class" { return "RateClass" }
    if yname == "logical-channel-assignments" { return "LogicalChannelAssignments" }
    if yname == "otn" { return "Otn" }
    return ""
}

func (channel *LogicalChannels_Channel) GetSegmentPath() string {
    return "channel" + "[channel-index='" + fmt.Sprintf("%v", channel.ChannelIndex) + "']"
}

func (channel *LogicalChannels_Channel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logical-channel-assignments" {
        return &channel.LogicalChannelAssignments
    }
    if childYangName == "otn" {
        return &channel.Otn
    }
    return nil
}

func (channel *LogicalChannels_Channel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["logical-channel-assignments"] = &channel.LogicalChannelAssignments
    children["otn"] = &channel.Otn
    return children
}

func (channel *LogicalChannels_Channel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["channel-index"] = channel.ChannelIndex
    leafs["trib-protocol"] = channel.TribProtocol
    leafs["description"] = channel.Description
    leafs["ingress-client-port"] = channel.IngressClientPort
    leafs["ingress-physical-channel"] = channel.IngressPhysicalChannel
    leafs["admin-state"] = channel.AdminState
    leafs["loopback-mode"] = channel.LoopbackMode
    leafs["logical-channel-type"] = channel.LogicalChannelType
    leafs["rate-class"] = channel.RateClass
    return leafs
}

func (channel *LogicalChannels_Channel) GetBundleName() string { return "cisco_ios_xr" }

func (channel *LogicalChannels_Channel) GetYangName() string { return "channel" }

func (channel *LogicalChannels_Channel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (channel *LogicalChannels_Channel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (channel *LogicalChannels_Channel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (channel *LogicalChannels_Channel) SetParent(parent types.Entity) { channel.parent = parent }

func (channel *LogicalChannels_Channel) GetParent() types.Entity { return channel.parent }

func (channel *LogicalChannels_Channel) GetParentYangName() string { return "logical-channels" }

// LogicalChannels_Channel_LogicalChannelAssignments
// Logical channel assignment for logical channel
type LogicalChannels_Channel_LogicalChannelAssignments struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logical Channel Assignment id. The type is slice of
    // LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment.
    LogicalChannelAssignment []LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment
}

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetFilter() yfilter.YFilter { return logicalChannelAssignments.YFilter }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) SetFilter(yf yfilter.YFilter) { logicalChannelAssignments.YFilter = yf }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetGoName(yname string) string {
    if yname == "logical-channel-assignment" { return "LogicalChannelAssignment" }
    return ""
}

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetSegmentPath() string {
    return "logical-channel-assignments"
}

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logical-channel-assignment" {
        for _, c := range logicalChannelAssignments.LogicalChannelAssignment {
            if logicalChannelAssignments.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment{}
        logicalChannelAssignments.LogicalChannelAssignment = append(logicalChannelAssignments.LogicalChannelAssignment, child)
        return &logicalChannelAssignments.LogicalChannelAssignment[len(logicalChannelAssignments.LogicalChannelAssignment)-1]
    }
    return nil
}

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range logicalChannelAssignments.LogicalChannelAssignment {
        children[logicalChannelAssignments.LogicalChannelAssignment[i].GetSegmentPath()] = &logicalChannelAssignments.LogicalChannelAssignment[i]
    }
    return children
}

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetBundleName() string { return "cisco_ios_xr" }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetYangName() string { return "logical-channel-assignments" }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) SetParent(parent types.Entity) { logicalChannelAssignments.parent = parent }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetParent() types.Entity { return logicalChannelAssignments.parent }

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetParentYangName() string { return "channel" }

// LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment
// Logical Channel Assignment id
type LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Logical channel assignment index. The type is
    // interface{} with range: -2147483648..2147483647.
    AssignmentIndex interface{}

    // Configure description for this assignment. The type is string with length:
    // 1..255.
    Description interface{}

    // Configure logical channel for this assignment. The type is interface{} with
    // range: -2147483648..2147483647.
    LogicalChannelId interface{}

    // Type of assignment for logical channel. The type is
    // LogicalChannelAssignment.
    AssignmentType interface{}

    // Configure Allocation for this assignment(10, 40 or 100G). The type is
    // interface{} with range: -2147483648..2147483647.
    Allocation interface{}

    // Configure optical channel for this assignment. The type is string.
    OpticalChannelId interface{}
}

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetFilter() yfilter.YFilter { return logicalChannelAssignment.YFilter }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) SetFilter(yf yfilter.YFilter) { logicalChannelAssignment.YFilter = yf }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetGoName(yname string) string {
    if yname == "assignment-index" { return "AssignmentIndex" }
    if yname == "description" { return "Description" }
    if yname == "logical-channel-id" { return "LogicalChannelId" }
    if yname == "assignment-type" { return "AssignmentType" }
    if yname == "allocation" { return "Allocation" }
    if yname == "optical-channel-id" { return "OpticalChannelId" }
    return ""
}

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetSegmentPath() string {
    return "logical-channel-assignment" + "[assignment-index='" + fmt.Sprintf("%v", logicalChannelAssignment.AssignmentIndex) + "']"
}

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["assignment-index"] = logicalChannelAssignment.AssignmentIndex
    leafs["description"] = logicalChannelAssignment.Description
    leafs["logical-channel-id"] = logicalChannelAssignment.LogicalChannelId
    leafs["assignment-type"] = logicalChannelAssignment.AssignmentType
    leafs["allocation"] = logicalChannelAssignment.Allocation
    leafs["optical-channel-id"] = logicalChannelAssignment.OpticalChannelId
    return leafs
}

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetBundleName() string { return "cisco_ios_xr" }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetYangName() string { return "logical-channel-assignment" }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) SetParent(parent types.Entity) { logicalChannelAssignment.parent = parent }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetParent() types.Entity { return logicalChannelAssignment.parent }

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetParentYangName() string { return "logical-channel-assignments" }

// LogicalChannels_Channel_Otn
// Otn Related configs for Logical channel
type LogicalChannels_Channel_Otn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trail trace identifier (TTI) transmit message automatically created. If
    // True, then setting a custom transmit message would be invalid. Trail trace
    // identifier (TTI) transmit message automatically created. The type is
    // LogicalChannelOtnTtiAuto.
    TtiMsgAuto interface{}

    // Trail trace identifier (TTI) message expectedTrail trace identifier (TTI)
    // message expected. The type is string with length: 1..255.
    TtiMsgExpected interface{}

    // Trail trace identifier (TTI) message transmittedTrail trace identifier
    // (TTI) message transmitted. The type is string with length: 1..255.
    TtiMsgTransmit interface{}
}

func (otn *LogicalChannels_Channel_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *LogicalChannels_Channel_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *LogicalChannels_Channel_Otn) GetGoName(yname string) string {
    if yname == "tti-msg-auto" { return "TtiMsgAuto" }
    if yname == "tti-msg-expected" { return "TtiMsgExpected" }
    if yname == "tti-msg-transmit" { return "TtiMsgTransmit" }
    return ""
}

func (otn *LogicalChannels_Channel_Otn) GetSegmentPath() string {
    return "otn"
}

func (otn *LogicalChannels_Channel_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otn *LogicalChannels_Channel_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otn *LogicalChannels_Channel_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tti-msg-auto"] = otn.TtiMsgAuto
    leafs["tti-msg-expected"] = otn.TtiMsgExpected
    leafs["tti-msg-transmit"] = otn.TtiMsgTransmit
    return leafs
}

func (otn *LogicalChannels_Channel_Otn) GetBundleName() string { return "cisco_ios_xr" }

func (otn *LogicalChannels_Channel_Otn) GetYangName() string { return "otn" }

func (otn *LogicalChannels_Channel_Otn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otn *LogicalChannels_Channel_Otn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otn *LogicalChannels_Channel_Otn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otn *LogicalChannels_Channel_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *LogicalChannels_Channel_Otn) GetParent() types.Entity { return otn.parent }

func (otn *LogicalChannels_Channel_Otn) GetParentYangName() string { return "channel" }

// OpticalChannels
// optical channels
type OpticalChannels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Optical Channel index. The type is slice of OpticalChannels_OpticalChannel.
    OpticalChannel []OpticalChannels_OpticalChannel
}

func (opticalChannels *OpticalChannels) GetFilter() yfilter.YFilter { return opticalChannels.YFilter }

func (opticalChannels *OpticalChannels) SetFilter(yf yfilter.YFilter) { opticalChannels.YFilter = yf }

func (opticalChannels *OpticalChannels) GetGoName(yname string) string {
    if yname == "optical-channel" { return "OpticalChannel" }
    return ""
}

func (opticalChannels *OpticalChannels) GetSegmentPath() string {
    return "Cisco-IOS-XR-terminal-device-cfg:optical-channels"
}

func (opticalChannels *OpticalChannels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optical-channel" {
        for _, c := range opticalChannels.OpticalChannel {
            if opticalChannels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticalChannels_OpticalChannel{}
        opticalChannels.OpticalChannel = append(opticalChannels.OpticalChannel, child)
        return &opticalChannels.OpticalChannel[len(opticalChannels.OpticalChannel)-1]
    }
    return nil
}

func (opticalChannels *OpticalChannels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opticalChannels.OpticalChannel {
        children[opticalChannels.OpticalChannel[i].GetSegmentPath()] = &opticalChannels.OpticalChannel[i]
    }
    return children
}

func (opticalChannels *OpticalChannels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticalChannels *OpticalChannels) GetBundleName() string { return "cisco_ios_xr" }

func (opticalChannels *OpticalChannels) GetYangName() string { return "optical-channels" }

func (opticalChannels *OpticalChannels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalChannels *OpticalChannels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalChannels *OpticalChannels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalChannels *OpticalChannels) SetParent(parent types.Entity) { opticalChannels.parent = parent }

func (opticalChannels *OpticalChannels) GetParent() types.Entity { return opticalChannels.parent }

func (opticalChannels *OpticalChannels) GetParentYangName() string { return "Cisco-IOS-XR-terminal-device-cfg" }

// OpticalChannels_OpticalChannel
// Optical Channel index
type OpticalChannels_OpticalChannel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Optical Channel Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ifname interface{}

    // Configure operational mode. The type is interface{} with range: 1..100000.
    OperationalMode interface{}

    // Specify R/S/I/P. The type is string with pattern: [a-zA-Z0-9./-]+.
    LinePort interface{}
}

func (opticalChannel *OpticalChannels_OpticalChannel) GetFilter() yfilter.YFilter { return opticalChannel.YFilter }

func (opticalChannel *OpticalChannels_OpticalChannel) SetFilter(yf yfilter.YFilter) { opticalChannel.YFilter = yf }

func (opticalChannel *OpticalChannels_OpticalChannel) GetGoName(yname string) string {
    if yname == "ifname" { return "Ifname" }
    if yname == "operational-mode" { return "OperationalMode" }
    if yname == "line-port" { return "LinePort" }
    return ""
}

func (opticalChannel *OpticalChannels_OpticalChannel) GetSegmentPath() string {
    return "optical-channel" + "[ifname='" + fmt.Sprintf("%v", opticalChannel.Ifname) + "']"
}

func (opticalChannel *OpticalChannels_OpticalChannel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalChannel *OpticalChannels_OpticalChannel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalChannel *OpticalChannels_OpticalChannel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifname"] = opticalChannel.Ifname
    leafs["operational-mode"] = opticalChannel.OperationalMode
    leafs["line-port"] = opticalChannel.LinePort
    return leafs
}

func (opticalChannel *OpticalChannels_OpticalChannel) GetBundleName() string { return "cisco_ios_xr" }

func (opticalChannel *OpticalChannels_OpticalChannel) GetYangName() string { return "optical-channel" }

func (opticalChannel *OpticalChannels_OpticalChannel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalChannel *OpticalChannels_OpticalChannel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalChannel *OpticalChannels_OpticalChannel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalChannel *OpticalChannels_OpticalChannel) SetParent(parent types.Entity) { opticalChannel.parent = parent }

func (opticalChannel *OpticalChannels_OpticalChannel) GetParent() types.Entity { return opticalChannel.parent }

func (opticalChannel *OpticalChannels_OpticalChannel) GetParentYangName() string { return "optical-channels" }

