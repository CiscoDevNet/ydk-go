// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-statsd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   infra-statistics: Statistics Infrastructure
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_statsd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_statsd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-statsd-oper infra-statistics}", reflect.TypeOf(InfraStatistics{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-statsd-oper:infra-statistics", reflect.TypeOf(InfraStatistics{}))
}

// InfraStatistics
// Statistics Infrastructure
type InfraStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of interfaces.
    Interfaces InfraStatistics_Interfaces
}

func (infraStatistics *InfraStatistics) GetEntityData() *types.CommonEntityData {
    infraStatistics.EntityData.YFilter = infraStatistics.YFilter
    infraStatistics.EntityData.YangName = "infra-statistics"
    infraStatistics.EntityData.BundleName = "cisco_ios_xr"
    infraStatistics.EntityData.ParentYangName = "Cisco-IOS-XR-infra-statsd-oper"
    infraStatistics.EntityData.SegmentPath = "Cisco-IOS-XR-infra-statsd-oper:infra-statistics"
    infraStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraStatistics.EntityData.Children = types.NewOrderedMap()
    infraStatistics.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &infraStatistics.Interfaces})
    infraStatistics.EntityData.Leafs = types.NewOrderedMap()

    infraStatistics.EntityData.YListKeys = []string {}

    return &(infraStatistics.EntityData)
}

// InfraStatistics_Interfaces
// List of interfaces
type InfraStatistics_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of an interface. The type is slice of
    // InfraStatistics_Interfaces_Interface.
    Interface []*InfraStatistics_Interfaces_Interface
}

func (interfaces *InfraStatistics_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "infra-statistics"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// InfraStatistics_Interfaces_Interface
// Statistics of an interface
type InfraStatistics_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Cached stats data of interfaces.
    Cache InfraStatistics_Interfaces_Interface_Cache

    // Latest stats data of interfaces.
    Latest InfraStatistics_Interfaces_Interface_Latest

    // Total stats data of interfaces.
    Total InfraStatistics_Interfaces_Interface_Total

    // List of protocols.
    Protocols InfraStatistics_Interfaces_Interface_Protocols

    // Set of interface counters as displayed by the InterfacesMIB.
    InterfacesMibCounters InfraStatistics_Interfaces_Interface_InterfacesMibCounters

    // Datarate information.
    DataRate InfraStatistics_Interfaces_Interface_DataRate

    // Generic set of interface counters.
    GenericCounters InfraStatistics_Interfaces_Interface_GenericCounters
}

func (self *InfraStatistics_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("cache", types.YChild{"Cache", &self.Cache})
    self.EntityData.Children.Append("latest", types.YChild{"Latest", &self.Latest})
    self.EntityData.Children.Append("total", types.YChild{"Total", &self.Total})
    self.EntityData.Children.Append("protocols", types.YChild{"Protocols", &self.Protocols})
    self.EntityData.Children.Append("interfaces-mib-counters", types.YChild{"InterfacesMibCounters", &self.InterfacesMibCounters})
    self.EntityData.Children.Append("data-rate", types.YChild{"DataRate", &self.DataRate})
    self.EntityData.Children.Append("generic-counters", types.YChild{"GenericCounters", &self.GenericCounters})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// InfraStatistics_Interfaces_Interface_Cache
// Cached stats data of interfaces
type InfraStatistics_Interfaces_Interface_Cache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of protocols.
    Protocols InfraStatistics_Interfaces_Interface_Cache_Protocols

    // Set of interface counters as displayed by the InterfacesMIB.
    InterfacesMibCounters InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters

    // Datarate information.
    DataRate InfraStatistics_Interfaces_Interface_Cache_DataRate

    // Generic set of interface counters.
    GenericCounters InfraStatistics_Interfaces_Interface_Cache_GenericCounters
}

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetEntityData() *types.CommonEntityData {
    cache.EntityData.YFilter = cache.YFilter
    cache.EntityData.YangName = "cache"
    cache.EntityData.BundleName = "cisco_ios_xr"
    cache.EntityData.ParentYangName = "interface"
    cache.EntityData.SegmentPath = "cache"
    cache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cache.EntityData.Children = types.NewOrderedMap()
    cache.EntityData.Children.Append("protocols", types.YChild{"Protocols", &cache.Protocols})
    cache.EntityData.Children.Append("interfaces-mib-counters", types.YChild{"InterfacesMibCounters", &cache.InterfacesMibCounters})
    cache.EntityData.Children.Append("data-rate", types.YChild{"DataRate", &cache.DataRate})
    cache.EntityData.Children.Append("generic-counters", types.YChild{"GenericCounters", &cache.GenericCounters})
    cache.EntityData.Leafs = types.NewOrderedMap()

    cache.EntityData.YListKeys = []string {}

    return &(cache.EntityData)
}

// InfraStatistics_Interfaces_Interface_Cache_Protocols
// List of protocols
type InfraStatistics_Interfaces_Interface_Cache_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface counters per protocol. The type is slice of
    // InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol.
    Protocol []*InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol
}

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "cache"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range protocols.Protocol {
        protocols.EntityData.Children.Append(types.GetSegmentPath(protocols.Protocol[i]), types.YChild{"Protocol", protocols.Protocol[i]})
    }
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol
// Interface counters per protocol
type InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the protocol. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProtocolName interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Protocol number. The type is interface{} with range: 0..4294967295.
    Protocol interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..18446744073709551615. Units are second.
    LastDataTime interface{}

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}
}

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "protocols"
    protocol.EntityData.SegmentPath = "protocol" + types.AddKeyToken(protocol.ProtocolName, "protocol-name")
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", protocol.ProtocolName})
    protocol.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", protocol.BytesReceived})
    protocol.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", protocol.PacketsReceived})
    protocol.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", protocol.BytesSent})
    protocol.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", protocol.PacketsSent})
    protocol.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", protocol.Protocol})
    protocol.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", protocol.LastDataTime})
    protocol.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", protocol.InputDataRate})
    protocol.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", protocol.InputPacketRate})
    protocol.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", protocol.OutputDataRate})
    protocol.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", protocol.OutputPacketRate})

    protocol.EntityData.YListKeys = []string {"ProtocolName"}

    return &(protocol.EntityData)
}

// InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters
// Set of interface counters as displayed by the
// InterfacesMIB
type InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetEntityData() *types.CommonEntityData {
    interfacesMibCounters.EntityData.YFilter = interfacesMibCounters.YFilter
    interfacesMibCounters.EntityData.YangName = "interfaces-mib-counters"
    interfacesMibCounters.EntityData.BundleName = "cisco_ios_xr"
    interfacesMibCounters.EntityData.ParentYangName = "cache"
    interfacesMibCounters.EntityData.SegmentPath = "interfaces-mib-counters"
    interfacesMibCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfacesMibCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfacesMibCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfacesMibCounters.EntityData.Children = types.NewOrderedMap()
    interfacesMibCounters.EntityData.Leafs = types.NewOrderedMap()
    interfacesMibCounters.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", interfacesMibCounters.PacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", interfacesMibCounters.BytesReceived})
    interfacesMibCounters.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", interfacesMibCounters.PacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", interfacesMibCounters.BytesSent})
    interfacesMibCounters.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", interfacesMibCounters.MulticastPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", interfacesMibCounters.BroadcastPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", interfacesMibCounters.MulticastPacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", interfacesMibCounters.BroadcastPacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", interfacesMibCounters.OutputDrops})
    interfacesMibCounters.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", interfacesMibCounters.OutputQueueDrops})
    interfacesMibCounters.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", interfacesMibCounters.InputDrops})
    interfacesMibCounters.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", interfacesMibCounters.InputQueueDrops})
    interfacesMibCounters.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", interfacesMibCounters.RuntPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", interfacesMibCounters.GiantPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", interfacesMibCounters.ThrottledPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", interfacesMibCounters.ParityPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", interfacesMibCounters.UnknownProtocolPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", interfacesMibCounters.InputErrors})
    interfacesMibCounters.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", interfacesMibCounters.CrcErrors})
    interfacesMibCounters.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", interfacesMibCounters.InputOverruns})
    interfacesMibCounters.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", interfacesMibCounters.FramingErrorsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", interfacesMibCounters.InputIgnoredPackets})
    interfacesMibCounters.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", interfacesMibCounters.InputAborts})
    interfacesMibCounters.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", interfacesMibCounters.OutputErrors})
    interfacesMibCounters.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", interfacesMibCounters.OutputUnderruns})
    interfacesMibCounters.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", interfacesMibCounters.OutputBufferFailures})
    interfacesMibCounters.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", interfacesMibCounters.OutputBuffersSwappedOut})
    interfacesMibCounters.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", interfacesMibCounters.Applique})
    interfacesMibCounters.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", interfacesMibCounters.Resets})
    interfacesMibCounters.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", interfacesMibCounters.CarrierTransitions})
    interfacesMibCounters.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", interfacesMibCounters.AvailabilityFlag})
    interfacesMibCounters.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", interfacesMibCounters.LastDataTime})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", interfacesMibCounters.SecondsSinceLastClearCounters})
    interfacesMibCounters.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", interfacesMibCounters.LastDiscontinuityTime})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", interfacesMibCounters.SecondsSincePacketReceived})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", interfacesMibCounters.SecondsSincePacketSent})

    interfacesMibCounters.EntityData.YListKeys = []string {}

    return &(interfacesMibCounters.EntityData)
}

// InfraStatistics_Interfaces_Interface_Cache_DataRate
// Datarate information
type InfraStatistics_Interfaces_Interface_Cache_DataRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}

    // Peak input data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputDataRate interface{}

    // Peak input packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputPacketRate interface{}

    // Peak output data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputDataRate interface{}

    // Peak output packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputPacketRate interface{}

    // Bandwidth (in kbps). The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    Bandwidth interface{}

    // Number of 30-sec intervals less one. The type is interface{} with range:
    // 0..4294967295.
    LoadInterval interface{}

    // Output load as fraction of 255. The type is interface{} with range: 0..255.
    OutputLoad interface{}

    // Input load as fraction of 255. The type is interface{} with range: 0..255.
    InputLoad interface{}

    // Reliability coefficient. The type is interface{} with range: 0..255.
    Reliability interface{}
}

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetEntityData() *types.CommonEntityData {
    dataRate.EntityData.YFilter = dataRate.YFilter
    dataRate.EntityData.YangName = "data-rate"
    dataRate.EntityData.BundleName = "cisco_ios_xr"
    dataRate.EntityData.ParentYangName = "cache"
    dataRate.EntityData.SegmentPath = "data-rate"
    dataRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRate.EntityData.Children = types.NewOrderedMap()
    dataRate.EntityData.Leafs = types.NewOrderedMap()
    dataRate.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", dataRate.InputDataRate})
    dataRate.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", dataRate.InputPacketRate})
    dataRate.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", dataRate.OutputDataRate})
    dataRate.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", dataRate.OutputPacketRate})
    dataRate.EntityData.Leafs.Append("peak-input-data-rate", types.YLeaf{"PeakInputDataRate", dataRate.PeakInputDataRate})
    dataRate.EntityData.Leafs.Append("peak-input-packet-rate", types.YLeaf{"PeakInputPacketRate", dataRate.PeakInputPacketRate})
    dataRate.EntityData.Leafs.Append("peak-output-data-rate", types.YLeaf{"PeakOutputDataRate", dataRate.PeakOutputDataRate})
    dataRate.EntityData.Leafs.Append("peak-output-packet-rate", types.YLeaf{"PeakOutputPacketRate", dataRate.PeakOutputPacketRate})
    dataRate.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", dataRate.Bandwidth})
    dataRate.EntityData.Leafs.Append("load-interval", types.YLeaf{"LoadInterval", dataRate.LoadInterval})
    dataRate.EntityData.Leafs.Append("output-load", types.YLeaf{"OutputLoad", dataRate.OutputLoad})
    dataRate.EntityData.Leafs.Append("input-load", types.YLeaf{"InputLoad", dataRate.InputLoad})
    dataRate.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", dataRate.Reliability})

    dataRate.EntityData.YListKeys = []string {}

    return &(dataRate.EntityData)
}

// InfraStatistics_Interfaces_Interface_Cache_GenericCounters
// Generic set of interface counters
type InfraStatistics_Interfaces_Interface_Cache_GenericCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetEntityData() *types.CommonEntityData {
    genericCounters.EntityData.YFilter = genericCounters.YFilter
    genericCounters.EntityData.YangName = "generic-counters"
    genericCounters.EntityData.BundleName = "cisco_ios_xr"
    genericCounters.EntityData.ParentYangName = "cache"
    genericCounters.EntityData.SegmentPath = "generic-counters"
    genericCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounters.EntityData.Children = types.NewOrderedMap()
    genericCounters.EntityData.Leafs = types.NewOrderedMap()
    genericCounters.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", genericCounters.PacketsReceived})
    genericCounters.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", genericCounters.BytesReceived})
    genericCounters.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", genericCounters.PacketsSent})
    genericCounters.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", genericCounters.BytesSent})
    genericCounters.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", genericCounters.MulticastPacketsReceived})
    genericCounters.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", genericCounters.BroadcastPacketsReceived})
    genericCounters.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", genericCounters.MulticastPacketsSent})
    genericCounters.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", genericCounters.BroadcastPacketsSent})
    genericCounters.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", genericCounters.OutputDrops})
    genericCounters.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", genericCounters.OutputQueueDrops})
    genericCounters.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", genericCounters.InputDrops})
    genericCounters.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", genericCounters.InputQueueDrops})
    genericCounters.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", genericCounters.RuntPacketsReceived})
    genericCounters.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", genericCounters.GiantPacketsReceived})
    genericCounters.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", genericCounters.ThrottledPacketsReceived})
    genericCounters.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", genericCounters.ParityPacketsReceived})
    genericCounters.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", genericCounters.UnknownProtocolPacketsReceived})
    genericCounters.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", genericCounters.InputErrors})
    genericCounters.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", genericCounters.CrcErrors})
    genericCounters.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", genericCounters.InputOverruns})
    genericCounters.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", genericCounters.FramingErrorsReceived})
    genericCounters.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", genericCounters.InputIgnoredPackets})
    genericCounters.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", genericCounters.InputAborts})
    genericCounters.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", genericCounters.OutputErrors})
    genericCounters.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", genericCounters.OutputUnderruns})
    genericCounters.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", genericCounters.OutputBufferFailures})
    genericCounters.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", genericCounters.OutputBuffersSwappedOut})
    genericCounters.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", genericCounters.Applique})
    genericCounters.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", genericCounters.Resets})
    genericCounters.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", genericCounters.CarrierTransitions})
    genericCounters.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", genericCounters.AvailabilityFlag})
    genericCounters.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", genericCounters.LastDataTime})
    genericCounters.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", genericCounters.SecondsSinceLastClearCounters})
    genericCounters.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", genericCounters.LastDiscontinuityTime})
    genericCounters.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", genericCounters.SecondsSincePacketReceived})
    genericCounters.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", genericCounters.SecondsSincePacketSent})

    genericCounters.EntityData.YListKeys = []string {}

    return &(genericCounters.EntityData)
}

// InfraStatistics_Interfaces_Interface_Latest
// Latest stats data of interfaces
type InfraStatistics_Interfaces_Interface_Latest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of protocols.
    Protocols InfraStatistics_Interfaces_Interface_Latest_Protocols

    // Set of interface counters as displayed by the InterfacesMIB.
    InterfacesMibCounters InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters

    // Datarate information.
    DataRate InfraStatistics_Interfaces_Interface_Latest_DataRate

    // Generic set of interface counters.
    GenericCounters InfraStatistics_Interfaces_Interface_Latest_GenericCounters
}

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetEntityData() *types.CommonEntityData {
    latest.EntityData.YFilter = latest.YFilter
    latest.EntityData.YangName = "latest"
    latest.EntityData.BundleName = "cisco_ios_xr"
    latest.EntityData.ParentYangName = "interface"
    latest.EntityData.SegmentPath = "latest"
    latest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    latest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    latest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    latest.EntityData.Children = types.NewOrderedMap()
    latest.EntityData.Children.Append("protocols", types.YChild{"Protocols", &latest.Protocols})
    latest.EntityData.Children.Append("interfaces-mib-counters", types.YChild{"InterfacesMibCounters", &latest.InterfacesMibCounters})
    latest.EntityData.Children.Append("data-rate", types.YChild{"DataRate", &latest.DataRate})
    latest.EntityData.Children.Append("generic-counters", types.YChild{"GenericCounters", &latest.GenericCounters})
    latest.EntityData.Leafs = types.NewOrderedMap()

    latest.EntityData.YListKeys = []string {}

    return &(latest.EntityData)
}

// InfraStatistics_Interfaces_Interface_Latest_Protocols
// List of protocols
type InfraStatistics_Interfaces_Interface_Latest_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface counters per protocol. The type is slice of
    // InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol.
    Protocol []*InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol
}

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "latest"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range protocols.Protocol {
        protocols.EntityData.Children.Append(types.GetSegmentPath(protocols.Protocol[i]), types.YChild{"Protocol", protocols.Protocol[i]})
    }
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol
// Interface counters per protocol
type InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the protocol. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProtocolName interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Protocol number. The type is interface{} with range: 0..4294967295.
    Protocol interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..18446744073709551615. Units are second.
    LastDataTime interface{}

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}
}

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "protocols"
    protocol.EntityData.SegmentPath = "protocol" + types.AddKeyToken(protocol.ProtocolName, "protocol-name")
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", protocol.ProtocolName})
    protocol.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", protocol.BytesReceived})
    protocol.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", protocol.PacketsReceived})
    protocol.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", protocol.BytesSent})
    protocol.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", protocol.PacketsSent})
    protocol.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", protocol.Protocol})
    protocol.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", protocol.LastDataTime})
    protocol.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", protocol.InputDataRate})
    protocol.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", protocol.InputPacketRate})
    protocol.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", protocol.OutputDataRate})
    protocol.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", protocol.OutputPacketRate})

    protocol.EntityData.YListKeys = []string {"ProtocolName"}

    return &(protocol.EntityData)
}

// InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters
// Set of interface counters as displayed by the
// InterfacesMIB
type InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetEntityData() *types.CommonEntityData {
    interfacesMibCounters.EntityData.YFilter = interfacesMibCounters.YFilter
    interfacesMibCounters.EntityData.YangName = "interfaces-mib-counters"
    interfacesMibCounters.EntityData.BundleName = "cisco_ios_xr"
    interfacesMibCounters.EntityData.ParentYangName = "latest"
    interfacesMibCounters.EntityData.SegmentPath = "interfaces-mib-counters"
    interfacesMibCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfacesMibCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfacesMibCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfacesMibCounters.EntityData.Children = types.NewOrderedMap()
    interfacesMibCounters.EntityData.Leafs = types.NewOrderedMap()
    interfacesMibCounters.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", interfacesMibCounters.PacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", interfacesMibCounters.BytesReceived})
    interfacesMibCounters.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", interfacesMibCounters.PacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", interfacesMibCounters.BytesSent})
    interfacesMibCounters.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", interfacesMibCounters.MulticastPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", interfacesMibCounters.BroadcastPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", interfacesMibCounters.MulticastPacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", interfacesMibCounters.BroadcastPacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", interfacesMibCounters.OutputDrops})
    interfacesMibCounters.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", interfacesMibCounters.OutputQueueDrops})
    interfacesMibCounters.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", interfacesMibCounters.InputDrops})
    interfacesMibCounters.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", interfacesMibCounters.InputQueueDrops})
    interfacesMibCounters.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", interfacesMibCounters.RuntPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", interfacesMibCounters.GiantPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", interfacesMibCounters.ThrottledPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", interfacesMibCounters.ParityPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", interfacesMibCounters.UnknownProtocolPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", interfacesMibCounters.InputErrors})
    interfacesMibCounters.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", interfacesMibCounters.CrcErrors})
    interfacesMibCounters.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", interfacesMibCounters.InputOverruns})
    interfacesMibCounters.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", interfacesMibCounters.FramingErrorsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", interfacesMibCounters.InputIgnoredPackets})
    interfacesMibCounters.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", interfacesMibCounters.InputAborts})
    interfacesMibCounters.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", interfacesMibCounters.OutputErrors})
    interfacesMibCounters.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", interfacesMibCounters.OutputUnderruns})
    interfacesMibCounters.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", interfacesMibCounters.OutputBufferFailures})
    interfacesMibCounters.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", interfacesMibCounters.OutputBuffersSwappedOut})
    interfacesMibCounters.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", interfacesMibCounters.Applique})
    interfacesMibCounters.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", interfacesMibCounters.Resets})
    interfacesMibCounters.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", interfacesMibCounters.CarrierTransitions})
    interfacesMibCounters.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", interfacesMibCounters.AvailabilityFlag})
    interfacesMibCounters.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", interfacesMibCounters.LastDataTime})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", interfacesMibCounters.SecondsSinceLastClearCounters})
    interfacesMibCounters.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", interfacesMibCounters.LastDiscontinuityTime})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", interfacesMibCounters.SecondsSincePacketReceived})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", interfacesMibCounters.SecondsSincePacketSent})

    interfacesMibCounters.EntityData.YListKeys = []string {}

    return &(interfacesMibCounters.EntityData)
}

// InfraStatistics_Interfaces_Interface_Latest_DataRate
// Datarate information
type InfraStatistics_Interfaces_Interface_Latest_DataRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}

    // Peak input data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputDataRate interface{}

    // Peak input packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputPacketRate interface{}

    // Peak output data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputDataRate interface{}

    // Peak output packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputPacketRate interface{}

    // Bandwidth (in kbps). The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    Bandwidth interface{}

    // Number of 30-sec intervals less one. The type is interface{} with range:
    // 0..4294967295.
    LoadInterval interface{}

    // Output load as fraction of 255. The type is interface{} with range: 0..255.
    OutputLoad interface{}

    // Input load as fraction of 255. The type is interface{} with range: 0..255.
    InputLoad interface{}

    // Reliability coefficient. The type is interface{} with range: 0..255.
    Reliability interface{}
}

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetEntityData() *types.CommonEntityData {
    dataRate.EntityData.YFilter = dataRate.YFilter
    dataRate.EntityData.YangName = "data-rate"
    dataRate.EntityData.BundleName = "cisco_ios_xr"
    dataRate.EntityData.ParentYangName = "latest"
    dataRate.EntityData.SegmentPath = "data-rate"
    dataRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRate.EntityData.Children = types.NewOrderedMap()
    dataRate.EntityData.Leafs = types.NewOrderedMap()
    dataRate.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", dataRate.InputDataRate})
    dataRate.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", dataRate.InputPacketRate})
    dataRate.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", dataRate.OutputDataRate})
    dataRate.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", dataRate.OutputPacketRate})
    dataRate.EntityData.Leafs.Append("peak-input-data-rate", types.YLeaf{"PeakInputDataRate", dataRate.PeakInputDataRate})
    dataRate.EntityData.Leafs.Append("peak-input-packet-rate", types.YLeaf{"PeakInputPacketRate", dataRate.PeakInputPacketRate})
    dataRate.EntityData.Leafs.Append("peak-output-data-rate", types.YLeaf{"PeakOutputDataRate", dataRate.PeakOutputDataRate})
    dataRate.EntityData.Leafs.Append("peak-output-packet-rate", types.YLeaf{"PeakOutputPacketRate", dataRate.PeakOutputPacketRate})
    dataRate.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", dataRate.Bandwidth})
    dataRate.EntityData.Leafs.Append("load-interval", types.YLeaf{"LoadInterval", dataRate.LoadInterval})
    dataRate.EntityData.Leafs.Append("output-load", types.YLeaf{"OutputLoad", dataRate.OutputLoad})
    dataRate.EntityData.Leafs.Append("input-load", types.YLeaf{"InputLoad", dataRate.InputLoad})
    dataRate.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", dataRate.Reliability})

    dataRate.EntityData.YListKeys = []string {}

    return &(dataRate.EntityData)
}

// InfraStatistics_Interfaces_Interface_Latest_GenericCounters
// Generic set of interface counters
type InfraStatistics_Interfaces_Interface_Latest_GenericCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetEntityData() *types.CommonEntityData {
    genericCounters.EntityData.YFilter = genericCounters.YFilter
    genericCounters.EntityData.YangName = "generic-counters"
    genericCounters.EntityData.BundleName = "cisco_ios_xr"
    genericCounters.EntityData.ParentYangName = "latest"
    genericCounters.EntityData.SegmentPath = "generic-counters"
    genericCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounters.EntityData.Children = types.NewOrderedMap()
    genericCounters.EntityData.Leafs = types.NewOrderedMap()
    genericCounters.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", genericCounters.PacketsReceived})
    genericCounters.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", genericCounters.BytesReceived})
    genericCounters.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", genericCounters.PacketsSent})
    genericCounters.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", genericCounters.BytesSent})
    genericCounters.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", genericCounters.MulticastPacketsReceived})
    genericCounters.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", genericCounters.BroadcastPacketsReceived})
    genericCounters.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", genericCounters.MulticastPacketsSent})
    genericCounters.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", genericCounters.BroadcastPacketsSent})
    genericCounters.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", genericCounters.OutputDrops})
    genericCounters.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", genericCounters.OutputQueueDrops})
    genericCounters.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", genericCounters.InputDrops})
    genericCounters.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", genericCounters.InputQueueDrops})
    genericCounters.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", genericCounters.RuntPacketsReceived})
    genericCounters.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", genericCounters.GiantPacketsReceived})
    genericCounters.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", genericCounters.ThrottledPacketsReceived})
    genericCounters.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", genericCounters.ParityPacketsReceived})
    genericCounters.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", genericCounters.UnknownProtocolPacketsReceived})
    genericCounters.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", genericCounters.InputErrors})
    genericCounters.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", genericCounters.CrcErrors})
    genericCounters.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", genericCounters.InputOverruns})
    genericCounters.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", genericCounters.FramingErrorsReceived})
    genericCounters.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", genericCounters.InputIgnoredPackets})
    genericCounters.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", genericCounters.InputAborts})
    genericCounters.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", genericCounters.OutputErrors})
    genericCounters.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", genericCounters.OutputUnderruns})
    genericCounters.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", genericCounters.OutputBufferFailures})
    genericCounters.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", genericCounters.OutputBuffersSwappedOut})
    genericCounters.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", genericCounters.Applique})
    genericCounters.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", genericCounters.Resets})
    genericCounters.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", genericCounters.CarrierTransitions})
    genericCounters.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", genericCounters.AvailabilityFlag})
    genericCounters.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", genericCounters.LastDataTime})
    genericCounters.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", genericCounters.SecondsSinceLastClearCounters})
    genericCounters.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", genericCounters.LastDiscontinuityTime})
    genericCounters.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", genericCounters.SecondsSincePacketReceived})
    genericCounters.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", genericCounters.SecondsSincePacketSent})

    genericCounters.EntityData.YListKeys = []string {}

    return &(genericCounters.EntityData)
}

// InfraStatistics_Interfaces_Interface_Total
// Total stats data of interfaces
type InfraStatistics_Interfaces_Interface_Total struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of protocols.
    Protocols InfraStatistics_Interfaces_Interface_Total_Protocols

    // Set of interface counters as displayed by the InterfacesMIB.
    InterfacesMibCounters InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters

    // Datarate information.
    DataRate InfraStatistics_Interfaces_Interface_Total_DataRate

    // Generic set of interface counters.
    GenericCounters InfraStatistics_Interfaces_Interface_Total_GenericCounters
}

func (total *InfraStatistics_Interfaces_Interface_Total) GetEntityData() *types.CommonEntityData {
    total.EntityData.YFilter = total.YFilter
    total.EntityData.YangName = "total"
    total.EntityData.BundleName = "cisco_ios_xr"
    total.EntityData.ParentYangName = "interface"
    total.EntityData.SegmentPath = "total"
    total.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    total.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    total.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    total.EntityData.Children = types.NewOrderedMap()
    total.EntityData.Children.Append("protocols", types.YChild{"Protocols", &total.Protocols})
    total.EntityData.Children.Append("interfaces-mib-counters", types.YChild{"InterfacesMibCounters", &total.InterfacesMibCounters})
    total.EntityData.Children.Append("data-rate", types.YChild{"DataRate", &total.DataRate})
    total.EntityData.Children.Append("generic-counters", types.YChild{"GenericCounters", &total.GenericCounters})
    total.EntityData.Leafs = types.NewOrderedMap()

    total.EntityData.YListKeys = []string {}

    return &(total.EntityData)
}

// InfraStatistics_Interfaces_Interface_Total_Protocols
// List of protocols
type InfraStatistics_Interfaces_Interface_Total_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface counters per protocol. The type is slice of
    // InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol.
    Protocol []*InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol
}

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "total"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range protocols.Protocol {
        protocols.EntityData.Children.Append(types.GetSegmentPath(protocols.Protocol[i]), types.YChild{"Protocol", protocols.Protocol[i]})
    }
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol
// Interface counters per protocol
type InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the protocol. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProtocolName interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Protocol number. The type is interface{} with range: 0..4294967295.
    Protocol interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..18446744073709551615. Units are second.
    LastDataTime interface{}

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}
}

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "protocols"
    protocol.EntityData.SegmentPath = "protocol" + types.AddKeyToken(protocol.ProtocolName, "protocol-name")
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", protocol.ProtocolName})
    protocol.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", protocol.BytesReceived})
    protocol.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", protocol.PacketsReceived})
    protocol.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", protocol.BytesSent})
    protocol.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", protocol.PacketsSent})
    protocol.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", protocol.Protocol})
    protocol.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", protocol.LastDataTime})
    protocol.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", protocol.InputDataRate})
    protocol.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", protocol.InputPacketRate})
    protocol.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", protocol.OutputDataRate})
    protocol.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", protocol.OutputPacketRate})

    protocol.EntityData.YListKeys = []string {"ProtocolName"}

    return &(protocol.EntityData)
}

// InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters
// Set of interface counters as displayed by the
// InterfacesMIB
type InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetEntityData() *types.CommonEntityData {
    interfacesMibCounters.EntityData.YFilter = interfacesMibCounters.YFilter
    interfacesMibCounters.EntityData.YangName = "interfaces-mib-counters"
    interfacesMibCounters.EntityData.BundleName = "cisco_ios_xr"
    interfacesMibCounters.EntityData.ParentYangName = "total"
    interfacesMibCounters.EntityData.SegmentPath = "interfaces-mib-counters"
    interfacesMibCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfacesMibCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfacesMibCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfacesMibCounters.EntityData.Children = types.NewOrderedMap()
    interfacesMibCounters.EntityData.Leafs = types.NewOrderedMap()
    interfacesMibCounters.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", interfacesMibCounters.PacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", interfacesMibCounters.BytesReceived})
    interfacesMibCounters.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", interfacesMibCounters.PacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", interfacesMibCounters.BytesSent})
    interfacesMibCounters.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", interfacesMibCounters.MulticastPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", interfacesMibCounters.BroadcastPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", interfacesMibCounters.MulticastPacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", interfacesMibCounters.BroadcastPacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", interfacesMibCounters.OutputDrops})
    interfacesMibCounters.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", interfacesMibCounters.OutputQueueDrops})
    interfacesMibCounters.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", interfacesMibCounters.InputDrops})
    interfacesMibCounters.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", interfacesMibCounters.InputQueueDrops})
    interfacesMibCounters.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", interfacesMibCounters.RuntPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", interfacesMibCounters.GiantPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", interfacesMibCounters.ThrottledPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", interfacesMibCounters.ParityPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", interfacesMibCounters.UnknownProtocolPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", interfacesMibCounters.InputErrors})
    interfacesMibCounters.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", interfacesMibCounters.CrcErrors})
    interfacesMibCounters.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", interfacesMibCounters.InputOverruns})
    interfacesMibCounters.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", interfacesMibCounters.FramingErrorsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", interfacesMibCounters.InputIgnoredPackets})
    interfacesMibCounters.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", interfacesMibCounters.InputAborts})
    interfacesMibCounters.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", interfacesMibCounters.OutputErrors})
    interfacesMibCounters.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", interfacesMibCounters.OutputUnderruns})
    interfacesMibCounters.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", interfacesMibCounters.OutputBufferFailures})
    interfacesMibCounters.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", interfacesMibCounters.OutputBuffersSwappedOut})
    interfacesMibCounters.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", interfacesMibCounters.Applique})
    interfacesMibCounters.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", interfacesMibCounters.Resets})
    interfacesMibCounters.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", interfacesMibCounters.CarrierTransitions})
    interfacesMibCounters.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", interfacesMibCounters.AvailabilityFlag})
    interfacesMibCounters.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", interfacesMibCounters.LastDataTime})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", interfacesMibCounters.SecondsSinceLastClearCounters})
    interfacesMibCounters.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", interfacesMibCounters.LastDiscontinuityTime})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", interfacesMibCounters.SecondsSincePacketReceived})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", interfacesMibCounters.SecondsSincePacketSent})

    interfacesMibCounters.EntityData.YListKeys = []string {}

    return &(interfacesMibCounters.EntityData)
}

// InfraStatistics_Interfaces_Interface_Total_DataRate
// Datarate information
type InfraStatistics_Interfaces_Interface_Total_DataRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}

    // Peak input data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputDataRate interface{}

    // Peak input packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputPacketRate interface{}

    // Peak output data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputDataRate interface{}

    // Peak output packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputPacketRate interface{}

    // Bandwidth (in kbps). The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    Bandwidth interface{}

    // Number of 30-sec intervals less one. The type is interface{} with range:
    // 0..4294967295.
    LoadInterval interface{}

    // Output load as fraction of 255. The type is interface{} with range: 0..255.
    OutputLoad interface{}

    // Input load as fraction of 255. The type is interface{} with range: 0..255.
    InputLoad interface{}

    // Reliability coefficient. The type is interface{} with range: 0..255.
    Reliability interface{}
}

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetEntityData() *types.CommonEntityData {
    dataRate.EntityData.YFilter = dataRate.YFilter
    dataRate.EntityData.YangName = "data-rate"
    dataRate.EntityData.BundleName = "cisco_ios_xr"
    dataRate.EntityData.ParentYangName = "total"
    dataRate.EntityData.SegmentPath = "data-rate"
    dataRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRate.EntityData.Children = types.NewOrderedMap()
    dataRate.EntityData.Leafs = types.NewOrderedMap()
    dataRate.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", dataRate.InputDataRate})
    dataRate.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", dataRate.InputPacketRate})
    dataRate.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", dataRate.OutputDataRate})
    dataRate.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", dataRate.OutputPacketRate})
    dataRate.EntityData.Leafs.Append("peak-input-data-rate", types.YLeaf{"PeakInputDataRate", dataRate.PeakInputDataRate})
    dataRate.EntityData.Leafs.Append("peak-input-packet-rate", types.YLeaf{"PeakInputPacketRate", dataRate.PeakInputPacketRate})
    dataRate.EntityData.Leafs.Append("peak-output-data-rate", types.YLeaf{"PeakOutputDataRate", dataRate.PeakOutputDataRate})
    dataRate.EntityData.Leafs.Append("peak-output-packet-rate", types.YLeaf{"PeakOutputPacketRate", dataRate.PeakOutputPacketRate})
    dataRate.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", dataRate.Bandwidth})
    dataRate.EntityData.Leafs.Append("load-interval", types.YLeaf{"LoadInterval", dataRate.LoadInterval})
    dataRate.EntityData.Leafs.Append("output-load", types.YLeaf{"OutputLoad", dataRate.OutputLoad})
    dataRate.EntityData.Leafs.Append("input-load", types.YLeaf{"InputLoad", dataRate.InputLoad})
    dataRate.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", dataRate.Reliability})

    dataRate.EntityData.YListKeys = []string {}

    return &(dataRate.EntityData)
}

// InfraStatistics_Interfaces_Interface_Total_GenericCounters
// Generic set of interface counters
type InfraStatistics_Interfaces_Interface_Total_GenericCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetEntityData() *types.CommonEntityData {
    genericCounters.EntityData.YFilter = genericCounters.YFilter
    genericCounters.EntityData.YangName = "generic-counters"
    genericCounters.EntityData.BundleName = "cisco_ios_xr"
    genericCounters.EntityData.ParentYangName = "total"
    genericCounters.EntityData.SegmentPath = "generic-counters"
    genericCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounters.EntityData.Children = types.NewOrderedMap()
    genericCounters.EntityData.Leafs = types.NewOrderedMap()
    genericCounters.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", genericCounters.PacketsReceived})
    genericCounters.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", genericCounters.BytesReceived})
    genericCounters.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", genericCounters.PacketsSent})
    genericCounters.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", genericCounters.BytesSent})
    genericCounters.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", genericCounters.MulticastPacketsReceived})
    genericCounters.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", genericCounters.BroadcastPacketsReceived})
    genericCounters.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", genericCounters.MulticastPacketsSent})
    genericCounters.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", genericCounters.BroadcastPacketsSent})
    genericCounters.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", genericCounters.OutputDrops})
    genericCounters.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", genericCounters.OutputQueueDrops})
    genericCounters.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", genericCounters.InputDrops})
    genericCounters.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", genericCounters.InputQueueDrops})
    genericCounters.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", genericCounters.RuntPacketsReceived})
    genericCounters.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", genericCounters.GiantPacketsReceived})
    genericCounters.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", genericCounters.ThrottledPacketsReceived})
    genericCounters.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", genericCounters.ParityPacketsReceived})
    genericCounters.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", genericCounters.UnknownProtocolPacketsReceived})
    genericCounters.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", genericCounters.InputErrors})
    genericCounters.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", genericCounters.CrcErrors})
    genericCounters.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", genericCounters.InputOverruns})
    genericCounters.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", genericCounters.FramingErrorsReceived})
    genericCounters.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", genericCounters.InputIgnoredPackets})
    genericCounters.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", genericCounters.InputAborts})
    genericCounters.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", genericCounters.OutputErrors})
    genericCounters.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", genericCounters.OutputUnderruns})
    genericCounters.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", genericCounters.OutputBufferFailures})
    genericCounters.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", genericCounters.OutputBuffersSwappedOut})
    genericCounters.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", genericCounters.Applique})
    genericCounters.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", genericCounters.Resets})
    genericCounters.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", genericCounters.CarrierTransitions})
    genericCounters.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", genericCounters.AvailabilityFlag})
    genericCounters.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", genericCounters.LastDataTime})
    genericCounters.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", genericCounters.SecondsSinceLastClearCounters})
    genericCounters.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", genericCounters.LastDiscontinuityTime})
    genericCounters.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", genericCounters.SecondsSincePacketReceived})
    genericCounters.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", genericCounters.SecondsSincePacketSent})

    genericCounters.EntityData.YListKeys = []string {}

    return &(genericCounters.EntityData)
}

// InfraStatistics_Interfaces_Interface_Protocols
// List of protocols
type InfraStatistics_Interfaces_Interface_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface counters per protocol. The type is slice of
    // InfraStatistics_Interfaces_Interface_Protocols_Protocol.
    Protocol []*InfraStatistics_Interfaces_Interface_Protocols_Protocol
}

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "interface"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range protocols.Protocol {
        protocols.EntityData.Children.Append(types.GetSegmentPath(protocols.Protocol[i]), types.YChild{"Protocol", protocols.Protocol[i]})
    }
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// InfraStatistics_Interfaces_Interface_Protocols_Protocol
// Interface counters per protocol
type InfraStatistics_Interfaces_Interface_Protocols_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the protocol. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProtocolName interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Protocol number. The type is interface{} with range: 0..4294967295.
    Protocol interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..18446744073709551615. Units are second.
    LastDataTime interface{}

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}
}

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "protocols"
    protocol.EntityData.SegmentPath = "protocol" + types.AddKeyToken(protocol.ProtocolName, "protocol-name")
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", protocol.ProtocolName})
    protocol.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", protocol.BytesReceived})
    protocol.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", protocol.PacketsReceived})
    protocol.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", protocol.BytesSent})
    protocol.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", protocol.PacketsSent})
    protocol.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", protocol.Protocol})
    protocol.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", protocol.LastDataTime})
    protocol.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", protocol.InputDataRate})
    protocol.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", protocol.InputPacketRate})
    protocol.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", protocol.OutputDataRate})
    protocol.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", protocol.OutputPacketRate})

    protocol.EntityData.YListKeys = []string {"ProtocolName"}

    return &(protocol.EntityData)
}

// InfraStatistics_Interfaces_Interface_InterfacesMibCounters
// Set of interface counters as displayed by the
// InterfacesMIB
type InfraStatistics_Interfaces_Interface_InterfacesMibCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetEntityData() *types.CommonEntityData {
    interfacesMibCounters.EntityData.YFilter = interfacesMibCounters.YFilter
    interfacesMibCounters.EntityData.YangName = "interfaces-mib-counters"
    interfacesMibCounters.EntityData.BundleName = "cisco_ios_xr"
    interfacesMibCounters.EntityData.ParentYangName = "interface"
    interfacesMibCounters.EntityData.SegmentPath = "interfaces-mib-counters"
    interfacesMibCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfacesMibCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfacesMibCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfacesMibCounters.EntityData.Children = types.NewOrderedMap()
    interfacesMibCounters.EntityData.Leafs = types.NewOrderedMap()
    interfacesMibCounters.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", interfacesMibCounters.PacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", interfacesMibCounters.BytesReceived})
    interfacesMibCounters.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", interfacesMibCounters.PacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", interfacesMibCounters.BytesSent})
    interfacesMibCounters.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", interfacesMibCounters.MulticastPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", interfacesMibCounters.BroadcastPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", interfacesMibCounters.MulticastPacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", interfacesMibCounters.BroadcastPacketsSent})
    interfacesMibCounters.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", interfacesMibCounters.OutputDrops})
    interfacesMibCounters.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", interfacesMibCounters.OutputQueueDrops})
    interfacesMibCounters.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", interfacesMibCounters.InputDrops})
    interfacesMibCounters.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", interfacesMibCounters.InputQueueDrops})
    interfacesMibCounters.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", interfacesMibCounters.RuntPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", interfacesMibCounters.GiantPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", interfacesMibCounters.ThrottledPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", interfacesMibCounters.ParityPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", interfacesMibCounters.UnknownProtocolPacketsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", interfacesMibCounters.InputErrors})
    interfacesMibCounters.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", interfacesMibCounters.CrcErrors})
    interfacesMibCounters.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", interfacesMibCounters.InputOverruns})
    interfacesMibCounters.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", interfacesMibCounters.FramingErrorsReceived})
    interfacesMibCounters.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", interfacesMibCounters.InputIgnoredPackets})
    interfacesMibCounters.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", interfacesMibCounters.InputAborts})
    interfacesMibCounters.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", interfacesMibCounters.OutputErrors})
    interfacesMibCounters.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", interfacesMibCounters.OutputUnderruns})
    interfacesMibCounters.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", interfacesMibCounters.OutputBufferFailures})
    interfacesMibCounters.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", interfacesMibCounters.OutputBuffersSwappedOut})
    interfacesMibCounters.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", interfacesMibCounters.Applique})
    interfacesMibCounters.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", interfacesMibCounters.Resets})
    interfacesMibCounters.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", interfacesMibCounters.CarrierTransitions})
    interfacesMibCounters.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", interfacesMibCounters.AvailabilityFlag})
    interfacesMibCounters.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", interfacesMibCounters.LastDataTime})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", interfacesMibCounters.SecondsSinceLastClearCounters})
    interfacesMibCounters.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", interfacesMibCounters.LastDiscontinuityTime})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", interfacesMibCounters.SecondsSincePacketReceived})
    interfacesMibCounters.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", interfacesMibCounters.SecondsSincePacketSent})

    interfacesMibCounters.EntityData.YListKeys = []string {}

    return &(interfacesMibCounters.EntityData)
}

// InfraStatistics_Interfaces_Interface_DataRate
// Datarate information
type InfraStatistics_Interfaces_Interface_DataRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}

    // Peak input data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputDataRate interface{}

    // Peak input packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputPacketRate interface{}

    // Peak output data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputDataRate interface{}

    // Peak output packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputPacketRate interface{}

    // Bandwidth (in kbps). The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    Bandwidth interface{}

    // Number of 30-sec intervals less one. The type is interface{} with range:
    // 0..4294967295.
    LoadInterval interface{}

    // Output load as fraction of 255. The type is interface{} with range: 0..255.
    OutputLoad interface{}

    // Input load as fraction of 255. The type is interface{} with range: 0..255.
    InputLoad interface{}

    // Reliability coefficient. The type is interface{} with range: 0..255.
    Reliability interface{}
}

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetEntityData() *types.CommonEntityData {
    dataRate.EntityData.YFilter = dataRate.YFilter
    dataRate.EntityData.YangName = "data-rate"
    dataRate.EntityData.BundleName = "cisco_ios_xr"
    dataRate.EntityData.ParentYangName = "interface"
    dataRate.EntityData.SegmentPath = "data-rate"
    dataRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRate.EntityData.Children = types.NewOrderedMap()
    dataRate.EntityData.Leafs = types.NewOrderedMap()
    dataRate.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", dataRate.InputDataRate})
    dataRate.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", dataRate.InputPacketRate})
    dataRate.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", dataRate.OutputDataRate})
    dataRate.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", dataRate.OutputPacketRate})
    dataRate.EntityData.Leafs.Append("peak-input-data-rate", types.YLeaf{"PeakInputDataRate", dataRate.PeakInputDataRate})
    dataRate.EntityData.Leafs.Append("peak-input-packet-rate", types.YLeaf{"PeakInputPacketRate", dataRate.PeakInputPacketRate})
    dataRate.EntityData.Leafs.Append("peak-output-data-rate", types.YLeaf{"PeakOutputDataRate", dataRate.PeakOutputDataRate})
    dataRate.EntityData.Leafs.Append("peak-output-packet-rate", types.YLeaf{"PeakOutputPacketRate", dataRate.PeakOutputPacketRate})
    dataRate.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", dataRate.Bandwidth})
    dataRate.EntityData.Leafs.Append("load-interval", types.YLeaf{"LoadInterval", dataRate.LoadInterval})
    dataRate.EntityData.Leafs.Append("output-load", types.YLeaf{"OutputLoad", dataRate.OutputLoad})
    dataRate.EntityData.Leafs.Append("input-load", types.YLeaf{"InputLoad", dataRate.InputLoad})
    dataRate.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", dataRate.Reliability})

    dataRate.EntityData.YListKeys = []string {}

    return &(dataRate.EntityData)
}

// InfraStatistics_Interfaces_Interface_GenericCounters
// Generic set of interface counters
type InfraStatistics_Interfaces_Interface_GenericCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetEntityData() *types.CommonEntityData {
    genericCounters.EntityData.YFilter = genericCounters.YFilter
    genericCounters.EntityData.YangName = "generic-counters"
    genericCounters.EntityData.BundleName = "cisco_ios_xr"
    genericCounters.EntityData.ParentYangName = "interface"
    genericCounters.EntityData.SegmentPath = "generic-counters"
    genericCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounters.EntityData.Children = types.NewOrderedMap()
    genericCounters.EntityData.Leafs = types.NewOrderedMap()
    genericCounters.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", genericCounters.PacketsReceived})
    genericCounters.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", genericCounters.BytesReceived})
    genericCounters.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", genericCounters.PacketsSent})
    genericCounters.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", genericCounters.BytesSent})
    genericCounters.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", genericCounters.MulticastPacketsReceived})
    genericCounters.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", genericCounters.BroadcastPacketsReceived})
    genericCounters.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", genericCounters.MulticastPacketsSent})
    genericCounters.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", genericCounters.BroadcastPacketsSent})
    genericCounters.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", genericCounters.OutputDrops})
    genericCounters.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", genericCounters.OutputQueueDrops})
    genericCounters.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", genericCounters.InputDrops})
    genericCounters.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", genericCounters.InputQueueDrops})
    genericCounters.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", genericCounters.RuntPacketsReceived})
    genericCounters.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", genericCounters.GiantPacketsReceived})
    genericCounters.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", genericCounters.ThrottledPacketsReceived})
    genericCounters.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", genericCounters.ParityPacketsReceived})
    genericCounters.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", genericCounters.UnknownProtocolPacketsReceived})
    genericCounters.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", genericCounters.InputErrors})
    genericCounters.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", genericCounters.CrcErrors})
    genericCounters.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", genericCounters.InputOverruns})
    genericCounters.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", genericCounters.FramingErrorsReceived})
    genericCounters.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", genericCounters.InputIgnoredPackets})
    genericCounters.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", genericCounters.InputAborts})
    genericCounters.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", genericCounters.OutputErrors})
    genericCounters.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", genericCounters.OutputUnderruns})
    genericCounters.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", genericCounters.OutputBufferFailures})
    genericCounters.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", genericCounters.OutputBuffersSwappedOut})
    genericCounters.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", genericCounters.Applique})
    genericCounters.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", genericCounters.Resets})
    genericCounters.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", genericCounters.CarrierTransitions})
    genericCounters.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", genericCounters.AvailabilityFlag})
    genericCounters.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", genericCounters.LastDataTime})
    genericCounters.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", genericCounters.SecondsSinceLastClearCounters})
    genericCounters.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", genericCounters.LastDiscontinuityTime})
    genericCounters.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", genericCounters.SecondsSincePacketReceived})
    genericCounters.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", genericCounters.SecondsSincePacketSent})

    genericCounters.EntityData.YListKeys = []string {}

    return &(genericCounters.EntityData)
}

