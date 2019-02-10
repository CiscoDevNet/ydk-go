// This module contains a collection of YANG definitions
// for Cisco IOS-XR terminal-device package configuration.
// 
// This module contains definitions
// for the following management objects:
//   logical-channels: Logical channel in mxp
//   optical-channels: optical channels
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    LogicalChannelOtnTtiAuto_false_ LogicalChannelOtnTtiAuto = "false"

    // Otn tti auto mode true
    LogicalChannelOtnTtiAuto_true_ LogicalChannelOtnTtiAuto = "true"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logical channel index. The type is slice of LogicalChannels_Channel.
    Channel []*LogicalChannels_Channel
}

func (logicalChannels *LogicalChannels) GetEntityData() *types.CommonEntityData {
    logicalChannels.EntityData.YFilter = logicalChannels.YFilter
    logicalChannels.EntityData.YangName = "logical-channels"
    logicalChannels.EntityData.BundleName = "cisco_ios_xr"
    logicalChannels.EntityData.ParentYangName = "Cisco-IOS-XR-terminal-device-cfg"
    logicalChannels.EntityData.SegmentPath = "Cisco-IOS-XR-terminal-device-cfg:logical-channels"
    logicalChannels.EntityData.AbsolutePath = logicalChannels.EntityData.SegmentPath
    logicalChannels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logicalChannels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logicalChannels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logicalChannels.EntityData.Children = types.NewOrderedMap()
    logicalChannels.EntityData.Children.Append("channel", types.YChild{"Channel", nil})
    for i := range logicalChannels.Channel {
        logicalChannels.EntityData.Children.Append(types.GetSegmentPath(logicalChannels.Channel[i]), types.YChild{"Channel", logicalChannels.Channel[i]})
    }
    logicalChannels.EntityData.Leafs = types.NewOrderedMap()

    logicalChannels.EntityData.YListKeys = []string {}

    return &(logicalChannels.EntityData)
}

// LogicalChannels_Channel
// Logical channel index
type LogicalChannels_Channel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Logical Channel Index. The type is interface{}
    // with range: 0..4294967295.
    ChannelIndex interface{}

    // Protocol framing of the tributary signal. The type is LogicalTribProtocol.
    TribProtocol interface{}

    // Description (Max 255 characters). The type is string with length: 1..255.
    Description interface{}

    // Configure ingress client port for this logical channel. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
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
    channel.EntityData.SegmentPath = "channel" + types.AddKeyToken(channel.ChannelIndex, "channel-index")
    channel.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-cfg:logical-channels/" + channel.EntityData.SegmentPath
    channel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channel.EntityData.Children = types.NewOrderedMap()
    channel.EntityData.Children.Append("logical-channel-assignments", types.YChild{"LogicalChannelAssignments", &channel.LogicalChannelAssignments})
    channel.EntityData.Children.Append("otn", types.YChild{"Otn", &channel.Otn})
    channel.EntityData.Leafs = types.NewOrderedMap()
    channel.EntityData.Leafs.Append("channel-index", types.YLeaf{"ChannelIndex", channel.ChannelIndex})
    channel.EntityData.Leafs.Append("trib-protocol", types.YLeaf{"TribProtocol", channel.TribProtocol})
    channel.EntityData.Leafs.Append("description", types.YLeaf{"Description", channel.Description})
    channel.EntityData.Leafs.Append("ingress-client-port", types.YLeaf{"IngressClientPort", channel.IngressClientPort})
    channel.EntityData.Leafs.Append("ingress-physical-channel", types.YLeaf{"IngressPhysicalChannel", channel.IngressPhysicalChannel})
    channel.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", channel.AdminState})
    channel.EntityData.Leafs.Append("loopback-mode", types.YLeaf{"LoopbackMode", channel.LoopbackMode})
    channel.EntityData.Leafs.Append("logical-channel-type", types.YLeaf{"LogicalChannelType", channel.LogicalChannelType})
    channel.EntityData.Leafs.Append("rate-class", types.YLeaf{"RateClass", channel.RateClass})

    channel.EntityData.YListKeys = []string {"ChannelIndex"}

    return &(channel.EntityData)
}

// LogicalChannels_Channel_LogicalChannelAssignments
// Logical channel assignment for logical channel
type LogicalChannels_Channel_LogicalChannelAssignments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logical Channel Assignment id. The type is slice of
    // LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment.
    LogicalChannelAssignment []*LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment
}

func (logicalChannelAssignments *LogicalChannels_Channel_LogicalChannelAssignments) GetEntityData() *types.CommonEntityData {
    logicalChannelAssignments.EntityData.YFilter = logicalChannelAssignments.YFilter
    logicalChannelAssignments.EntityData.YangName = "logical-channel-assignments"
    logicalChannelAssignments.EntityData.BundleName = "cisco_ios_xr"
    logicalChannelAssignments.EntityData.ParentYangName = "channel"
    logicalChannelAssignments.EntityData.SegmentPath = "logical-channel-assignments"
    logicalChannelAssignments.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-cfg:logical-channels/channel/" + logicalChannelAssignments.EntityData.SegmentPath
    logicalChannelAssignments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logicalChannelAssignments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logicalChannelAssignments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logicalChannelAssignments.EntityData.Children = types.NewOrderedMap()
    logicalChannelAssignments.EntityData.Children.Append("logical-channel-assignment", types.YChild{"LogicalChannelAssignment", nil})
    for i := range logicalChannelAssignments.LogicalChannelAssignment {
        logicalChannelAssignments.EntityData.Children.Append(types.GetSegmentPath(logicalChannelAssignments.LogicalChannelAssignment[i]), types.YChild{"LogicalChannelAssignment", logicalChannelAssignments.LogicalChannelAssignment[i]})
    }
    logicalChannelAssignments.EntityData.Leafs = types.NewOrderedMap()

    logicalChannelAssignments.EntityData.YListKeys = []string {}

    return &(logicalChannelAssignments.EntityData)
}

// LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment
// Logical Channel Assignment id
type LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Logical channel assignment index. The type is
    // interface{} with range: 0..4294967295.
    AssignmentIndex interface{}

    // Configure description for this assignment. The type is string with length:
    // 1..255.
    Description interface{}

    // Configure logical channel for this assignment. The type is interface{} with
    // range: 0..4294967295.
    LogicalChannelId interface{}

    // Type of assignment for logical channel. The type is
    // LogicalChannelAssignment.
    AssignmentType interface{}

    // Configure Allocation for this assignment(10, 40 or 100G). The type is
    // interface{} with range: 0..4294967295.
    Allocation interface{}

    // Configure optical channel for this assignment. The type is string.
    OpticalChannelId interface{}
}

func (logicalChannelAssignment *LogicalChannels_Channel_LogicalChannelAssignments_LogicalChannelAssignment) GetEntityData() *types.CommonEntityData {
    logicalChannelAssignment.EntityData.YFilter = logicalChannelAssignment.YFilter
    logicalChannelAssignment.EntityData.YangName = "logical-channel-assignment"
    logicalChannelAssignment.EntityData.BundleName = "cisco_ios_xr"
    logicalChannelAssignment.EntityData.ParentYangName = "logical-channel-assignments"
    logicalChannelAssignment.EntityData.SegmentPath = "logical-channel-assignment" + types.AddKeyToken(logicalChannelAssignment.AssignmentIndex, "assignment-index")
    logicalChannelAssignment.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-cfg:logical-channels/channel/logical-channel-assignments/" + logicalChannelAssignment.EntityData.SegmentPath
    logicalChannelAssignment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logicalChannelAssignment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logicalChannelAssignment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logicalChannelAssignment.EntityData.Children = types.NewOrderedMap()
    logicalChannelAssignment.EntityData.Leafs = types.NewOrderedMap()
    logicalChannelAssignment.EntityData.Leafs.Append("assignment-index", types.YLeaf{"AssignmentIndex", logicalChannelAssignment.AssignmentIndex})
    logicalChannelAssignment.EntityData.Leafs.Append("description", types.YLeaf{"Description", logicalChannelAssignment.Description})
    logicalChannelAssignment.EntityData.Leafs.Append("logical-channel-id", types.YLeaf{"LogicalChannelId", logicalChannelAssignment.LogicalChannelId})
    logicalChannelAssignment.EntityData.Leafs.Append("assignment-type", types.YLeaf{"AssignmentType", logicalChannelAssignment.AssignmentType})
    logicalChannelAssignment.EntityData.Leafs.Append("allocation", types.YLeaf{"Allocation", logicalChannelAssignment.Allocation})
    logicalChannelAssignment.EntityData.Leafs.Append("optical-channel-id", types.YLeaf{"OpticalChannelId", logicalChannelAssignment.OpticalChannelId})

    logicalChannelAssignment.EntityData.YListKeys = []string {"AssignmentIndex"}

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
    otn.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-cfg:logical-channels/channel/" + otn.EntityData.SegmentPath
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("tti-msg-auto", types.YLeaf{"TtiMsgAuto", otn.TtiMsgAuto})
    otn.EntityData.Leafs.Append("tti-msg-expected", types.YLeaf{"TtiMsgExpected", otn.TtiMsgExpected})
    otn.EntityData.Leafs.Append("tti-msg-transmit", types.YLeaf{"TtiMsgTransmit", otn.TtiMsgTransmit})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// OpticalChannels
// optical channels
type OpticalChannels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optical Channel index. The type is slice of OpticalChannels_OpticalChannel.
    OpticalChannel []*OpticalChannels_OpticalChannel
}

func (opticalChannels *OpticalChannels) GetEntityData() *types.CommonEntityData {
    opticalChannels.EntityData.YFilter = opticalChannels.YFilter
    opticalChannels.EntityData.YangName = "optical-channels"
    opticalChannels.EntityData.BundleName = "cisco_ios_xr"
    opticalChannels.EntityData.ParentYangName = "Cisco-IOS-XR-terminal-device-cfg"
    opticalChannels.EntityData.SegmentPath = "Cisco-IOS-XR-terminal-device-cfg:optical-channels"
    opticalChannels.EntityData.AbsolutePath = opticalChannels.EntityData.SegmentPath
    opticalChannels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalChannels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalChannels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalChannels.EntityData.Children = types.NewOrderedMap()
    opticalChannels.EntityData.Children.Append("optical-channel", types.YChild{"OpticalChannel", nil})
    for i := range opticalChannels.OpticalChannel {
        opticalChannels.EntityData.Children.Append(types.GetSegmentPath(opticalChannels.OpticalChannel[i]), types.YChild{"OpticalChannel", opticalChannels.OpticalChannel[i]})
    }
    opticalChannels.EntityData.Leafs = types.NewOrderedMap()

    opticalChannels.EntityData.YListKeys = []string {}

    return &(opticalChannels.EntityData)
}

// OpticalChannels_OpticalChannel
// Optical Channel index
type OpticalChannels_OpticalChannel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Optical Channel Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ifname interface{}

    // Configure operational mode. The type is interface{} with range: 1..100000.
    OperationalMode interface{}

    // Specify R/S/I/P. The type is string with pattern: [a-zA-Z0-9._/-]+.
    LinePort interface{}
}

func (opticalChannel *OpticalChannels_OpticalChannel) GetEntityData() *types.CommonEntityData {
    opticalChannel.EntityData.YFilter = opticalChannel.YFilter
    opticalChannel.EntityData.YangName = "optical-channel"
    opticalChannel.EntityData.BundleName = "cisco_ios_xr"
    opticalChannel.EntityData.ParentYangName = "optical-channels"
    opticalChannel.EntityData.SegmentPath = "optical-channel" + types.AddKeyToken(opticalChannel.Ifname, "ifname")
    opticalChannel.EntityData.AbsolutePath = "Cisco-IOS-XR-terminal-device-cfg:optical-channels/" + opticalChannel.EntityData.SegmentPath
    opticalChannel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalChannel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalChannel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalChannel.EntityData.Children = types.NewOrderedMap()
    opticalChannel.EntityData.Leafs = types.NewOrderedMap()
    opticalChannel.EntityData.Leafs.Append("ifname", types.YLeaf{"Ifname", opticalChannel.Ifname})
    opticalChannel.EntityData.Leafs.Append("operational-mode", types.YLeaf{"OperationalMode", opticalChannel.OperationalMode})
    opticalChannel.EntityData.Leafs.Append("line-port", types.YLeaf{"LinePort", opticalChannel.LinePort})

    opticalChannel.EntityData.YListKeys = []string {"Ifname"}

    return &(opticalChannel.EntityData)
}

