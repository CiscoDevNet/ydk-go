// This module contains a collection of YANG definitions
// for Cisco IOS-XR icpe-infra package configuration.
// 
// This module contains definitions
// for the following management objects:
//   nv-satellite-global: nV Satellite Global configuration
//   nv-satellites: nv satellites
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-rgmgr-cfg,
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package icpe_infra_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package icpe_infra_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-icpe-infra-cfg nv-satellite-global}", reflect.TypeOf(NvSatelliteGlobal{}))
    ydk.RegisterEntity("Cisco-IOS-XR-icpe-infra-cfg:nv-satellite-global", reflect.TypeOf(NvSatelliteGlobal{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-icpe-infra-cfg nv-satellites}", reflect.TypeOf(NvSatellites{}))
    ydk.RegisterEntity("Cisco-IOS-XR-icpe-infra-cfg:nv-satellites", reflect.TypeOf(NvSatellites{}))
}

// NvSatelliteGlobal
// nV Satellite Global configuration
type NvSatelliteGlobal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Chassis MAC address.
    ChassisMac NvSatelliteGlobal_ChassisMac
}

func (nvSatelliteGlobal *NvSatelliteGlobal) GetFilter() yfilter.YFilter { return nvSatelliteGlobal.YFilter }

func (nvSatelliteGlobal *NvSatelliteGlobal) SetFilter(yf yfilter.YFilter) { nvSatelliteGlobal.YFilter = yf }

func (nvSatelliteGlobal *NvSatelliteGlobal) GetGoName(yname string) string {
    if yname == "chassis-mac" { return "ChassisMac" }
    return ""
}

func (nvSatelliteGlobal *NvSatelliteGlobal) GetSegmentPath() string {
    return "Cisco-IOS-XR-icpe-infra-cfg:nv-satellite-global"
}

func (nvSatelliteGlobal *NvSatelliteGlobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "chassis-mac" {
        return &nvSatelliteGlobal.ChassisMac
    }
    return nil
}

func (nvSatelliteGlobal *NvSatelliteGlobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["chassis-mac"] = &nvSatelliteGlobal.ChassisMac
    return children
}

func (nvSatelliteGlobal *NvSatelliteGlobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nvSatelliteGlobal *NvSatelliteGlobal) GetBundleName() string { return "cisco_ios_xr" }

func (nvSatelliteGlobal *NvSatelliteGlobal) GetYangName() string { return "nv-satellite-global" }

func (nvSatelliteGlobal *NvSatelliteGlobal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nvSatelliteGlobal *NvSatelliteGlobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nvSatelliteGlobal *NvSatelliteGlobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nvSatelliteGlobal *NvSatelliteGlobal) SetParent(parent types.Entity) { nvSatelliteGlobal.parent = parent }

func (nvSatelliteGlobal *NvSatelliteGlobal) GetParent() types.Entity { return nvSatelliteGlobal.parent }

func (nvSatelliteGlobal *NvSatelliteGlobal) GetParentYangName() string { return "Cisco-IOS-XR-icpe-infra-cfg" }

// NvSatelliteGlobal_ChassisMac
// Chassis MAC address
type NvSatelliteGlobal_ChassisMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // First two bytes of MAC address. The type is interface{} with range:
    // 0..2147483647. Units are byte.
    Mac1 interface{}

    // Second two bytes of MAC address. The type is interface{} with range:
    // 0..2147483647. Units are byte.
    Mac2 interface{}

    // Third two bytes of MAC address. The type is interface{} with range:
    // 0..2147483647. Units are byte.
    Mac3 interface{}
}

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetFilter() yfilter.YFilter { return chassisMac.YFilter }

func (chassisMac *NvSatelliteGlobal_ChassisMac) SetFilter(yf yfilter.YFilter) { chassisMac.YFilter = yf }

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetGoName(yname string) string {
    if yname == "mac1" { return "Mac1" }
    if yname == "mac2" { return "Mac2" }
    if yname == "mac3" { return "Mac3" }
    return ""
}

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetSegmentPath() string {
    return "chassis-mac"
}

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac1"] = chassisMac.Mac1
    leafs["mac2"] = chassisMac.Mac2
    leafs["mac3"] = chassisMac.Mac3
    return leafs
}

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetBundleName() string { return "cisco_ios_xr" }

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetYangName() string { return "chassis-mac" }

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassisMac *NvSatelliteGlobal_ChassisMac) SetParent(parent types.Entity) { chassisMac.parent = parent }

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetParent() types.Entity { return chassisMac.parent }

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetParentYangName() string { return "nv-satellite-global" }

// NvSatellites
// nv satellites
type NvSatellites struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite Configuration. The type is slice of NvSatellites_NvSatellite.
    NvSatellite []NvSatellites_NvSatellite
}

func (nvSatellites *NvSatellites) GetFilter() yfilter.YFilter { return nvSatellites.YFilter }

func (nvSatellites *NvSatellites) SetFilter(yf yfilter.YFilter) { nvSatellites.YFilter = yf }

func (nvSatellites *NvSatellites) GetGoName(yname string) string {
    if yname == "nv-satellite" { return "NvSatellite" }
    return ""
}

func (nvSatellites *NvSatellites) GetSegmentPath() string {
    return "Cisco-IOS-XR-icpe-infra-cfg:nv-satellites"
}

func (nvSatellites *NvSatellites) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nv-satellite" {
        for _, c := range nvSatellites.NvSatellite {
            if nvSatellites.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellites_NvSatellite{}
        nvSatellites.NvSatellite = append(nvSatellites.NvSatellite, child)
        return &nvSatellites.NvSatellite[len(nvSatellites.NvSatellite)-1]
    }
    return nil
}

func (nvSatellites *NvSatellites) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nvSatellites.NvSatellite {
        children[nvSatellites.NvSatellite[i].GetSegmentPath()] = &nvSatellites.NvSatellite[i]
    }
    return children
}

func (nvSatellites *NvSatellites) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nvSatellites *NvSatellites) GetBundleName() string { return "cisco_ios_xr" }

func (nvSatellites *NvSatellites) GetYangName() string { return "nv-satellites" }

func (nvSatellites *NvSatellites) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nvSatellites *NvSatellites) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nvSatellites *NvSatellites) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nvSatellites *NvSatellites) SetParent(parent types.Entity) { nvSatellites.parent = parent }

func (nvSatellites *NvSatellites) GetParent() types.Entity { return nvSatellites.parent }

func (nvSatellites *NvSatellites) GetParentYangName() string { return "Cisco-IOS-XR-icpe-infra-cfg" }

// NvSatellites_NvSatellite
// Satellite Configuration
type NvSatellites_NvSatellite struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite ID. The type is interface{} with range:
    // 100..65534.
    SatelliteId interface{}

    // VRF for Satellite IP Address. The type is string.
    Vrf interface{}

    // Discovery timeout warning for the satellite. The type is interface{} with
    // range: 0..4294967295.
    TimeoutWarning interface{}

    // Satellite Name. The type is string.
    DeviceName interface{}

    // Satellite Description. The type is string.
    Description interface{}

    // Satellite Type. The type is string.
    Type interface{}

    // Enable. The type is interface{}.
    Enable interface{}

    // Discovery timeout for the satellite. The type is interface{} with range:
    // 0..4294967295.
    DiscTimeout interface{}

    // Timer (in seconds) for delaying switchback in a dual home setup. The type
    // is interface{} with range: 0..4294967295. Units are second.
    DelayedSwitchback interface{}

    // Satellite Serial Number. The type is string.
    SerialNumber interface{}

    // Encrypted password for the Satellite. The type is string with pattern:
    // (!.+)|([^!].+).
    Secret interface{}

    // Satellite IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Satellite auto-upgrade capability.
    UpgradeOnConnect NvSatellites_NvSatellite_UpgradeOnConnect

    // Enable interfaces on the satellite to be used as fabric ports table.
    CandidateFabricPorts NvSatellites_NvSatellite_CandidateFabricPorts

    // Satellite User.
    ConnectionInfo NvSatellites_NvSatellite_ConnectionInfo

    // Redundancy submode.
    Redundancy NvSatellites_NvSatellite_Redundancy
}

func (nvSatellite *NvSatellites_NvSatellite) GetFilter() yfilter.YFilter { return nvSatellite.YFilter }

func (nvSatellite *NvSatellites_NvSatellite) SetFilter(yf yfilter.YFilter) { nvSatellite.YFilter = yf }

func (nvSatellite *NvSatellites_NvSatellite) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "vrf" { return "Vrf" }
    if yname == "timeout-warning" { return "TimeoutWarning" }
    if yname == "device-name" { return "DeviceName" }
    if yname == "description" { return "Description" }
    if yname == "type" { return "Type" }
    if yname == "enable" { return "Enable" }
    if yname == "disc-timeout" { return "DiscTimeout" }
    if yname == "delayed-switchback" { return "DelayedSwitchback" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "secret" { return "Secret" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "upgrade-on-connect" { return "UpgradeOnConnect" }
    if yname == "candidate-fabric-ports" { return "CandidateFabricPorts" }
    if yname == "connection-info" { return "ConnectionInfo" }
    if yname == "redundancy" { return "Redundancy" }
    return ""
}

func (nvSatellite *NvSatellites_NvSatellite) GetSegmentPath() string {
    return "nv-satellite" + "[satellite-id='" + fmt.Sprintf("%v", nvSatellite.SatelliteId) + "']"
}

func (nvSatellite *NvSatellites_NvSatellite) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "upgrade-on-connect" {
        return &nvSatellite.UpgradeOnConnect
    }
    if childYangName == "candidate-fabric-ports" {
        return &nvSatellite.CandidateFabricPorts
    }
    if childYangName == "connection-info" {
        return &nvSatellite.ConnectionInfo
    }
    if childYangName == "redundancy" {
        return &nvSatellite.Redundancy
    }
    return nil
}

func (nvSatellite *NvSatellites_NvSatellite) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["upgrade-on-connect"] = &nvSatellite.UpgradeOnConnect
    children["candidate-fabric-ports"] = &nvSatellite.CandidateFabricPorts
    children["connection-info"] = &nvSatellite.ConnectionInfo
    children["redundancy"] = &nvSatellite.Redundancy
    return children
}

func (nvSatellite *NvSatellites_NvSatellite) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = nvSatellite.SatelliteId
    leafs["vrf"] = nvSatellite.Vrf
    leafs["timeout-warning"] = nvSatellite.TimeoutWarning
    leafs["device-name"] = nvSatellite.DeviceName
    leafs["description"] = nvSatellite.Description
    leafs["type"] = nvSatellite.Type
    leafs["enable"] = nvSatellite.Enable
    leafs["disc-timeout"] = nvSatellite.DiscTimeout
    leafs["delayed-switchback"] = nvSatellite.DelayedSwitchback
    leafs["serial-number"] = nvSatellite.SerialNumber
    leafs["secret"] = nvSatellite.Secret
    leafs["ip-address"] = nvSatellite.IpAddress
    return leafs
}

func (nvSatellite *NvSatellites_NvSatellite) GetBundleName() string { return "cisco_ios_xr" }

func (nvSatellite *NvSatellites_NvSatellite) GetYangName() string { return "nv-satellite" }

func (nvSatellite *NvSatellites_NvSatellite) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nvSatellite *NvSatellites_NvSatellite) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nvSatellite *NvSatellites_NvSatellite) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nvSatellite *NvSatellites_NvSatellite) SetParent(parent types.Entity) { nvSatellite.parent = parent }

func (nvSatellite *NvSatellites_NvSatellite) GetParent() types.Entity { return nvSatellite.parent }

func (nvSatellite *NvSatellites_NvSatellite) GetParentYangName() string { return "nv-satellites" }

// NvSatellites_NvSatellite_UpgradeOnConnect
// Satellite auto-upgrade capability
type NvSatellites_NvSatellite_UpgradeOnConnect struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // When to upgrade the satellite. The type is ConnectType.
    ConnectType interface{}

    // Reference name. The type is string.
    Reference interface{}
}

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetFilter() yfilter.YFilter { return upgradeOnConnect.YFilter }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) SetFilter(yf yfilter.YFilter) { upgradeOnConnect.YFilter = yf }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetGoName(yname string) string {
    if yname == "connect-type" { return "ConnectType" }
    if yname == "reference" { return "Reference" }
    return ""
}

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetSegmentPath() string {
    return "upgrade-on-connect"
}

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["connect-type"] = upgradeOnConnect.ConnectType
    leafs["reference"] = upgradeOnConnect.Reference
    return leafs
}

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetBundleName() string { return "cisco_ios_xr" }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetYangName() string { return "upgrade-on-connect" }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) SetParent(parent types.Entity) { upgradeOnConnect.parent = parent }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetParent() types.Entity { return upgradeOnConnect.parent }

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetParentYangName() string { return "nv-satellite" }

// NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType represents When to upgrade the satellite
type NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType string

const (
    // Satellite Upgrade on Connection
    NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType_on_connection NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType = "on-connection"

    // Satellite Upgrade on First Connection after
    // configuration or host boot
    NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType_on_first_connection NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType = "on-first-connection"
)

// NvSatellites_NvSatellite_CandidateFabricPorts
// Enable interfaces on the satellite to be used
// as fabric ports table
type NvSatellites_NvSatellite_CandidateFabricPorts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable interfaces on the satellite to be used as fabric ports. The type is
    // slice of NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort.
    CandidateFabricPort []NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort
}

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetFilter() yfilter.YFilter { return candidateFabricPorts.YFilter }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) SetFilter(yf yfilter.YFilter) { candidateFabricPorts.YFilter = yf }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetGoName(yname string) string {
    if yname == "candidate-fabric-port" { return "CandidateFabricPort" }
    return ""
}

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetSegmentPath() string {
    return "candidate-fabric-ports"
}

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-fabric-port" {
        for _, c := range candidateFabricPorts.CandidateFabricPort {
            if candidateFabricPorts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort{}
        candidateFabricPorts.CandidateFabricPort = append(candidateFabricPorts.CandidateFabricPort, child)
        return &candidateFabricPorts.CandidateFabricPort[len(candidateFabricPorts.CandidateFabricPort)-1]
    }
    return nil
}

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range candidateFabricPorts.CandidateFabricPort {
        children[candidateFabricPorts.CandidateFabricPort[i].GetSegmentPath()] = &candidateFabricPorts.CandidateFabricPort[i]
    }
    return children
}

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetBundleName() string { return "cisco_ios_xr" }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetYangName() string { return "candidate-fabric-ports" }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) SetParent(parent types.Entity) { candidateFabricPorts.parent = parent }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetParent() types.Entity { return candidateFabricPorts.parent }

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetParentYangName() string { return "nv-satellite" }

// NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort
// Enable interfaces on the satellite to be used
// as fabric ports
type NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PortType interface{}

    // This attribute is a key. Slot. The type is interface{} with range: 0..8.
    Slot interface{}

    // This attribute is a key. Sub slot. The type is interface{} with range:
    // 0..8.
    SubSlot interface{}

    // Port range. The type is string. This attribute is mandatory.
    PortRange interface{}
}

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetFilter() yfilter.YFilter { return candidateFabricPort.YFilter }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) SetFilter(yf yfilter.YFilter) { candidateFabricPort.YFilter = yf }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetGoName(yname string) string {
    if yname == "port-type" { return "PortType" }
    if yname == "slot" { return "Slot" }
    if yname == "sub-slot" { return "SubSlot" }
    if yname == "port-range" { return "PortRange" }
    return ""
}

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetSegmentPath() string {
    return "candidate-fabric-port" + "[port-type='" + fmt.Sprintf("%v", candidateFabricPort.PortType) + "']" + "[slot='" + fmt.Sprintf("%v", candidateFabricPort.Slot) + "']" + "[sub-slot='" + fmt.Sprintf("%v", candidateFabricPort.SubSlot) + "']"
}

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-type"] = candidateFabricPort.PortType
    leafs["slot"] = candidateFabricPort.Slot
    leafs["sub-slot"] = candidateFabricPort.SubSlot
    leafs["port-range"] = candidateFabricPort.PortRange
    return leafs
}

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetBundleName() string { return "cisco_ios_xr" }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetYangName() string { return "candidate-fabric-port" }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) SetParent(parent types.Entity) { candidateFabricPort.parent = parent }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetParent() types.Entity { return candidateFabricPort.parent }

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetParentYangName() string { return "candidate-fabric-ports" }

// NvSatellites_NvSatellite_ConnectionInfo
// Satellite User
type NvSatellites_NvSatellite_ConnectionInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite Username. The type is string.
    Username interface{}

    // Encrypted password for the user. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}
}

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetFilter() yfilter.YFilter { return connectionInfo.YFilter }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) SetFilter(yf yfilter.YFilter) { connectionInfo.YFilter = yf }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    if yname == "password" { return "Password" }
    return ""
}

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetSegmentPath() string {
    return "connection-info"
}

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["username"] = connectionInfo.Username
    leafs["password"] = connectionInfo.Password
    return leafs
}

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetBundleName() string { return "cisco_ios_xr" }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetYangName() string { return "connection-info" }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) SetParent(parent types.Entity) { connectionInfo.parent = parent }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetParent() types.Entity { return connectionInfo.parent }

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetParentYangName() string { return "nv-satellite" }

// NvSatellites_NvSatellite_Redundancy
// Redundancy submode
type NvSatellites_NvSatellite_Redundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Priority for this host for the given satellite. The type is interface{}
    // with range: 0..255.
    HostPriority interface{}
}

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetFilter() yfilter.YFilter { return redundancy.YFilter }

func (redundancy *NvSatellites_NvSatellite_Redundancy) SetFilter(yf yfilter.YFilter) { redundancy.YFilter = yf }

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetGoName(yname string) string {
    if yname == "host-priority" { return "HostPriority" }
    return ""
}

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetSegmentPath() string {
    return "redundancy"
}

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-priority"] = redundancy.HostPriority
    return leafs
}

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetBundleName() string { return "cisco_ios_xr" }

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetYangName() string { return "redundancy" }

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redundancy *NvSatellites_NvSatellite_Redundancy) SetParent(parent types.Entity) { redundancy.parent = parent }

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetParent() types.Entity { return redundancy.parent }

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetParentYangName() string { return "nv-satellite" }

