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

// LogicalChannelAssignment represents Logical channel assignment
type LogicalChannelAssignment string

const (
    // Type Logical channel
    LogicalChannelAssignment_type_logical_channel LogicalChannelAssignment = "type-logical-channel"

    // Type Optical channel
    LogicalChannelAssignment_type_optical_channel LogicalChannelAssignment = "type-optical-channel"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logical channel index. The type is slice of LogicalChannels_Channel.
    Channel []LogicalChannels_Channel
}

func (logicalChannels *LogicalChannels) GetEntityData() *types.CommonEntityData {
    logicalChannels.EntityData.YFilter = logicalChannels.YFilter
    logicalChannels.EntityData.YangName = "logical-channels"
    logicalChannels.EntityData.BundleName = "cisco_ios_xr"
    logicalChannels.EntityData.ParentYangName = "Cisco-IOS-XR-terminal-device-cfg"
    logicalChannels.EntityData.SegmentPath = "Cisco-IOS-XR-terminal-device-cfg:logical-channels"
    logicalChannels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logicalChannels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logicalChannels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logicalChannels.EntityData.Children = make(map[string]types.YChild)
    logicalChannels.EntityData.Children["channel"] = types.YChild{"Channel", nil}
    for i := range logicalChannels.Channel {
        logicalChannels.EntityData.Children[types.GetSegmentPath(&logicalChannels.Channel[i])] = types.YChild{"Channel", &logicalChannels.Channel[i]}
    }
    logicalChannels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(logicalChannels.EntityData)
}

// LogicalChannels_Channel
// Logical channel index
type LogicalChannels_Channel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Logical Channel Index. The type is interface{}
    // with range: -2147483648..2147483647.
    ChannelIndex interface{}

    // Protocol framing of the tributary signal. The type is LogicalTribProtocol.
    TribProtocol interface{}

    // Description (Max 255 characters). The type is string with length: 1..255.
    Description interface{}

    // Configure ingress client port for this logical channel. The type is string
    // with pattern: b'[a-zA-Z0-9./-]+'.
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

func (channel *LogicalChannels_Channel) GetEntityData() *types.CommonEntityData {
    channel.EntityData.YFilter = channel.YFilter
    channel.EntityData.YangName = "channel"
    channel.EntityData.BundleName = "cisco_ios_xr"
    channel.EntityData.ParentYangName = "logical-channels"
    channel.EntityData.SegmentPath = "channel" + "[channel-index='" + fmt.Sprintf("%v", channel.ChannelIndex) + "']"
    channel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channel.EntityData.Children = make(map[string]types.YChild)
    channel.EntityData.Children["logical-channel-assignments"] = types.YChild{"LogicalChannelAssignments", &channel.LogicalChannelAssignments}
    channel.EntityData.Children["otn"] = types.YChild{"Otn", &channel.Otn}
    channel.EntityData.Leafs = make(map[string]types.YLeaf)
    channel.EntityData.Leafs["channel-index"] = types.YLeaf{"ChannelIndex", channel.ChannelIndex}
    channel.EntityData.Leafs["trib-protocol"] = types.YLeaf{"TribProtocol", channel.TribProtocol}
    channel.EntityData.Leafs["description"] = types.YLeaf{"Description", channel.Description}
    channel.EntityData.Leafs["ingress-client-port"] = types.YLeaf{"IngressClientPort", channel.IngressClientPort}
    channel.EntityData.Leafs["ingress-physical-channel"] = types.YLeaf{"IngressPhysicalChannel", channel.IngressPhysicalChannel}
    channel.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", channel.AdminState}
    channel.EntityData.Leafs["loopback-mode"] = types.YLeaf{"LoopbackMode", channel.LoopbackMode}
    channel.EntityData.Leafs["logical-channel-type"] = types.YLeaf{"LogicalChannelType", channel.LogicalChannelType}
    channel.EntityData.Leafs["rate-class"] = types.YLeaf{"RateClass", channel.RateClass}
    return &(channel.EntityData)
}

// LogicalChannels_Channel_LogicalChannelAssignments
// Logical channel assignment for logical channel
type LogicalChannels_Channel_LogicalChannelAssignments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logical Channel Assignment id. The type is slice of
    // LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment.
    LogicalChannelAssignment []LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment
}

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetEntityData() *types.CommonEntityData {
    logicalChannelAssignments.EntityData.YFilter = logicalChannelAssignments.YFilter
    logicalChannelAssignments.EntityData.YangName = "logical-channel-assignments"
    logicalChannelAssignments.EntityData.BundleName = "cisco_ios_xr"
    logicalChannelAssignments.EntityData.ParentYangName = "channel"
    logicalChannelAssignments.EntityData.SegmentPath = "logical-channel-assignments"
    logicalChannelAssignments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logicalChannelAssignments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logicalChannelAssignments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logicalChannelAssignments.EntityData.Children = make(map[string]types.YChild)
    logicalChannelAssignments.EntityData.Children["logical-channel-assignment"] = types.YChild{"LogicalChannelAssignment", nil}
    for i := range logicalChannelAssignments.LogicalChannelAssignment {
        logicalChannelAssignments.EntityData.Children[types.GetSegmentPath(&logicalChannelAssignments.LogicalChannelAssignment[i])] = types.YChild{"LogicalChannelAssignment", &logicalChannelAssignments.LogicalChannelAssignment[i]}
    }
    logicalChannelAssignments.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(logicalChannelAssignments.EntityData)
}

// LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment
// Logical Channel Assignment id
type LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment struct {
    EntityData types.CommonEntityData
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

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetEntityData() *types.CommonEntityData {
    logicalChannelAssignment.EntityData.YFilter = logicalChannelAssignment.YFilter
    logicalChannelAssignment.EntityData.YangName = "logical-channel-assignment"
    logicalChannelAssignment.EntityData.BundleName = "cisco_ios_xr"
    logicalChannelAssignment.EntityData.ParentYangName = "logical-channel-assignments"
    logicalChannelAssignment.EntityData.SegmentPath = "logical-channel-assignment" + "[assignment-index='" + fmt.Sprintf("%v", logicalChannelAssignment.AssignmentIndex) + "']"
    logicalChannelAssignment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logicalChannelAssignment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logicalChannelAssignment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logicalChannelAssignment.EntityData.Children = make(map[string]types.YChild)
    logicalChannelAssignment.EntityData.Leafs = make(map[string]types.YLeaf)
    logicalChannelAssignment.EntityData.Leafs["assignment-index"] = types.YLeaf{"AssignmentIndex", logicalChannelAssignment.AssignmentIndex}
    logicalChannelAssignment.EntityData.Leafs["description"] = types.YLeaf{"Description", logicalChannelAssignment.Description}
    logicalChannelAssignment.EntityData.Leafs["logical-channel-id"] = types.YLeaf{"LogicalChannelId", logicalChannelAssignment.LogicalChannelId}
    logicalChannelAssignment.EntityData.Leafs["assignment-type"] = types.YLeaf{"AssignmentType", logicalChannelAssignment.AssignmentType}
    logicalChannelAssignment.EntityData.Leafs["allocation"] = types.YLeaf{"Allocation", logicalChannelAssignment.Allocation}
    logicalChannelAssignment.EntityData.Leafs["optical-channel-id"] = types.YLeaf{"OpticalChannelId", logicalChannelAssignment.OpticalChannelId}
    return &(logicalChannelAssignment.EntityData)
}

// LogicalChannels_Channel_Otn
// Otn Related configs for Logical channel
type LogicalChannels_Channel_Otn struct {
    EntityData types.CommonEntityData
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

func (otn *LogicalChannels_Channel_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "channel"
    otn.EntityData.SegmentPath = "otn"
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = make(map[string]types.YChild)
    otn.EntityData.Leafs = make(map[string]types.YLeaf)
    otn.EntityData.Leafs["tti-msg-auto"] = types.YLeaf{"TtiMsgAuto", otn.TtiMsgAuto}
    otn.EntityData.Leafs["tti-msg-expected"] = types.YLeaf{"TtiMsgExpected", otn.TtiMsgExpected}
    otn.EntityData.Leafs["tti-msg-transmit"] = types.YLeaf{"TtiMsgTransmit", otn.TtiMsgTransmit}
    return &(otn.EntityData)
}

// OpticalChannels
// optical channels
type OpticalChannels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optical Channel index. The type is slice of OpticalChannels_OpticalChannel.
    OpticalChannel []OpticalChannels_OpticalChannel
}

func (opticalChannels *OpticalChannels) GetEntityData() *types.CommonEntityData {
    opticalChannels.EntityData.YFilter = opticalChannels.YFilter
    opticalChannels.EntityData.YangName = "optical-channels"
    opticalChannels.EntityData.BundleName = "cisco_ios_xr"
    opticalChannels.EntityData.ParentYangName = "Cisco-IOS-XR-terminal-device-cfg"
    opticalChannels.EntityData.SegmentPath = "Cisco-IOS-XR-terminal-device-cfg:optical-channels"
    opticalChannels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalChannels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalChannels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalChannels.EntityData.Children = make(map[string]types.YChild)
    opticalChannels.EntityData.Children["optical-channel"] = types.YChild{"OpticalChannel", nil}
    for i := range opticalChannels.OpticalChannel {
        opticalChannels.EntityData.Children[types.GetSegmentPath(&opticalChannels.OpticalChannel[i])] = types.YChild{"OpticalChannel", &opticalChannels.OpticalChannel[i]}
    }
    opticalChannels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(opticalChannels.EntityData)
}

// OpticalChannels_OpticalChannel
// Optical Channel index
type OpticalChannels_OpticalChannel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Optical Channel Name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Ifname interface{}

    // Configure operational mode. The type is interface{} with range: 1..100000.
    OperationalMode interface{}

    // Specify R/S/I/P. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    LinePort interface{}
}

func (opticalChannel *OpticalChannels_OpticalChannel) GetEntityData() *types.CommonEntityData {
    opticalChannel.EntityData.YFilter = opticalChannel.YFilter
    opticalChannel.EntityData.YangName = "optical-channel"
    opticalChannel.EntityData.BundleName = "cisco_ios_xr"
    opticalChannel.EntityData.ParentYangName = "optical-channels"
    opticalChannel.EntityData.SegmentPath = "optical-channel" + "[ifname='" + fmt.Sprintf("%v", opticalChannel.Ifname) + "']"
    opticalChannel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalChannel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalChannel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalChannel.EntityData.Children = make(map[string]types.YChild)
    opticalChannel.EntityData.Leafs = make(map[string]types.YLeaf)
    opticalChannel.EntityData.Leafs["ifname"] = types.YLeaf{"Ifname", opticalChannel.Ifname}
    opticalChannel.EntityData.Leafs["operational-mode"] = types.YLeaf{"OperationalMode", opticalChannel.OperationalMode}
    opticalChannel.EntityData.Leafs["line-port"] = types.YLeaf{"LinePort", opticalChannel.LinePort}
    return &(opticalChannel.EntityData)
}

