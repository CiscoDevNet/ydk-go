// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package snmp_test_trap_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_test_trap_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act snmp-cold-start}", reflect.TypeOf(SnmpColdStart{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:snmp-cold-start", reflect.TypeOf(SnmpColdStart{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act snmp-warm-start}", reflect.TypeOf(SnmpWarmStart{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:snmp-warm-start", reflect.TypeOf(SnmpWarmStart{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act interface-link-up}", reflect.TypeOf(InterfaceLinkUp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:interface-link-up", reflect.TypeOf(InterfaceLinkUp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act interface-link-down}", reflect.TypeOf(InterfaceLinkDown{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:interface-link-down", reflect.TypeOf(InterfaceLinkDown{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act sonet-section-status}", reflect.TypeOf(SonetSectionStatus{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:sonet-section-status", reflect.TypeOf(SonetSectionStatus{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act sonet-line-status}", reflect.TypeOf(SonetLineStatus{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:sonet-line-status", reflect.TypeOf(SonetLineStatus{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act sonet-path-status}", reflect.TypeOf(SonetPathStatus{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:sonet-path-status", reflect.TypeOf(SonetPathStatus{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act infra-syslog-message-generated}", reflect.TypeOf(InfraSyslogMessageGenerated{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:infra-syslog-message-generated", reflect.TypeOf(InfraSyslogMessageGenerated{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act infra-flash-device-inserted}", reflect.TypeOf(InfraFlashDeviceInserted{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:infra-flash-device-inserted", reflect.TypeOf(InfraFlashDeviceInserted{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act infra-flash-device-removed}", reflect.TypeOf(InfraFlashDeviceRemoved{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:infra-flash-device-removed", reflect.TypeOf(InfraFlashDeviceRemoved{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act infra-redundancy-progression}", reflect.TypeOf(InfraRedundancyProgression{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:infra-redundancy-progression", reflect.TypeOf(InfraRedundancyProgression{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act infra-redundancy-switch}", reflect.TypeOf(InfraRedundancySwitch{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:infra-redundancy-switch", reflect.TypeOf(InfraRedundancySwitch{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act infra-bridge-new-root}", reflect.TypeOf(InfraBridgeNewRoot{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:infra-bridge-new-root", reflect.TypeOf(InfraBridgeNewRoot{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act infra-bridge-topology-change}", reflect.TypeOf(InfraBridgeTopologyChange{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:infra-bridge-topology-change", reflect.TypeOf(InfraBridgeTopologyChange{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act infra-config-event}", reflect.TypeOf(InfraConfigEvent{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:infra-config-event", reflect.TypeOf(InfraConfigEvent{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act entity-sensor-threshold-notification}", reflect.TypeOf(EntitySensorThresholdNotification{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:entity-sensor-threshold-notification", reflect.TypeOf(EntitySensorThresholdNotification{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act entity-fru-power-status-change-failed}", reflect.TypeOf(EntityFruPowerStatusChangeFailed{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:entity-fru-power-status-change-failed", reflect.TypeOf(EntityFruPowerStatusChangeFailed{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act entity-fru-module-status-change-up}", reflect.TypeOf(EntityFruModuleStatusChangeUp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:entity-fru-module-status-change-up", reflect.TypeOf(EntityFruModuleStatusChangeUp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act entity-fru-module-status-change-down}", reflect.TypeOf(EntityFruModuleStatusChangeDown{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:entity-fru-module-status-change-down", reflect.TypeOf(EntityFruModuleStatusChangeDown{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act entity-fru-fan-tray-oper-status-up}", reflect.TypeOf(EntityFruFanTrayOperStatusUp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-oper-status-up", reflect.TypeOf(EntityFruFanTrayOperStatusUp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act entity-fru-fan-tray-inserted}", reflect.TypeOf(EntityFruFanTrayInserted{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-inserted", reflect.TypeOf(EntityFruFanTrayInserted{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act entity-fru-fan-tray-removed}", reflect.TypeOf(EntityFruFanTrayRemoved{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-removed", reflect.TypeOf(EntityFruFanTrayRemoved{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act platform-hfr-bundle-downed-link}", reflect.TypeOf(PlatformHfrBundleDownedLink{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-bundle-downed-link", reflect.TypeOf(PlatformHfrBundleDownedLink{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act platform-hfr-bundle-state}", reflect.TypeOf(PlatformHfrBundleState{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-bundle-state", reflect.TypeOf(PlatformHfrBundleState{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act platform-hfr-plane-state}", reflect.TypeOf(PlatformHfrPlaneState{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-plane-state", reflect.TypeOf(PlatformHfrPlaneState{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-bgp-established}", reflect.TypeOf(RoutingBgpEstablished{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-established", reflect.TypeOf(RoutingBgpEstablished{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-bgp-established-remote-peer}", reflect.TypeOf(RoutingBgpEstablishedRemotePeer{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-established-remote-peer", reflect.TypeOf(RoutingBgpEstablishedRemotePeer{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-bgp-state-change}", reflect.TypeOf(RoutingBgpStateChange{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-state-change", reflect.TypeOf(RoutingBgpStateChange{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-bgp-state-change-remote-peer}", reflect.TypeOf(RoutingBgpStateChangeRemotePeer{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-state-change-remote-peer", reflect.TypeOf(RoutingBgpStateChangeRemotePeer{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-ospf-neighbor-state-change}", reflect.TypeOf(RoutingOspfNeighborStateChange{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-ospf-neighbor-state-change", reflect.TypeOf(RoutingOspfNeighborStateChange{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-ospf-neighbor-state-change-address}", reflect.TypeOf(RoutingOspfNeighborStateChangeAddress{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-ospf-neighbor-state-change-address", reflect.TypeOf(RoutingOspfNeighborStateChangeAddress{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-mpls-ldp-session-down}", reflect.TypeOf(RoutingMplsLdpSessionDown{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-ldp-session-down", reflect.TypeOf(RoutingMplsLdpSessionDown{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-mpls-ldp-session-down-entity-id}", reflect.TypeOf(RoutingMplsLdpSessionDownEntityId{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-ldp-session-down-entity-id", reflect.TypeOf(RoutingMplsLdpSessionDownEntityId{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-mpls-tunnel-re-routed}", reflect.TypeOf(RoutingMplsTunnelReRouted{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-routed", reflect.TypeOf(RoutingMplsTunnelReRouted{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-mpls-tunnel-re-routed-index}", reflect.TypeOf(RoutingMplsTunnelReRoutedIndex{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-routed-index", reflect.TypeOf(RoutingMplsTunnelReRoutedIndex{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-mpls-tunnel-re-optimized}", reflect.TypeOf(RoutingMplsTunnelReOptimized{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-optimized", reflect.TypeOf(RoutingMplsTunnelReOptimized{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-mpls-tunnel-re-optimized-index}", reflect.TypeOf(RoutingMplsTunnelReOptimizedIndex{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-optimized-index", reflect.TypeOf(RoutingMplsTunnelReOptimizedIndex{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-mpls-tunnel-down}", reflect.TypeOf(RoutingMplsTunnelDown{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-down", reflect.TypeOf(RoutingMplsTunnelDown{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act routing-mpls-tunnel-down-index}", reflect.TypeOf(RoutingMplsTunnelDownIndex{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-down-index", reflect.TypeOf(RoutingMplsTunnelDownIndex{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-test-trap-act all}", reflect.TypeOf(All{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-test-trap-act:all", reflect.TypeOf(All{}))
}

// SnmpColdStart
// Generate SNMPv2-MIB::coldStart
type SnmpColdStart struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (snmpColdStart *SnmpColdStart) GetFilter() yfilter.YFilter { return snmpColdStart.YFilter }

func (snmpColdStart *SnmpColdStart) SetFilter(yf yfilter.YFilter) { snmpColdStart.YFilter = yf }

func (snmpColdStart *SnmpColdStart) GetGoName(yname string) string {
    return ""
}

func (snmpColdStart *SnmpColdStart) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:snmp-cold-start"
}

func (snmpColdStart *SnmpColdStart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmpColdStart *SnmpColdStart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmpColdStart *SnmpColdStart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snmpColdStart *SnmpColdStart) GetBundleName() string { return "cisco_ios_xr" }

func (snmpColdStart *SnmpColdStart) GetYangName() string { return "snmp-cold-start" }

func (snmpColdStart *SnmpColdStart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snmpColdStart *SnmpColdStart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snmpColdStart *SnmpColdStart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snmpColdStart *SnmpColdStart) SetParent(parent types.Entity) { snmpColdStart.parent = parent }

func (snmpColdStart *SnmpColdStart) GetParent() types.Entity { return snmpColdStart.parent }

func (snmpColdStart *SnmpColdStart) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// SnmpWarmStart
// Generate SNMPv2-MIB::warmStart
type SnmpWarmStart struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (snmpWarmStart *SnmpWarmStart) GetFilter() yfilter.YFilter { return snmpWarmStart.YFilter }

func (snmpWarmStart *SnmpWarmStart) SetFilter(yf yfilter.YFilter) { snmpWarmStart.YFilter = yf }

func (snmpWarmStart *SnmpWarmStart) GetGoName(yname string) string {
    return ""
}

func (snmpWarmStart *SnmpWarmStart) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:snmp-warm-start"
}

func (snmpWarmStart *SnmpWarmStart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmpWarmStart *SnmpWarmStart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmpWarmStart *SnmpWarmStart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snmpWarmStart *SnmpWarmStart) GetBundleName() string { return "cisco_ios_xr" }

func (snmpWarmStart *SnmpWarmStart) GetYangName() string { return "snmp-warm-start" }

func (snmpWarmStart *SnmpWarmStart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snmpWarmStart *SnmpWarmStart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snmpWarmStart *SnmpWarmStart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snmpWarmStart *SnmpWarmStart) SetParent(parent types.Entity) { snmpWarmStart.parent = parent }

func (snmpWarmStart *SnmpWarmStart) GetParent() types.Entity { return snmpWarmStart.parent }

func (snmpWarmStart *SnmpWarmStart) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InterfaceLinkUp
// Generate IF-MIB::linkUp
type InterfaceLinkUp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input InterfaceLinkUp_Input
}

func (interfaceLinkUp *InterfaceLinkUp) GetFilter() yfilter.YFilter { return interfaceLinkUp.YFilter }

func (interfaceLinkUp *InterfaceLinkUp) SetFilter(yf yfilter.YFilter) { interfaceLinkUp.YFilter = yf }

func (interfaceLinkUp *InterfaceLinkUp) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (interfaceLinkUp *InterfaceLinkUp) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:interface-link-up"
}

func (interfaceLinkUp *InterfaceLinkUp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &interfaceLinkUp.Input
    }
    return nil
}

func (interfaceLinkUp *InterfaceLinkUp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &interfaceLinkUp.Input
    return children
}

func (interfaceLinkUp *InterfaceLinkUp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceLinkUp *InterfaceLinkUp) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceLinkUp *InterfaceLinkUp) GetYangName() string { return "interface-link-up" }

func (interfaceLinkUp *InterfaceLinkUp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceLinkUp *InterfaceLinkUp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceLinkUp *InterfaceLinkUp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceLinkUp *InterfaceLinkUp) SetParent(parent types.Entity) { interfaceLinkUp.parent = parent }

func (interfaceLinkUp *InterfaceLinkUp) GetParent() types.Entity { return interfaceLinkUp.parent }

func (interfaceLinkUp *InterfaceLinkUp) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InterfaceLinkUp_Input
type InterfaceLinkUp_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface index for which to generate LinkUp trap. The type is interface{}
    // with range: 1..2147483647.
    Ifindex interface{}
}

func (input *InterfaceLinkUp_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *InterfaceLinkUp_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *InterfaceLinkUp_Input) GetGoName(yname string) string {
    if yname == "ifindex" { return "Ifindex" }
    return ""
}

func (input *InterfaceLinkUp_Input) GetSegmentPath() string {
    return "input"
}

func (input *InterfaceLinkUp_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *InterfaceLinkUp_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *InterfaceLinkUp_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifindex"] = input.Ifindex
    return leafs
}

func (input *InterfaceLinkUp_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *InterfaceLinkUp_Input) GetYangName() string { return "input" }

func (input *InterfaceLinkUp_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *InterfaceLinkUp_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *InterfaceLinkUp_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *InterfaceLinkUp_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *InterfaceLinkUp_Input) GetParent() types.Entity { return input.parent }

func (input *InterfaceLinkUp_Input) GetParentYangName() string { return "interface-link-up" }

// InterfaceLinkDown
// Generate IF-MIB::linkDown
type InterfaceLinkDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input InterfaceLinkDown_Input
}

func (interfaceLinkDown *InterfaceLinkDown) GetFilter() yfilter.YFilter { return interfaceLinkDown.YFilter }

func (interfaceLinkDown *InterfaceLinkDown) SetFilter(yf yfilter.YFilter) { interfaceLinkDown.YFilter = yf }

func (interfaceLinkDown *InterfaceLinkDown) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (interfaceLinkDown *InterfaceLinkDown) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:interface-link-down"
}

func (interfaceLinkDown *InterfaceLinkDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &interfaceLinkDown.Input
    }
    return nil
}

func (interfaceLinkDown *InterfaceLinkDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &interfaceLinkDown.Input
    return children
}

func (interfaceLinkDown *InterfaceLinkDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceLinkDown *InterfaceLinkDown) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceLinkDown *InterfaceLinkDown) GetYangName() string { return "interface-link-down" }

func (interfaceLinkDown *InterfaceLinkDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceLinkDown *InterfaceLinkDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceLinkDown *InterfaceLinkDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceLinkDown *InterfaceLinkDown) SetParent(parent types.Entity) { interfaceLinkDown.parent = parent }

func (interfaceLinkDown *InterfaceLinkDown) GetParent() types.Entity { return interfaceLinkDown.parent }

func (interfaceLinkDown *InterfaceLinkDown) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InterfaceLinkDown_Input
type InterfaceLinkDown_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface index for which to generate LinkDown trap. The type is
    // interface{} with range: 1..2147483647.
    Ifindex interface{}
}

func (input *InterfaceLinkDown_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *InterfaceLinkDown_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *InterfaceLinkDown_Input) GetGoName(yname string) string {
    if yname == "ifindex" { return "Ifindex" }
    return ""
}

func (input *InterfaceLinkDown_Input) GetSegmentPath() string {
    return "input"
}

func (input *InterfaceLinkDown_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *InterfaceLinkDown_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *InterfaceLinkDown_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifindex"] = input.Ifindex
    return leafs
}

func (input *InterfaceLinkDown_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *InterfaceLinkDown_Input) GetYangName() string { return "input" }

func (input *InterfaceLinkDown_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *InterfaceLinkDown_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *InterfaceLinkDown_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *InterfaceLinkDown_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *InterfaceLinkDown_Input) GetParent() types.Entity { return input.parent }

func (input *InterfaceLinkDown_Input) GetParentYangName() string { return "interface-link-down" }

// SonetSectionStatus
// Generate CISCO-SONET-MIB::ciscoSonetSectionStatusChange
type SonetSectionStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input SonetSectionStatus_Input
}

func (sonetSectionStatus *SonetSectionStatus) GetFilter() yfilter.YFilter { return sonetSectionStatus.YFilter }

func (sonetSectionStatus *SonetSectionStatus) SetFilter(yf yfilter.YFilter) { sonetSectionStatus.YFilter = yf }

func (sonetSectionStatus *SonetSectionStatus) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (sonetSectionStatus *SonetSectionStatus) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:sonet-section-status"
}

func (sonetSectionStatus *SonetSectionStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &sonetSectionStatus.Input
    }
    return nil
}

func (sonetSectionStatus *SonetSectionStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &sonetSectionStatus.Input
    return children
}

func (sonetSectionStatus *SonetSectionStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetSectionStatus *SonetSectionStatus) GetBundleName() string { return "cisco_ios_xr" }

func (sonetSectionStatus *SonetSectionStatus) GetYangName() string { return "sonet-section-status" }

func (sonetSectionStatus *SonetSectionStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sonetSectionStatus *SonetSectionStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sonetSectionStatus *SonetSectionStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sonetSectionStatus *SonetSectionStatus) SetParent(parent types.Entity) { sonetSectionStatus.parent = parent }

func (sonetSectionStatus *SonetSectionStatus) GetParent() types.Entity { return sonetSectionStatus.parent }

func (sonetSectionStatus *SonetSectionStatus) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// SonetSectionStatus_Input
type SonetSectionStatus_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface index for which to generate ciscoSonetSectionStatusChange trap.
    // The type is interface{} with range: 1..2147483647.
    Ifindex interface{}
}

func (input *SonetSectionStatus_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *SonetSectionStatus_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *SonetSectionStatus_Input) GetGoName(yname string) string {
    if yname == "ifindex" { return "Ifindex" }
    return ""
}

func (input *SonetSectionStatus_Input) GetSegmentPath() string {
    return "input"
}

func (input *SonetSectionStatus_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *SonetSectionStatus_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *SonetSectionStatus_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifindex"] = input.Ifindex
    return leafs
}

func (input *SonetSectionStatus_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *SonetSectionStatus_Input) GetYangName() string { return "input" }

func (input *SonetSectionStatus_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *SonetSectionStatus_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *SonetSectionStatus_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *SonetSectionStatus_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *SonetSectionStatus_Input) GetParent() types.Entity { return input.parent }

func (input *SonetSectionStatus_Input) GetParentYangName() string { return "sonet-section-status" }

// SonetLineStatus
// Generate CISCO-SONET-MIB::ciscoSonetLineStatusChange
type SonetLineStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input SonetLineStatus_Input
}

func (sonetLineStatus *SonetLineStatus) GetFilter() yfilter.YFilter { return sonetLineStatus.YFilter }

func (sonetLineStatus *SonetLineStatus) SetFilter(yf yfilter.YFilter) { sonetLineStatus.YFilter = yf }

func (sonetLineStatus *SonetLineStatus) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (sonetLineStatus *SonetLineStatus) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:sonet-line-status"
}

func (sonetLineStatus *SonetLineStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &sonetLineStatus.Input
    }
    return nil
}

func (sonetLineStatus *SonetLineStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &sonetLineStatus.Input
    return children
}

func (sonetLineStatus *SonetLineStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetLineStatus *SonetLineStatus) GetBundleName() string { return "cisco_ios_xr" }

func (sonetLineStatus *SonetLineStatus) GetYangName() string { return "sonet-line-status" }

func (sonetLineStatus *SonetLineStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sonetLineStatus *SonetLineStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sonetLineStatus *SonetLineStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sonetLineStatus *SonetLineStatus) SetParent(parent types.Entity) { sonetLineStatus.parent = parent }

func (sonetLineStatus *SonetLineStatus) GetParent() types.Entity { return sonetLineStatus.parent }

func (sonetLineStatus *SonetLineStatus) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// SonetLineStatus_Input
type SonetLineStatus_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface index for which to generate ciscoSonetLineStatusChange trap. The
    // type is interface{} with range: 1..2147483647.
    Ifindex interface{}
}

func (input *SonetLineStatus_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *SonetLineStatus_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *SonetLineStatus_Input) GetGoName(yname string) string {
    if yname == "ifindex" { return "Ifindex" }
    return ""
}

func (input *SonetLineStatus_Input) GetSegmentPath() string {
    return "input"
}

func (input *SonetLineStatus_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *SonetLineStatus_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *SonetLineStatus_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifindex"] = input.Ifindex
    return leafs
}

func (input *SonetLineStatus_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *SonetLineStatus_Input) GetYangName() string { return "input" }

func (input *SonetLineStatus_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *SonetLineStatus_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *SonetLineStatus_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *SonetLineStatus_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *SonetLineStatus_Input) GetParent() types.Entity { return input.parent }

func (input *SonetLineStatus_Input) GetParentYangName() string { return "sonet-line-status" }

// SonetPathStatus
// Generate CISCO-SONET-MIB::ciscoSonetPathStatusChange
type SonetPathStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input SonetPathStatus_Input
}

func (sonetPathStatus *SonetPathStatus) GetFilter() yfilter.YFilter { return sonetPathStatus.YFilter }

func (sonetPathStatus *SonetPathStatus) SetFilter(yf yfilter.YFilter) { sonetPathStatus.YFilter = yf }

func (sonetPathStatus *SonetPathStatus) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (sonetPathStatus *SonetPathStatus) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:sonet-path-status"
}

func (sonetPathStatus *SonetPathStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &sonetPathStatus.Input
    }
    return nil
}

func (sonetPathStatus *SonetPathStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &sonetPathStatus.Input
    return children
}

func (sonetPathStatus *SonetPathStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonetPathStatus *SonetPathStatus) GetBundleName() string { return "cisco_ios_xr" }

func (sonetPathStatus *SonetPathStatus) GetYangName() string { return "sonet-path-status" }

func (sonetPathStatus *SonetPathStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sonetPathStatus *SonetPathStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sonetPathStatus *SonetPathStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sonetPathStatus *SonetPathStatus) SetParent(parent types.Entity) { sonetPathStatus.parent = parent }

func (sonetPathStatus *SonetPathStatus) GetParent() types.Entity { return sonetPathStatus.parent }

func (sonetPathStatus *SonetPathStatus) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// SonetPathStatus_Input
type SonetPathStatus_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface index for which to generate ciscoSonetPathStatusChange trap. The
    // type is interface{} with range: 1..2147483647.
    Ifindex interface{}
}

func (input *SonetPathStatus_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *SonetPathStatus_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *SonetPathStatus_Input) GetGoName(yname string) string {
    if yname == "ifindex" { return "Ifindex" }
    return ""
}

func (input *SonetPathStatus_Input) GetSegmentPath() string {
    return "input"
}

func (input *SonetPathStatus_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *SonetPathStatus_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *SonetPathStatus_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifindex"] = input.Ifindex
    return leafs
}

func (input *SonetPathStatus_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *SonetPathStatus_Input) GetYangName() string { return "input" }

func (input *SonetPathStatus_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *SonetPathStatus_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *SonetPathStatus_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *SonetPathStatus_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *SonetPathStatus_Input) GetParent() types.Entity { return input.parent }

func (input *SonetPathStatus_Input) GetParentYangName() string { return "sonet-path-status" }

// InfraSyslogMessageGenerated
// Generate CISCO-SYSLOG-MIB::clogMessageGenerated
type InfraSyslogMessageGenerated struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetFilter() yfilter.YFilter { return infraSyslogMessageGenerated.YFilter }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) SetFilter(yf yfilter.YFilter) { infraSyslogMessageGenerated.YFilter = yf }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetGoName(yname string) string {
    return ""
}

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:infra-syslog-message-generated"
}

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetBundleName() string { return "cisco_ios_xr" }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetYangName() string { return "infra-syslog-message-generated" }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) SetParent(parent types.Entity) { infraSyslogMessageGenerated.parent = parent }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetParent() types.Entity { return infraSyslogMessageGenerated.parent }

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InfraFlashDeviceInserted
// Generate CISCO-FLASH-MIB::ciscoFlashDeviceInsertedNotif
type InfraFlashDeviceInserted struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetFilter() yfilter.YFilter { return infraFlashDeviceInserted.YFilter }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) SetFilter(yf yfilter.YFilter) { infraFlashDeviceInserted.YFilter = yf }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetGoName(yname string) string {
    return ""
}

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:infra-flash-device-inserted"
}

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetBundleName() string { return "cisco_ios_xr" }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetYangName() string { return "infra-flash-device-inserted" }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) SetParent(parent types.Entity) { infraFlashDeviceInserted.parent = parent }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetParent() types.Entity { return infraFlashDeviceInserted.parent }

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InfraFlashDeviceRemoved
// Generate CISCO-FLASH-MIB::ciscoFlashDeviceRemovedNotif
type InfraFlashDeviceRemoved struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetFilter() yfilter.YFilter { return infraFlashDeviceRemoved.YFilter }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) SetFilter(yf yfilter.YFilter) { infraFlashDeviceRemoved.YFilter = yf }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetGoName(yname string) string {
    return ""
}

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:infra-flash-device-removed"
}

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetBundleName() string { return "cisco_ios_xr" }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetYangName() string { return "infra-flash-device-removed" }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) SetParent(parent types.Entity) { infraFlashDeviceRemoved.parent = parent }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetParent() types.Entity { return infraFlashDeviceRemoved.parent }

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InfraRedundancyProgression
// Generate CISCO-RF-MIB::ciscoRFProgressionNotif
type InfraRedundancyProgression struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (infraRedundancyProgression *InfraRedundancyProgression) GetFilter() yfilter.YFilter { return infraRedundancyProgression.YFilter }

func (infraRedundancyProgression *InfraRedundancyProgression) SetFilter(yf yfilter.YFilter) { infraRedundancyProgression.YFilter = yf }

func (infraRedundancyProgression *InfraRedundancyProgression) GetGoName(yname string) string {
    return ""
}

func (infraRedundancyProgression *InfraRedundancyProgression) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:infra-redundancy-progression"
}

func (infraRedundancyProgression *InfraRedundancyProgression) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (infraRedundancyProgression *InfraRedundancyProgression) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (infraRedundancyProgression *InfraRedundancyProgression) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraRedundancyProgression *InfraRedundancyProgression) GetBundleName() string { return "cisco_ios_xr" }

func (infraRedundancyProgression *InfraRedundancyProgression) GetYangName() string { return "infra-redundancy-progression" }

func (infraRedundancyProgression *InfraRedundancyProgression) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraRedundancyProgression *InfraRedundancyProgression) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraRedundancyProgression *InfraRedundancyProgression) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraRedundancyProgression *InfraRedundancyProgression) SetParent(parent types.Entity) { infraRedundancyProgression.parent = parent }

func (infraRedundancyProgression *InfraRedundancyProgression) GetParent() types.Entity { return infraRedundancyProgression.parent }

func (infraRedundancyProgression *InfraRedundancyProgression) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InfraRedundancySwitch
// Generate CISCO-RF-MIB::ciscoRFSwactNotif
type InfraRedundancySwitch struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (infraRedundancySwitch *InfraRedundancySwitch) GetFilter() yfilter.YFilter { return infraRedundancySwitch.YFilter }

func (infraRedundancySwitch *InfraRedundancySwitch) SetFilter(yf yfilter.YFilter) { infraRedundancySwitch.YFilter = yf }

func (infraRedundancySwitch *InfraRedundancySwitch) GetGoName(yname string) string {
    return ""
}

func (infraRedundancySwitch *InfraRedundancySwitch) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:infra-redundancy-switch"
}

func (infraRedundancySwitch *InfraRedundancySwitch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (infraRedundancySwitch *InfraRedundancySwitch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (infraRedundancySwitch *InfraRedundancySwitch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraRedundancySwitch *InfraRedundancySwitch) GetBundleName() string { return "cisco_ios_xr" }

func (infraRedundancySwitch *InfraRedundancySwitch) GetYangName() string { return "infra-redundancy-switch" }

func (infraRedundancySwitch *InfraRedundancySwitch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraRedundancySwitch *InfraRedundancySwitch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraRedundancySwitch *InfraRedundancySwitch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraRedundancySwitch *InfraRedundancySwitch) SetParent(parent types.Entity) { infraRedundancySwitch.parent = parent }

func (infraRedundancySwitch *InfraRedundancySwitch) GetParent() types.Entity { return infraRedundancySwitch.parent }

func (infraRedundancySwitch *InfraRedundancySwitch) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InfraBridgeNewRoot
// Generate BRIDGE-MIB::newRoot
type InfraBridgeNewRoot struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetFilter() yfilter.YFilter { return infraBridgeNewRoot.YFilter }

func (infraBridgeNewRoot *InfraBridgeNewRoot) SetFilter(yf yfilter.YFilter) { infraBridgeNewRoot.YFilter = yf }

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetGoName(yname string) string {
    return ""
}

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:infra-bridge-new-root"
}

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetBundleName() string { return "cisco_ios_xr" }

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetYangName() string { return "infra-bridge-new-root" }

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraBridgeNewRoot *InfraBridgeNewRoot) SetParent(parent types.Entity) { infraBridgeNewRoot.parent = parent }

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetParent() types.Entity { return infraBridgeNewRoot.parent }

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InfraBridgeTopologyChange
// Generate BRIDGE-MIB::topologyChange
type InfraBridgeTopologyChange struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetFilter() yfilter.YFilter { return infraBridgeTopologyChange.YFilter }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) SetFilter(yf yfilter.YFilter) { infraBridgeTopologyChange.YFilter = yf }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetGoName(yname string) string {
    return ""
}

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:infra-bridge-topology-change"
}

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetBundleName() string { return "cisco_ios_xr" }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetYangName() string { return "infra-bridge-topology-change" }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) SetParent(parent types.Entity) { infraBridgeTopologyChange.parent = parent }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetParent() types.Entity { return infraBridgeTopologyChange.parent }

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// InfraConfigEvent
// Generate CISCO-CONFIG-MAN-MIB::ciscoConfigManEvent
type InfraConfigEvent struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (infraConfigEvent *InfraConfigEvent) GetFilter() yfilter.YFilter { return infraConfigEvent.YFilter }

func (infraConfigEvent *InfraConfigEvent) SetFilter(yf yfilter.YFilter) { infraConfigEvent.YFilter = yf }

func (infraConfigEvent *InfraConfigEvent) GetGoName(yname string) string {
    return ""
}

func (infraConfigEvent *InfraConfigEvent) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:infra-config-event"
}

func (infraConfigEvent *InfraConfigEvent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (infraConfigEvent *InfraConfigEvent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (infraConfigEvent *InfraConfigEvent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infraConfigEvent *InfraConfigEvent) GetBundleName() string { return "cisco_ios_xr" }

func (infraConfigEvent *InfraConfigEvent) GetYangName() string { return "infra-config-event" }

func (infraConfigEvent *InfraConfigEvent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infraConfigEvent *InfraConfigEvent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infraConfigEvent *InfraConfigEvent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infraConfigEvent *InfraConfigEvent) SetParent(parent types.Entity) { infraConfigEvent.parent = parent }

func (infraConfigEvent *InfraConfigEvent) GetParent() types.Entity { return infraConfigEvent.parent }

func (infraConfigEvent *InfraConfigEvent) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// EntitySensorThresholdNotification
// Generate CISCO-ENTITY-SENSOR-MIB::entSensorThresholdNotification
type EntitySensorThresholdNotification struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EntitySensorThresholdNotification_Input
}

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetFilter() yfilter.YFilter { return entitySensorThresholdNotification.YFilter }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) SetFilter(yf yfilter.YFilter) { entitySensorThresholdNotification.YFilter = yf }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:entity-sensor-threshold-notification"
}

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &entitySensorThresholdNotification.Input
    }
    return nil
}

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &entitySensorThresholdNotification.Input
    return children
}

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetBundleName() string { return "cisco_ios_xr" }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetYangName() string { return "entity-sensor-threshold-notification" }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) SetParent(parent types.Entity) { entitySensorThresholdNotification.parent = parent }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetParent() types.Entity { return entitySensorThresholdNotification.parent }

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// EntitySensorThresholdNotification_Input
type EntitySensorThresholdNotification_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntitySensorThresholdNotification_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EntitySensorThresholdNotification_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EntitySensorThresholdNotification_Input) GetGoName(yname string) string {
    if yname == "entindex" { return "Entindex" }
    return ""
}

func (input *EntitySensorThresholdNotification_Input) GetSegmentPath() string {
    return "input"
}

func (input *EntitySensorThresholdNotification_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *EntitySensorThresholdNotification_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *EntitySensorThresholdNotification_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entindex"] = input.Entindex
    return leafs
}

func (input *EntitySensorThresholdNotification_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *EntitySensorThresholdNotification_Input) GetYangName() string { return "input" }

func (input *EntitySensorThresholdNotification_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *EntitySensorThresholdNotification_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *EntitySensorThresholdNotification_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *EntitySensorThresholdNotification_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EntitySensorThresholdNotification_Input) GetParent() types.Entity { return input.parent }

func (input *EntitySensorThresholdNotification_Input) GetParentYangName() string { return "entity-sensor-threshold-notification" }

// EntityFruPowerStatusChangeFailed
// oper status changed to failed
type EntityFruPowerStatusChangeFailed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EntityFruPowerStatusChangeFailed_Input
}

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetFilter() yfilter.YFilter { return entityFruPowerStatusChangeFailed.YFilter }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) SetFilter(yf yfilter.YFilter) { entityFruPowerStatusChangeFailed.YFilter = yf }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-power-status-change-failed"
}

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &entityFruPowerStatusChangeFailed.Input
    }
    return nil
}

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &entityFruPowerStatusChangeFailed.Input
    return children
}

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetBundleName() string { return "cisco_ios_xr" }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetYangName() string { return "entity-fru-power-status-change-failed" }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) SetParent(parent types.Entity) { entityFruPowerStatusChangeFailed.parent = parent }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetParent() types.Entity { return entityFruPowerStatusChangeFailed.parent }

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// EntityFruPowerStatusChangeFailed_Input
type EntityFruPowerStatusChangeFailed_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruPowerStatusChangeFailed_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EntityFruPowerStatusChangeFailed_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EntityFruPowerStatusChangeFailed_Input) GetGoName(yname string) string {
    if yname == "entindex" { return "Entindex" }
    return ""
}

func (input *EntityFruPowerStatusChangeFailed_Input) GetSegmentPath() string {
    return "input"
}

func (input *EntityFruPowerStatusChangeFailed_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *EntityFruPowerStatusChangeFailed_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *EntityFruPowerStatusChangeFailed_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entindex"] = input.Entindex
    return leafs
}

func (input *EntityFruPowerStatusChangeFailed_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *EntityFruPowerStatusChangeFailed_Input) GetYangName() string { return "input" }

func (input *EntityFruPowerStatusChangeFailed_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *EntityFruPowerStatusChangeFailed_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *EntityFruPowerStatusChangeFailed_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *EntityFruPowerStatusChangeFailed_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EntityFruPowerStatusChangeFailed_Input) GetParent() types.Entity { return input.parent }

func (input *EntityFruPowerStatusChangeFailed_Input) GetParentYangName() string { return "entity-fru-power-status-change-failed" }

// EntityFruModuleStatusChangeUp
// fu trap module status changed as ok
type EntityFruModuleStatusChangeUp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EntityFruModuleStatusChangeUp_Input
}

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetFilter() yfilter.YFilter { return entityFruModuleStatusChangeUp.YFilter }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) SetFilter(yf yfilter.YFilter) { entityFruModuleStatusChangeUp.YFilter = yf }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-module-status-change-up"
}

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &entityFruModuleStatusChangeUp.Input
    }
    return nil
}

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &entityFruModuleStatusChangeUp.Input
    return children
}

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetBundleName() string { return "cisco_ios_xr" }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetYangName() string { return "entity-fru-module-status-change-up" }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) SetParent(parent types.Entity) { entityFruModuleStatusChangeUp.parent = parent }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetParent() types.Entity { return entityFruModuleStatusChangeUp.parent }

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// EntityFruModuleStatusChangeUp_Input
type EntityFruModuleStatusChangeUp_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruModuleStatusChangeUp_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EntityFruModuleStatusChangeUp_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EntityFruModuleStatusChangeUp_Input) GetGoName(yname string) string {
    if yname == "entindex" { return "Entindex" }
    return ""
}

func (input *EntityFruModuleStatusChangeUp_Input) GetSegmentPath() string {
    return "input"
}

func (input *EntityFruModuleStatusChangeUp_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *EntityFruModuleStatusChangeUp_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *EntityFruModuleStatusChangeUp_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entindex"] = input.Entindex
    return leafs
}

func (input *EntityFruModuleStatusChangeUp_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *EntityFruModuleStatusChangeUp_Input) GetYangName() string { return "input" }

func (input *EntityFruModuleStatusChangeUp_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *EntityFruModuleStatusChangeUp_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *EntityFruModuleStatusChangeUp_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *EntityFruModuleStatusChangeUp_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EntityFruModuleStatusChangeUp_Input) GetParent() types.Entity { return input.parent }

func (input *EntityFruModuleStatusChangeUp_Input) GetParentYangName() string { return "entity-fru-module-status-change-up" }

// EntityFruModuleStatusChangeDown
// fu trap module status changed as failed
type EntityFruModuleStatusChangeDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EntityFruModuleStatusChangeDown_Input
}

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetFilter() yfilter.YFilter { return entityFruModuleStatusChangeDown.YFilter }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) SetFilter(yf yfilter.YFilter) { entityFruModuleStatusChangeDown.YFilter = yf }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-module-status-change-down"
}

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &entityFruModuleStatusChangeDown.Input
    }
    return nil
}

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &entityFruModuleStatusChangeDown.Input
    return children
}

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetBundleName() string { return "cisco_ios_xr" }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetYangName() string { return "entity-fru-module-status-change-down" }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) SetParent(parent types.Entity) { entityFruModuleStatusChangeDown.parent = parent }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetParent() types.Entity { return entityFruModuleStatusChangeDown.parent }

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// EntityFruModuleStatusChangeDown_Input
type EntityFruModuleStatusChangeDown_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruModuleStatusChangeDown_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EntityFruModuleStatusChangeDown_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EntityFruModuleStatusChangeDown_Input) GetGoName(yname string) string {
    if yname == "entindex" { return "Entindex" }
    return ""
}

func (input *EntityFruModuleStatusChangeDown_Input) GetSegmentPath() string {
    return "input"
}

func (input *EntityFruModuleStatusChangeDown_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *EntityFruModuleStatusChangeDown_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *EntityFruModuleStatusChangeDown_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entindex"] = input.Entindex
    return leafs
}

func (input *EntityFruModuleStatusChangeDown_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *EntityFruModuleStatusChangeDown_Input) GetYangName() string { return "input" }

func (input *EntityFruModuleStatusChangeDown_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *EntityFruModuleStatusChangeDown_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *EntityFruModuleStatusChangeDown_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *EntityFruModuleStatusChangeDown_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EntityFruModuleStatusChangeDown_Input) GetParent() types.Entity { return input.parent }

func (input *EntityFruModuleStatusChangeDown_Input) GetParentYangName() string { return "entity-fru-module-status-change-down" }

// EntityFruFanTrayOperStatusUp
// Generate CISCO-ENTITY-FRU-CONTROL-MIB::cefcFanTrayStatusChange
type EntityFruFanTrayOperStatusUp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EntityFruFanTrayOperStatusUp_Input
}

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetFilter() yfilter.YFilter { return entityFruFanTrayOperStatusUp.YFilter }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) SetFilter(yf yfilter.YFilter) { entityFruFanTrayOperStatusUp.YFilter = yf }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-oper-status-up"
}

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &entityFruFanTrayOperStatusUp.Input
    }
    return nil
}

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &entityFruFanTrayOperStatusUp.Input
    return children
}

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetBundleName() string { return "cisco_ios_xr" }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetYangName() string { return "entity-fru-fan-tray-oper-status-up" }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) SetParent(parent types.Entity) { entityFruFanTrayOperStatusUp.parent = parent }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetParent() types.Entity { return entityFruFanTrayOperStatusUp.parent }

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// EntityFruFanTrayOperStatusUp_Input
type EntityFruFanTrayOperStatusUp_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruFanTrayOperStatusUp_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EntityFruFanTrayOperStatusUp_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EntityFruFanTrayOperStatusUp_Input) GetGoName(yname string) string {
    if yname == "entindex" { return "Entindex" }
    return ""
}

func (input *EntityFruFanTrayOperStatusUp_Input) GetSegmentPath() string {
    return "input"
}

func (input *EntityFruFanTrayOperStatusUp_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *EntityFruFanTrayOperStatusUp_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *EntityFruFanTrayOperStatusUp_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entindex"] = input.Entindex
    return leafs
}

func (input *EntityFruFanTrayOperStatusUp_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *EntityFruFanTrayOperStatusUp_Input) GetYangName() string { return "input" }

func (input *EntityFruFanTrayOperStatusUp_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *EntityFruFanTrayOperStatusUp_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *EntityFruFanTrayOperStatusUp_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *EntityFruFanTrayOperStatusUp_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EntityFruFanTrayOperStatusUp_Input) GetParent() types.Entity { return input.parent }

func (input *EntityFruFanTrayOperStatusUp_Input) GetParentYangName() string { return "entity-fru-fan-tray-oper-status-up" }

// EntityFruFanTrayInserted
// Generate CISCO-ENTITY-FRU-CONTROL-MIB::cefcFRUInserted
type EntityFruFanTrayInserted struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EntityFruFanTrayInserted_Input
}

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetFilter() yfilter.YFilter { return entityFruFanTrayInserted.YFilter }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) SetFilter(yf yfilter.YFilter) { entityFruFanTrayInserted.YFilter = yf }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-inserted"
}

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &entityFruFanTrayInserted.Input
    }
    return nil
}

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &entityFruFanTrayInserted.Input
    return children
}

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetBundleName() string { return "cisco_ios_xr" }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetYangName() string { return "entity-fru-fan-tray-inserted" }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) SetParent(parent types.Entity) { entityFruFanTrayInserted.parent = parent }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetParent() types.Entity { return entityFruFanTrayInserted.parent }

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// EntityFruFanTrayInserted_Input
type EntityFruFanTrayInserted_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruFanTrayInserted_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EntityFruFanTrayInserted_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EntityFruFanTrayInserted_Input) GetGoName(yname string) string {
    if yname == "entindex" { return "Entindex" }
    return ""
}

func (input *EntityFruFanTrayInserted_Input) GetSegmentPath() string {
    return "input"
}

func (input *EntityFruFanTrayInserted_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *EntityFruFanTrayInserted_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *EntityFruFanTrayInserted_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entindex"] = input.Entindex
    return leafs
}

func (input *EntityFruFanTrayInserted_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *EntityFruFanTrayInserted_Input) GetYangName() string { return "input" }

func (input *EntityFruFanTrayInserted_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *EntityFruFanTrayInserted_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *EntityFruFanTrayInserted_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *EntityFruFanTrayInserted_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EntityFruFanTrayInserted_Input) GetParent() types.Entity { return input.parent }

func (input *EntityFruFanTrayInserted_Input) GetParentYangName() string { return "entity-fru-fan-tray-inserted" }

// EntityFruFanTrayRemoved
// Generate CISCO-ENTITY-FRU-CONTROL-MIB::cefcFRURemoved
type EntityFruFanTrayRemoved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EntityFruFanTrayRemoved_Input
}

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetFilter() yfilter.YFilter { return entityFruFanTrayRemoved.YFilter }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) SetFilter(yf yfilter.YFilter) { entityFruFanTrayRemoved.YFilter = yf }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-removed"
}

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &entityFruFanTrayRemoved.Input
    }
    return nil
}

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &entityFruFanTrayRemoved.Input
    return children
}

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetBundleName() string { return "cisco_ios_xr" }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetYangName() string { return "entity-fru-fan-tray-removed" }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) SetParent(parent types.Entity) { entityFruFanTrayRemoved.parent = parent }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetParent() types.Entity { return entityFruFanTrayRemoved.parent }

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// EntityFruFanTrayRemoved_Input
type EntityFruFanTrayRemoved_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruFanTrayRemoved_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EntityFruFanTrayRemoved_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EntityFruFanTrayRemoved_Input) GetGoName(yname string) string {
    if yname == "entindex" { return "Entindex" }
    return ""
}

func (input *EntityFruFanTrayRemoved_Input) GetSegmentPath() string {
    return "input"
}

func (input *EntityFruFanTrayRemoved_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *EntityFruFanTrayRemoved_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *EntityFruFanTrayRemoved_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entindex"] = input.Entindex
    return leafs
}

func (input *EntityFruFanTrayRemoved_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *EntityFruFanTrayRemoved_Input) GetYangName() string { return "input" }

func (input *EntityFruFanTrayRemoved_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *EntityFruFanTrayRemoved_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *EntityFruFanTrayRemoved_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *EntityFruFanTrayRemoved_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EntityFruFanTrayRemoved_Input) GetParent() types.Entity { return input.parent }

func (input *EntityFruFanTrayRemoved_Input) GetParentYangName() string { return "entity-fru-fan-tray-removed" }

// PlatformHfrBundleDownedLink
// Generate CISCO-FABRIC-HFR-MIB::cfhBundleDownedLinkNotification
type PlatformHfrBundleDownedLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input PlatformHfrBundleDownedLink_Input
}

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetFilter() yfilter.YFilter { return platformHfrBundleDownedLink.YFilter }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) SetFilter(yf yfilter.YFilter) { platformHfrBundleDownedLink.YFilter = yf }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-bundle-downed-link"
}

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &platformHfrBundleDownedLink.Input
    }
    return nil
}

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &platformHfrBundleDownedLink.Input
    return children
}

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetBundleName() string { return "cisco_ios_xr" }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetYangName() string { return "platform-hfr-bundle-downed-link" }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) SetParent(parent types.Entity) { platformHfrBundleDownedLink.parent = parent }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetParent() types.Entity { return platformHfrBundleDownedLink.parent }

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// PlatformHfrBundleDownedLink_Input
type PlatformHfrBundleDownedLink_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bundle name for which to generate the trap. The type is string.
    BundleName interface{}
}

func (input *PlatformHfrBundleDownedLink_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *PlatformHfrBundleDownedLink_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *PlatformHfrBundleDownedLink_Input) GetGoName(yname string) string {
    if yname == "bundle-name" { return "BundleName" }
    return ""
}

func (input *PlatformHfrBundleDownedLink_Input) GetSegmentPath() string {
    return "input"
}

func (input *PlatformHfrBundleDownedLink_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *PlatformHfrBundleDownedLink_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *PlatformHfrBundleDownedLink_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bundle-name"] = input.BundleName
    return leafs
}

func (input *PlatformHfrBundleDownedLink_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *PlatformHfrBundleDownedLink_Input) GetYangName() string { return "input" }

func (input *PlatformHfrBundleDownedLink_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *PlatformHfrBundleDownedLink_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *PlatformHfrBundleDownedLink_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *PlatformHfrBundleDownedLink_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *PlatformHfrBundleDownedLink_Input) GetParent() types.Entity { return input.parent }

func (input *PlatformHfrBundleDownedLink_Input) GetParentYangName() string { return "platform-hfr-bundle-downed-link" }

// PlatformHfrBundleState
// Generate CISCO-FABRIC-HFR-MIB::cfhBundleStateNotification
type PlatformHfrBundleState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input PlatformHfrBundleState_Input
}

func (platformHfrBundleState *PlatformHfrBundleState) GetFilter() yfilter.YFilter { return platformHfrBundleState.YFilter }

func (platformHfrBundleState *PlatformHfrBundleState) SetFilter(yf yfilter.YFilter) { platformHfrBundleState.YFilter = yf }

func (platformHfrBundleState *PlatformHfrBundleState) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (platformHfrBundleState *PlatformHfrBundleState) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-bundle-state"
}

func (platformHfrBundleState *PlatformHfrBundleState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &platformHfrBundleState.Input
    }
    return nil
}

func (platformHfrBundleState *PlatformHfrBundleState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &platformHfrBundleState.Input
    return children
}

func (platformHfrBundleState *PlatformHfrBundleState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformHfrBundleState *PlatformHfrBundleState) GetBundleName() string { return "cisco_ios_xr" }

func (platformHfrBundleState *PlatformHfrBundleState) GetYangName() string { return "platform-hfr-bundle-state" }

func (platformHfrBundleState *PlatformHfrBundleState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformHfrBundleState *PlatformHfrBundleState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformHfrBundleState *PlatformHfrBundleState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformHfrBundleState *PlatformHfrBundleState) SetParent(parent types.Entity) { platformHfrBundleState.parent = parent }

func (platformHfrBundleState *PlatformHfrBundleState) GetParent() types.Entity { return platformHfrBundleState.parent }

func (platformHfrBundleState *PlatformHfrBundleState) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// PlatformHfrBundleState_Input
type PlatformHfrBundleState_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bundle name for which to generate the trap. The type is string.
    BundleName interface{}
}

func (input *PlatformHfrBundleState_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *PlatformHfrBundleState_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *PlatformHfrBundleState_Input) GetGoName(yname string) string {
    if yname == "bundle-name" { return "BundleName" }
    return ""
}

func (input *PlatformHfrBundleState_Input) GetSegmentPath() string {
    return "input"
}

func (input *PlatformHfrBundleState_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *PlatformHfrBundleState_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *PlatformHfrBundleState_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bundle-name"] = input.BundleName
    return leafs
}

func (input *PlatformHfrBundleState_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *PlatformHfrBundleState_Input) GetYangName() string { return "input" }

func (input *PlatformHfrBundleState_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *PlatformHfrBundleState_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *PlatformHfrBundleState_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *PlatformHfrBundleState_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *PlatformHfrBundleState_Input) GetParent() types.Entity { return input.parent }

func (input *PlatformHfrBundleState_Input) GetParentYangName() string { return "platform-hfr-bundle-state" }

// PlatformHfrPlaneState
// Generate CISCO-FABRIC-HFR-MIB::cfhPlaneStateNotification
type PlatformHfrPlaneState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input PlatformHfrPlaneState_Input
}

func (platformHfrPlaneState *PlatformHfrPlaneState) GetFilter() yfilter.YFilter { return platformHfrPlaneState.YFilter }

func (platformHfrPlaneState *PlatformHfrPlaneState) SetFilter(yf yfilter.YFilter) { platformHfrPlaneState.YFilter = yf }

func (platformHfrPlaneState *PlatformHfrPlaneState) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (platformHfrPlaneState *PlatformHfrPlaneState) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-plane-state"
}

func (platformHfrPlaneState *PlatformHfrPlaneState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &platformHfrPlaneState.Input
    }
    return nil
}

func (platformHfrPlaneState *PlatformHfrPlaneState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &platformHfrPlaneState.Input
    return children
}

func (platformHfrPlaneState *PlatformHfrPlaneState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformHfrPlaneState *PlatformHfrPlaneState) GetBundleName() string { return "cisco_ios_xr" }

func (platformHfrPlaneState *PlatformHfrPlaneState) GetYangName() string { return "platform-hfr-plane-state" }

func (platformHfrPlaneState *PlatformHfrPlaneState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformHfrPlaneState *PlatformHfrPlaneState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformHfrPlaneState *PlatformHfrPlaneState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformHfrPlaneState *PlatformHfrPlaneState) SetParent(parent types.Entity) { platformHfrPlaneState.parent = parent }

func (platformHfrPlaneState *PlatformHfrPlaneState) GetParent() types.Entity { return platformHfrPlaneState.parent }

func (platformHfrPlaneState *PlatformHfrPlaneState) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// PlatformHfrPlaneState_Input
type PlatformHfrPlaneState_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // plane identifier for which to generate the trap. The type is interface{}
    // with range: 0..4294967295.
    PlaneId interface{}
}

func (input *PlatformHfrPlaneState_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *PlatformHfrPlaneState_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *PlatformHfrPlaneState_Input) GetGoName(yname string) string {
    if yname == "plane-id" { return "PlaneId" }
    return ""
}

func (input *PlatformHfrPlaneState_Input) GetSegmentPath() string {
    return "input"
}

func (input *PlatformHfrPlaneState_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *PlatformHfrPlaneState_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *PlatformHfrPlaneState_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["plane-id"] = input.PlaneId
    return leafs
}

func (input *PlatformHfrPlaneState_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *PlatformHfrPlaneState_Input) GetYangName() string { return "input" }

func (input *PlatformHfrPlaneState_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *PlatformHfrPlaneState_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *PlatformHfrPlaneState_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *PlatformHfrPlaneState_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *PlatformHfrPlaneState_Input) GetParent() types.Entity { return input.parent }

func (input *PlatformHfrPlaneState_Input) GetParentYangName() string { return "platform-hfr-plane-state" }

// RoutingBgpEstablished
// Generate BGP4-MIB::bglEstablishedNotification
type RoutingBgpEstablished struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (routingBgpEstablished *RoutingBgpEstablished) GetFilter() yfilter.YFilter { return routingBgpEstablished.YFilter }

func (routingBgpEstablished *RoutingBgpEstablished) SetFilter(yf yfilter.YFilter) { routingBgpEstablished.YFilter = yf }

func (routingBgpEstablished *RoutingBgpEstablished) GetGoName(yname string) string {
    return ""
}

func (routingBgpEstablished *RoutingBgpEstablished) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-established"
}

func (routingBgpEstablished *RoutingBgpEstablished) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routingBgpEstablished *RoutingBgpEstablished) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routingBgpEstablished *RoutingBgpEstablished) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingBgpEstablished *RoutingBgpEstablished) GetBundleName() string { return "cisco_ios_xr" }

func (routingBgpEstablished *RoutingBgpEstablished) GetYangName() string { return "routing-bgp-established" }

func (routingBgpEstablished *RoutingBgpEstablished) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingBgpEstablished *RoutingBgpEstablished) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingBgpEstablished *RoutingBgpEstablished) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingBgpEstablished *RoutingBgpEstablished) SetParent(parent types.Entity) { routingBgpEstablished.parent = parent }

func (routingBgpEstablished *RoutingBgpEstablished) GetParent() types.Entity { return routingBgpEstablished.parent }

func (routingBgpEstablished *RoutingBgpEstablished) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingBgpEstablishedRemotePeer
// Generate BGP4-MIB::bglEstablishedNotification remote peer
type RoutingBgpEstablishedRemotePeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RoutingBgpEstablishedRemotePeer_Input
}

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetFilter() yfilter.YFilter { return routingBgpEstablishedRemotePeer.YFilter }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) SetFilter(yf yfilter.YFilter) { routingBgpEstablishedRemotePeer.YFilter = yf }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-established-remote-peer"
}

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &routingBgpEstablishedRemotePeer.Input
    }
    return nil
}

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &routingBgpEstablishedRemotePeer.Input
    return children
}

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetBundleName() string { return "cisco_ios_xr" }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetYangName() string { return "routing-bgp-established-remote-peer" }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) SetParent(parent types.Entity) { routingBgpEstablishedRemotePeer.parent = parent }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetParent() types.Entity { return routingBgpEstablishedRemotePeer.parent }

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingBgpEstablishedRemotePeer_Input
type RoutingBgpEstablishedRemotePeer_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP remote peer IP address for which to generate the trap. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}
}

func (input *RoutingBgpEstablishedRemotePeer_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RoutingBgpEstablishedRemotePeer_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RoutingBgpEstablishedRemotePeer_Input) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (input *RoutingBgpEstablishedRemotePeer_Input) GetSegmentPath() string {
    return "input"
}

func (input *RoutingBgpEstablishedRemotePeer_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RoutingBgpEstablishedRemotePeer_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RoutingBgpEstablishedRemotePeer_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = input.Address
    return leafs
}

func (input *RoutingBgpEstablishedRemotePeer_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RoutingBgpEstablishedRemotePeer_Input) GetYangName() string { return "input" }

func (input *RoutingBgpEstablishedRemotePeer_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RoutingBgpEstablishedRemotePeer_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RoutingBgpEstablishedRemotePeer_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RoutingBgpEstablishedRemotePeer_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RoutingBgpEstablishedRemotePeer_Input) GetParent() types.Entity { return input.parent }

func (input *RoutingBgpEstablishedRemotePeer_Input) GetParentYangName() string { return "routing-bgp-established-remote-peer" }

// RoutingBgpStateChange
// Generate CISCO-BGP-MIB::cbgpBackwardTransition
type RoutingBgpStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (routingBgpStateChange *RoutingBgpStateChange) GetFilter() yfilter.YFilter { return routingBgpStateChange.YFilter }

func (routingBgpStateChange *RoutingBgpStateChange) SetFilter(yf yfilter.YFilter) { routingBgpStateChange.YFilter = yf }

func (routingBgpStateChange *RoutingBgpStateChange) GetGoName(yname string) string {
    return ""
}

func (routingBgpStateChange *RoutingBgpStateChange) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-state-change"
}

func (routingBgpStateChange *RoutingBgpStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routingBgpStateChange *RoutingBgpStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routingBgpStateChange *RoutingBgpStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingBgpStateChange *RoutingBgpStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (routingBgpStateChange *RoutingBgpStateChange) GetYangName() string { return "routing-bgp-state-change" }

func (routingBgpStateChange *RoutingBgpStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingBgpStateChange *RoutingBgpStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingBgpStateChange *RoutingBgpStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingBgpStateChange *RoutingBgpStateChange) SetParent(parent types.Entity) { routingBgpStateChange.parent = parent }

func (routingBgpStateChange *RoutingBgpStateChange) GetParent() types.Entity { return routingBgpStateChange.parent }

func (routingBgpStateChange *RoutingBgpStateChange) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingBgpStateChangeRemotePeer
// Generate CISCO-BGP-MIB::cbgpBackwardTransition remote peer
type RoutingBgpStateChangeRemotePeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RoutingBgpStateChangeRemotePeer_Input
}

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetFilter() yfilter.YFilter { return routingBgpStateChangeRemotePeer.YFilter }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) SetFilter(yf yfilter.YFilter) { routingBgpStateChangeRemotePeer.YFilter = yf }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-state-change-remote-peer"
}

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &routingBgpStateChangeRemotePeer.Input
    }
    return nil
}

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &routingBgpStateChangeRemotePeer.Input
    return children
}

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetBundleName() string { return "cisco_ios_xr" }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetYangName() string { return "routing-bgp-state-change-remote-peer" }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) SetParent(parent types.Entity) { routingBgpStateChangeRemotePeer.parent = parent }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetParent() types.Entity { return routingBgpStateChangeRemotePeer.parent }

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingBgpStateChangeRemotePeer_Input
type RoutingBgpStateChangeRemotePeer_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP remote peer IP address for which to generate the trap. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}
}

func (input *RoutingBgpStateChangeRemotePeer_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RoutingBgpStateChangeRemotePeer_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RoutingBgpStateChangeRemotePeer_Input) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (input *RoutingBgpStateChangeRemotePeer_Input) GetSegmentPath() string {
    return "input"
}

func (input *RoutingBgpStateChangeRemotePeer_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RoutingBgpStateChangeRemotePeer_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RoutingBgpStateChangeRemotePeer_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = input.Address
    return leafs
}

func (input *RoutingBgpStateChangeRemotePeer_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RoutingBgpStateChangeRemotePeer_Input) GetYangName() string { return "input" }

func (input *RoutingBgpStateChangeRemotePeer_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RoutingBgpStateChangeRemotePeer_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RoutingBgpStateChangeRemotePeer_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RoutingBgpStateChangeRemotePeer_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RoutingBgpStateChangeRemotePeer_Input) GetParent() types.Entity { return input.parent }

func (input *RoutingBgpStateChangeRemotePeer_Input) GetParentYangName() string { return "routing-bgp-state-change-remote-peer" }

// RoutingOspfNeighborStateChange
// Generate OSPF-TRAP-MIB::ospfNbrStateChange
type RoutingOspfNeighborStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetFilter() yfilter.YFilter { return routingOspfNeighborStateChange.YFilter }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) SetFilter(yf yfilter.YFilter) { routingOspfNeighborStateChange.YFilter = yf }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetGoName(yname string) string {
    return ""
}

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-ospf-neighbor-state-change"
}

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetYangName() string { return "routing-ospf-neighbor-state-change" }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) SetParent(parent types.Entity) { routingOspfNeighborStateChange.parent = parent }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetParent() types.Entity { return routingOspfNeighborStateChange.parent }

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingOspfNeighborStateChangeAddress
// Generate OSPF-TRAP-MIB::ospfNbrStateChange address
type RoutingOspfNeighborStateChangeAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RoutingOspfNeighborStateChangeAddress_Input
}

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetFilter() yfilter.YFilter { return routingOspfNeighborStateChangeAddress.YFilter }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) SetFilter(yf yfilter.YFilter) { routingOspfNeighborStateChangeAddress.YFilter = yf }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-ospf-neighbor-state-change-address"
}

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &routingOspfNeighborStateChangeAddress.Input
    }
    return nil
}

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &routingOspfNeighborStateChangeAddress.Input
    return children
}

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetBundleName() string { return "cisco_ios_xr" }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetYangName() string { return "routing-ospf-neighbor-state-change-address" }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) SetParent(parent types.Entity) { routingOspfNeighborStateChangeAddress.parent = parent }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetParent() types.Entity { return routingOspfNeighborStateChangeAddress.parent }

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingOspfNeighborStateChangeAddress_Input
type RoutingOspfNeighborStateChangeAddress_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // neighbor's IP source address for which to generate the trap. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // 0 for interfaces having IP addresses or IF-MIB ifindex of addressless
    // interface. The type is interface{} with range: 0..2147483647. This
    // attribute is mandatory.
    Ifindex interface{}
}

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RoutingOspfNeighborStateChangeAddress_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "ifindex" { return "Ifindex" }
    return ""
}

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetSegmentPath() string {
    return "input"
}

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = input.Address
    leafs["ifindex"] = input.Ifindex
    return leafs
}

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetYangName() string { return "input" }

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RoutingOspfNeighborStateChangeAddress_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetParent() types.Entity { return input.parent }

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetParentYangName() string { return "routing-ospf-neighbor-state-change-address" }

// RoutingMplsLdpSessionDown
// Generate MPLS-LDP-STD-MIB::mplsLdpSessionDown
type RoutingMplsLdpSessionDown struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetFilter() yfilter.YFilter { return routingMplsLdpSessionDown.YFilter }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) SetFilter(yf yfilter.YFilter) { routingMplsLdpSessionDown.YFilter = yf }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetGoName(yname string) string {
    return ""
}

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-ldp-session-down"
}

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetBundleName() string { return "cisco_ios_xr" }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetYangName() string { return "routing-mpls-ldp-session-down" }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) SetParent(parent types.Entity) { routingMplsLdpSessionDown.parent = parent }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetParent() types.Entity { return routingMplsLdpSessionDown.parent }

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingMplsLdpSessionDownEntityId
// Generate MPLS-LDP-STD-MIB::mplsLdpSessionDown entity-id
type RoutingMplsLdpSessionDownEntityId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RoutingMplsLdpSessionDownEntityId_Input
}

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetFilter() yfilter.YFilter { return routingMplsLdpSessionDownEntityId.YFilter }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) SetFilter(yf yfilter.YFilter) { routingMplsLdpSessionDownEntityId.YFilter = yf }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-ldp-session-down-entity-id"
}

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &routingMplsLdpSessionDownEntityId.Input
    }
    return nil
}

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &routingMplsLdpSessionDownEntityId.Input
    return children
}

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetBundleName() string { return "cisco_ios_xr" }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetYangName() string { return "routing-mpls-ldp-session-down-entity-id" }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) SetParent(parent types.Entity) { routingMplsLdpSessionDownEntityId.parent = parent }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetParent() types.Entity { return routingMplsLdpSessionDownEntityId.parent }

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingMplsLdpSessionDownEntityId_Input
type RoutingMplsLdpSessionDownEntityId_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // entity ldp-id in x.x.x.x.y.y format where x.x.x.x is the entity IP address
    // and y.y is the label space. The type is string with length: 23. This
    // attribute is mandatory.
    EntityId interface{}

    // entity index for which to generate the trap. The type is interface{} with
    // range: 1..4294967295. This attribute is mandatory.
    EntityIndex interface{}

    // peer ldp-id in x.x.x.x.y.y format where x.x.x.x is the entity IP address
    // and y.y is the label space. The type is string with length: 23. This
    // attribute is mandatory.
    PeerId interface{}
}

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RoutingMplsLdpSessionDownEntityId_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetGoName(yname string) string {
    if yname == "entity-id" { return "EntityId" }
    if yname == "entity-index" { return "EntityIndex" }
    if yname == "peer-id" { return "PeerId" }
    return ""
}

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetSegmentPath() string {
    return "input"
}

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entity-id"] = input.EntityId
    leafs["entity-index"] = input.EntityIndex
    leafs["peer-id"] = input.PeerId
    return leafs
}

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetYangName() string { return "input" }

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RoutingMplsLdpSessionDownEntityId_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetParent() types.Entity { return input.parent }

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetParentYangName() string { return "routing-mpls-ldp-session-down-entity-id" }

// RoutingMplsTunnelReRouted
// Generate MPLS-TE-STD-MIB::mplsTunnelRerouted
type RoutingMplsTunnelReRouted struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetFilter() yfilter.YFilter { return routingMplsTunnelReRouted.YFilter }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) SetFilter(yf yfilter.YFilter) { routingMplsTunnelReRouted.YFilter = yf }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetGoName(yname string) string {
    return ""
}

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-routed"
}

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetBundleName() string { return "cisco_ios_xr" }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetYangName() string { return "routing-mpls-tunnel-re-routed" }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) SetParent(parent types.Entity) { routingMplsTunnelReRouted.parent = parent }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetParent() types.Entity { return routingMplsTunnelReRouted.parent }

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingMplsTunnelReRoutedIndex
// Generate MPLS-TE-STD-MIB::mplsTunnelRerouted index
type RoutingMplsTunnelReRoutedIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RoutingMplsTunnelReRoutedIndex_Input
}

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetFilter() yfilter.YFilter { return routingMplsTunnelReRoutedIndex.YFilter }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) SetFilter(yf yfilter.YFilter) { routingMplsTunnelReRoutedIndex.YFilter = yf }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-routed-index"
}

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &routingMplsTunnelReRoutedIndex.Input
    }
    return nil
}

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &routingMplsTunnelReRoutedIndex.Input
    return children
}

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetBundleName() string { return "cisco_ios_xr" }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetYangName() string { return "routing-mpls-tunnel-re-routed-index" }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) SetParent(parent types.Entity) { routingMplsTunnelReRoutedIndex.parent = parent }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetParent() types.Entity { return routingMplsTunnelReRoutedIndex.parent }

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingMplsTunnelReRoutedIndex_Input
type RoutingMplsTunnelReRoutedIndex_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // tunnel index for which to generate the trap. The type is interface{} with
    // range: 0..65535. This attribute is mandatory.
    Index interface{}

    // tunnel instance for which to generate the trap. The type is interface{}
    // with range: 0..65535. This attribute is mandatory.
    Instance interface{}

    // source address for which to generate the trap. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Source interface{}

    // destination address for which to generate the trap. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Destination interface{}
}

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RoutingMplsTunnelReRoutedIndex_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "instance" { return "Instance" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetSegmentPath() string {
    return "input"
}

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = input.Index
    leafs["instance"] = input.Instance
    leafs["source"] = input.Source
    leafs["destination"] = input.Destination
    return leafs
}

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetYangName() string { return "input" }

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RoutingMplsTunnelReRoutedIndex_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetParent() types.Entity { return input.parent }

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetParentYangName() string { return "routing-mpls-tunnel-re-routed-index" }

// RoutingMplsTunnelReOptimized
// Generate MPLS-TE-STD-MIB::mplsTunnelReoptimized
type RoutingMplsTunnelReOptimized struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetFilter() yfilter.YFilter { return routingMplsTunnelReOptimized.YFilter }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) SetFilter(yf yfilter.YFilter) { routingMplsTunnelReOptimized.YFilter = yf }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetGoName(yname string) string {
    return ""
}

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-optimized"
}

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetBundleName() string { return "cisco_ios_xr" }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetYangName() string { return "routing-mpls-tunnel-re-optimized" }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) SetParent(parent types.Entity) { routingMplsTunnelReOptimized.parent = parent }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetParent() types.Entity { return routingMplsTunnelReOptimized.parent }

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingMplsTunnelReOptimizedIndex
// Generate MPLS-TE-STD-MIB::mplsTunnelReoptimized index
type RoutingMplsTunnelReOptimizedIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RoutingMplsTunnelReOptimizedIndex_Input
}

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetFilter() yfilter.YFilter { return routingMplsTunnelReOptimizedIndex.YFilter }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) SetFilter(yf yfilter.YFilter) { routingMplsTunnelReOptimizedIndex.YFilter = yf }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-optimized-index"
}

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &routingMplsTunnelReOptimizedIndex.Input
    }
    return nil
}

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &routingMplsTunnelReOptimizedIndex.Input
    return children
}

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetBundleName() string { return "cisco_ios_xr" }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetYangName() string { return "routing-mpls-tunnel-re-optimized-index" }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) SetParent(parent types.Entity) { routingMplsTunnelReOptimizedIndex.parent = parent }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetParent() types.Entity { return routingMplsTunnelReOptimizedIndex.parent }

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingMplsTunnelReOptimizedIndex_Input
type RoutingMplsTunnelReOptimizedIndex_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // tunnel index for which to generate the trap. The type is interface{} with
    // range: 0..65535. This attribute is mandatory.
    Index interface{}

    // tunnel instance for which to generate the trap. The type is interface{}
    // with range: 0..65535. This attribute is mandatory.
    Instance interface{}

    // source address for which to generate the trap. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Source interface{}

    // destination address for which to generate the trap. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Destination interface{}
}

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "instance" { return "Instance" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetSegmentPath() string {
    return "input"
}

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = input.Index
    leafs["instance"] = input.Instance
    leafs["source"] = input.Source
    leafs["destination"] = input.Destination
    return leafs
}

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetYangName() string { return "input" }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetParent() types.Entity { return input.parent }

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetParentYangName() string { return "routing-mpls-tunnel-re-optimized-index" }

// RoutingMplsTunnelDown
// Generate MPLS-TE-STD-MIB::mplsTunnelDown
type RoutingMplsTunnelDown struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetFilter() yfilter.YFilter { return routingMplsTunnelDown.YFilter }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) SetFilter(yf yfilter.YFilter) { routingMplsTunnelDown.YFilter = yf }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetGoName(yname string) string {
    return ""
}

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-down"
}

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetBundleName() string { return "cisco_ios_xr" }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetYangName() string { return "routing-mpls-tunnel-down" }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) SetParent(parent types.Entity) { routingMplsTunnelDown.parent = parent }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetParent() types.Entity { return routingMplsTunnelDown.parent }

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingMplsTunnelDownIndex
// Generate MPLS-TE-STD-MIB::mplsTunnelDown index
type RoutingMplsTunnelDownIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RoutingMplsTunnelDownIndex_Input
}

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetFilter() yfilter.YFilter { return routingMplsTunnelDownIndex.YFilter }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) SetFilter(yf yfilter.YFilter) { routingMplsTunnelDownIndex.YFilter = yf }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-down-index"
}

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &routingMplsTunnelDownIndex.Input
    }
    return nil
}

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &routingMplsTunnelDownIndex.Input
    return children
}

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetBundleName() string { return "cisco_ios_xr" }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetYangName() string { return "routing-mpls-tunnel-down-index" }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) SetParent(parent types.Entity) { routingMplsTunnelDownIndex.parent = parent }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetParent() types.Entity { return routingMplsTunnelDownIndex.parent }

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

// RoutingMplsTunnelDownIndex_Input
type RoutingMplsTunnelDownIndex_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // tunnel index for which to generate the trap. The type is interface{} with
    // range: 0..65535. This attribute is mandatory.
    Index interface{}

    // tunnel instance for which to generate the trap. The type is interface{}
    // with range: 0..65535. This attribute is mandatory.
    Instance interface{}

    // src address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Source interface{}

    // destination address for which to generate the trap. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Destination interface{}
}

func (input *RoutingMplsTunnelDownIndex_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RoutingMplsTunnelDownIndex_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RoutingMplsTunnelDownIndex_Input) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "instance" { return "Instance" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (input *RoutingMplsTunnelDownIndex_Input) GetSegmentPath() string {
    return "input"
}

func (input *RoutingMplsTunnelDownIndex_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RoutingMplsTunnelDownIndex_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RoutingMplsTunnelDownIndex_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = input.Index
    leafs["instance"] = input.Instance
    leafs["source"] = input.Source
    leafs["destination"] = input.Destination
    return leafs
}

func (input *RoutingMplsTunnelDownIndex_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RoutingMplsTunnelDownIndex_Input) GetYangName() string { return "input" }

func (input *RoutingMplsTunnelDownIndex_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RoutingMplsTunnelDownIndex_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RoutingMplsTunnelDownIndex_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RoutingMplsTunnelDownIndex_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RoutingMplsTunnelDownIndex_Input) GetParent() types.Entity { return input.parent }

func (input *RoutingMplsTunnelDownIndex_Input) GetParentYangName() string { return "routing-mpls-tunnel-down-index" }

// All
// generate all the supported traps
type All struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (all *All) GetFilter() yfilter.YFilter { return all.YFilter }

func (all *All) SetFilter(yf yfilter.YFilter) { all.YFilter = yf }

func (all *All) GetGoName(yname string) string {
    return ""
}

func (all *All) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-test-trap-act:all"
}

func (all *All) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (all *All) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (all *All) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (all *All) GetBundleName() string { return "cisco_ios_xr" }

func (all *All) GetYangName() string { return "all" }

func (all *All) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (all *All) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (all *All) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (all *All) SetParent(parent types.Entity) { all.parent = parent }

func (all *All) GetParent() types.Entity { return all.parent }

func (all *All) GetParentYangName() string { return "Cisco-IOS-XR-snmp-test-trap-act" }

