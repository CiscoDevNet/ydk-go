// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-statsd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   infra-statistics: Statistics Infrastructure
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of interfaces.
    Interfaces InfraStatistics_Interfaces
}

func (infraStatistics *InfraStatistics) GetFilter() yfilter.YFilter { return infraStatistics.YFilter }

func (infraStatistics *InfraStatistics) SetFilter(yf yfilter.YFilter) { infraStatistics.YFilter = yf }

func (infraStatistics *InfraStatistics) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (infraStatistics *InfraStatistics) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-statsd-oper:infra-statistics"
}

func (infraStatistics *InfraStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &infraStatistics.Interfaces
    }
    return nil
}

func (infraStatistics *InfraStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &infraStatistics.Interfaces
    return children
}

func (infraStatistics *InfraStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraStatistics *InfraStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (infraStatistics *InfraStatistics) GetYangName() string { return "infra-statistics" }

func (infraStatistics *InfraStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraStatistics *InfraStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraStatistics *InfraStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraStatistics *InfraStatistics) SetParent(parent types.Entity) { infraStatistics.parent = parent }

func (infraStatistics *InfraStatistics) GetParent() types.Entity { return infraStatistics.parent }

func (infraStatistics *InfraStatistics) GetParentYangName() string { return "Cisco-IOS-XR-infra-statsd-oper" }

// InfraStatistics_Interfaces
// List of interfaces
type InfraStatistics_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics of an interface. The type is slice of
    // InfraStatistics_Interfaces_Interface.
    Interface []InfraStatistics_Interfaces_Interface
}

func (interfaces *InfraStatistics_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *InfraStatistics_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *InfraStatistics_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *InfraStatistics_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *InfraStatistics_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InfraStatistics_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *InfraStatistics_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *InfraStatistics_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *InfraStatistics_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *InfraStatistics_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *InfraStatistics_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *InfraStatistics_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *InfraStatistics_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *InfraStatistics_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *InfraStatistics_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *InfraStatistics_Interfaces) GetParentYangName() string { return "infra-statistics" }

// InfraStatistics_Interfaces_Interface
// Statistics of an interface
type InfraStatistics_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (self *InfraStatistics_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *InfraStatistics_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *InfraStatistics_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "cache" { return "Cache" }
    if yname == "latest" { return "Latest" }
    if yname == "total" { return "Total" }
    if yname == "protocols" { return "Protocols" }
    if yname == "interfaces-mib-counters" { return "InterfacesMibCounters" }
    if yname == "data-rate" { return "DataRate" }
    if yname == "generic-counters" { return "GenericCounters" }
    return ""
}

func (self *InfraStatistics_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *InfraStatistics_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cache" {
        return &self.Cache
    }
    if childYangName == "latest" {
        return &self.Latest
    }
    if childYangName == "total" {
        return &self.Total
    }
    if childYangName == "protocols" {
        return &self.Protocols
    }
    if childYangName == "interfaces-mib-counters" {
        return &self.InterfacesMibCounters
    }
    if childYangName == "data-rate" {
        return &self.DataRate
    }
    if childYangName == "generic-counters" {
        return &self.GenericCounters
    }
    return nil
}

func (self *InfraStatistics_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cache"] = &self.Cache
    children["latest"] = &self.Latest
    children["total"] = &self.Total
    children["protocols"] = &self.Protocols
    children["interfaces-mib-counters"] = &self.InterfacesMibCounters
    children["data-rate"] = &self.DataRate
    children["generic-counters"] = &self.GenericCounters
    return children
}

func (self *InfraStatistics_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *InfraStatistics_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *InfraStatistics_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *InfraStatistics_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *InfraStatistics_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *InfraStatistics_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *InfraStatistics_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *InfraStatistics_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *InfraStatistics_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// InfraStatistics_Interfaces_Interface_Cache
// Cached stats data of interfaces
type InfraStatistics_Interfaces_Interface_Cache struct {
    parent types.Entity
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

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetFilter() yfilter.YFilter { return cache.YFilter }

func (cache *InfraStatistics_Interfaces_Interface_Cache) SetFilter(yf yfilter.YFilter) { cache.YFilter = yf }

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetGoName(yname string) string {
    if yname == "protocols" { return "Protocols" }
    if yname == "interfaces-mib-counters" { return "InterfacesMibCounters" }
    if yname == "data-rate" { return "DataRate" }
    if yname == "generic-counters" { return "GenericCounters" }
    return ""
}

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetSegmentPath() string {
    return "cache"
}

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocols" {
        return &cache.Protocols
    }
    if childYangName == "interfaces-mib-counters" {
        return &cache.InterfacesMibCounters
    }
    if childYangName == "data-rate" {
        return &cache.DataRate
    }
    if childYangName == "generic-counters" {
        return &cache.GenericCounters
    }
    return nil
}

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocols"] = &cache.Protocols
    children["interfaces-mib-counters"] = &cache.InterfacesMibCounters
    children["data-rate"] = &cache.DataRate
    children["generic-counters"] = &cache.GenericCounters
    return children
}

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetBundleName() string { return "cisco_ios_xr" }

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetYangName() string { return "cache" }

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cache *InfraStatistics_Interfaces_Interface_Cache) SetParent(parent types.Entity) { cache.parent = parent }

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetParent() types.Entity { return cache.parent }

func (cache *InfraStatistics_Interfaces_Interface_Cache) GetParentYangName() string { return "interface" }

// InfraStatistics_Interfaces_Interface_Cache_Protocols
// List of protocols
type InfraStatistics_Interfaces_Interface_Cache_Protocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface counters per protocol. The type is slice of
    // InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol.
    Protocol []InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol
}

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetFilter() yfilter.YFilter { return protocols.YFilter }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) SetFilter(yf yfilter.YFilter) { protocols.YFilter = yf }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetSegmentPath() string {
    return "protocols"
}

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        for _, c := range protocols.Protocol {
            if protocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol{}
        protocols.Protocol = append(protocols.Protocol, child)
        return &protocols.Protocol[len(protocols.Protocol)-1]
    }
    return nil
}

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocols.Protocol {
        children[protocols.Protocol[i].GetSegmentPath()] = &protocols.Protocol[i]
    }
    return children
}

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetBundleName() string { return "cisco_ios_xr" }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetYangName() string { return "protocols" }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) SetParent(parent types.Entity) { protocols.parent = parent }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetParent() types.Entity { return protocols.parent }

func (protocols *InfraStatistics_Interfaces_Interface_Cache_Protocols) GetParentYangName() string { return "cache" }

// InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol
// Interface counters per protocol
type InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol struct {
    parent types.Entity
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
    // with range: 0..4294967295. Units are second.
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

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "protocol" { return "Protocol" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    return ""
}

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetSegmentPath() string {
    return "protocol" + "[protocol-name='" + fmt.Sprintf("%v", protocol.ProtocolName) + "']"
}

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = protocol.ProtocolName
    leafs["bytes-received"] = protocol.BytesReceived
    leafs["packets-received"] = protocol.PacketsReceived
    leafs["bytes-sent"] = protocol.BytesSent
    leafs["packets-sent"] = protocol.PacketsSent
    leafs["protocol"] = protocol.Protocol
    leafs["last-data-time"] = protocol.LastDataTime
    leafs["input-data-rate"] = protocol.InputDataRate
    leafs["input-packet-rate"] = protocol.InputPacketRate
    leafs["output-data-rate"] = protocol.OutputDataRate
    leafs["output-packet-rate"] = protocol.OutputPacketRate
    return leafs
}

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetYangName() string { return "protocol" }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *InfraStatistics_Interfaces_Interface_Cache_Protocols_Protocol) GetParentYangName() string { return "protocols" }

// InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters
// Set of interface counters as displayed by the
// InterfacesMIB
type InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters struct {
    parent types.Entity
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

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetFilter() yfilter.YFilter { return interfacesMibCounters.YFilter }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) SetFilter(yf yfilter.YFilter) { interfacesMibCounters.YFilter = yf }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetSegmentPath() string {
    return "interfaces-mib-counters"
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = interfacesMibCounters.PacketsReceived
    leafs["bytes-received"] = interfacesMibCounters.BytesReceived
    leafs["packets-sent"] = interfacesMibCounters.PacketsSent
    leafs["bytes-sent"] = interfacesMibCounters.BytesSent
    leafs["multicast-packets-received"] = interfacesMibCounters.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = interfacesMibCounters.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = interfacesMibCounters.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = interfacesMibCounters.BroadcastPacketsSent
    leafs["output-drops"] = interfacesMibCounters.OutputDrops
    leafs["output-queue-drops"] = interfacesMibCounters.OutputQueueDrops
    leafs["input-drops"] = interfacesMibCounters.InputDrops
    leafs["input-queue-drops"] = interfacesMibCounters.InputQueueDrops
    leafs["runt-packets-received"] = interfacesMibCounters.RuntPacketsReceived
    leafs["giant-packets-received"] = interfacesMibCounters.GiantPacketsReceived
    leafs["throttled-packets-received"] = interfacesMibCounters.ThrottledPacketsReceived
    leafs["parity-packets-received"] = interfacesMibCounters.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = interfacesMibCounters.UnknownProtocolPacketsReceived
    leafs["input-errors"] = interfacesMibCounters.InputErrors
    leafs["crc-errors"] = interfacesMibCounters.CrcErrors
    leafs["input-overruns"] = interfacesMibCounters.InputOverruns
    leafs["framing-errors-received"] = interfacesMibCounters.FramingErrorsReceived
    leafs["input-ignored-packets"] = interfacesMibCounters.InputIgnoredPackets
    leafs["input-aborts"] = interfacesMibCounters.InputAborts
    leafs["output-errors"] = interfacesMibCounters.OutputErrors
    leafs["output-underruns"] = interfacesMibCounters.OutputUnderruns
    leafs["output-buffer-failures"] = interfacesMibCounters.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = interfacesMibCounters.OutputBuffersSwappedOut
    leafs["applique"] = interfacesMibCounters.Applique
    leafs["resets"] = interfacesMibCounters.Resets
    leafs["carrier-transitions"] = interfacesMibCounters.CarrierTransitions
    leafs["availability-flag"] = interfacesMibCounters.AvailabilityFlag
    leafs["last-data-time"] = interfacesMibCounters.LastDataTime
    leafs["seconds-since-last-clear-counters"] = interfacesMibCounters.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = interfacesMibCounters.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = interfacesMibCounters.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = interfacesMibCounters.SecondsSincePacketSent
    return leafs
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetBundleName() string { return "cisco_ios_xr" }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetYangName() string { return "interfaces-mib-counters" }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) SetParent(parent types.Entity) { interfacesMibCounters.parent = parent }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetParent() types.Entity { return interfacesMibCounters.parent }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Cache_InterfacesMibCounters) GetParentYangName() string { return "cache" }

// InfraStatistics_Interfaces_Interface_Cache_DataRate
// Datarate information
type InfraStatistics_Interfaces_Interface_Cache_DataRate struct {
    parent types.Entity
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

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetFilter() yfilter.YFilter { return dataRate.YFilter }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) SetFilter(yf yfilter.YFilter) { dataRate.YFilter = yf }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetGoName(yname string) string {
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    if yname == "peak-input-data-rate" { return "PeakInputDataRate" }
    if yname == "peak-input-packet-rate" { return "PeakInputPacketRate" }
    if yname == "peak-output-data-rate" { return "PeakOutputDataRate" }
    if yname == "peak-output-packet-rate" { return "PeakOutputPacketRate" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "load-interval" { return "LoadInterval" }
    if yname == "output-load" { return "OutputLoad" }
    if yname == "input-load" { return "InputLoad" }
    if yname == "reliability" { return "Reliability" }
    return ""
}

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetSegmentPath() string {
    return "data-rate"
}

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input-data-rate"] = dataRate.InputDataRate
    leafs["input-packet-rate"] = dataRate.InputPacketRate
    leafs["output-data-rate"] = dataRate.OutputDataRate
    leafs["output-packet-rate"] = dataRate.OutputPacketRate
    leafs["peak-input-data-rate"] = dataRate.PeakInputDataRate
    leafs["peak-input-packet-rate"] = dataRate.PeakInputPacketRate
    leafs["peak-output-data-rate"] = dataRate.PeakOutputDataRate
    leafs["peak-output-packet-rate"] = dataRate.PeakOutputPacketRate
    leafs["bandwidth"] = dataRate.Bandwidth
    leafs["load-interval"] = dataRate.LoadInterval
    leafs["output-load"] = dataRate.OutputLoad
    leafs["input-load"] = dataRate.InputLoad
    leafs["reliability"] = dataRate.Reliability
    return leafs
}

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetBundleName() string { return "cisco_ios_xr" }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetYangName() string { return "data-rate" }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) SetParent(parent types.Entity) { dataRate.parent = parent }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetParent() types.Entity { return dataRate.parent }

func (dataRate *InfraStatistics_Interfaces_Interface_Cache_DataRate) GetParentYangName() string { return "cache" }

// InfraStatistics_Interfaces_Interface_Cache_GenericCounters
// Generic set of interface counters
type InfraStatistics_Interfaces_Interface_Cache_GenericCounters struct {
    parent types.Entity
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

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetFilter() yfilter.YFilter { return genericCounters.YFilter }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) SetFilter(yf yfilter.YFilter) { genericCounters.YFilter = yf }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetSegmentPath() string {
    return "generic-counters"
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = genericCounters.PacketsReceived
    leafs["bytes-received"] = genericCounters.BytesReceived
    leafs["packets-sent"] = genericCounters.PacketsSent
    leafs["bytes-sent"] = genericCounters.BytesSent
    leafs["multicast-packets-received"] = genericCounters.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = genericCounters.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = genericCounters.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = genericCounters.BroadcastPacketsSent
    leafs["output-drops"] = genericCounters.OutputDrops
    leafs["output-queue-drops"] = genericCounters.OutputQueueDrops
    leafs["input-drops"] = genericCounters.InputDrops
    leafs["input-queue-drops"] = genericCounters.InputQueueDrops
    leafs["runt-packets-received"] = genericCounters.RuntPacketsReceived
    leafs["giant-packets-received"] = genericCounters.GiantPacketsReceived
    leafs["throttled-packets-received"] = genericCounters.ThrottledPacketsReceived
    leafs["parity-packets-received"] = genericCounters.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = genericCounters.UnknownProtocolPacketsReceived
    leafs["input-errors"] = genericCounters.InputErrors
    leafs["crc-errors"] = genericCounters.CrcErrors
    leafs["input-overruns"] = genericCounters.InputOverruns
    leafs["framing-errors-received"] = genericCounters.FramingErrorsReceived
    leafs["input-ignored-packets"] = genericCounters.InputIgnoredPackets
    leafs["input-aborts"] = genericCounters.InputAborts
    leafs["output-errors"] = genericCounters.OutputErrors
    leafs["output-underruns"] = genericCounters.OutputUnderruns
    leafs["output-buffer-failures"] = genericCounters.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = genericCounters.OutputBuffersSwappedOut
    leafs["applique"] = genericCounters.Applique
    leafs["resets"] = genericCounters.Resets
    leafs["carrier-transitions"] = genericCounters.CarrierTransitions
    leafs["availability-flag"] = genericCounters.AvailabilityFlag
    leafs["last-data-time"] = genericCounters.LastDataTime
    leafs["seconds-since-last-clear-counters"] = genericCounters.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = genericCounters.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = genericCounters.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = genericCounters.SecondsSincePacketSent
    return leafs
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetYangName() string { return "generic-counters" }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) SetParent(parent types.Entity) { genericCounters.parent = parent }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetParent() types.Entity { return genericCounters.parent }

func (genericCounters *InfraStatistics_Interfaces_Interface_Cache_GenericCounters) GetParentYangName() string { return "cache" }

// InfraStatistics_Interfaces_Interface_Latest
// Latest stats data of interfaces
type InfraStatistics_Interfaces_Interface_Latest struct {
    parent types.Entity
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

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetFilter() yfilter.YFilter { return latest.YFilter }

func (latest *InfraStatistics_Interfaces_Interface_Latest) SetFilter(yf yfilter.YFilter) { latest.YFilter = yf }

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetGoName(yname string) string {
    if yname == "protocols" { return "Protocols" }
    if yname == "interfaces-mib-counters" { return "InterfacesMibCounters" }
    if yname == "data-rate" { return "DataRate" }
    if yname == "generic-counters" { return "GenericCounters" }
    return ""
}

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetSegmentPath() string {
    return "latest"
}

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocols" {
        return &latest.Protocols
    }
    if childYangName == "interfaces-mib-counters" {
        return &latest.InterfacesMibCounters
    }
    if childYangName == "data-rate" {
        return &latest.DataRate
    }
    if childYangName == "generic-counters" {
        return &latest.GenericCounters
    }
    return nil
}

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocols"] = &latest.Protocols
    children["interfaces-mib-counters"] = &latest.InterfacesMibCounters
    children["data-rate"] = &latest.DataRate
    children["generic-counters"] = &latest.GenericCounters
    return children
}

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetBundleName() string { return "cisco_ios_xr" }

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetYangName() string { return "latest" }

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (latest *InfraStatistics_Interfaces_Interface_Latest) SetParent(parent types.Entity) { latest.parent = parent }

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetParent() types.Entity { return latest.parent }

func (latest *InfraStatistics_Interfaces_Interface_Latest) GetParentYangName() string { return "interface" }

// InfraStatistics_Interfaces_Interface_Latest_Protocols
// List of protocols
type InfraStatistics_Interfaces_Interface_Latest_Protocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface counters per protocol. The type is slice of
    // InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol.
    Protocol []InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol
}

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetFilter() yfilter.YFilter { return protocols.YFilter }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) SetFilter(yf yfilter.YFilter) { protocols.YFilter = yf }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetSegmentPath() string {
    return "protocols"
}

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        for _, c := range protocols.Protocol {
            if protocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol{}
        protocols.Protocol = append(protocols.Protocol, child)
        return &protocols.Protocol[len(protocols.Protocol)-1]
    }
    return nil
}

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocols.Protocol {
        children[protocols.Protocol[i].GetSegmentPath()] = &protocols.Protocol[i]
    }
    return children
}

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetBundleName() string { return "cisco_ios_xr" }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetYangName() string { return "protocols" }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) SetParent(parent types.Entity) { protocols.parent = parent }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetParent() types.Entity { return protocols.parent }

func (protocols *InfraStatistics_Interfaces_Interface_Latest_Protocols) GetParentYangName() string { return "latest" }

// InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol
// Interface counters per protocol
type InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol struct {
    parent types.Entity
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
    // with range: 0..4294967295. Units are second.
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

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "protocol" { return "Protocol" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    return ""
}

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetSegmentPath() string {
    return "protocol" + "[protocol-name='" + fmt.Sprintf("%v", protocol.ProtocolName) + "']"
}

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = protocol.ProtocolName
    leafs["bytes-received"] = protocol.BytesReceived
    leafs["packets-received"] = protocol.PacketsReceived
    leafs["bytes-sent"] = protocol.BytesSent
    leafs["packets-sent"] = protocol.PacketsSent
    leafs["protocol"] = protocol.Protocol
    leafs["last-data-time"] = protocol.LastDataTime
    leafs["input-data-rate"] = protocol.InputDataRate
    leafs["input-packet-rate"] = protocol.InputPacketRate
    leafs["output-data-rate"] = protocol.OutputDataRate
    leafs["output-packet-rate"] = protocol.OutputPacketRate
    return leafs
}

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetYangName() string { return "protocol" }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *InfraStatistics_Interfaces_Interface_Latest_Protocols_Protocol) GetParentYangName() string { return "protocols" }

// InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters
// Set of interface counters as displayed by the
// InterfacesMIB
type InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters struct {
    parent types.Entity
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

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetFilter() yfilter.YFilter { return interfacesMibCounters.YFilter }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) SetFilter(yf yfilter.YFilter) { interfacesMibCounters.YFilter = yf }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetSegmentPath() string {
    return "interfaces-mib-counters"
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = interfacesMibCounters.PacketsReceived
    leafs["bytes-received"] = interfacesMibCounters.BytesReceived
    leafs["packets-sent"] = interfacesMibCounters.PacketsSent
    leafs["bytes-sent"] = interfacesMibCounters.BytesSent
    leafs["multicast-packets-received"] = interfacesMibCounters.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = interfacesMibCounters.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = interfacesMibCounters.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = interfacesMibCounters.BroadcastPacketsSent
    leafs["output-drops"] = interfacesMibCounters.OutputDrops
    leafs["output-queue-drops"] = interfacesMibCounters.OutputQueueDrops
    leafs["input-drops"] = interfacesMibCounters.InputDrops
    leafs["input-queue-drops"] = interfacesMibCounters.InputQueueDrops
    leafs["runt-packets-received"] = interfacesMibCounters.RuntPacketsReceived
    leafs["giant-packets-received"] = interfacesMibCounters.GiantPacketsReceived
    leafs["throttled-packets-received"] = interfacesMibCounters.ThrottledPacketsReceived
    leafs["parity-packets-received"] = interfacesMibCounters.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = interfacesMibCounters.UnknownProtocolPacketsReceived
    leafs["input-errors"] = interfacesMibCounters.InputErrors
    leafs["crc-errors"] = interfacesMibCounters.CrcErrors
    leafs["input-overruns"] = interfacesMibCounters.InputOverruns
    leafs["framing-errors-received"] = interfacesMibCounters.FramingErrorsReceived
    leafs["input-ignored-packets"] = interfacesMibCounters.InputIgnoredPackets
    leafs["input-aborts"] = interfacesMibCounters.InputAborts
    leafs["output-errors"] = interfacesMibCounters.OutputErrors
    leafs["output-underruns"] = interfacesMibCounters.OutputUnderruns
    leafs["output-buffer-failures"] = interfacesMibCounters.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = interfacesMibCounters.OutputBuffersSwappedOut
    leafs["applique"] = interfacesMibCounters.Applique
    leafs["resets"] = interfacesMibCounters.Resets
    leafs["carrier-transitions"] = interfacesMibCounters.CarrierTransitions
    leafs["availability-flag"] = interfacesMibCounters.AvailabilityFlag
    leafs["last-data-time"] = interfacesMibCounters.LastDataTime
    leafs["seconds-since-last-clear-counters"] = interfacesMibCounters.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = interfacesMibCounters.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = interfacesMibCounters.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = interfacesMibCounters.SecondsSincePacketSent
    return leafs
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetBundleName() string { return "cisco_ios_xr" }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetYangName() string { return "interfaces-mib-counters" }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) SetParent(parent types.Entity) { interfacesMibCounters.parent = parent }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetParent() types.Entity { return interfacesMibCounters.parent }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Latest_InterfacesMibCounters) GetParentYangName() string { return "latest" }

// InfraStatistics_Interfaces_Interface_Latest_DataRate
// Datarate information
type InfraStatistics_Interfaces_Interface_Latest_DataRate struct {
    parent types.Entity
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

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetFilter() yfilter.YFilter { return dataRate.YFilter }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) SetFilter(yf yfilter.YFilter) { dataRate.YFilter = yf }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetGoName(yname string) string {
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    if yname == "peak-input-data-rate" { return "PeakInputDataRate" }
    if yname == "peak-input-packet-rate" { return "PeakInputPacketRate" }
    if yname == "peak-output-data-rate" { return "PeakOutputDataRate" }
    if yname == "peak-output-packet-rate" { return "PeakOutputPacketRate" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "load-interval" { return "LoadInterval" }
    if yname == "output-load" { return "OutputLoad" }
    if yname == "input-load" { return "InputLoad" }
    if yname == "reliability" { return "Reliability" }
    return ""
}

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetSegmentPath() string {
    return "data-rate"
}

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input-data-rate"] = dataRate.InputDataRate
    leafs["input-packet-rate"] = dataRate.InputPacketRate
    leafs["output-data-rate"] = dataRate.OutputDataRate
    leafs["output-packet-rate"] = dataRate.OutputPacketRate
    leafs["peak-input-data-rate"] = dataRate.PeakInputDataRate
    leafs["peak-input-packet-rate"] = dataRate.PeakInputPacketRate
    leafs["peak-output-data-rate"] = dataRate.PeakOutputDataRate
    leafs["peak-output-packet-rate"] = dataRate.PeakOutputPacketRate
    leafs["bandwidth"] = dataRate.Bandwidth
    leafs["load-interval"] = dataRate.LoadInterval
    leafs["output-load"] = dataRate.OutputLoad
    leafs["input-load"] = dataRate.InputLoad
    leafs["reliability"] = dataRate.Reliability
    return leafs
}

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetBundleName() string { return "cisco_ios_xr" }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetYangName() string { return "data-rate" }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) SetParent(parent types.Entity) { dataRate.parent = parent }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetParent() types.Entity { return dataRate.parent }

func (dataRate *InfraStatistics_Interfaces_Interface_Latest_DataRate) GetParentYangName() string { return "latest" }

// InfraStatistics_Interfaces_Interface_Latest_GenericCounters
// Generic set of interface counters
type InfraStatistics_Interfaces_Interface_Latest_GenericCounters struct {
    parent types.Entity
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

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetFilter() yfilter.YFilter { return genericCounters.YFilter }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) SetFilter(yf yfilter.YFilter) { genericCounters.YFilter = yf }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetSegmentPath() string {
    return "generic-counters"
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = genericCounters.PacketsReceived
    leafs["bytes-received"] = genericCounters.BytesReceived
    leafs["packets-sent"] = genericCounters.PacketsSent
    leafs["bytes-sent"] = genericCounters.BytesSent
    leafs["multicast-packets-received"] = genericCounters.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = genericCounters.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = genericCounters.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = genericCounters.BroadcastPacketsSent
    leafs["output-drops"] = genericCounters.OutputDrops
    leafs["output-queue-drops"] = genericCounters.OutputQueueDrops
    leafs["input-drops"] = genericCounters.InputDrops
    leafs["input-queue-drops"] = genericCounters.InputQueueDrops
    leafs["runt-packets-received"] = genericCounters.RuntPacketsReceived
    leafs["giant-packets-received"] = genericCounters.GiantPacketsReceived
    leafs["throttled-packets-received"] = genericCounters.ThrottledPacketsReceived
    leafs["parity-packets-received"] = genericCounters.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = genericCounters.UnknownProtocolPacketsReceived
    leafs["input-errors"] = genericCounters.InputErrors
    leafs["crc-errors"] = genericCounters.CrcErrors
    leafs["input-overruns"] = genericCounters.InputOverruns
    leafs["framing-errors-received"] = genericCounters.FramingErrorsReceived
    leafs["input-ignored-packets"] = genericCounters.InputIgnoredPackets
    leafs["input-aborts"] = genericCounters.InputAborts
    leafs["output-errors"] = genericCounters.OutputErrors
    leafs["output-underruns"] = genericCounters.OutputUnderruns
    leafs["output-buffer-failures"] = genericCounters.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = genericCounters.OutputBuffersSwappedOut
    leafs["applique"] = genericCounters.Applique
    leafs["resets"] = genericCounters.Resets
    leafs["carrier-transitions"] = genericCounters.CarrierTransitions
    leafs["availability-flag"] = genericCounters.AvailabilityFlag
    leafs["last-data-time"] = genericCounters.LastDataTime
    leafs["seconds-since-last-clear-counters"] = genericCounters.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = genericCounters.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = genericCounters.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = genericCounters.SecondsSincePacketSent
    return leafs
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetYangName() string { return "generic-counters" }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) SetParent(parent types.Entity) { genericCounters.parent = parent }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetParent() types.Entity { return genericCounters.parent }

func (genericCounters *InfraStatistics_Interfaces_Interface_Latest_GenericCounters) GetParentYangName() string { return "latest" }

// InfraStatistics_Interfaces_Interface_Total
// Total stats data of interfaces
type InfraStatistics_Interfaces_Interface_Total struct {
    parent types.Entity
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

func (total *InfraStatistics_Interfaces_Interface_Total) GetFilter() yfilter.YFilter { return total.YFilter }

func (total *InfraStatistics_Interfaces_Interface_Total) SetFilter(yf yfilter.YFilter) { total.YFilter = yf }

func (total *InfraStatistics_Interfaces_Interface_Total) GetGoName(yname string) string {
    if yname == "protocols" { return "Protocols" }
    if yname == "interfaces-mib-counters" { return "InterfacesMibCounters" }
    if yname == "data-rate" { return "DataRate" }
    if yname == "generic-counters" { return "GenericCounters" }
    return ""
}

func (total *InfraStatistics_Interfaces_Interface_Total) GetSegmentPath() string {
    return "total"
}

func (total *InfraStatistics_Interfaces_Interface_Total) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocols" {
        return &total.Protocols
    }
    if childYangName == "interfaces-mib-counters" {
        return &total.InterfacesMibCounters
    }
    if childYangName == "data-rate" {
        return &total.DataRate
    }
    if childYangName == "generic-counters" {
        return &total.GenericCounters
    }
    return nil
}

func (total *InfraStatistics_Interfaces_Interface_Total) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocols"] = &total.Protocols
    children["interfaces-mib-counters"] = &total.InterfacesMibCounters
    children["data-rate"] = &total.DataRate
    children["generic-counters"] = &total.GenericCounters
    return children
}

func (total *InfraStatistics_Interfaces_Interface_Total) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (total *InfraStatistics_Interfaces_Interface_Total) GetBundleName() string { return "cisco_ios_xr" }

func (total *InfraStatistics_Interfaces_Interface_Total) GetYangName() string { return "total" }

func (total *InfraStatistics_Interfaces_Interface_Total) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (total *InfraStatistics_Interfaces_Interface_Total) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (total *InfraStatistics_Interfaces_Interface_Total) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (total *InfraStatistics_Interfaces_Interface_Total) SetParent(parent types.Entity) { total.parent = parent }

func (total *InfraStatistics_Interfaces_Interface_Total) GetParent() types.Entity { return total.parent }

func (total *InfraStatistics_Interfaces_Interface_Total) GetParentYangName() string { return "interface" }

// InfraStatistics_Interfaces_Interface_Total_Protocols
// List of protocols
type InfraStatistics_Interfaces_Interface_Total_Protocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface counters per protocol. The type is slice of
    // InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol.
    Protocol []InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol
}

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetFilter() yfilter.YFilter { return protocols.YFilter }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) SetFilter(yf yfilter.YFilter) { protocols.YFilter = yf }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetSegmentPath() string {
    return "protocols"
}

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        for _, c := range protocols.Protocol {
            if protocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol{}
        protocols.Protocol = append(protocols.Protocol, child)
        return &protocols.Protocol[len(protocols.Protocol)-1]
    }
    return nil
}

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocols.Protocol {
        children[protocols.Protocol[i].GetSegmentPath()] = &protocols.Protocol[i]
    }
    return children
}

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetBundleName() string { return "cisco_ios_xr" }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetYangName() string { return "protocols" }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) SetParent(parent types.Entity) { protocols.parent = parent }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetParent() types.Entity { return protocols.parent }

func (protocols *InfraStatistics_Interfaces_Interface_Total_Protocols) GetParentYangName() string { return "total" }

// InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol
// Interface counters per protocol
type InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol struct {
    parent types.Entity
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
    // with range: 0..4294967295. Units are second.
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

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "protocol" { return "Protocol" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    return ""
}

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetSegmentPath() string {
    return "protocol" + "[protocol-name='" + fmt.Sprintf("%v", protocol.ProtocolName) + "']"
}

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = protocol.ProtocolName
    leafs["bytes-received"] = protocol.BytesReceived
    leafs["packets-received"] = protocol.PacketsReceived
    leafs["bytes-sent"] = protocol.BytesSent
    leafs["packets-sent"] = protocol.PacketsSent
    leafs["protocol"] = protocol.Protocol
    leafs["last-data-time"] = protocol.LastDataTime
    leafs["input-data-rate"] = protocol.InputDataRate
    leafs["input-packet-rate"] = protocol.InputPacketRate
    leafs["output-data-rate"] = protocol.OutputDataRate
    leafs["output-packet-rate"] = protocol.OutputPacketRate
    return leafs
}

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetYangName() string { return "protocol" }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *InfraStatistics_Interfaces_Interface_Total_Protocols_Protocol) GetParentYangName() string { return "protocols" }

// InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters
// Set of interface counters as displayed by the
// InterfacesMIB
type InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters struct {
    parent types.Entity
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

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetFilter() yfilter.YFilter { return interfacesMibCounters.YFilter }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) SetFilter(yf yfilter.YFilter) { interfacesMibCounters.YFilter = yf }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetSegmentPath() string {
    return "interfaces-mib-counters"
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = interfacesMibCounters.PacketsReceived
    leafs["bytes-received"] = interfacesMibCounters.BytesReceived
    leafs["packets-sent"] = interfacesMibCounters.PacketsSent
    leafs["bytes-sent"] = interfacesMibCounters.BytesSent
    leafs["multicast-packets-received"] = interfacesMibCounters.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = interfacesMibCounters.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = interfacesMibCounters.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = interfacesMibCounters.BroadcastPacketsSent
    leafs["output-drops"] = interfacesMibCounters.OutputDrops
    leafs["output-queue-drops"] = interfacesMibCounters.OutputQueueDrops
    leafs["input-drops"] = interfacesMibCounters.InputDrops
    leafs["input-queue-drops"] = interfacesMibCounters.InputQueueDrops
    leafs["runt-packets-received"] = interfacesMibCounters.RuntPacketsReceived
    leafs["giant-packets-received"] = interfacesMibCounters.GiantPacketsReceived
    leafs["throttled-packets-received"] = interfacesMibCounters.ThrottledPacketsReceived
    leafs["parity-packets-received"] = interfacesMibCounters.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = interfacesMibCounters.UnknownProtocolPacketsReceived
    leafs["input-errors"] = interfacesMibCounters.InputErrors
    leafs["crc-errors"] = interfacesMibCounters.CrcErrors
    leafs["input-overruns"] = interfacesMibCounters.InputOverruns
    leafs["framing-errors-received"] = interfacesMibCounters.FramingErrorsReceived
    leafs["input-ignored-packets"] = interfacesMibCounters.InputIgnoredPackets
    leafs["input-aborts"] = interfacesMibCounters.InputAborts
    leafs["output-errors"] = interfacesMibCounters.OutputErrors
    leafs["output-underruns"] = interfacesMibCounters.OutputUnderruns
    leafs["output-buffer-failures"] = interfacesMibCounters.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = interfacesMibCounters.OutputBuffersSwappedOut
    leafs["applique"] = interfacesMibCounters.Applique
    leafs["resets"] = interfacesMibCounters.Resets
    leafs["carrier-transitions"] = interfacesMibCounters.CarrierTransitions
    leafs["availability-flag"] = interfacesMibCounters.AvailabilityFlag
    leafs["last-data-time"] = interfacesMibCounters.LastDataTime
    leafs["seconds-since-last-clear-counters"] = interfacesMibCounters.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = interfacesMibCounters.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = interfacesMibCounters.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = interfacesMibCounters.SecondsSincePacketSent
    return leafs
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetBundleName() string { return "cisco_ios_xr" }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetYangName() string { return "interfaces-mib-counters" }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) SetParent(parent types.Entity) { interfacesMibCounters.parent = parent }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetParent() types.Entity { return interfacesMibCounters.parent }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_Total_InterfacesMibCounters) GetParentYangName() string { return "total" }

// InfraStatistics_Interfaces_Interface_Total_DataRate
// Datarate information
type InfraStatistics_Interfaces_Interface_Total_DataRate struct {
    parent types.Entity
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

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetFilter() yfilter.YFilter { return dataRate.YFilter }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) SetFilter(yf yfilter.YFilter) { dataRate.YFilter = yf }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetGoName(yname string) string {
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    if yname == "peak-input-data-rate" { return "PeakInputDataRate" }
    if yname == "peak-input-packet-rate" { return "PeakInputPacketRate" }
    if yname == "peak-output-data-rate" { return "PeakOutputDataRate" }
    if yname == "peak-output-packet-rate" { return "PeakOutputPacketRate" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "load-interval" { return "LoadInterval" }
    if yname == "output-load" { return "OutputLoad" }
    if yname == "input-load" { return "InputLoad" }
    if yname == "reliability" { return "Reliability" }
    return ""
}

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetSegmentPath() string {
    return "data-rate"
}

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input-data-rate"] = dataRate.InputDataRate
    leafs["input-packet-rate"] = dataRate.InputPacketRate
    leafs["output-data-rate"] = dataRate.OutputDataRate
    leafs["output-packet-rate"] = dataRate.OutputPacketRate
    leafs["peak-input-data-rate"] = dataRate.PeakInputDataRate
    leafs["peak-input-packet-rate"] = dataRate.PeakInputPacketRate
    leafs["peak-output-data-rate"] = dataRate.PeakOutputDataRate
    leafs["peak-output-packet-rate"] = dataRate.PeakOutputPacketRate
    leafs["bandwidth"] = dataRate.Bandwidth
    leafs["load-interval"] = dataRate.LoadInterval
    leafs["output-load"] = dataRate.OutputLoad
    leafs["input-load"] = dataRate.InputLoad
    leafs["reliability"] = dataRate.Reliability
    return leafs
}

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetBundleName() string { return "cisco_ios_xr" }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetYangName() string { return "data-rate" }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) SetParent(parent types.Entity) { dataRate.parent = parent }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetParent() types.Entity { return dataRate.parent }

func (dataRate *InfraStatistics_Interfaces_Interface_Total_DataRate) GetParentYangName() string { return "total" }

// InfraStatistics_Interfaces_Interface_Total_GenericCounters
// Generic set of interface counters
type InfraStatistics_Interfaces_Interface_Total_GenericCounters struct {
    parent types.Entity
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

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetFilter() yfilter.YFilter { return genericCounters.YFilter }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) SetFilter(yf yfilter.YFilter) { genericCounters.YFilter = yf }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetSegmentPath() string {
    return "generic-counters"
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = genericCounters.PacketsReceived
    leafs["bytes-received"] = genericCounters.BytesReceived
    leafs["packets-sent"] = genericCounters.PacketsSent
    leafs["bytes-sent"] = genericCounters.BytesSent
    leafs["multicast-packets-received"] = genericCounters.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = genericCounters.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = genericCounters.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = genericCounters.BroadcastPacketsSent
    leafs["output-drops"] = genericCounters.OutputDrops
    leafs["output-queue-drops"] = genericCounters.OutputQueueDrops
    leafs["input-drops"] = genericCounters.InputDrops
    leafs["input-queue-drops"] = genericCounters.InputQueueDrops
    leafs["runt-packets-received"] = genericCounters.RuntPacketsReceived
    leafs["giant-packets-received"] = genericCounters.GiantPacketsReceived
    leafs["throttled-packets-received"] = genericCounters.ThrottledPacketsReceived
    leafs["parity-packets-received"] = genericCounters.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = genericCounters.UnknownProtocolPacketsReceived
    leafs["input-errors"] = genericCounters.InputErrors
    leafs["crc-errors"] = genericCounters.CrcErrors
    leafs["input-overruns"] = genericCounters.InputOverruns
    leafs["framing-errors-received"] = genericCounters.FramingErrorsReceived
    leafs["input-ignored-packets"] = genericCounters.InputIgnoredPackets
    leafs["input-aborts"] = genericCounters.InputAborts
    leafs["output-errors"] = genericCounters.OutputErrors
    leafs["output-underruns"] = genericCounters.OutputUnderruns
    leafs["output-buffer-failures"] = genericCounters.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = genericCounters.OutputBuffersSwappedOut
    leafs["applique"] = genericCounters.Applique
    leafs["resets"] = genericCounters.Resets
    leafs["carrier-transitions"] = genericCounters.CarrierTransitions
    leafs["availability-flag"] = genericCounters.AvailabilityFlag
    leafs["last-data-time"] = genericCounters.LastDataTime
    leafs["seconds-since-last-clear-counters"] = genericCounters.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = genericCounters.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = genericCounters.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = genericCounters.SecondsSincePacketSent
    return leafs
}

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetYangName() string { return "generic-counters" }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) SetParent(parent types.Entity) { genericCounters.parent = parent }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetParent() types.Entity { return genericCounters.parent }

func (genericCounters *InfraStatistics_Interfaces_Interface_Total_GenericCounters) GetParentYangName() string { return "total" }

// InfraStatistics_Interfaces_Interface_Protocols
// List of protocols
type InfraStatistics_Interfaces_Interface_Protocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface counters per protocol. The type is slice of
    // InfraStatistics_Interfaces_Interface_Protocols_Protocol.
    Protocol []InfraStatistics_Interfaces_Interface_Protocols_Protocol
}

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetFilter() yfilter.YFilter { return protocols.YFilter }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) SetFilter(yf yfilter.YFilter) { protocols.YFilter = yf }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetSegmentPath() string {
    return "protocols"
}

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        for _, c := range protocols.Protocol {
            if protocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InfraStatistics_Interfaces_Interface_Protocols_Protocol{}
        protocols.Protocol = append(protocols.Protocol, child)
        return &protocols.Protocol[len(protocols.Protocol)-1]
    }
    return nil
}

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocols.Protocol {
        children[protocols.Protocol[i].GetSegmentPath()] = &protocols.Protocol[i]
    }
    return children
}

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetBundleName() string { return "cisco_ios_xr" }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetYangName() string { return "protocols" }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) SetParent(parent types.Entity) { protocols.parent = parent }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetParent() types.Entity { return protocols.parent }

func (protocols *InfraStatistics_Interfaces_Interface_Protocols) GetParentYangName() string { return "interface" }

// InfraStatistics_Interfaces_Interface_Protocols_Protocol
// Interface counters per protocol
type InfraStatistics_Interfaces_Interface_Protocols_Protocol struct {
    parent types.Entity
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
    // with range: 0..4294967295. Units are second.
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

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "protocol" { return "Protocol" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    return ""
}

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetSegmentPath() string {
    return "protocol" + "[protocol-name='" + fmt.Sprintf("%v", protocol.ProtocolName) + "']"
}

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = protocol.ProtocolName
    leafs["bytes-received"] = protocol.BytesReceived
    leafs["packets-received"] = protocol.PacketsReceived
    leafs["bytes-sent"] = protocol.BytesSent
    leafs["packets-sent"] = protocol.PacketsSent
    leafs["protocol"] = protocol.Protocol
    leafs["last-data-time"] = protocol.LastDataTime
    leafs["input-data-rate"] = protocol.InputDataRate
    leafs["input-packet-rate"] = protocol.InputPacketRate
    leafs["output-data-rate"] = protocol.OutputDataRate
    leafs["output-packet-rate"] = protocol.OutputPacketRate
    return leafs
}

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetYangName() string { return "protocol" }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *InfraStatistics_Interfaces_Interface_Protocols_Protocol) GetParentYangName() string { return "protocols" }

// InfraStatistics_Interfaces_Interface_InterfacesMibCounters
// Set of interface counters as displayed by the
// InterfacesMIB
type InfraStatistics_Interfaces_Interface_InterfacesMibCounters struct {
    parent types.Entity
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

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetFilter() yfilter.YFilter { return interfacesMibCounters.YFilter }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) SetFilter(yf yfilter.YFilter) { interfacesMibCounters.YFilter = yf }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetSegmentPath() string {
    return "interfaces-mib-counters"
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = interfacesMibCounters.PacketsReceived
    leafs["bytes-received"] = interfacesMibCounters.BytesReceived
    leafs["packets-sent"] = interfacesMibCounters.PacketsSent
    leafs["bytes-sent"] = interfacesMibCounters.BytesSent
    leafs["multicast-packets-received"] = interfacesMibCounters.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = interfacesMibCounters.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = interfacesMibCounters.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = interfacesMibCounters.BroadcastPacketsSent
    leafs["output-drops"] = interfacesMibCounters.OutputDrops
    leafs["output-queue-drops"] = interfacesMibCounters.OutputQueueDrops
    leafs["input-drops"] = interfacesMibCounters.InputDrops
    leafs["input-queue-drops"] = interfacesMibCounters.InputQueueDrops
    leafs["runt-packets-received"] = interfacesMibCounters.RuntPacketsReceived
    leafs["giant-packets-received"] = interfacesMibCounters.GiantPacketsReceived
    leafs["throttled-packets-received"] = interfacesMibCounters.ThrottledPacketsReceived
    leafs["parity-packets-received"] = interfacesMibCounters.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = interfacesMibCounters.UnknownProtocolPacketsReceived
    leafs["input-errors"] = interfacesMibCounters.InputErrors
    leafs["crc-errors"] = interfacesMibCounters.CrcErrors
    leafs["input-overruns"] = interfacesMibCounters.InputOverruns
    leafs["framing-errors-received"] = interfacesMibCounters.FramingErrorsReceived
    leafs["input-ignored-packets"] = interfacesMibCounters.InputIgnoredPackets
    leafs["input-aborts"] = interfacesMibCounters.InputAborts
    leafs["output-errors"] = interfacesMibCounters.OutputErrors
    leafs["output-underruns"] = interfacesMibCounters.OutputUnderruns
    leafs["output-buffer-failures"] = interfacesMibCounters.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = interfacesMibCounters.OutputBuffersSwappedOut
    leafs["applique"] = interfacesMibCounters.Applique
    leafs["resets"] = interfacesMibCounters.Resets
    leafs["carrier-transitions"] = interfacesMibCounters.CarrierTransitions
    leafs["availability-flag"] = interfacesMibCounters.AvailabilityFlag
    leafs["last-data-time"] = interfacesMibCounters.LastDataTime
    leafs["seconds-since-last-clear-counters"] = interfacesMibCounters.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = interfacesMibCounters.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = interfacesMibCounters.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = interfacesMibCounters.SecondsSincePacketSent
    return leafs
}

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetBundleName() string { return "cisco_ios_xr" }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetYangName() string { return "interfaces-mib-counters" }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) SetParent(parent types.Entity) { interfacesMibCounters.parent = parent }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetParent() types.Entity { return interfacesMibCounters.parent }

func (interfacesMibCounters *InfraStatistics_Interfaces_Interface_InterfacesMibCounters) GetParentYangName() string { return "interface" }

// InfraStatistics_Interfaces_Interface_DataRate
// Datarate information
type InfraStatistics_Interfaces_Interface_DataRate struct {
    parent types.Entity
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

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetFilter() yfilter.YFilter { return dataRate.YFilter }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) SetFilter(yf yfilter.YFilter) { dataRate.YFilter = yf }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetGoName(yname string) string {
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    if yname == "peak-input-data-rate" { return "PeakInputDataRate" }
    if yname == "peak-input-packet-rate" { return "PeakInputPacketRate" }
    if yname == "peak-output-data-rate" { return "PeakOutputDataRate" }
    if yname == "peak-output-packet-rate" { return "PeakOutputPacketRate" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "load-interval" { return "LoadInterval" }
    if yname == "output-load" { return "OutputLoad" }
    if yname == "input-load" { return "InputLoad" }
    if yname == "reliability" { return "Reliability" }
    return ""
}

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetSegmentPath() string {
    return "data-rate"
}

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input-data-rate"] = dataRate.InputDataRate
    leafs["input-packet-rate"] = dataRate.InputPacketRate
    leafs["output-data-rate"] = dataRate.OutputDataRate
    leafs["output-packet-rate"] = dataRate.OutputPacketRate
    leafs["peak-input-data-rate"] = dataRate.PeakInputDataRate
    leafs["peak-input-packet-rate"] = dataRate.PeakInputPacketRate
    leafs["peak-output-data-rate"] = dataRate.PeakOutputDataRate
    leafs["peak-output-packet-rate"] = dataRate.PeakOutputPacketRate
    leafs["bandwidth"] = dataRate.Bandwidth
    leafs["load-interval"] = dataRate.LoadInterval
    leafs["output-load"] = dataRate.OutputLoad
    leafs["input-load"] = dataRate.InputLoad
    leafs["reliability"] = dataRate.Reliability
    return leafs
}

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetBundleName() string { return "cisco_ios_xr" }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetYangName() string { return "data-rate" }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) SetParent(parent types.Entity) { dataRate.parent = parent }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetParent() types.Entity { return dataRate.parent }

func (dataRate *InfraStatistics_Interfaces_Interface_DataRate) GetParentYangName() string { return "interface" }

// InfraStatistics_Interfaces_Interface_GenericCounters
// Generic set of interface counters
type InfraStatistics_Interfaces_Interface_GenericCounters struct {
    parent types.Entity
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

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetFilter() yfilter.YFilter { return genericCounters.YFilter }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) SetFilter(yf yfilter.YFilter) { genericCounters.YFilter = yf }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetSegmentPath() string {
    return "generic-counters"
}

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = genericCounters.PacketsReceived
    leafs["bytes-received"] = genericCounters.BytesReceived
    leafs["packets-sent"] = genericCounters.PacketsSent
    leafs["bytes-sent"] = genericCounters.BytesSent
    leafs["multicast-packets-received"] = genericCounters.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = genericCounters.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = genericCounters.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = genericCounters.BroadcastPacketsSent
    leafs["output-drops"] = genericCounters.OutputDrops
    leafs["output-queue-drops"] = genericCounters.OutputQueueDrops
    leafs["input-drops"] = genericCounters.InputDrops
    leafs["input-queue-drops"] = genericCounters.InputQueueDrops
    leafs["runt-packets-received"] = genericCounters.RuntPacketsReceived
    leafs["giant-packets-received"] = genericCounters.GiantPacketsReceived
    leafs["throttled-packets-received"] = genericCounters.ThrottledPacketsReceived
    leafs["parity-packets-received"] = genericCounters.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = genericCounters.UnknownProtocolPacketsReceived
    leafs["input-errors"] = genericCounters.InputErrors
    leafs["crc-errors"] = genericCounters.CrcErrors
    leafs["input-overruns"] = genericCounters.InputOverruns
    leafs["framing-errors-received"] = genericCounters.FramingErrorsReceived
    leafs["input-ignored-packets"] = genericCounters.InputIgnoredPackets
    leafs["input-aborts"] = genericCounters.InputAborts
    leafs["output-errors"] = genericCounters.OutputErrors
    leafs["output-underruns"] = genericCounters.OutputUnderruns
    leafs["output-buffer-failures"] = genericCounters.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = genericCounters.OutputBuffersSwappedOut
    leafs["applique"] = genericCounters.Applique
    leafs["resets"] = genericCounters.Resets
    leafs["carrier-transitions"] = genericCounters.CarrierTransitions
    leafs["availability-flag"] = genericCounters.AvailabilityFlag
    leafs["last-data-time"] = genericCounters.LastDataTime
    leafs["seconds-since-last-clear-counters"] = genericCounters.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = genericCounters.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = genericCounters.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = genericCounters.SecondsSincePacketSent
    return leafs
}

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetYangName() string { return "generic-counters" }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) SetParent(parent types.Entity) { genericCounters.parent = parent }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetParent() types.Entity { return genericCounters.parent }

func (genericCounters *InfraStatistics_Interfaces_Interface_GenericCounters) GetParentYangName() string { return "interface" }

