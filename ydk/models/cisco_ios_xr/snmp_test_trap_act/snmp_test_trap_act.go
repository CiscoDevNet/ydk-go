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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (snmpColdStart *SnmpColdStart) GetEntityData() *types.CommonEntityData {
    snmpColdStart.EntityData.YFilter = snmpColdStart.YFilter
    snmpColdStart.EntityData.YangName = "snmp-cold-start"
    snmpColdStart.EntityData.BundleName = "cisco_ios_xr"
    snmpColdStart.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    snmpColdStart.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:snmp-cold-start"
    snmpColdStart.EntityData.AbsolutePath = snmpColdStart.EntityData.SegmentPath
    snmpColdStart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpColdStart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpColdStart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpColdStart.EntityData.Children = types.NewOrderedMap()
    snmpColdStart.EntityData.Leafs = types.NewOrderedMap()

    snmpColdStart.EntityData.YListKeys = []string {}

    return &(snmpColdStart.EntityData)
}

// SnmpWarmStart
// Generate SNMPv2-MIB::warmStart
type SnmpWarmStart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (snmpWarmStart *SnmpWarmStart) GetEntityData() *types.CommonEntityData {
    snmpWarmStart.EntityData.YFilter = snmpWarmStart.YFilter
    snmpWarmStart.EntityData.YangName = "snmp-warm-start"
    snmpWarmStart.EntityData.BundleName = "cisco_ios_xr"
    snmpWarmStart.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    snmpWarmStart.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:snmp-warm-start"
    snmpWarmStart.EntityData.AbsolutePath = snmpWarmStart.EntityData.SegmentPath
    snmpWarmStart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpWarmStart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpWarmStart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpWarmStart.EntityData.Children = types.NewOrderedMap()
    snmpWarmStart.EntityData.Leafs = types.NewOrderedMap()

    snmpWarmStart.EntityData.YListKeys = []string {}

    return &(snmpWarmStart.EntityData)
}

// InterfaceLinkUp
// Generate IF-MIB::linkUp
type InterfaceLinkUp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InterfaceLinkUp_Input
}

func (interfaceLinkUp *InterfaceLinkUp) GetEntityData() *types.CommonEntityData {
    interfaceLinkUp.EntityData.YFilter = interfaceLinkUp.YFilter
    interfaceLinkUp.EntityData.YangName = "interface-link-up"
    interfaceLinkUp.EntityData.BundleName = "cisco_ios_xr"
    interfaceLinkUp.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    interfaceLinkUp.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:interface-link-up"
    interfaceLinkUp.EntityData.AbsolutePath = interfaceLinkUp.EntityData.SegmentPath
    interfaceLinkUp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLinkUp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLinkUp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLinkUp.EntityData.Children = types.NewOrderedMap()
    interfaceLinkUp.EntityData.Children.Append("input", types.YChild{"Input", &interfaceLinkUp.Input})
    interfaceLinkUp.EntityData.Leafs = types.NewOrderedMap()

    interfaceLinkUp.EntityData.YListKeys = []string {}

    return &(interfaceLinkUp.EntityData)
}

// InterfaceLinkUp_Input
type InterfaceLinkUp_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface index for which to generate LinkUp trap. The type is interface{}
    // with range: 1..2147483647.
    Ifindex interface{}
}

func (input *InterfaceLinkUp_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "interface-link-up"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:interface-link-up/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", input.Ifindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InterfaceLinkDown
// Generate IF-MIB::linkDown
type InterfaceLinkDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InterfaceLinkDown_Input
}

func (interfaceLinkDown *InterfaceLinkDown) GetEntityData() *types.CommonEntityData {
    interfaceLinkDown.EntityData.YFilter = interfaceLinkDown.YFilter
    interfaceLinkDown.EntityData.YangName = "interface-link-down"
    interfaceLinkDown.EntityData.BundleName = "cisco_ios_xr"
    interfaceLinkDown.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    interfaceLinkDown.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:interface-link-down"
    interfaceLinkDown.EntityData.AbsolutePath = interfaceLinkDown.EntityData.SegmentPath
    interfaceLinkDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLinkDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLinkDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLinkDown.EntityData.Children = types.NewOrderedMap()
    interfaceLinkDown.EntityData.Children.Append("input", types.YChild{"Input", &interfaceLinkDown.Input})
    interfaceLinkDown.EntityData.Leafs = types.NewOrderedMap()

    interfaceLinkDown.EntityData.YListKeys = []string {}

    return &(interfaceLinkDown.EntityData)
}

// InterfaceLinkDown_Input
type InterfaceLinkDown_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface index for which to generate LinkDown trap. The type is
    // interface{} with range: 1..2147483647.
    Ifindex interface{}
}

func (input *InterfaceLinkDown_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "interface-link-down"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:interface-link-down/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", input.Ifindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// SonetSectionStatus
// Generate CISCO-SONET-MIB::ciscoSonetSectionStatusChange
type SonetSectionStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input SonetSectionStatus_Input
}

func (sonetSectionStatus *SonetSectionStatus) GetEntityData() *types.CommonEntityData {
    sonetSectionStatus.EntityData.YFilter = sonetSectionStatus.YFilter
    sonetSectionStatus.EntityData.YangName = "sonet-section-status"
    sonetSectionStatus.EntityData.BundleName = "cisco_ios_xr"
    sonetSectionStatus.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    sonetSectionStatus.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:sonet-section-status"
    sonetSectionStatus.EntityData.AbsolutePath = sonetSectionStatus.EntityData.SegmentPath
    sonetSectionStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sonetSectionStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sonetSectionStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sonetSectionStatus.EntityData.Children = types.NewOrderedMap()
    sonetSectionStatus.EntityData.Children.Append("input", types.YChild{"Input", &sonetSectionStatus.Input})
    sonetSectionStatus.EntityData.Leafs = types.NewOrderedMap()

    sonetSectionStatus.EntityData.YListKeys = []string {}

    return &(sonetSectionStatus.EntityData)
}

// SonetSectionStatus_Input
type SonetSectionStatus_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface index for which to generate ciscoSonetSectionStatusChange trap.
    // The type is interface{} with range: 1..2147483647.
    Ifindex interface{}
}

func (input *SonetSectionStatus_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "sonet-section-status"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:sonet-section-status/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", input.Ifindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// SonetLineStatus
// Generate CISCO-SONET-MIB::ciscoSonetLineStatusChange
type SonetLineStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input SonetLineStatus_Input
}

func (sonetLineStatus *SonetLineStatus) GetEntityData() *types.CommonEntityData {
    sonetLineStatus.EntityData.YFilter = sonetLineStatus.YFilter
    sonetLineStatus.EntityData.YangName = "sonet-line-status"
    sonetLineStatus.EntityData.BundleName = "cisco_ios_xr"
    sonetLineStatus.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    sonetLineStatus.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:sonet-line-status"
    sonetLineStatus.EntityData.AbsolutePath = sonetLineStatus.EntityData.SegmentPath
    sonetLineStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sonetLineStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sonetLineStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sonetLineStatus.EntityData.Children = types.NewOrderedMap()
    sonetLineStatus.EntityData.Children.Append("input", types.YChild{"Input", &sonetLineStatus.Input})
    sonetLineStatus.EntityData.Leafs = types.NewOrderedMap()

    sonetLineStatus.EntityData.YListKeys = []string {}

    return &(sonetLineStatus.EntityData)
}

// SonetLineStatus_Input
type SonetLineStatus_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface index for which to generate ciscoSonetLineStatusChange trap. The
    // type is interface{} with range: 1..2147483647.
    Ifindex interface{}
}

func (input *SonetLineStatus_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "sonet-line-status"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:sonet-line-status/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", input.Ifindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// SonetPathStatus
// Generate CISCO-SONET-MIB::ciscoSonetPathStatusChange
type SonetPathStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input SonetPathStatus_Input
}

func (sonetPathStatus *SonetPathStatus) GetEntityData() *types.CommonEntityData {
    sonetPathStatus.EntityData.YFilter = sonetPathStatus.YFilter
    sonetPathStatus.EntityData.YangName = "sonet-path-status"
    sonetPathStatus.EntityData.BundleName = "cisco_ios_xr"
    sonetPathStatus.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    sonetPathStatus.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:sonet-path-status"
    sonetPathStatus.EntityData.AbsolutePath = sonetPathStatus.EntityData.SegmentPath
    sonetPathStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sonetPathStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sonetPathStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sonetPathStatus.EntityData.Children = types.NewOrderedMap()
    sonetPathStatus.EntityData.Children.Append("input", types.YChild{"Input", &sonetPathStatus.Input})
    sonetPathStatus.EntityData.Leafs = types.NewOrderedMap()

    sonetPathStatus.EntityData.YListKeys = []string {}

    return &(sonetPathStatus.EntityData)
}

// SonetPathStatus_Input
type SonetPathStatus_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface index for which to generate ciscoSonetPathStatusChange trap. The
    // type is interface{} with range: 1..2147483647.
    Ifindex interface{}
}

func (input *SonetPathStatus_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "sonet-path-status"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:sonet-path-status/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", input.Ifindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InfraSyslogMessageGenerated
// Generate CISCO-SYSLOG-MIB::clogMessageGenerated
type InfraSyslogMessageGenerated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (infraSyslogMessageGenerated *InfraSyslogMessageGenerated) GetEntityData() *types.CommonEntityData {
    infraSyslogMessageGenerated.EntityData.YFilter = infraSyslogMessageGenerated.YFilter
    infraSyslogMessageGenerated.EntityData.YangName = "infra-syslog-message-generated"
    infraSyslogMessageGenerated.EntityData.BundleName = "cisco_ios_xr"
    infraSyslogMessageGenerated.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    infraSyslogMessageGenerated.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:infra-syslog-message-generated"
    infraSyslogMessageGenerated.EntityData.AbsolutePath = infraSyslogMessageGenerated.EntityData.SegmentPath
    infraSyslogMessageGenerated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraSyslogMessageGenerated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraSyslogMessageGenerated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraSyslogMessageGenerated.EntityData.Children = types.NewOrderedMap()
    infraSyslogMessageGenerated.EntityData.Leafs = types.NewOrderedMap()

    infraSyslogMessageGenerated.EntityData.YListKeys = []string {}

    return &(infraSyslogMessageGenerated.EntityData)
}

// InfraFlashDeviceInserted
// Generate CISCO-FLASH-MIB::ciscoFlashDeviceInsertedNotif
type InfraFlashDeviceInserted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (infraFlashDeviceInserted *InfraFlashDeviceInserted) GetEntityData() *types.CommonEntityData {
    infraFlashDeviceInserted.EntityData.YFilter = infraFlashDeviceInserted.YFilter
    infraFlashDeviceInserted.EntityData.YangName = "infra-flash-device-inserted"
    infraFlashDeviceInserted.EntityData.BundleName = "cisco_ios_xr"
    infraFlashDeviceInserted.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    infraFlashDeviceInserted.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:infra-flash-device-inserted"
    infraFlashDeviceInserted.EntityData.AbsolutePath = infraFlashDeviceInserted.EntityData.SegmentPath
    infraFlashDeviceInserted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraFlashDeviceInserted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraFlashDeviceInserted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraFlashDeviceInserted.EntityData.Children = types.NewOrderedMap()
    infraFlashDeviceInserted.EntityData.Leafs = types.NewOrderedMap()

    infraFlashDeviceInserted.EntityData.YListKeys = []string {}

    return &(infraFlashDeviceInserted.EntityData)
}

// InfraFlashDeviceRemoved
// Generate CISCO-FLASH-MIB::ciscoFlashDeviceRemovedNotif
type InfraFlashDeviceRemoved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (infraFlashDeviceRemoved *InfraFlashDeviceRemoved) GetEntityData() *types.CommonEntityData {
    infraFlashDeviceRemoved.EntityData.YFilter = infraFlashDeviceRemoved.YFilter
    infraFlashDeviceRemoved.EntityData.YangName = "infra-flash-device-removed"
    infraFlashDeviceRemoved.EntityData.BundleName = "cisco_ios_xr"
    infraFlashDeviceRemoved.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    infraFlashDeviceRemoved.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:infra-flash-device-removed"
    infraFlashDeviceRemoved.EntityData.AbsolutePath = infraFlashDeviceRemoved.EntityData.SegmentPath
    infraFlashDeviceRemoved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraFlashDeviceRemoved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraFlashDeviceRemoved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraFlashDeviceRemoved.EntityData.Children = types.NewOrderedMap()
    infraFlashDeviceRemoved.EntityData.Leafs = types.NewOrderedMap()

    infraFlashDeviceRemoved.EntityData.YListKeys = []string {}

    return &(infraFlashDeviceRemoved.EntityData)
}

// InfraRedundancyProgression
// Generate CISCO-RF-MIB::ciscoRFProgressionNotif
type InfraRedundancyProgression struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (infraRedundancyProgression *InfraRedundancyProgression) GetEntityData() *types.CommonEntityData {
    infraRedundancyProgression.EntityData.YFilter = infraRedundancyProgression.YFilter
    infraRedundancyProgression.EntityData.YangName = "infra-redundancy-progression"
    infraRedundancyProgression.EntityData.BundleName = "cisco_ios_xr"
    infraRedundancyProgression.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    infraRedundancyProgression.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:infra-redundancy-progression"
    infraRedundancyProgression.EntityData.AbsolutePath = infraRedundancyProgression.EntityData.SegmentPath
    infraRedundancyProgression.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraRedundancyProgression.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraRedundancyProgression.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraRedundancyProgression.EntityData.Children = types.NewOrderedMap()
    infraRedundancyProgression.EntityData.Leafs = types.NewOrderedMap()

    infraRedundancyProgression.EntityData.YListKeys = []string {}

    return &(infraRedundancyProgression.EntityData)
}

// InfraRedundancySwitch
// Generate CISCO-RF-MIB::ciscoRFSwactNotif
type InfraRedundancySwitch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (infraRedundancySwitch *InfraRedundancySwitch) GetEntityData() *types.CommonEntityData {
    infraRedundancySwitch.EntityData.YFilter = infraRedundancySwitch.YFilter
    infraRedundancySwitch.EntityData.YangName = "infra-redundancy-switch"
    infraRedundancySwitch.EntityData.BundleName = "cisco_ios_xr"
    infraRedundancySwitch.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    infraRedundancySwitch.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:infra-redundancy-switch"
    infraRedundancySwitch.EntityData.AbsolutePath = infraRedundancySwitch.EntityData.SegmentPath
    infraRedundancySwitch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraRedundancySwitch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraRedundancySwitch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraRedundancySwitch.EntityData.Children = types.NewOrderedMap()
    infraRedundancySwitch.EntityData.Leafs = types.NewOrderedMap()

    infraRedundancySwitch.EntityData.YListKeys = []string {}

    return &(infraRedundancySwitch.EntityData)
}

// InfraBridgeNewRoot
// Generate BRIDGE-MIB::newRoot
type InfraBridgeNewRoot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (infraBridgeNewRoot *InfraBridgeNewRoot) GetEntityData() *types.CommonEntityData {
    infraBridgeNewRoot.EntityData.YFilter = infraBridgeNewRoot.YFilter
    infraBridgeNewRoot.EntityData.YangName = "infra-bridge-new-root"
    infraBridgeNewRoot.EntityData.BundleName = "cisco_ios_xr"
    infraBridgeNewRoot.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    infraBridgeNewRoot.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:infra-bridge-new-root"
    infraBridgeNewRoot.EntityData.AbsolutePath = infraBridgeNewRoot.EntityData.SegmentPath
    infraBridgeNewRoot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraBridgeNewRoot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraBridgeNewRoot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraBridgeNewRoot.EntityData.Children = types.NewOrderedMap()
    infraBridgeNewRoot.EntityData.Leafs = types.NewOrderedMap()

    infraBridgeNewRoot.EntityData.YListKeys = []string {}

    return &(infraBridgeNewRoot.EntityData)
}

// InfraBridgeTopologyChange
// Generate BRIDGE-MIB::topologyChange
type InfraBridgeTopologyChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (infraBridgeTopologyChange *InfraBridgeTopologyChange) GetEntityData() *types.CommonEntityData {
    infraBridgeTopologyChange.EntityData.YFilter = infraBridgeTopologyChange.YFilter
    infraBridgeTopologyChange.EntityData.YangName = "infra-bridge-topology-change"
    infraBridgeTopologyChange.EntityData.BundleName = "cisco_ios_xr"
    infraBridgeTopologyChange.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    infraBridgeTopologyChange.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:infra-bridge-topology-change"
    infraBridgeTopologyChange.EntityData.AbsolutePath = infraBridgeTopologyChange.EntityData.SegmentPath
    infraBridgeTopologyChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraBridgeTopologyChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraBridgeTopologyChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraBridgeTopologyChange.EntityData.Children = types.NewOrderedMap()
    infraBridgeTopologyChange.EntityData.Leafs = types.NewOrderedMap()

    infraBridgeTopologyChange.EntityData.YListKeys = []string {}

    return &(infraBridgeTopologyChange.EntityData)
}

// InfraConfigEvent
// Generate CISCO-CONFIG-MAN-MIB::ciscoConfigManEvent
type InfraConfigEvent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (infraConfigEvent *InfraConfigEvent) GetEntityData() *types.CommonEntityData {
    infraConfigEvent.EntityData.YFilter = infraConfigEvent.YFilter
    infraConfigEvent.EntityData.YangName = "infra-config-event"
    infraConfigEvent.EntityData.BundleName = "cisco_ios_xr"
    infraConfigEvent.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    infraConfigEvent.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:infra-config-event"
    infraConfigEvent.EntityData.AbsolutePath = infraConfigEvent.EntityData.SegmentPath
    infraConfigEvent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infraConfigEvent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infraConfigEvent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infraConfigEvent.EntityData.Children = types.NewOrderedMap()
    infraConfigEvent.EntityData.Leafs = types.NewOrderedMap()

    infraConfigEvent.EntityData.YListKeys = []string {}

    return &(infraConfigEvent.EntityData)
}

// EntitySensorThresholdNotification
// Generate CISCO-ENTITY-SENSOR-MIB::entSensorThresholdNotification
type EntitySensorThresholdNotification struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EntitySensorThresholdNotification_Input
}

func (entitySensorThresholdNotification *EntitySensorThresholdNotification) GetEntityData() *types.CommonEntityData {
    entitySensorThresholdNotification.EntityData.YFilter = entitySensorThresholdNotification.YFilter
    entitySensorThresholdNotification.EntityData.YangName = "entity-sensor-threshold-notification"
    entitySensorThresholdNotification.EntityData.BundleName = "cisco_ios_xr"
    entitySensorThresholdNotification.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    entitySensorThresholdNotification.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:entity-sensor-threshold-notification"
    entitySensorThresholdNotification.EntityData.AbsolutePath = entitySensorThresholdNotification.EntityData.SegmentPath
    entitySensorThresholdNotification.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entitySensorThresholdNotification.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entitySensorThresholdNotification.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entitySensorThresholdNotification.EntityData.Children = types.NewOrderedMap()
    entitySensorThresholdNotification.EntityData.Children.Append("input", types.YChild{"Input", &entitySensorThresholdNotification.Input})
    entitySensorThresholdNotification.EntityData.Leafs = types.NewOrderedMap()

    entitySensorThresholdNotification.EntityData.YListKeys = []string {}

    return &(entitySensorThresholdNotification.EntityData)
}

// EntitySensorThresholdNotification_Input
type EntitySensorThresholdNotification_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntitySensorThresholdNotification_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "entity-sensor-threshold-notification"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:entity-sensor-threshold-notification/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("entindex", types.YLeaf{"Entindex", input.Entindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// EntityFruPowerStatusChangeFailed
// oper status changed to failed
type EntityFruPowerStatusChangeFailed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EntityFruPowerStatusChangeFailed_Input
}

func (entityFruPowerStatusChangeFailed *EntityFruPowerStatusChangeFailed) GetEntityData() *types.CommonEntityData {
    entityFruPowerStatusChangeFailed.EntityData.YFilter = entityFruPowerStatusChangeFailed.YFilter
    entityFruPowerStatusChangeFailed.EntityData.YangName = "entity-fru-power-status-change-failed"
    entityFruPowerStatusChangeFailed.EntityData.BundleName = "cisco_ios_xr"
    entityFruPowerStatusChangeFailed.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    entityFruPowerStatusChangeFailed.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-power-status-change-failed"
    entityFruPowerStatusChangeFailed.EntityData.AbsolutePath = entityFruPowerStatusChangeFailed.EntityData.SegmentPath
    entityFruPowerStatusChangeFailed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityFruPowerStatusChangeFailed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityFruPowerStatusChangeFailed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityFruPowerStatusChangeFailed.EntityData.Children = types.NewOrderedMap()
    entityFruPowerStatusChangeFailed.EntityData.Children.Append("input", types.YChild{"Input", &entityFruPowerStatusChangeFailed.Input})
    entityFruPowerStatusChangeFailed.EntityData.Leafs = types.NewOrderedMap()

    entityFruPowerStatusChangeFailed.EntityData.YListKeys = []string {}

    return &(entityFruPowerStatusChangeFailed.EntityData)
}

// EntityFruPowerStatusChangeFailed_Input
type EntityFruPowerStatusChangeFailed_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruPowerStatusChangeFailed_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "entity-fru-power-status-change-failed"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-power-status-change-failed/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("entindex", types.YLeaf{"Entindex", input.Entindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// EntityFruModuleStatusChangeUp
// fu trap module status changed as ok
type EntityFruModuleStatusChangeUp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EntityFruModuleStatusChangeUp_Input
}

func (entityFruModuleStatusChangeUp *EntityFruModuleStatusChangeUp) GetEntityData() *types.CommonEntityData {
    entityFruModuleStatusChangeUp.EntityData.YFilter = entityFruModuleStatusChangeUp.YFilter
    entityFruModuleStatusChangeUp.EntityData.YangName = "entity-fru-module-status-change-up"
    entityFruModuleStatusChangeUp.EntityData.BundleName = "cisco_ios_xr"
    entityFruModuleStatusChangeUp.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    entityFruModuleStatusChangeUp.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-module-status-change-up"
    entityFruModuleStatusChangeUp.EntityData.AbsolutePath = entityFruModuleStatusChangeUp.EntityData.SegmentPath
    entityFruModuleStatusChangeUp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityFruModuleStatusChangeUp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityFruModuleStatusChangeUp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityFruModuleStatusChangeUp.EntityData.Children = types.NewOrderedMap()
    entityFruModuleStatusChangeUp.EntityData.Children.Append("input", types.YChild{"Input", &entityFruModuleStatusChangeUp.Input})
    entityFruModuleStatusChangeUp.EntityData.Leafs = types.NewOrderedMap()

    entityFruModuleStatusChangeUp.EntityData.YListKeys = []string {}

    return &(entityFruModuleStatusChangeUp.EntityData)
}

// EntityFruModuleStatusChangeUp_Input
type EntityFruModuleStatusChangeUp_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruModuleStatusChangeUp_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "entity-fru-module-status-change-up"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-module-status-change-up/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("entindex", types.YLeaf{"Entindex", input.Entindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// EntityFruModuleStatusChangeDown
// fu trap module status changed as failed
type EntityFruModuleStatusChangeDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EntityFruModuleStatusChangeDown_Input
}

func (entityFruModuleStatusChangeDown *EntityFruModuleStatusChangeDown) GetEntityData() *types.CommonEntityData {
    entityFruModuleStatusChangeDown.EntityData.YFilter = entityFruModuleStatusChangeDown.YFilter
    entityFruModuleStatusChangeDown.EntityData.YangName = "entity-fru-module-status-change-down"
    entityFruModuleStatusChangeDown.EntityData.BundleName = "cisco_ios_xr"
    entityFruModuleStatusChangeDown.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    entityFruModuleStatusChangeDown.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-module-status-change-down"
    entityFruModuleStatusChangeDown.EntityData.AbsolutePath = entityFruModuleStatusChangeDown.EntityData.SegmentPath
    entityFruModuleStatusChangeDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityFruModuleStatusChangeDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityFruModuleStatusChangeDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityFruModuleStatusChangeDown.EntityData.Children = types.NewOrderedMap()
    entityFruModuleStatusChangeDown.EntityData.Children.Append("input", types.YChild{"Input", &entityFruModuleStatusChangeDown.Input})
    entityFruModuleStatusChangeDown.EntityData.Leafs = types.NewOrderedMap()

    entityFruModuleStatusChangeDown.EntityData.YListKeys = []string {}

    return &(entityFruModuleStatusChangeDown.EntityData)
}

// EntityFruModuleStatusChangeDown_Input
type EntityFruModuleStatusChangeDown_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruModuleStatusChangeDown_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "entity-fru-module-status-change-down"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-module-status-change-down/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("entindex", types.YLeaf{"Entindex", input.Entindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// EntityFruFanTrayOperStatusUp
// Generate CISCO-ENTITY-FRU-CONTROL-MIB::cefcFanTrayStatusChange
type EntityFruFanTrayOperStatusUp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EntityFruFanTrayOperStatusUp_Input
}

func (entityFruFanTrayOperStatusUp *EntityFruFanTrayOperStatusUp) GetEntityData() *types.CommonEntityData {
    entityFruFanTrayOperStatusUp.EntityData.YFilter = entityFruFanTrayOperStatusUp.YFilter
    entityFruFanTrayOperStatusUp.EntityData.YangName = "entity-fru-fan-tray-oper-status-up"
    entityFruFanTrayOperStatusUp.EntityData.BundleName = "cisco_ios_xr"
    entityFruFanTrayOperStatusUp.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    entityFruFanTrayOperStatusUp.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-oper-status-up"
    entityFruFanTrayOperStatusUp.EntityData.AbsolutePath = entityFruFanTrayOperStatusUp.EntityData.SegmentPath
    entityFruFanTrayOperStatusUp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityFruFanTrayOperStatusUp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityFruFanTrayOperStatusUp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityFruFanTrayOperStatusUp.EntityData.Children = types.NewOrderedMap()
    entityFruFanTrayOperStatusUp.EntityData.Children.Append("input", types.YChild{"Input", &entityFruFanTrayOperStatusUp.Input})
    entityFruFanTrayOperStatusUp.EntityData.Leafs = types.NewOrderedMap()

    entityFruFanTrayOperStatusUp.EntityData.YListKeys = []string {}

    return &(entityFruFanTrayOperStatusUp.EntityData)
}

// EntityFruFanTrayOperStatusUp_Input
type EntityFruFanTrayOperStatusUp_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruFanTrayOperStatusUp_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "entity-fru-fan-tray-oper-status-up"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-oper-status-up/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("entindex", types.YLeaf{"Entindex", input.Entindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// EntityFruFanTrayInserted
// Generate CISCO-ENTITY-FRU-CONTROL-MIB::cefcFRUInserted
type EntityFruFanTrayInserted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EntityFruFanTrayInserted_Input
}

func (entityFruFanTrayInserted *EntityFruFanTrayInserted) GetEntityData() *types.CommonEntityData {
    entityFruFanTrayInserted.EntityData.YFilter = entityFruFanTrayInserted.YFilter
    entityFruFanTrayInserted.EntityData.YangName = "entity-fru-fan-tray-inserted"
    entityFruFanTrayInserted.EntityData.BundleName = "cisco_ios_xr"
    entityFruFanTrayInserted.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    entityFruFanTrayInserted.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-inserted"
    entityFruFanTrayInserted.EntityData.AbsolutePath = entityFruFanTrayInserted.EntityData.SegmentPath
    entityFruFanTrayInserted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityFruFanTrayInserted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityFruFanTrayInserted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityFruFanTrayInserted.EntityData.Children = types.NewOrderedMap()
    entityFruFanTrayInserted.EntityData.Children.Append("input", types.YChild{"Input", &entityFruFanTrayInserted.Input})
    entityFruFanTrayInserted.EntityData.Leafs = types.NewOrderedMap()

    entityFruFanTrayInserted.EntityData.YListKeys = []string {}

    return &(entityFruFanTrayInserted.EntityData)
}

// EntityFruFanTrayInserted_Input
type EntityFruFanTrayInserted_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruFanTrayInserted_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "entity-fru-fan-tray-inserted"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-inserted/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("entindex", types.YLeaf{"Entindex", input.Entindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// EntityFruFanTrayRemoved
// Generate CISCO-ENTITY-FRU-CONTROL-MIB::cefcFRURemoved
type EntityFruFanTrayRemoved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EntityFruFanTrayRemoved_Input
}

func (entityFruFanTrayRemoved *EntityFruFanTrayRemoved) GetEntityData() *types.CommonEntityData {
    entityFruFanTrayRemoved.EntityData.YFilter = entityFruFanTrayRemoved.YFilter
    entityFruFanTrayRemoved.EntityData.YangName = "entity-fru-fan-tray-removed"
    entityFruFanTrayRemoved.EntityData.BundleName = "cisco_ios_xr"
    entityFruFanTrayRemoved.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    entityFruFanTrayRemoved.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-removed"
    entityFruFanTrayRemoved.EntityData.AbsolutePath = entityFruFanTrayRemoved.EntityData.SegmentPath
    entityFruFanTrayRemoved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityFruFanTrayRemoved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityFruFanTrayRemoved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityFruFanTrayRemoved.EntityData.Children = types.NewOrderedMap()
    entityFruFanTrayRemoved.EntityData.Children.Append("input", types.YChild{"Input", &entityFruFanTrayRemoved.Input})
    entityFruFanTrayRemoved.EntityData.Leafs = types.NewOrderedMap()

    entityFruFanTrayRemoved.EntityData.YListKeys = []string {}

    return &(entityFruFanTrayRemoved.EntityData)
}

// EntityFruFanTrayRemoved_Input
type EntityFruFanTrayRemoved_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // entity Physical Index for which to generate trap. The type is interface{}
    // with range: 1..2147483647.
    Entindex interface{}
}

func (input *EntityFruFanTrayRemoved_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "entity-fru-fan-tray-removed"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:entity-fru-fan-tray-removed/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("entindex", types.YLeaf{"Entindex", input.Entindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// PlatformHfrBundleDownedLink
// Generate CISCO-FABRIC-HFR-MIB::cfhBundleDownedLinkNotification
type PlatformHfrBundleDownedLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input PlatformHfrBundleDownedLink_Input
}

func (platformHfrBundleDownedLink *PlatformHfrBundleDownedLink) GetEntityData() *types.CommonEntityData {
    platformHfrBundleDownedLink.EntityData.YFilter = platformHfrBundleDownedLink.YFilter
    platformHfrBundleDownedLink.EntityData.YangName = "platform-hfr-bundle-downed-link"
    platformHfrBundleDownedLink.EntityData.BundleName = "cisco_ios_xr"
    platformHfrBundleDownedLink.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    platformHfrBundleDownedLink.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-bundle-downed-link"
    platformHfrBundleDownedLink.EntityData.AbsolutePath = platformHfrBundleDownedLink.EntityData.SegmentPath
    platformHfrBundleDownedLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformHfrBundleDownedLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformHfrBundleDownedLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformHfrBundleDownedLink.EntityData.Children = types.NewOrderedMap()
    platformHfrBundleDownedLink.EntityData.Children.Append("input", types.YChild{"Input", &platformHfrBundleDownedLink.Input})
    platformHfrBundleDownedLink.EntityData.Leafs = types.NewOrderedMap()

    platformHfrBundleDownedLink.EntityData.YListKeys = []string {}

    return &(platformHfrBundleDownedLink.EntityData)
}

// PlatformHfrBundleDownedLink_Input
type PlatformHfrBundleDownedLink_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bundle name for which to generate the trap. The type is string.
    BundleName interface{}
}

func (input *PlatformHfrBundleDownedLink_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "platform-hfr-bundle-downed-link"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-bundle-downed-link/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("bundle-name", types.YLeaf{"BundleName", input.BundleName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// PlatformHfrBundleState
// Generate CISCO-FABRIC-HFR-MIB::cfhBundleStateNotification
type PlatformHfrBundleState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input PlatformHfrBundleState_Input
}

func (platformHfrBundleState *PlatformHfrBundleState) GetEntityData() *types.CommonEntityData {
    platformHfrBundleState.EntityData.YFilter = platformHfrBundleState.YFilter
    platformHfrBundleState.EntityData.YangName = "platform-hfr-bundle-state"
    platformHfrBundleState.EntityData.BundleName = "cisco_ios_xr"
    platformHfrBundleState.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    platformHfrBundleState.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-bundle-state"
    platformHfrBundleState.EntityData.AbsolutePath = platformHfrBundleState.EntityData.SegmentPath
    platformHfrBundleState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformHfrBundleState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformHfrBundleState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformHfrBundleState.EntityData.Children = types.NewOrderedMap()
    platformHfrBundleState.EntityData.Children.Append("input", types.YChild{"Input", &platformHfrBundleState.Input})
    platformHfrBundleState.EntityData.Leafs = types.NewOrderedMap()

    platformHfrBundleState.EntityData.YListKeys = []string {}

    return &(platformHfrBundleState.EntityData)
}

// PlatformHfrBundleState_Input
type PlatformHfrBundleState_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bundle name for which to generate the trap. The type is string.
    BundleName interface{}
}

func (input *PlatformHfrBundleState_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "platform-hfr-bundle-state"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-bundle-state/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("bundle-name", types.YLeaf{"BundleName", input.BundleName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// PlatformHfrPlaneState
// Generate CISCO-FABRIC-HFR-MIB::cfhPlaneStateNotification
type PlatformHfrPlaneState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input PlatformHfrPlaneState_Input
}

func (platformHfrPlaneState *PlatformHfrPlaneState) GetEntityData() *types.CommonEntityData {
    platformHfrPlaneState.EntityData.YFilter = platformHfrPlaneState.YFilter
    platformHfrPlaneState.EntityData.YangName = "platform-hfr-plane-state"
    platformHfrPlaneState.EntityData.BundleName = "cisco_ios_xr"
    platformHfrPlaneState.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    platformHfrPlaneState.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-plane-state"
    platformHfrPlaneState.EntityData.AbsolutePath = platformHfrPlaneState.EntityData.SegmentPath
    platformHfrPlaneState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformHfrPlaneState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformHfrPlaneState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformHfrPlaneState.EntityData.Children = types.NewOrderedMap()
    platformHfrPlaneState.EntityData.Children.Append("input", types.YChild{"Input", &platformHfrPlaneState.Input})
    platformHfrPlaneState.EntityData.Leafs = types.NewOrderedMap()

    platformHfrPlaneState.EntityData.YListKeys = []string {}

    return &(platformHfrPlaneState.EntityData)
}

// PlatformHfrPlaneState_Input
type PlatformHfrPlaneState_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // plane identifier for which to generate the trap. The type is interface{}
    // with range: 0..4294967295.
    PlaneId interface{}
}

func (input *PlatformHfrPlaneState_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "platform-hfr-plane-state"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:platform-hfr-plane-state/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("plane-id", types.YLeaf{"PlaneId", input.PlaneId})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RoutingBgpEstablished
// Generate BGP4-MIB::bglEstablishedNotification
type RoutingBgpEstablished struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (routingBgpEstablished *RoutingBgpEstablished) GetEntityData() *types.CommonEntityData {
    routingBgpEstablished.EntityData.YFilter = routingBgpEstablished.YFilter
    routingBgpEstablished.EntityData.YangName = "routing-bgp-established"
    routingBgpEstablished.EntityData.BundleName = "cisco_ios_xr"
    routingBgpEstablished.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingBgpEstablished.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-established"
    routingBgpEstablished.EntityData.AbsolutePath = routingBgpEstablished.EntityData.SegmentPath
    routingBgpEstablished.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingBgpEstablished.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingBgpEstablished.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingBgpEstablished.EntityData.Children = types.NewOrderedMap()
    routingBgpEstablished.EntityData.Leafs = types.NewOrderedMap()

    routingBgpEstablished.EntityData.YListKeys = []string {}

    return &(routingBgpEstablished.EntityData)
}

// RoutingBgpEstablishedRemotePeer
// Generate BGP4-MIB::bglEstablishedNotification remote peer
type RoutingBgpEstablishedRemotePeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RoutingBgpEstablishedRemotePeer_Input
}

func (routingBgpEstablishedRemotePeer *RoutingBgpEstablishedRemotePeer) GetEntityData() *types.CommonEntityData {
    routingBgpEstablishedRemotePeer.EntityData.YFilter = routingBgpEstablishedRemotePeer.YFilter
    routingBgpEstablishedRemotePeer.EntityData.YangName = "routing-bgp-established-remote-peer"
    routingBgpEstablishedRemotePeer.EntityData.BundleName = "cisco_ios_xr"
    routingBgpEstablishedRemotePeer.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingBgpEstablishedRemotePeer.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-established-remote-peer"
    routingBgpEstablishedRemotePeer.EntityData.AbsolutePath = routingBgpEstablishedRemotePeer.EntityData.SegmentPath
    routingBgpEstablishedRemotePeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingBgpEstablishedRemotePeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingBgpEstablishedRemotePeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingBgpEstablishedRemotePeer.EntityData.Children = types.NewOrderedMap()
    routingBgpEstablishedRemotePeer.EntityData.Children.Append("input", types.YChild{"Input", &routingBgpEstablishedRemotePeer.Input})
    routingBgpEstablishedRemotePeer.EntityData.Leafs = types.NewOrderedMap()

    routingBgpEstablishedRemotePeer.EntityData.YListKeys = []string {}

    return &(routingBgpEstablishedRemotePeer.EntityData)
}

// RoutingBgpEstablishedRemotePeer_Input
type RoutingBgpEstablishedRemotePeer_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP remote peer IP address for which to generate the trap. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}
}

func (input *RoutingBgpEstablishedRemotePeer_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "routing-bgp-established-remote-peer"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-established-remote-peer/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("address", types.YLeaf{"Address", input.Address})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RoutingBgpStateChange
// Generate CISCO-BGP-MIB::cbgpBackwardTransition
type RoutingBgpStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (routingBgpStateChange *RoutingBgpStateChange) GetEntityData() *types.CommonEntityData {
    routingBgpStateChange.EntityData.YFilter = routingBgpStateChange.YFilter
    routingBgpStateChange.EntityData.YangName = "routing-bgp-state-change"
    routingBgpStateChange.EntityData.BundleName = "cisco_ios_xr"
    routingBgpStateChange.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingBgpStateChange.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-state-change"
    routingBgpStateChange.EntityData.AbsolutePath = routingBgpStateChange.EntityData.SegmentPath
    routingBgpStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingBgpStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingBgpStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingBgpStateChange.EntityData.Children = types.NewOrderedMap()
    routingBgpStateChange.EntityData.Leafs = types.NewOrderedMap()

    routingBgpStateChange.EntityData.YListKeys = []string {}

    return &(routingBgpStateChange.EntityData)
}

// RoutingBgpStateChangeRemotePeer
// Generate CISCO-BGP-MIB::cbgpBackwardTransition remote peer
type RoutingBgpStateChangeRemotePeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RoutingBgpStateChangeRemotePeer_Input
}

func (routingBgpStateChangeRemotePeer *RoutingBgpStateChangeRemotePeer) GetEntityData() *types.CommonEntityData {
    routingBgpStateChangeRemotePeer.EntityData.YFilter = routingBgpStateChangeRemotePeer.YFilter
    routingBgpStateChangeRemotePeer.EntityData.YangName = "routing-bgp-state-change-remote-peer"
    routingBgpStateChangeRemotePeer.EntityData.BundleName = "cisco_ios_xr"
    routingBgpStateChangeRemotePeer.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingBgpStateChangeRemotePeer.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-state-change-remote-peer"
    routingBgpStateChangeRemotePeer.EntityData.AbsolutePath = routingBgpStateChangeRemotePeer.EntityData.SegmentPath
    routingBgpStateChangeRemotePeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingBgpStateChangeRemotePeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingBgpStateChangeRemotePeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingBgpStateChangeRemotePeer.EntityData.Children = types.NewOrderedMap()
    routingBgpStateChangeRemotePeer.EntityData.Children.Append("input", types.YChild{"Input", &routingBgpStateChangeRemotePeer.Input})
    routingBgpStateChangeRemotePeer.EntityData.Leafs = types.NewOrderedMap()

    routingBgpStateChangeRemotePeer.EntityData.YListKeys = []string {}

    return &(routingBgpStateChangeRemotePeer.EntityData)
}

// RoutingBgpStateChangeRemotePeer_Input
type RoutingBgpStateChangeRemotePeer_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP remote peer IP address for which to generate the trap. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}
}

func (input *RoutingBgpStateChangeRemotePeer_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "routing-bgp-state-change-remote-peer"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:routing-bgp-state-change-remote-peer/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("address", types.YLeaf{"Address", input.Address})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RoutingOspfNeighborStateChange
// Generate OSPF-TRAP-MIB::ospfNbrStateChange
type RoutingOspfNeighborStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (routingOspfNeighborStateChange *RoutingOspfNeighborStateChange) GetEntityData() *types.CommonEntityData {
    routingOspfNeighborStateChange.EntityData.YFilter = routingOspfNeighborStateChange.YFilter
    routingOspfNeighborStateChange.EntityData.YangName = "routing-ospf-neighbor-state-change"
    routingOspfNeighborStateChange.EntityData.BundleName = "cisco_ios_xr"
    routingOspfNeighborStateChange.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingOspfNeighborStateChange.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-ospf-neighbor-state-change"
    routingOspfNeighborStateChange.EntityData.AbsolutePath = routingOspfNeighborStateChange.EntityData.SegmentPath
    routingOspfNeighborStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingOspfNeighborStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingOspfNeighborStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingOspfNeighborStateChange.EntityData.Children = types.NewOrderedMap()
    routingOspfNeighborStateChange.EntityData.Leafs = types.NewOrderedMap()

    routingOspfNeighborStateChange.EntityData.YListKeys = []string {}

    return &(routingOspfNeighborStateChange.EntityData)
}

// RoutingOspfNeighborStateChangeAddress
// Generate OSPF-TRAP-MIB::ospfNbrStateChange address
type RoutingOspfNeighborStateChangeAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RoutingOspfNeighborStateChangeAddress_Input
}

func (routingOspfNeighborStateChangeAddress *RoutingOspfNeighborStateChangeAddress) GetEntityData() *types.CommonEntityData {
    routingOspfNeighborStateChangeAddress.EntityData.YFilter = routingOspfNeighborStateChangeAddress.YFilter
    routingOspfNeighborStateChangeAddress.EntityData.YangName = "routing-ospf-neighbor-state-change-address"
    routingOspfNeighborStateChangeAddress.EntityData.BundleName = "cisco_ios_xr"
    routingOspfNeighborStateChangeAddress.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingOspfNeighborStateChangeAddress.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-ospf-neighbor-state-change-address"
    routingOspfNeighborStateChangeAddress.EntityData.AbsolutePath = routingOspfNeighborStateChangeAddress.EntityData.SegmentPath
    routingOspfNeighborStateChangeAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingOspfNeighborStateChangeAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingOspfNeighborStateChangeAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingOspfNeighborStateChangeAddress.EntityData.Children = types.NewOrderedMap()
    routingOspfNeighborStateChangeAddress.EntityData.Children.Append("input", types.YChild{"Input", &routingOspfNeighborStateChangeAddress.Input})
    routingOspfNeighborStateChangeAddress.EntityData.Leafs = types.NewOrderedMap()

    routingOspfNeighborStateChangeAddress.EntityData.YListKeys = []string {}

    return &(routingOspfNeighborStateChangeAddress.EntityData)
}

// RoutingOspfNeighborStateChangeAddress_Input
type RoutingOspfNeighborStateChangeAddress_Input struct {
    EntityData types.CommonEntityData
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

func (input *RoutingOspfNeighborStateChangeAddress_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "routing-ospf-neighbor-state-change-address"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:routing-ospf-neighbor-state-change-address/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("address", types.YLeaf{"Address", input.Address})
    input.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", input.Ifindex})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RoutingMplsLdpSessionDown
// Generate MPLS-LDP-STD-MIB::mplsLdpSessionDown
type RoutingMplsLdpSessionDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (routingMplsLdpSessionDown *RoutingMplsLdpSessionDown) GetEntityData() *types.CommonEntityData {
    routingMplsLdpSessionDown.EntityData.YFilter = routingMplsLdpSessionDown.YFilter
    routingMplsLdpSessionDown.EntityData.YangName = "routing-mpls-ldp-session-down"
    routingMplsLdpSessionDown.EntityData.BundleName = "cisco_ios_xr"
    routingMplsLdpSessionDown.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingMplsLdpSessionDown.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-ldp-session-down"
    routingMplsLdpSessionDown.EntityData.AbsolutePath = routingMplsLdpSessionDown.EntityData.SegmentPath
    routingMplsLdpSessionDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingMplsLdpSessionDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingMplsLdpSessionDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingMplsLdpSessionDown.EntityData.Children = types.NewOrderedMap()
    routingMplsLdpSessionDown.EntityData.Leafs = types.NewOrderedMap()

    routingMplsLdpSessionDown.EntityData.YListKeys = []string {}

    return &(routingMplsLdpSessionDown.EntityData)
}

// RoutingMplsLdpSessionDownEntityId
// Generate MPLS-LDP-STD-MIB::mplsLdpSessionDown entity-id
type RoutingMplsLdpSessionDownEntityId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RoutingMplsLdpSessionDownEntityId_Input
}

func (routingMplsLdpSessionDownEntityId *RoutingMplsLdpSessionDownEntityId) GetEntityData() *types.CommonEntityData {
    routingMplsLdpSessionDownEntityId.EntityData.YFilter = routingMplsLdpSessionDownEntityId.YFilter
    routingMplsLdpSessionDownEntityId.EntityData.YangName = "routing-mpls-ldp-session-down-entity-id"
    routingMplsLdpSessionDownEntityId.EntityData.BundleName = "cisco_ios_xr"
    routingMplsLdpSessionDownEntityId.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingMplsLdpSessionDownEntityId.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-ldp-session-down-entity-id"
    routingMplsLdpSessionDownEntityId.EntityData.AbsolutePath = routingMplsLdpSessionDownEntityId.EntityData.SegmentPath
    routingMplsLdpSessionDownEntityId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingMplsLdpSessionDownEntityId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingMplsLdpSessionDownEntityId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingMplsLdpSessionDownEntityId.EntityData.Children = types.NewOrderedMap()
    routingMplsLdpSessionDownEntityId.EntityData.Children.Append("input", types.YChild{"Input", &routingMplsLdpSessionDownEntityId.Input})
    routingMplsLdpSessionDownEntityId.EntityData.Leafs = types.NewOrderedMap()

    routingMplsLdpSessionDownEntityId.EntityData.YListKeys = []string {}

    return &(routingMplsLdpSessionDownEntityId.EntityData)
}

// RoutingMplsLdpSessionDownEntityId_Input
type RoutingMplsLdpSessionDownEntityId_Input struct {
    EntityData types.CommonEntityData
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

func (input *RoutingMplsLdpSessionDownEntityId_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "routing-mpls-ldp-session-down-entity-id"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-ldp-session-down-entity-id/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("entity-id", types.YLeaf{"EntityId", input.EntityId})
    input.EntityData.Leafs.Append("entity-index", types.YLeaf{"EntityIndex", input.EntityIndex})
    input.EntityData.Leafs.Append("peer-id", types.YLeaf{"PeerId", input.PeerId})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RoutingMplsTunnelReRouted
// Generate MPLS-TE-STD-MIB::mplsTunnelRerouted
type RoutingMplsTunnelReRouted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (routingMplsTunnelReRouted *RoutingMplsTunnelReRouted) GetEntityData() *types.CommonEntityData {
    routingMplsTunnelReRouted.EntityData.YFilter = routingMplsTunnelReRouted.YFilter
    routingMplsTunnelReRouted.EntityData.YangName = "routing-mpls-tunnel-re-routed"
    routingMplsTunnelReRouted.EntityData.BundleName = "cisco_ios_xr"
    routingMplsTunnelReRouted.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingMplsTunnelReRouted.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-routed"
    routingMplsTunnelReRouted.EntityData.AbsolutePath = routingMplsTunnelReRouted.EntityData.SegmentPath
    routingMplsTunnelReRouted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingMplsTunnelReRouted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingMplsTunnelReRouted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingMplsTunnelReRouted.EntityData.Children = types.NewOrderedMap()
    routingMplsTunnelReRouted.EntityData.Leafs = types.NewOrderedMap()

    routingMplsTunnelReRouted.EntityData.YListKeys = []string {}

    return &(routingMplsTunnelReRouted.EntityData)
}

// RoutingMplsTunnelReRoutedIndex
// Generate MPLS-TE-STD-MIB::mplsTunnelRerouted index
type RoutingMplsTunnelReRoutedIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RoutingMplsTunnelReRoutedIndex_Input
}

func (routingMplsTunnelReRoutedIndex *RoutingMplsTunnelReRoutedIndex) GetEntityData() *types.CommonEntityData {
    routingMplsTunnelReRoutedIndex.EntityData.YFilter = routingMplsTunnelReRoutedIndex.YFilter
    routingMplsTunnelReRoutedIndex.EntityData.YangName = "routing-mpls-tunnel-re-routed-index"
    routingMplsTunnelReRoutedIndex.EntityData.BundleName = "cisco_ios_xr"
    routingMplsTunnelReRoutedIndex.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingMplsTunnelReRoutedIndex.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-routed-index"
    routingMplsTunnelReRoutedIndex.EntityData.AbsolutePath = routingMplsTunnelReRoutedIndex.EntityData.SegmentPath
    routingMplsTunnelReRoutedIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingMplsTunnelReRoutedIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingMplsTunnelReRoutedIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingMplsTunnelReRoutedIndex.EntityData.Children = types.NewOrderedMap()
    routingMplsTunnelReRoutedIndex.EntityData.Children.Append("input", types.YChild{"Input", &routingMplsTunnelReRoutedIndex.Input})
    routingMplsTunnelReRoutedIndex.EntityData.Leafs = types.NewOrderedMap()

    routingMplsTunnelReRoutedIndex.EntityData.YListKeys = []string {}

    return &(routingMplsTunnelReRoutedIndex.EntityData)
}

// RoutingMplsTunnelReRoutedIndex_Input
type RoutingMplsTunnelReRoutedIndex_Input struct {
    EntityData types.CommonEntityData
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

func (input *RoutingMplsTunnelReRoutedIndex_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "routing-mpls-tunnel-re-routed-index"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-routed-index/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("index", types.YLeaf{"Index", input.Index})
    input.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", input.Instance})
    input.EntityData.Leafs.Append("source", types.YLeaf{"Source", input.Source})
    input.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", input.Destination})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RoutingMplsTunnelReOptimized
// Generate MPLS-TE-STD-MIB::mplsTunnelReoptimized
type RoutingMplsTunnelReOptimized struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (routingMplsTunnelReOptimized *RoutingMplsTunnelReOptimized) GetEntityData() *types.CommonEntityData {
    routingMplsTunnelReOptimized.EntityData.YFilter = routingMplsTunnelReOptimized.YFilter
    routingMplsTunnelReOptimized.EntityData.YangName = "routing-mpls-tunnel-re-optimized"
    routingMplsTunnelReOptimized.EntityData.BundleName = "cisco_ios_xr"
    routingMplsTunnelReOptimized.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingMplsTunnelReOptimized.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-optimized"
    routingMplsTunnelReOptimized.EntityData.AbsolutePath = routingMplsTunnelReOptimized.EntityData.SegmentPath
    routingMplsTunnelReOptimized.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingMplsTunnelReOptimized.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingMplsTunnelReOptimized.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingMplsTunnelReOptimized.EntityData.Children = types.NewOrderedMap()
    routingMplsTunnelReOptimized.EntityData.Leafs = types.NewOrderedMap()

    routingMplsTunnelReOptimized.EntityData.YListKeys = []string {}

    return &(routingMplsTunnelReOptimized.EntityData)
}

// RoutingMplsTunnelReOptimizedIndex
// Generate MPLS-TE-STD-MIB::mplsTunnelReoptimized index
type RoutingMplsTunnelReOptimizedIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RoutingMplsTunnelReOptimizedIndex_Input
}

func (routingMplsTunnelReOptimizedIndex *RoutingMplsTunnelReOptimizedIndex) GetEntityData() *types.CommonEntityData {
    routingMplsTunnelReOptimizedIndex.EntityData.YFilter = routingMplsTunnelReOptimizedIndex.YFilter
    routingMplsTunnelReOptimizedIndex.EntityData.YangName = "routing-mpls-tunnel-re-optimized-index"
    routingMplsTunnelReOptimizedIndex.EntityData.BundleName = "cisco_ios_xr"
    routingMplsTunnelReOptimizedIndex.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingMplsTunnelReOptimizedIndex.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-optimized-index"
    routingMplsTunnelReOptimizedIndex.EntityData.AbsolutePath = routingMplsTunnelReOptimizedIndex.EntityData.SegmentPath
    routingMplsTunnelReOptimizedIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingMplsTunnelReOptimizedIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingMplsTunnelReOptimizedIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingMplsTunnelReOptimizedIndex.EntityData.Children = types.NewOrderedMap()
    routingMplsTunnelReOptimizedIndex.EntityData.Children.Append("input", types.YChild{"Input", &routingMplsTunnelReOptimizedIndex.Input})
    routingMplsTunnelReOptimizedIndex.EntityData.Leafs = types.NewOrderedMap()

    routingMplsTunnelReOptimizedIndex.EntityData.YListKeys = []string {}

    return &(routingMplsTunnelReOptimizedIndex.EntityData)
}

// RoutingMplsTunnelReOptimizedIndex_Input
type RoutingMplsTunnelReOptimizedIndex_Input struct {
    EntityData types.CommonEntityData
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

func (input *RoutingMplsTunnelReOptimizedIndex_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "routing-mpls-tunnel-re-optimized-index"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-re-optimized-index/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("index", types.YLeaf{"Index", input.Index})
    input.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", input.Instance})
    input.EntityData.Leafs.Append("source", types.YLeaf{"Source", input.Source})
    input.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", input.Destination})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RoutingMplsTunnelDown
// Generate MPLS-TE-STD-MIB::mplsTunnelDown
type RoutingMplsTunnelDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (routingMplsTunnelDown *RoutingMplsTunnelDown) GetEntityData() *types.CommonEntityData {
    routingMplsTunnelDown.EntityData.YFilter = routingMplsTunnelDown.YFilter
    routingMplsTunnelDown.EntityData.YangName = "routing-mpls-tunnel-down"
    routingMplsTunnelDown.EntityData.BundleName = "cisco_ios_xr"
    routingMplsTunnelDown.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingMplsTunnelDown.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-down"
    routingMplsTunnelDown.EntityData.AbsolutePath = routingMplsTunnelDown.EntityData.SegmentPath
    routingMplsTunnelDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingMplsTunnelDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingMplsTunnelDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingMplsTunnelDown.EntityData.Children = types.NewOrderedMap()
    routingMplsTunnelDown.EntityData.Leafs = types.NewOrderedMap()

    routingMplsTunnelDown.EntityData.YListKeys = []string {}

    return &(routingMplsTunnelDown.EntityData)
}

// RoutingMplsTunnelDownIndex
// Generate MPLS-TE-STD-MIB::mplsTunnelDown index
type RoutingMplsTunnelDownIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RoutingMplsTunnelDownIndex_Input
}

func (routingMplsTunnelDownIndex *RoutingMplsTunnelDownIndex) GetEntityData() *types.CommonEntityData {
    routingMplsTunnelDownIndex.EntityData.YFilter = routingMplsTunnelDownIndex.YFilter
    routingMplsTunnelDownIndex.EntityData.YangName = "routing-mpls-tunnel-down-index"
    routingMplsTunnelDownIndex.EntityData.BundleName = "cisco_ios_xr"
    routingMplsTunnelDownIndex.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    routingMplsTunnelDownIndex.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-down-index"
    routingMplsTunnelDownIndex.EntityData.AbsolutePath = routingMplsTunnelDownIndex.EntityData.SegmentPath
    routingMplsTunnelDownIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingMplsTunnelDownIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingMplsTunnelDownIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingMplsTunnelDownIndex.EntityData.Children = types.NewOrderedMap()
    routingMplsTunnelDownIndex.EntityData.Children.Append("input", types.YChild{"Input", &routingMplsTunnelDownIndex.Input})
    routingMplsTunnelDownIndex.EntityData.Leafs = types.NewOrderedMap()

    routingMplsTunnelDownIndex.EntityData.YListKeys = []string {}

    return &(routingMplsTunnelDownIndex.EntityData)
}

// RoutingMplsTunnelDownIndex_Input
type RoutingMplsTunnelDownIndex_Input struct {
    EntityData types.CommonEntityData
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

func (input *RoutingMplsTunnelDownIndex_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "routing-mpls-tunnel-down-index"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-test-trap-act:routing-mpls-tunnel-down-index/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("index", types.YLeaf{"Index", input.Index})
    input.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", input.Instance})
    input.EntityData.Leafs.Append("source", types.YLeaf{"Source", input.Source})
    input.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", input.Destination})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// All
// generate all the supported traps
type All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (all *All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-test-trap-act"
    all.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-test-trap-act:all"
    all.EntityData.AbsolutePath = all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

